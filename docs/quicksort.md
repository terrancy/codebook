# 快速排序

## 一、核心思想

分治策略：选一个基准值 pivot，将数组划分为「小于基准 | 基准 | 大于基准」三部分，然后对左右两个子数组递归执行同样的操作。

时间复杂度平均 O(nlogn)，最坏 O(n²)（已排序数组且选首元素为基准），空间复杂度 O(1)（原地排序），不稳定。

---

## 二、朴素版 & 问题分析

### 朴素版代码

```go
func QuickSortOld(nums []int) []int {
    n := len(nums)
    if n == 0 {
        return nil
    }
    left, right := make([]int, 0), make([]int, 0)
    for i := 1; i < n; i++ {
        if nums[i] < nums[0] {
            left = append(left, nums[i])
        } else {
            right = append(right, nums[i])
        }
    }
    left = QuickSortOld(left)
    left = append(left, nums[0])
    right = QuickSortOld(right)
    return append(left, right...)
}
```

### 朴素版存在的问题

| # | 问题 | 严重程度 | 具体表现 |
|---|---|---|---|
| 1 | **空间复杂度 O(n)** | 高 | 每次递归都新建 left/right 切片，n 总元素反复 append，峰值空间 ≈ O(n)。应该原地排序。 |
| 2 | **基准固定选首元素** | 高 | 已排序/接近有序数组上每次 partition 后基准都在最左，左子数组为空，右子数组长度 n-1，退化为 O(n²)。 |
| 3 | **双递归无保护** | 高 | Go 的 gc 编译器不自动做尾递归优化。已排序数组上递归嵌套 n 层 → 栈溢出。 |
| 4 | **小数组慢** | 中 | 递归的函数调用开销 + partition 常数因子，在子数组长度 < 10 时不如插入排序。 |
| 5 | **大量重复值低效** | 中 | 等于 pivot 的元素被打散到左右两边，下一轮还要重复处理同样的值。 |

---

## 三、优化路线图

### 问题 → 方案映射

| 问题 | 优化方案 | 权衡考量 |
|---|---|---|
| ① 空间 O(n) → 原地排序 | 用 Lomuto 双向扫描分区，原地交换代替 append | 需要把基准先"藏"起来作为中转站，但换来了空间 O(1) |
| ② 基准固定 → O(n²) 退化 | 三数取中选基准（取 left/right/mid 的中位数） | 多了几次读取和一次交换，但避免了最坏情况 |
| ③ 双递归栈溢出 | 手写尾递归优化：选短子数组递归，长子数组用 for 循环处理 | Go gc 不自动做 TCO，必须手写；代价是调度逻辑比纯双递归多了几行 if-else |
| ④ 小数组慢 | 子数组长度 < 10 时切换插入排序 | 引入另一个排序算法，但阈值 10 是实测平衡点（常数因子交叉点），无副作用 |
| ⑤ 重复值低效 | 三路快排：把数组分成「小于 | 等于 | 大于」三段，只递归左右两段 | DNF 算法多了一个指针，但在重复值场景下效率成倍提升；无重复值时和双路快排开销相当 |

### 最终方案概览

调度层：**for 循环 + 选短递归 + 小数组插入排序**
分区层：**三数取中基准 + Lomuto 双向扫描**（双路） / **DNF 三指针一次遍历**（三路）

```go
// 双路快排（默认）
func QuickSort(nums []int, left int, right int) {
    for left < right {                        // ③ for 循环消掉尾递归
        if right-left < 10 {                 // ④ 小数组切换插入排序
            InsertSort(nums[left : right+1])
            return
        }
        pivot := selectPivot(nums, left, right)        // ② 三数取中选基准
        mid := partitionLomuto(nums, left, right, pivot) // ① Lomuto 原地分区
        if mid-left < right-mid {            // ③ 选短的递归，保证栈深度 O(logn)
            QuickSort(nums, left, mid-1)
            left = mid + 1                   // 长的留给 for 循环
        } else {
            QuickSort(nums, mid+1, right)
            right = mid - 1
        }
    }
}
```

---

## 四、调度层优化详解

### 4.1 尾递归优化：为什么用 `for` 循环

标准快排是双递归：
```go
QuickSort(nums, left, mid-1)    // 递归 A
QuickSort(nums, mid+1, right)   // 递归 B（尾调用）
```

**尾递归**：函数最后一步调用自身，调用结果直接返回，无后续计算。理想情况下编译器可以复用当前栈帧，把递归变成循环（尾递归优化 TCO）。

**Go 不做 TCO 的权衡**：Go 的设计哲学是显式优于隐式，gc 编译器选择不自动优化尾递归。这意味着我们必须手写。

