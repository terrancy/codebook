package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/IBM/sarama"
)

// ========================================
// 使用示例
// ========================================

func ExampleProducer() {
	brokers := []string{"localhost:9092"}
	topic := "test-topic"

	producer, err := NewProducer(brokers, topic)
	if err != nil {
		fmt.Println("创建生产者失败:", err)
		return
	}
	defer producer.Close()

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key-%d", i)
		value := fmt.Sprintf("value-%d", i)
		partition, offset, err := producer.Send(key, value)
		if err != nil {
			fmt.Printf("发送失败: %v\n", err)
			continue
		}
		fmt.Printf("发送成功: partition=%d, offset=%d\n", partition, offset)
	}
}

func ExampleConsumer() {
	brokers := []string{"localhost:9092"}
	topic := "test-topic"

	consumer, err := NewConsumer(brokers, topic)
	if err != nil {
		fmt.Println("创建消费者失败:", err)
		return
	}
	defer consumer.Close()

	err = consumer.Consume(0, sarama.OffsetNewest, func(msg *sarama.ConsumerMessage) error {
		fmt.Printf("收到消息: key=%s, value=%s\n", string(msg.Key), string(msg.Value))
		return nil
	})
	if err != nil {
		fmt.Println("消费失败:", err)
	}
}

func ExampleConsumerGroup() {
	brokers := []string{"localhost:9092"}
	topic := "test-topic"
	groupID := "test-group"

	cg, err := NewConsumerGroup(brokers, topic, groupID)
	if err != nil {
		fmt.Println("创建消费者组失败:", err)
		return
	}
	defer cg.Close()

	handler := &ExampleHandler{Ready: make(chan bool)}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	go func() {
		<-handler.Ready
		fmt.Println("消费者组已就绪")
	}()

	if err := cg.Consume(ctx, handler); err != nil {
		fmt.Println("消费失败:", err)
	}
}
