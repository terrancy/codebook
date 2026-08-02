package creational

import "sync"

// Singleton 单例模式：保证一个类型全局只有一个实例，并提供统一访问点。
//
// Go 推荐用 sync.Once 实现懒加载且并发安全的单例，比双重检查锁更简洁可靠。

type singleton struct {
	name string
}

var (
	instance *singleton
	once     sync.Once
)

// GetInstance 返回全局唯一实例；多次调用只会初始化一次。
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{name: "unique"}
	})
	return instance
}

func (s *singleton) Name() string { return s.name }
