## 七大设计原则（Go 视角）

设计原则是设计模式的「底层哲学」。掌握原则，模式是水到渠成的工具。

---

### 1. 单一职责原则 SRP（Single Responsibility Principle）

> 一个类（在 Go 中是一个类型 / 包）应当只有一个引起它变化的原因。

- **核心**：把「会因不同原因而变」的逻辑拆开。职责越多，修改的诱因越多，越易引入 Bug。
- **Go 落地**：一个 `struct` 只承担一类职责；一个 `package` 对外只暴露一个内聚的能力。

```go
// ❌ 违反：用户既要管业务，又要管持久化
type User struct {
    ID   int
    Name string
}
func (u User) Save()        { /* 写数据库 */ }
func (u User) Validate()    { /* 校验规则 */ }

// ✅ 符合：拆分职责
type User struct {
    ID   int
    Name string
}
type UserValidator struct{}
func (UserValidator) Validate(u User) error { /* 只管校验 */ }

type UserRepository struct{}
func (UserRepository) Save(u User) error { /* 只管存储 */ }
```

---

### 2. 开闭原则 OCP（Open/Closed Principle）

> 软件实体应当对扩展开放，对修改关闭。

- **核心**：新功能通过新增代码实现，而非改旧代码。靠「抽象 + 多态」把变化点隔离。
- **Go 落地**：面向接口编程，用实现接口的新类型扩展行为。

```go
type Discount interface { Calc(price float64) float64 }

type NoDiscount struct{}
func (NoDiscount) Calc(p float64) float64 { return p }

type VipDiscount struct{}
func (VipDiscount) Calc(p float64) float64 { return p * 0.8 }

// 新增活动折扣：只加类型，不改 PriceService
type PromoDiscount struct{}
func (PromoDiscount) Calc(p float64) float64 { return p - 10 }

type PriceService struct{ d Discount }
func (s PriceService) Final(p float64) float64 { return s.d.Calc(p) }
```

---

### 3. 里氏替换原则 LSP（Liskov Substitution Principle）

> 子类型必须能够替换掉它们的基类型，且不破坏程序正确性。

- **核心**：如果 `B` 是 `A` 的实现，那么任何用 `A` 的地方换成 `B` 都不能出错。
- **Go 落地**：实现同一接口的不同类型，对外行为（前置/后置条件、异常语义）应一致。不要「为了凑接口」让某个实现偷偷违反约定。

```go
// 矩形接口
type Rect interface {
    SetW(int); SetH(int); Area() int
}
type Rectangle struct{ w, h int }
func (r *Rectangle) SetW(w int) { r.w = w }
func (r *Rectangle) SetH(h int) { r.h = h }
func (r *Rectangle) Area() int  { return r.w * r.h }

// Square 若实现 Rect 但 SetW 同时改 h，就会破坏“宽高独立”的约定 -> 违反 LSP
// 正确做法：Square 不实现 Rect，或保证行为语义一致
```

---

### 4. 接口隔离原则 ISP（Interface Segregation Principle）

> 客户端不应被迫依赖它不需要的接口。

- **核心**：大而全的接口会让实现者被迫实现一堆无用方法。拆成小而专的接口。
- **Go 落地**：Go 的接口天然鼓励「小接口」（`io.Reader`/`io.Writer` 只有一个方法）。定义接口时尽量贴近调用方需求。

```go
// ❌ 违反：一个臃肿接口，调用方只用 Read
type ReaderWriter interface {
    Read(p []byte) (int, error)
    Write(p []byte) (int, error)
    Close() error
}
// ✅ 符合：拆成小接口，按需组合
type Reader interface { Read(p []byte) (int, error) }
type Writer interface { Write(p []byte) (int, error) }
```

---

### 5. 依赖倒置原则 DIP（Dependency Inversion Principle）

> 高层模块不应依赖低层模块，二者都应依赖抽象；抽象不应依赖细节，细节应依赖抽象。

- **核心**：别在业务代码里 `new` 具体实现，而是通过接口注入依赖。
- **Go 落地**：构造函数接收接口参数（依赖注入），便于替换实现与单测 mock。

```go
type Storage interface { Save(k, v string) error }

type MemoryStorage struct{ m map[string]string }
func (m *MemoryStorage) Save(k, v string) error { m.m[k] = v; return nil }

// 高层业务依赖抽象 Storage，而非具体 MemoryStorage
type Service struct{ store Storage }
func NewService(s Storage) *Service { return &Service{store: s} } // 依赖注入
```

---

### 6. 迪米特法则 LoD（Law of Demeter）

> 一个对象应尽可能少地了解其他对象，只与「直接朋友」通信。

- **核心**：不要 `a.getB().getC().doSomething()` 这种「火车残骸」式调用——它把内部结构暴露给调用方，耦合爆表。
- **Go 落地**：通过方法暴露「要做什么」，而不是「怎么拿到的」。

```go
// ❌ 违反：链式穿透，强耦合内部对象关系
customer.GetAccount().GetWallet().Deduct(100)

// ✅ 符合：在 customer 上提供意图方法，内部自己协调
func (c *Customer) Pay(amount int) error {
    return c.account.wallet.Deduct(amount) // 内部封装
}
```

---

### 7. 合成复用原则 CRP（Composite Reuse Principle）

> 尽量使用对象组合 / 委托，而不是类继承来达到复用目的。

- **核心**：继承是「白盒复用」（父类实现对子类可见，耦合强）；组合是「黑盒复用」（只依赖接口，耦合弱）。
- **Go 落地**：Go 没有继承，天然靠组合（结构体嵌入）和委托；这正是它避免「继承爆炸」的优势。

```go
// 通过组合复用日志能力，而非继承 Logger
type Logger struct{}
func (Logger) Log(msg string) {}

type Service struct {
    Logger // 嵌入组合，复用 Log 方法
}
// Service 自动拥有 Log，但二者是“has-a”关系，松耦合
```

---

### 原则之间的关系

- **OCP 是目标**，**DIP / ISP 是实现手段**（靠抽象与注入隔离变化）。
- **SRP 是基础**：职责不清，接口就拆不干净，OCP 也无从谈起。
- **LSP 是约束**：保证「用抽象替换实现」不会翻车。
- **CRP / LoD 降耦合**：减少模块间的了解面，让系统更易改、易测。
