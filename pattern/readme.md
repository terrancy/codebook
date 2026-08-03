## 面向对象设计模式与七大原则

本目录收录面向对象的经典设计模式（Go 语言实现）与七大设计原则，按 GoF 三大类分到子包下：

```
pattern/
├── readme.md          # 本文件：总览
├── principles.md      # 七大设计原则详解（含 Go 示例）
├── creational/        # 创建型（5）：singleton / factory_method / abstract_factory / builder / prototype
├── structural/        # 结构型（7）：adapter / decorator / proxy / bridge / composite / facade / flyweight
└── behavioral/        # 行为型（11）：strategy / observer / template_method / chain_of_responsibility /
                       #              command / iterator / mediator / memento / state / visitor / interpreter
```

> 设计模式不是银弹：它解决的是「变化点」带来的重复修改问题。先识别变化，再套模式，别为了用而用。

### 一、七大设计原则（SOLID + 2）

| 原则 | 英文 | 一句话 | 说明文档 |
| --- | --- | --- | --- |
| 单一职责 | SRP | 一个类只因一个理由变化 | [principles.md](principles.md) |
| 开闭原则 | OCP | 对扩展开放、对修改关闭 | [principles.md](principles.md) |
| 里氏替换 | LSP | 子类必须能替换父类 | [principles.md](principles.md) |
| 接口隔离 | ISP | 不要强迫依赖用不到的接口 | [principles.md](principles.md) |
| 依赖倒置 | DIP | 依赖抽象，不依赖细节 | [principles.md](principles.md) |
| 迪米特法则 | LoD | 只与直接朋友通信，降低耦合 | [principles.md](principles.md) |
| 合成复用 | CRP | 优先组合/委托，而非继承 | [principles.md](principles.md) |

### 二、23 种 GoF 设计模式总览

#### 创建型 `creational/`（5）—— 解决对象创建的问题

创工原, 单抽建

| 模式 | 意图 | 文件 |
| --- | --- | --- |
| 单例 Singleton | 全局唯一实例 | `creational/singleton.go` |
| 工厂方法 Factory Method | 子类决定创建哪种产品 | `creational/factory_method.go` |
| 抽象工厂 Abstract Factory | 创建一组相关/依赖的产品族 | `creational/abstract_factory.go` |
| 建造者 Builder | 分步构造复杂对象，分离构建与表示 | `creational/builder.go` |
| 原型 Prototype | 通过拷贝原型创建新对象 | `creational/prototype.go` |

#### 结构型 `structural/`（7）—— 解决类/对象的组合问题

结享外组, 适代装桥

| 模式 | 意图 | 文件 |
| --- | --- | --- |
| 适配器 Adapter | 转换接口，使不兼容的类协同工作 | `structural/adapter.go` |
| 装饰器 Decorator | 动态给对象加职责，替代继承膨胀 | `structural/decorator.go` |
| 代理 Proxy | 为对象提供替身以控制访问 | `structural/proxy.go` |
| 桥接 Bridge | 抽象与实现分离，各自独立变化 | `structural/bridge.go` |
| 组合 Composite | 树形结构，统一对待叶子与容器 | `structural/composite.go` |
| 外观 Facade | 为复杂子系统提供统一入口 | `structural/facade.go` |
| 享元 Flyweight | 共享细粒度对象，节省内存 | `structural/flyweight.go` |

#### 行为型 `behavioral/`（11）—— 解决对象间的职责与通信问题

形状责中模访,解备观策命迭

| 模式 | 意图 | 文件 |
| --- | --- | --- |
| 策略 Strategy | 封装可互换算法，运行时选择 | `behavioral/strategy.go` |
| 观察者 Observer | 一对多通知，状态变更自动推播 | `behavioral/observer.go` |
| 模板方法 Template Method | 固定算法骨架，步骤延迟到子类 | `behavioral/template_method.go` |
| 责任链 Chain of Responsibility | 请求沿链传递，直到被处理 | `behavioral/chain_of_responsibility.go` |
| 命令 Command | 把请求封装成对象，支持撤销/队列 | `behavioral/command.go` |
| 迭代器 Iterator | 顺序访问聚合对象而不暴露结构 | `behavioral/iterator.go` |
| 中介者 Mediator | 用中介对象解耦一组对象 | `behavioral/mediator.go` |
| 备忘录 Memento | 捕获并恢复对象内部状态 | `behavioral/memento.go` |
| 状态 State | 对象状态改变时改变行为 | `behavioral/state.go` |
| 访问者 Visitor | 在不改类的前提下新增操作 | `behavioral/visitor.go` |
| 解释器 Interpreter | 定义文法并解释语言 | `behavioral/interpreter.go` |

### 三、Go 语言实现模式的几点提醒

- **没有继承，只有组合**：Go 用嵌入（embedding）和结构体内嵌模拟「is-a」，用接口模拟多态。
- **接口是隐式实现**：类型只要方法集满足接口即可赋值，无需 `implements` 声明。这让「依赖倒置」「接口隔离」在 Go 中成本极低。
- **没有类，只有结构体 + 方法**：职责划分落在包（package）与接收者方法上，单一职责等价于「包职责单一」。
- **多态靠接口**：真正多态用具体接口而非空接口（`any`），避免运行时类型断言满天飞。

### 参考

- GoF《设计模式：可复用面向对象软件的基础》
- refactoring.guru 设计模式图解（含多语言示例）
