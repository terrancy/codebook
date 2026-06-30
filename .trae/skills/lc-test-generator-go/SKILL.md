---
name: "lc-test-generator-go"
description: "Adds LeetCode problems (Go source code + table-driven unit tests) to the Go project. Invoke when the current project is a Go project (has go.mod), or user mentions '添加力扣题go', '写测试go', 'go单元测试', or explicitly asks for Go."
---

# LeetCode 题目添加与测试生成器（Go 版）

根据指定的力扣题号或题目，自动完成：Go 源码函数编写 + table-driven 单元测试生成。

项目模块路径：`terrancy/awesome`

## 触发条件

- 用户要求添加某道力扣题目（如"添加力扣第6题"、"添加力扣题150"）
- 用户要求为某个力扣题目编写单元测试
- 用户要求为某个算法函数添加测试用例
- 用户提到"写测试"、"加测试"、"单元测试"等关键词
- 用户提到"重构测试"、"规范测试"、"改造测试"、"测试不符合规范"、"测试规范化"等关键词
- 用户提到"重构源码注释"、"规范注释"、"注释不符合规范"、"注释规范化"等关键词
- 当前项目为 Go 项目时自动适用此 skill

## 题目分类映射

根据力扣题目类型，确定应放入哪个源码目录和测试文件：

| 题目类型 | 源码目录 | 源码包名 | 测试文件 |
|---------|---------|---------|--------|
| 数组 | `array/` | `array` | `test/array_test.go` |
| 回溯 | `backtrack/` | `backtrack` | `test/backtrack_test.go` |
| 堆 | `heap/` | `heap` | `test/heap_test.go` |
| 链表 | `link/` | `link` | `test/link_test.go` |
| 树 | `tree/` | `trees` | `test/tree_test.go` |
| 图 | `graph/` | `graph` | `test/graph_test.go` |
| 动态规划 | `dp/` | `dp` | `test/dp_test.go` |
| 字符串 | `string/` | `strings`（导入时别名） | `test/string_test.go` |
| 栈 | `stack/` | `stack` | `test/stack_test.go` |
| 队列 | `queue/` | `queue` | `test/queue_test.go`（注：当前无此文件） |
| 滑动窗口 | `slide/` | `slide` | `test/slide_test.go` |
| 其他/位运算/数学/哈希 | `others/` | `others` | `test/other_test.go` |
| 挑战题 | `challenge/` | `challenge` | `test/challenge_*_test.go` |

如果用户指定了分类（如"添加力扣题6到数组分类"），以用户指定为准。

## 源码函数规范

### 1. 函数签名

```go
func FunctionName(param1 type1, param2 type2) returnType {
```

- 函数名使用 PascalCase（Go 导出规范），从力扣题目英文名转换而来
- 首字母大写表示导出函数
- 参数和返回值必须带类型声明
- 左花括号 `{` 不换行，与函数签名同行（Go 强制规范）

### 2. 注释格式

```go
// FunctionName
// @Title: LC<题号>.<中文题目名>
// @Description: <题目简述>
// @Link: https://leetcode.cn/problems/<slug>/
// @param param1
// @param param2
// @return returnType
func FunctionName(param1 type1, param2 type2) returnType {
```

- `@Title:` 格式固定为 `LC<题号>.<中文题目名>`
- `@Link:` 使用力扣中国站链接
- `@Description:` 可多行，每行以 `// @Description:` 开头
- `@param` 和 `@return` 遵循 Go doc 规范

### 3. 插入位置

- 函数追加到对应源码文件的**末尾**
- 如果源码文件不存在，在对应目录下创建新文件，文件名使用 camelCase，如 `findPath.go`、`twoSum.go`
- 同一题目有多个解法时，可放在同一文件中，函数名加后缀区分（如 `FindPathIII` 和 `FindPathIIIPrefixSum`）

### 4. 函数实现

- 如果用户未提供实现，函数体写 `panic("not implemented")`，等待后续补充
- 如果用户提供了实现，按用户代码写入

## 测试规范