做法：
- 递归 B 是尾调用，用 `for` 循环手动替代：`left = mid+1` 或 `right = mid-1` 把长的那边交给循环下一轮处理
- 递归 A 不是尾调用（它之后还要执行递归 B），保留

**为什么这样做是等价的**：原来的递归 B 执行完后函数就返回了，现在把边界写回循环变量，下一轮循环处理的就是原来递归 B 要处理的区间，执行顺序一模一样，只是省了一次压栈。

### 4.2 选短递归：为什么保证栈深度 O(logn)

```go
if mid-left < right-mid {   // mid-left = 左子数组长度, right-mid = 右子数组长度
    QuickSort(nums, left, mid-1)   // 递归短的
    left = mid + 1                 // 长的交给循环
} else {
    QuickSort(nums, mid+1, right)  // 递归短的
    right = mid - 1                // 长的交给循环
}
```

**原理**：当前区间长度 n，两个子数组长度之和 = n-1（基准占一个）。较短子数组长度 `min(L, R) ≤ (n-1)/2 < n/2`。每递归一次长度至少减半，嵌套层数 = log₂n。

**对比不选短的（普通双递归）**：在**已排序数组**（比如 `[1,2,3,...,n]`）上，每次 partition 后基准都在最左，右子数组长度 = n-1，递归嵌套 n 层 → 栈溢出。优化版在同样场景下短子数组为空（长度 0），递归直接返回，栈深度始终 O(1)。

**为什么不把两次递归都消掉**：硬要消掉可以用手动栈（slice 模拟调用栈），但代码复杂度上去了，手动栈大小也是 O(logn)，和递归方案一样。性价比不高。

### 4.3 小数组优化：为什么阈值选 10

见之前讨论的完整分析。核心结论：

| | 快排（10 元素） | 插入排序（10 元素） |
|---|---|---|
| 函数调用次数 | 4~6 次 | 0 |
| 常数因子 | ~50~100ns / 操作 | ~5~10ns / 操作 |
| CPU 缓存命中 | 较差（双向扫描跳跃） | 极好（连续访问） |

虽然插入排序比较次数略多（~45 vs ~30），但快排常数因子是它的 5~10 倍，最终**插入排序快 5~7 倍**。

阈值 10 是 Go 里的经验平衡点：
- 太小（如 3）：频繁切换算法，插入排序还没发挥优势就被快排接管了
- 太大（如 50）：插入排序 O(n²) 的劣势开始显现（50 元素要 1250 次比较，快排只需 ~300 次）
- 跨语言参考：Java 取 7，C++ std::sort 取 15~20，Go 通常取 10~15

额外好处：加了阈值后递归深度直接减半（子数组到 10 就提前 return，不必一直递归到长度 1）。

---

## 五、基准选择（三数取中）

### 代码

```go
func selectPivot(nums []int, left int, right int) int {
    mid := left + (right-left)>>1
    pivot := threeSumMedian(nums[left], nums[right], nums[mid])
    idx := getPivotIdx(nums, left, right, mid, pivot)
    nums[left], nums[idx] = nums[idx], nums[left]
    return pivot
}
```

### 为什么要移到首位

后续的 Lomuto 分区算法假设 pivot 在 `nums[left]` 位置（它用"藏基准"技巧，把 `nums[left]` 当作中转站空位）。所以选完基准后必须交换到首位。

### 为什么三数取中而不是随机

- **随机基准**：也能避免最坏情况，但每轮都要 rand()，开销大。而且存在极大概率连续选到差基准。
- **三数取中**：额外开销只是 3 次读取 + 2 次比较 + 1 次交换，确定性保证不会退化为 O(n²)（只要数组不是极端特殊构造的）。

---

## 六、双路分区（Lomuto 双向扫描）

### 完整代码

```go
func partitionLomuto(nums []int, left int, right int, pivot int) int {
    for left < right {
        for left < right && pivot <= nums[right] {
            right--
        }
        nums[left] = nums[right]
        for left < right && pivot >= nums[left] {
            left++
        }
        nums[right] = nums[left]
    }
    nums[left] = pivot
    return left
}
```

### 算法思想：藏基准 → 搬运 → 归位

1. **藏基准**：pivot 已保存在局部变量，`nums[left]` 位置"空出"当作中转站
2. **双向扫描搬运**：
   - 右指针从右向左跳过 ≥ pivot 的元素，遇到 < pivot 的搬到左空位
   - 左指针从左向右跳过 ≤ pivot 的元素，遇到 > pivot 的搬到右空位
   - 交替执行直到 `left == right`
3. **归位**：此时 `left` 指向的就是 pivot 的最终位置，把局部变量里的 pivot 写回去

### 为什么最后要 `nums[left] = pivot`

扫描过程中 pivot 保存在局部变量里，数组中原位已被覆盖。`left` 最终指向的位置就是它该在的地方，写回即完成归位。

