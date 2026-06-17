# Go 并发同步原语指南

## WaitGroup

### 一、基本用法

```go
var wg sync.WaitGroup

for i := 0; i < n; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        // 执行任务
    }()
}
wg.Wait()
```

### 二、通用写法

#### 1. 基本并发收集结果

```go
func CollectResults(taskNum int) []interface{} {
    var (
        dataCh = make(chan interface{})
        resp   = make([]interface{}, 0, taskNum)
    )

    go func() {
        var wg sync.WaitGroup
        for i := 0; i < taskNum; i++ {
            wg.Add(1)
            go func(ch chan<- interface{}, val int) {
                defer wg.Done()
                ch <- val
            }(dataCh, i)
        }
        wg.Wait()
        close(dataCh)
    }()

    for val := range dataCh {
        resp = append(resp, val)
    }
    return resp
}
```

**要点：** 在独立 goroutine 中 `wg.Wait()` + `close(channel)`，主 goroutine 通过 `range` 消费数据。

#### 2. WaitGroup + 超时控制

```go
func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
    done := make(chan struct{})
    go func() {
        wg.Wait()
        close(done)
    }()

    timer := time.NewTimer(timeout)
    defer timer.Stop()

    select {
    case <-done:
        return false
    case <-timer.C:
        return true
    }
}
```

**要点：** 使用 `time.NewTimer` + `defer timer.Stop()` 避免资源泄漏，不要用 `time.After`。

#### 3. WaitGroup + Context 取消

```go
func ProcessWithCancel(ctx context.Context, tasks []Task) {
    var wg sync.WaitGroup
    for _, task := range tasks {
        wg.Add(1)
        go func(t Task) {
            defer wg.Done()
            select {
            case <-ctx.Done():
                return
            default:
                // 执行任务
            }
        }(task)
    }
    wg.Wait()
}
```

**要点：** `ctx.Done()` 仅用于通知提前退出，不能阻塞工作流程本身。

### 三、常见错误

#### ❌ 错误1：Add 在 goroutine 内部调用

```go
// 错误：可能 wg.Wait() 先于 wg.Add(1) 执行
for i := 0; i < n; i++ {
    go func() {
        wg.Add(1)
        defer wg.Done()
    }()
}
wg.Wait()
```

```go
// 正确：Add 必须在启动 goroutine 之前调用
for i := 0; i < n; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
    }()
}
wg.Wait()
```

#### ❌ 错误2：WaitGroup 值拷贝

```go
// 错误：WaitGroup 不能拷贝，拷贝后内部计数器不共享
func doWork(wg sync.WaitGroup) {
    defer wg.Done()
}
```

```go
// 正确：必须传指针
func doWork(wg *sync.WaitGroup) {
    defer wg.Done()
}
```

#### ❌ 错误3：Done 调用次数超过 Add

```go
// 错误：Add(1) 但 Done() 被调用两次，计数器变为负数 → panic
wg.Add(1)
go func() {
    defer wg.Done()
    wg.Done() // panic: sync: WaitGroup is reused before previous Wait has returned
}()
```

#### ❌ 错误4：WaitGroup 与 Channel 阻塞形成死锁

```go
// 错误：goroutine 阻塞在 <-ch，无法执行 Done，Wait 永远不返回
wg.Add(1)
go func(ch <-chan struct{}) {
    defer wg.Done()
    <-ch        // 阻塞在这里
    fmt.Println("work")
}(ch)
wg.Wait()      // 永远阻塞
close(ch)      // 永远执行不到
```

```go
// 正确：先执行工作，再检查取消信号
wg.Add(1)
go func() {
    defer wg.Done()
    // 先做工作
    fmt.Println("work")
    // 再检查取消
    select {
    case <-ctx.Done():
        return
    default:
    }
}()
wg.Wait()
```

**核心原则：** `WaitGroup` 计数器必须在工作完成后递减，取消信号不能阻塞工作本身。

#### ❌ 错误5：time.After 在循环中使用导致泄漏

```go
// 错误：每次循环创建 Timer，未超时的 Timer 不会被回收
for {
    select {
    case <-done:
        return
    case <-time.After(5 * time.Second):
        return
    }
}
```

```go
// 正确：使用 time.NewTimer + defer Stop
timer := time.NewTimer(5 * time.Second)
defer timer.Stop()
select {
case <-done:
    return
case <-timer.C:
    return
}
```

#### ❌ 错误6：重复 Add 导致计数器不准

```go
// 错误：循环中多次 Add，但 goroutine 内部又嵌套启动 goroutine 并 Add
wg.Add(1)
go func() {
    defer wg.Done()
    wg.Add(1) // 嵌套 Add，容易导致计数器管理混乱
    go func() {
        defer wg.Done()
    }()
}()
```

```go
// 正确：一次性计算好任务数量，统一 Add
wg.Add(2)
go func() {
    defer wg.Done()
    go func() {
        defer wg.Done()
    }()
}()
```

### 四、要点速查

| 规则 | 说明 |
|------|------|
| `Add` 在 `go` 之前 | 确保 Wait 能正确等待 |
| `Done` 在 `defer` 中调用 | 防止 panic 导致计数器不递减 |
| 传指针不传值 | WaitGroup 内部状态不可拷贝 |
| `Add` 与 `Done` 数量一致 | 否则 panic 或死锁 |
| 取消信号不阻塞工作 | 否则形成死锁 |
| 循环中用 `NewTimer` | 避免资源泄漏 |