### 1. 测试数据格式（table-driven）

测试数据使用**包级变量**，命名格式为 `<functionName>Cases`，类型为匿名结构体切片：

```go
var findPathIIICases = []struct {
	name     string
	data     []int
	k        int
	expected int
}{
	{
		name:     "test_case_1",
		data:     []int{10, 5, -3, 3, 2, awesome.INF, 11, 3, -2, awesome.INF, 1},
		k:        8,
		expected: 3,
	},
	{
		name:     "test_case_2",
		data:     []int{5, 4, 8, 11, awesome.INF, 13, 4, 7, 2, awesome.INF, awesome.INF, 5, 1},
		k:        22,
		expected: 3,
	},
}
```

关键规则：

- 变量名格式：`<functionName>Cases`（CamelCase + Cases 后缀）
- 结构体字段：`name` 固定 + 与源码函数参数名一致的字段 + `expected` 固定表示期望输出
- `name` 字段固定使用 `"test_case_1"`, `"test_case_2"`, ... 格式
- 树类型题目用 `[]int` 表示层序遍历数组，`awesome.INF`（值 10001）表示空节点
- 链表类型题目用 `[]int` 表示链表值序列

### 2. 测试函数格式

```go
// TestFindPathIII LC437.路径总和III
func TestFindPathIII(t *testing.T) {
	for _, tt := range findPathIIICases {
		t.Run(tt.name, func(t *testing.T) {
			root := trees.BuildTreeNode(tt.data)
			res := trees.FindPathIII(root, tt.k)
			assert.Equal(t, tt.expected, res)
		})
	}
}
```

关键规则：

- 函数名格式：`Test<FunctionName>`
- 注释格式：`// Test<FunctionName> LC<题号>.<中文题目描述>`
- 使用 `for _, tt := range cases` + `t.Run(tt.name, ...)` 模式
- 使用 `assert.Equal(t, tt.expected, res)` 断言（来自 `github.com/stretchr/testify/assert`）
- 同一题目的多个解法共享同一组 cases，各自写独立 Test 函数

### 3. 不同数据类型的测试构造

**树类型题目：**
```go
root := trees.BuildTreeNode(tt.data)
res := trees.FunctionName(root, tt.param)
```

**链表类型题目：**
```go
head := link.BuildListNode(tt.data)
res := link.FunctionName(head, tt.param)
```

**数组/基本类型题目：**
```go
res := array.FunctionName(tt.data, tt.param)
```

**字符串类型题目：**
```go
res := strings.FunctionName(tt.str, tt.param)
```
注意：`string` 包导入时需用别名 `strings`，即 `import strings "terrancy/awesome/string"`

### 4. 无序比较

当返回值为无序切片（如两数之和、字母异位词分组等）时，先排序再比较：

```go
import "sort"

sort.Ints(res)
sort.Ints(tt.expected)
assert.Equal(t, tt.expected, res)
```

或对于二维切片：
```go
sort.Slice(res, func(i, j int) bool { /* 排序逻辑 */ })
```

### 5. 测试数据来源

- 优先使用力扣官方示例作为测试数据
- 至少包含 2-3 个测试用例
- 包含边界情况（空输入、单元素等）

### 6. 测试文件结构

- 测试文件统一放在 `test/` 目录下
- 包名固定为 `package test`
- import 区域需包含：
  - `"testing"`
  - `"github.com/stretchr/testify/assert"`（使用 assert 时）
  - `"fmt"`（使用 fmt.Println 时）
  - `"terrancy/awesome"`（使用 awesome.INF 时）
  - 对应源码包的导入路径
- 新增测试追加到对应测试文件**末尾**
- 同一题目的 cases 变量紧跟在相关注释/分隔之后，Test 函数紧跟 cases 之后

## 执行步骤

### 场景A：添加新力扣题目（源码 + 测试）