### 单向扫描 vs 双向扫描（权衡）

| | 单向扫描（Nico Lomuto 原实现） | 双向扫描（Hoare 版） |
|---|---|---|
| 做法 | 维护一个 border 指针，从左到右扫一遍，遇到 < pivot 的就交换到 border 左边 | 左右双向交替扫描，原地搬运 |
| 遍历次数 | 1 次 | 也是 1 次（同时从两端扫） |
| 交换次数 | 较多（border 每次移动都要交换） | 较少（搬运赋值代替交换，最后 pivot 只移动一次） |
| 实现复杂度 | 简单 | 稍复杂（要处理空位中转站） |
| 实际速度 | 略慢 | 略快 |

选双向扫描，牺牲一点代码简洁，换来更少的内存写入。

### 举例走一遍

```
初始: [3, 1, 4, 1, 5, 9, 2, 6]  pivot=3 在 nums[0]
       ↑left=0              ↑right=7

第一轮:
  右扫描: 6≥3 ✓ right=6; 2<3 ✗ → nums[0]=2 → [2,1,4,1,5,9,2,6]
  左扫描: 2≤3 ✓ left=1; 1≤3 ✓ left=2; 4>3 ✗ → nums[6]=4 → [2,1,4,1,5,9,4,6]

第二轮 (left=2, right=6):
  右扫描: 4≥3 ✓ right=5; 9≥3 ✓ right=4; 5≥3 ✓ right=3; 1<3 ✗ → nums[2]=1 → [2,1,1,1,5,9,4,6]
  左扫描: 1≤3 ✓ left=3 → left==right 退出

归位: nums[3]=3 → [2,1,1,3,5,9,4,6]  返回 mid=3
                    ↑基准最终位置
```

### 为什么最后要 `nums[left] = pivot`

扫描过程中 pivot 保存在局部变量里，数组中原位已被覆盖。`left` 最终指向的位置就是它该在的地方，写回即完成归位。

---

## 七、扩展：三路快排（处理重复值）

当数组存在大量重复值时，普通双路分区会把等于 pivot 的元素分散到两边，导致下一轮还要重新处理。三路快排把数组分成三段：

```
[小于 pivot | 等于 pivot | 大于 pivot]
```

只递归左右两段，中间相等的直接跳过，大量重复值场景下效率显著提升。

### 调度代码

```go
func QuickSortII(nums []int, left int, right int) {
    for left < right {
        if right-left < 10 {
            InsertSort(nums[left : right+1])
            return
        }
        pivot := selectPivot(nums, left, right)
        leftMid, rightMid := threeWayPartition(nums, left, right, pivot)
        if leftMid-left < right-rightMid {
            QuickSortII(nums, left, leftMid)
            left = rightMid
        } else {
            QuickSortII(nums, rightMid, right)
            right = leftMid
        }
    }
}
```

调度层和双路快排完全一致的优化思路：小数组插入排序 + 选短递归 + 长留给循环。

### 三路分区（DNF 荷兰国旗算法）

#### 完整代码

```go
func threeWayPartition(nums []int, left int, right int, pivot int) (int, int) {
    lt, i, gt := left, left, right
    for i <= gt {
        if nums[i] < pivot {
            nums[lt], nums[i] = nums[i], nums[lt]
            lt++
            i++
        } else if nums[i] > pivot {
            nums[i], nums[gt] = nums[gt], nums[i]
            gt--
        } else {
            i++
        }
    }
    return lt - 1, gt + 1
}
```

三个指针的含义：

```
[left ......... lt-1] [lt ......... i-1] [i ......... gt] [gt+1 ......... right]
       < pivot              == pivot        待扫描            > pivot
       (已确定)              (已确定)       (当前扫描点)      (已确定)
```

- `lt`：**小于段的下一个写入位置**。`[left, lt-1]` 全是 `< pivot` 的
- `i`：**当前扫描位置**。`[lt, i-1]` 全是 `== pivot` 的
- `gt`：**大于段的上一个写入位置**。`[gt+1, right]` 全是 `> pivot` 的

#### 三种情况的处理逻辑

| 扫描到 nums[i] | 操作 | 指针移动 | 原因 |
|---|---|---|---|
| `< pivot` | 和 `nums[lt]` 交换（塞到小于段末尾） | `lt++`, `i++` | lt 位置原来要么是等于的，要么是刚交换过来的 i 自己，都处理过了 |
| `> pivot` | 和 `nums[gt]` 交换（塞到大于段开头） | **只有 `gt--`** | 交换来的 gt 位置的元素还没扫描过，不能 `i++` |
| `== pivot` | 不动 | 只有 `i++` | 等于段自然向右扩展一格 |

#### 举例走一遍

