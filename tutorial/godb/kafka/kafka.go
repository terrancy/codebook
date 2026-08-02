package kafka

import (
	"context"
	"fmt"
	"sync"

	"github.com/IBM/sarama"
)

// Kafka基础

// ========================================
// 配置相关
// ========================================

// NewKafkaConfig 创建Kafka配置
func NewKafkaConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Consumer.Return.Errors = true
	return config
}

// ========================================
// 生产者
// ========================================

// Producer Kafka生产者封装
type Producer struct {
	producer sarama.SyncProducer
	topic    string
}

// NewProducer 创建生产者
func NewProducer(brokers []string, topic string) (*Producer, error) {
	config := NewKafkaConfig()
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, fmt.Errorf("创建生产者失败: %w", err)
	}
	return &Producer{producer: producer, topic: topic}, nil
}

// Send 发送消息
func (p *Producer) Send(key, value string) (int32, int64, error) {
	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(value),
	}
	return p.producer.SendMessage(msg)
}

// SendWithPartition 发送消息到指定分区
func (p *Producer) SendWithPartition(key, value string, partition int32) (int32, int64, error) {
	msg := &sarama.ProducerMessage{
		Topic:     p.topic,
		Key:       sarama.StringEncoder(key),
		Value:     sarama.ByteEncoder(value),
		Partition: partition,
	}
	return p.producer.SendMessage(msg)
}

// Close 关闭生产者
func (p *Producer) Close() error {
	return p.producer.Close()
}

// ========================================
// 消费者
// ========================================

// Consumer Kafka消费者封装
type Consumer struct {
	consumer sarama.Consumer
	topic    string
}

// NewConsumer 创建消费者
func NewConsumer(brokers []string, topic string) (*Consumer, error) {
	config := NewKafkaConfig()
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, fmt.Errorf("创建消费者失败: %w", err)
	}
	return &Consumer{consumer: consumer, topic: topic}, nil
}

// Consume 消费消息
func (c *Consumer) Consume(partition int32, offset int64, handler func(*sarama.ConsumerMessage) error) error {
	partitionConsumer, err := c.consumer.ConsumePartition(c.topic, partition, offset)
	if err != nil {
		return fmt.Errorf("创建分区消费者失败: %w", err)
	}
	defer partitionConsumer.Close()

	for msg := range partitionConsumer.Messages() {
		if err := handler(msg); err != nil {
			return err
		}
	}
	return nil
}

// Close 关闭消费者
func (c *Consumer) Close() error {
	return c.consumer.Close()
}

// ========================================
// 消费者组
// ========================================

// ConsumerGroup 消费者组封装
type ConsumerGroup struct {
	group sarama.ConsumerGroup
	topic string
}

// NewConsumerGroup 创建消费者组
func NewConsumerGroup(brokers []string, topic, groupID string) (*ConsumerGroup, error) {
	config := NewKafkaConfig()
	group, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, fmt.Errorf("创建消费者组失败: %w", err)
	}
	return &ConsumerGroup{group: group, topic: topic}, nil
}

// Consume 消费消息
func (cg *ConsumerGroup) Consume(ctx context.Context, handler sarama.ConsumerGroupHandler) error {
	topics := []string{cg.topic}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := cg.group.Consume(ctx, topics, handler); err != nil {
				return err
			}
		}
	}
}

// Close 关闭消费者组
func (cg *ConsumerGroup) Close() error {
	return cg.group.Close()
}

// ========================================
// 消费者组处理器示例
// ========================================

// ExampleHandler 消费者组处理器示例
type ExampleHandler struct {
	Ready chan bool
}

// Setup 在会话开始前运行
func (h *ExampleHandler) Setup(sarama.ConsumerGroupSession) error {
	close(h.Ready)
	return nil
}

// Cleanup 在会话结束后运行
func (h *ExampleHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim 消费消息
func (h *ExampleHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("消息: topic=%s, partition=%d, offset=%d, key=%s, value=%s\n",
			msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		session.MarkMessage(msg, "")
	}
	return nil
}

// ========================================
// 异步生产者
// ========================================

// AsyncProducer 异步生产者封装
type AsyncProducer struct {
	producer sarama.AsyncProducer
	topic    string
	wg       sync.WaitGroup
}

// NewAsyncProducer 创建异步生产者
func NewAsyncProducer(brokers []string, topic string) (*AsyncProducer, error) {
	config := NewKafkaConfig()
	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, fmt.Errorf("创建异步生产者失败: %w", err)
	}
	ap := &AsyncProducer{producer: producer, topic: topic}

	ap.wg.Add(1)
	go func() {
		defer ap.wg.Done()
		for success := range producer.Successes() {
			fmt.Printf("发送成功: partition=%d, offset=%d\n", success.Partition, success.Offset)
		}
	}()

	ap.wg.Add(1)
	go func() {
		defer ap.wg.Done()
		for err := range producer.Errors() {
			fmt.Printf("发送失败: %v\n", err)
		}
	}()

	return ap, nil
}

// Send 异步发送消息
func (ap *AsyncProducer) Send(key, value string) {
	msg := &sarama.ProducerMessage{
		Topic: ap.topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(value),
	}
	ap.producer.Input() <- msg
}

// Close 关闭异步生产者
func (ap *AsyncProducer) Close() error {
	err := ap.producer.Close()
	ap.wg.Wait()
	return err
}