1. 根据题号查询力扣题目信息（题目名、描述、分类、示例）
2. 根据题目类型或用户指定分类，确定目标源码目录和测试文件
3. 在源码目录中创建或追加函数（含注释），函数追加到文件末尾
4. 在对应测试文件末尾添加 cases 变量和 Test 函数
5. 确保测试文件的 import 区域包含所需依赖
6. 测试数据优先使用力扣官方示例，至少 2-3 个用例

### 场景B：仅为已有函数添加测试

1. 读取目标函数的源码，提取函数签名、`@Title:` 字段、参数名和类型
2. 确定函数所属的源码包和对应的测试文件
3. 读取对应测试文件，确认 import 是否完整
4. 根据力扣题目生成测试数据（至少包含力扣官方示例）
5. 按照测试规范生成 cases 变量和 Test 函数
6. 将测试代码追加到测试文件末尾
7. 如需新增 import，确保不破坏已有 import 声明
8. 测试数据优先使用力扣官方示例，至少 2-3 个用例

### 场景C：重构已有测试为规范格式

将不符合 table-driven 规范的旧测试（如 `fmt.Println` 输出式、硬编码单断言式等）重构为标准格式。

1. 读取目标测试文件，识别所有不符合规范的测试函数
2. 不符合规范的典型特征：
   - 使用 `fmt.Println` 输出结果而非 `assert` 断言
   - 测试数据硬编码在函数体内，未提取为 `cases` 变量
   - 缺少 `t.Run` 子测试结构
   - 缺少 `name` 字段和 `expected` 字段
3. 读取对应源码函数，提取函数签名、参数名和返回类型
4. 根据力扣题目信息或源码注释中的 `@Title:` / `@Link:` 获取题目详情和官方示例
5. 将硬编码的测试数据 + 力扣官方示例合并，生成规范的 `cases` 变量和 `Test` 函数
6. 替换原有不规范测试，确保新测试覆盖原有测试的所有场景
7. 测试数据优先使用力扣官方示例，至少 2-3 个用例
8. 清理不再需要的 import（如移除仅被旧测试使用的 `"fmt"`）
9. 确保重构后 `go test ./test/...` 全部通过

### 场景D：重构源码注释为规范格式

将不符合注释规范的源码函数调整为标准格式，并优先使用力扣链接。

1. 读取目标源码文件，识别所有不符合注释规范的函数
2. 不符合规范的典型特征：
   - 缺少 `@Title:` 字段
   - `@Link` 缺少冒号（应为 `@Link:`）
   - 使用牛客网等非力扣链接，但该题在力扣上存在
   - `@Description:` 格式不统一或缺失
   - 函数名与注释不一致
   - 左花括号 `{` 换行（Go 规范要求 `{` 与签名同行）
3. 规范注释格式：
   ```go
   // FunctionName
   // @Title: LC<题号>.<中文题目名>
   // @Description: <题目简述>
   // @Link: https://leetcode.cn/problems/<slug>/
   // @param param1
   // @param param2
   // @return returnType
   ```
4. 链接替换规则：
   - 优先查找力扣中国站链接（`https://leetcode.cn/problems/<slug>/`）
   - 如果该题在力扣上存在，替换为力扣链接
   - 如果该题在力扣上不存在（如纯牛客题），保留原链接
   - 常见对应关系：牛客 NC 题号 → 力扣 LC 题号，需搜索确认
5. `@Title:` 补充规则：
   - 如果原注释无 `@Title:`，根据题目信息补充，格式 `LC<题号>.<中文题目名>`
   - 如果原注释已有 `@Title:`，如果找到力扣题目信息，更新 `@Title:` 字段，否则保持不变
6. `@Description:` 如果找到力扣题目信息，更新为力扣相应内容，否则保留原有描述内容，仅修正格式
7. 修正代码格式问题（如左花括号换行等 Go 规范违规）
8. 检查对应测试文件中是否已有该函数的测试用例：
   - 如果缺少测试用例，按照场景B的规范补充 cases 变量和 Test 函数
   - 测试数据优先使用力扣官方示例，至少 2-3 个用例
   - 追加到对应测试文件末尾
9. 确保重构后代码仍可正常编译，测试全部通过