# go.uber.org/goleak 使用手册

## 概述

`go.uber.org/goleak` 是 Uber 开源的用于检测 Go 程序中 goroutine 泄漏的测试工具库。它能够在测试结束时自动检查是否存在未正确清理的 goroutine，帮助开发者及时发现并发代码中的资源泄漏问题。

## 安装

```bash
go get go.uber.org/goleak@latest
```

## 基本使用方式

### 方式一：包装测试主函数（推荐）

在 `_test.go` 文件中添加 `TestMain` 函数：

```go
package mypackage_test

import (
    "testing"
    "go.uber.org/goleak"
)

func TestMain(m *testing.M) {
    goleak.VerifyTestMain(m)
}
```

这种方式会在所有测试执行完毕后统一检查 goroutine 泄漏。

### 方式二：在单个测试中使用

```go
func TestMyFunction(t *testing.T) {
    defer goleak.VerifyNone(t)
    
    // 你的测试代码
    go someBackgroundTask()
}
```

这种方式适用于需要单独检测的特定测试用例。

## 核心 API

### 1. VerifyNone

```go
func VerifyNone(t testing.TB, options ...Option)
```

检查当前 goroutine 是否有泄漏，检测到泄漏时会调用 `t.Fatal` 终止测试。

### 2. VerifyTestMain

```go
func VerifyTestMain(m *testing.M, options ...Option)
```

包装 `testing.M`，在所有测试完成后检查泄漏并返回退出码。

### 3. Find

```go
func Find(options ...Option) []string
```

返回所有泄漏的 goroutine 堆栈信息，适用于需要自定义处理逻辑的场景。

## 配置选项

### IgnoreTopFunction

忽略特定顶层函数创建的 goroutine：

```go
func TestMain(m *testing.M) {
    goleak.VerifyTestMain(m,
        goleak.IgnoreTopFunction("net/http.(*Server).Serve"),
        goleak.IgnoreTopFunction("github.com/example/package.backgroundWorker"),
    )
}
```

### IgnoreCurrent

忽略当前正在运行的 goroutine：

```go
func TestWithIgnoredGoroutine(t *testing.T) {
    defer goleak.VerifyNone(t, goleak.IgnoreCurrent())
    
    // 当前测试 goroutine 不会被检测
}
```

### WithTimeout

设置等待 goroutine 退出的超时时间（默认 1 秒）：

```go
import "time"

func TestMain(m *testing.M) {
    goleak.VerifyTestMain(m,
        goleak.WithTimeout(5*time.Second),
    )
}
```

## 实际案例

### 检测 HTTP Server 泄漏

```go
package server_test

import (
    "net/http"
    "testing"
    "go.uber.org/goleak"
)

func TestMain(m *testing.M) {
    goleak.VerifyTestMain(m,
        goleak.IgnoreTopFunction("net/http.(*Server).Serve"),
    )
}

func TestHTTPServer(t *testing.T) {
    defer goleak.VerifyNone(t)
    
    server := &http.Server{Addr: ":8080"}
    go server.ListenAndServe()
    
    // 测试逻辑
    // ...
    
    // 正确关闭服务器
    server.Close()
}
```

### 检测定时器泄漏

```go
func TestTimerLeak(t *testing.T) {
    defer goleak.VerifyNone(t)
    
    // 错误示例：未停止的定时器会导致泄漏
    ticker := time.NewTicker(100 * time.Millisecond)
    
    // 正确做法：使用 defer 停止定时器
    // defer ticker.Stop()
    
    <-ticker.C
}
```

## 输出示例

检测到泄漏时的典型输出：

```
=== RUN   TestTimerLeak
    goroutine leak detected:
    goroutine 42 [running]:
    time.Sleep(0x3b9aca00)
        /usr/local/go/src/runtime/time.go:195 +0x15
    created by time.newTicker.func1
        /usr/local/go/src/time/tick.go:23 +0x45
FAIL    example.com/package  TestTimerLeak  0.102s
```

## 最佳实践

1. **全局检测**：在 `TestMain` 中使用 `VerifyTestMain` 进行全局检测
2. **精确控制**：对特殊场景使用 `VerifyNone` 配合 `IgnoreTopFunction`
3. **排除合法 goroutine**：某些后台服务（如 metrics collector）需要显式排除
4. **与 CI 集成**：确保所有测试通过 goroutine 泄漏检测

## 注意事项

- **性能影响**：检测会等待 goroutine 退出，可能增加测试时间
- **误报处理**：使用 `IgnoreTopFunction` 排除已知的后台 goroutine
- **生产环境**：仅用于测试阶段，不应在生产代码中使用

## 参考链接

- GitHub 仓库: https://github.com/uber-go/goleak
- GoDoc: https://pkg.go.dev/go.uber.org/goleak
