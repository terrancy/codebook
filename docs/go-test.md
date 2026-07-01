# Go Test 常见命令

## 基本用法

```bash
go test ./...                    # 运行当前项目所有测试
go test ./test/...               # 运行 test 目录下所有测试
go test ./tree/                  # 运行 tree 包的测试（仅限包内测试）
go test ./test/ -v               # 详细输出，显示每个用例的 PASS/FAIL
```

## 指定测试函数

```bash
go test ./test/ -run TestWidthOfBinaryTree -v          # 精确匹配
go test ./test/ -run TestFind -v                       # 正则匹配，运行所有 TestFind* 测试
go test ./test/ -run "TestFindPath|TestSerialize" -v   # 匹配多个测试（或逻辑）
```

`-run` 参数使用正则表达式匹配测试函数名，不是精确匹配。

## 运行单个子测试

```bash
go test ./test/ -run "TestWidthOfBinaryTree/test_case_1" -v
```

## 测试控制

```bash
go test ./test/ -count=1         # 禁用缓存，强制重新运行
go test ./test/ -timeout 30s     # 设置超时时间
go test ./test/ -short           # 跳过耗时较长的测试（代码中调用 testing.Short() 判断）
go test ./test/ -failfast        # 遇到失败立即停止
```

## 并发与性能

```bash
go test ./test/ -cpu 1,2,4       # 分别用 1、2、4 个 CPU 运行测试
go test ./test/ -race             # 启用竞态检测
go test ./test/ -cover            # 输出覆盖率概览
go test ./test/ -coverprofile=coverage.out   # 生成覆盖率文件
go tool cover -html=coverage.out             # 浏览器查看覆盖率报告
go tool cover -func=coverage.out             # 终端查看函数级覆盖率
```

## 基准测试

```bash
go test ./test/ -bench BenchmarkXxx -benchmem   # 运行基准测试并显示内存分配
go test ./test/ -bench . -benchtime 3s           # 所有基准测试，每个运行 3 秒
go test ./test/ -bench . -run ^$                 # 只跑基准测试，跳过普通测试
```

## 编译检查

```bash
go test -c ./test/               # 仅编译测试，不运行（生成 .test 二进制）
go test -v -list Test.* ./test/  # 列出匹配的测试函数名，不执行
```

## 常见问题

| 问题 | 原因 | 解决 |
|------|------|------|
| `go test TestFoo` 报错 | 缺少包路径 | 改为 `go test ./test/ -run TestFoo` |
| 测试结果被缓存 | Go 默认缓存测试结果 | 加 `-count=1` |
| `_test.go` 文件不在 `xxx_test` 包 | 测试文件包名与源码相同 | 包内测试，可访问未导出符号 |
| `undefined: xxx` | 测试文件 import 缺失 | 确认 import 路径正确 |