```
nums = [2, 1, 3, 2, 1, 3, 2], pivot = 2
初始: lt=0, i=0, gt=6

i=0: nums[0]=2 == pivot → i++ → lt=0 i=1 gt=6
i=1: nums[1]=1 < pivot → 交换 nums[0]↔nums[1] → [1,2,3,2,1,3,2]
                         lt=1, i=2
i=2: nums[2]=3 > pivot → 交换 nums[2]↔nums[6] → [1,2,2,2,1,3,3]
                         gt=5 (注意 i 不动！交换来的 2 还没处理)
i=2: nums[2]=2 == pivot → i++ → lt=1 i=3 gt=5
i=3: nums[3]=2 == pivot → i++ → lt=1 i=4 gt=5
i=4: nums[4]=1 < pivot → 交换 nums[1]↔nums[4] → [1,1,2,2,2,3,3]
                         lt=2, i=5
i=5: nums[5]=3 > pivot → 交换 nums[5]↔nums[5] → gt=4
                         (i=5 > gt=4, 循环退出！)

结果: [1, 1, 2, 2, 2, 3, 3]
       ^lt=2      ^gt=4
       等于pivot区间 [2,4]
```

返回值 `(lt-1, gt+1) = (1, 5)`：
- `QuickSortII(nums, 0, 1)` → 递归小于段 `[1, 1]`
- `QuickSortII(nums, 5, 6)` → 递归大于段 `[3, 3]`
- 中间 `[2, 2, 2]` 直接跳过 ✓

#### 为什么选 DNF 而不是"暂存+二次归位"

旧实现（当前代码库已替换）有 60 行，要做两次扫描（一次暂存到两端，一次把暂存的等于 pivot 元素归位到基准两侧）。DNF 只用 15 行，一次遍历就完成了分区，辅助变量从 6 个降到 3 个，还不需要依赖 pivot 预先移到首位。代价：没有（DNF 在任何场景下都不比旧实现慢）。

---

## 八、整体流程图

### 双路快排流程图

```
QuickSort(nums, left, right)
│
├─ for left < right:
│   │
│   ├─ right-left < 10? → InsertSort 处理, return  ← 小数组优化
│   │
│   ├─ selectPivot(nums, left, right)   ← 三数取中选基准, 交换到首位
│   │
│   ├─ partitionLomuto(nums, left, right, pivot)   ← Lomuto 双向扫描
│   │   ├─ 右指针左移跳过 ≥ pivot
│   │   ├─ 搬运到左空位
│   │   ├─ 左指针右移跳过 ≤ pivot
│   │   ├─ 搬运到右空位
│   │   └─ nums[left] = pivot 归位, 返回 mid
│   │
│   ├─ mid-left < right-mid?   ← 选短的递归, 保证栈深度 O(logn)
│   │   ├─ YES: QuickSort(nums, left, mid-1) 递归短的左边
│   │   │        left = mid + 1               长的右边留给 for
│   │   └─ NO:  QuickSort(nums, mid+1, right) 递归短的右边
│   │            right = mid - 1               长的左边留给 for
│   │
│   └─ 继续 for 循环...直到 left >= right, 排序完成
```

### 三路快排流程图

```
QuickSortII(nums, left, right)
│
├─ for left < right:
│   │
│   ├─ right-left < 10? → InsertSort 处理, return  ← 小数组优化
│   │
│   ├─ selectPivot(nums, left, right)   ← 三数取中选基准, 交换到首位
│   │
│   ├─ threeWayPartition(nums, left, right, pivot)   ← DNF 一次遍历
│   │   │
│   │   ├─ 初始化: lt=left, i=left, gt=right
│   │   │
│   │   ├─ for i <= gt:
│   │   │   ├─ nums[i] < pivot:
│   │   │   │   └─ 交换 nums[lt]↔nums[i], lt++, i++
│   │   │   ├─ nums[i] > pivot:
│   │   │   │   └─ 交换 nums[i]↔nums[gt], gt-- (i 不动!)
│   │   │   └─ nums[i] == pivot:
│   │   │       └─ i++ (自然归入等于段)
│   │   │
│   │   └─ 返回 (lt-1, gt+1)  ← 小于段右端、大于段左端, 中间 [lt,gt] 全等于 pivot
│   │
│   ├─ leftMid-left < right-rightMid?   ← 选短的递归
│   │   ├─ YES: QuickSortII(nums, left, leftMid)  递归短的左边 (< pivot)
│   │   │        left = rightMid                  长的右边留给 for
│   │   └─ NO:  QuickSortII(nums, rightMid, right) 递归短的右边 (> pivot)
│   │            right = leftMid                   长的左边留给 for
│   │
│   └─ 中间 [leftMid+1, rightMid-1] 全等于 pivot, 直接跳过 ✓
│
└─ for 循环直到 left >= right, 排序完成
```