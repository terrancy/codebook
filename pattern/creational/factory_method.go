package creational

// 工厂方法模式（Factory Method）：定义一个创建对象的接口，让子类决定实例化哪一个产品类。
//
// 关键点：把“new 哪个具体类型”的抉择延迟到子类，调用方只依赖抽象 Product。

// Product 抽象产品：所有日志输出都实现 Write 方法。
type Product interface {
	Write(msg string) string
}

// 具体产品 A：控制台日志
type ConsoleLogger struct{}

func (ConsoleLogger) Write(msg string) string { return "[console] " + msg }

// 具体产品 B：文件日志（简化，仅返回标记）
type FileLogger struct{}

func (FileLogger) Write(msg string) string { return "[file] " + msg }

// Creator 抽象创建者，声明工厂方法。
type Creator interface {
	Create() Product
	Log(msg string) string // 业务方法，依赖工厂方法拿到产品
}

// 具体创建者 A：生产 ConsoleLogger
type ConsoleCreator struct{}

func (ConsoleCreator) Create() Product { return ConsoleLogger{} }
func (c ConsoleCreator) Log(msg string) string {
	return c.Create().Write(msg)
}

// 具体创建者 B：生产 FileLogger
type FileCreator struct{}

func (FileCreator) Create() Product { return FileLogger{} }
func (f FileCreator) Log(msg string) string {
	return f.Create().Write(msg)
}
