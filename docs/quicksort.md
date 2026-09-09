# 快速排序

## 一、核心思想

分治策略：选一个基准值 pivot，将数组划分为「小于基准 | 基准 | 大于基准」三部分，然后对左右两个子数组递归执行同样的操作。

时间复杂度平均 O(nlogn)，最坏 O(n²)（已排序数组且选首元素为基准），空间复杂度 O(1)（原地排序），不稳定。

---

## 二、朴素版（便于理解）

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

用额外两个切片承接，直观但空间复杂度 O(n)。

---

## 三、优化版核心调度

### 完整代码

```go
func QuickSort(nums []int, left int, right int) {
    for left < right {
        pivot := selectPivot(nums, left, right)
        mid := partitionLomuto(nums, left, right, pivot)
        if mid-left < right-mid {
            QuickSort(nums, left, mid-1)
            left = mid + 1
        } else {
            QuickSort(nums, mid+1, right)
            right = mid - 1
        }
    }
}
```

### 关键设计解读

#### 1. 尾递归优化：为什么用 `for` 循环

标准快排是双递归：
```go
QuickSort(nums, left, mid-1)    // 递归 A
QuickSort(nums, mid+1, right)   // 递归 B（尾调用）
```

**尾递归**：函数最后一步调用自身，调用结果直接返回，无后续计算。理论上编译器可以复用当前栈帧，把递归变成循环（尾递归优化 TCO）。

但 **Go 的 gc 编译器不自动做 TCO**，所以手写优化：
- 递归 B 是尾调用，用 `for` 循环手动替代：`left = mid+1` 或 `right = mid-1` 把长的那边交给循环下一轮处理
- 递归 A 不是尾调用（它之后还要执行递归 B），保留

#### 2. 为什么选短子数组递归

```go
if mid-left < right-mid {   // mid-left = 左子数组长度, right-mid = 右子数组长度
    QuickSort(nums, left, mid-1)   // 递归短的
    left = mid + 1                 // 长的交给循环
} else {
    QuickSort(nums, mid+1, right)  // 递归短的
    right = mid - 1                // 长的交给循环
}
```

**保证栈深度 O(logn) 的关键：**

当前区间长度为 n，两个子数组长度之和 = n-1（基准占一个）。较短子数组长度 `min(L, R) ≤ (n-1)/2 < n/2`。每递归一次长度至少减半，嵌套层数 = log₂n。

对比普通双递归在**已排序数组**的最坏情况：每次 partition 后基准都在最左，右子数组长度 = n-1，递归嵌套 n 层 → 栈溢出。优化版在同样场景下短子数组为空，递归直接返回，栈深度始终 O(1)。

#### 3. 总结

| 方案 | 栈深度（平均） | 栈深度（最坏） |
|---|---|---|
| 双递归不优化 | O(logn) | O(n) → 栈溢出 |
| 短递归 + 循环长 | O(logn) | O(1) |

---

## 四、组件一：基准选择（三数取中）

```go
func selectPivot(nums []int, left int, right int) int {
    mid := left + (right-left)>>1
    pivot := threeSumMedian(nums[left], nums[right], nums[mid])
    idx := getPivotIdx(nums, left, right, mid, pivot)
    nums[left], nums[idx] = nums[idx], nums[left]
    return pivot
}
```

取 `nums[left]`、`nums[right]`、`nums[mid]` 三者的中位数作为基准，并把基准交换到 `nums[left]` 位置（给后续 partition 用）。

作用：避免在已排序/接近有序数组上退化为 O(n²)。

---

## 五、组件二：双路分区（Lomuto 双向扫描）

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

## 六、扩展：三路快排（处理重复值）

当数组存在大量重复值时，普通双路分区会把等于 pivot 的元素分散到两边，导致下一轮还要重新处理。三路快排把数组分成三段：

```
[小于 pivot | 等于 pivot | 大于 pivot]
```

### 调度代码

```go
func QuickSortII(nums []int, left int, right int) {
    if left >= right {
        return
    }
    pivot := selectPivot(nums, left, right)
    leftMid, rightMid := threeWayPartition(nums, left, right, pivot)
    QuickSortII(nums, left, leftMid)
    QuickSortII(nums, rightMid, right)
}
```

只递归左右两段，中间相等的直接跳过，大量重复值场景下效率显著提升。

### 三路分区核心思路

1. 双向扫描时，遇到等于 pivot 的元素暂存到两端
2. 扫描结束后基准归位，再把两端暂存的等于 pivot 的元素收拢到基准两侧
3. 返回等于 pivot 区间的左右边界 `(leftMid, rightMid)`

---

## 七、整体流程图

```
QuickSort(nums, left, right)
│
├─ for left < right:
│   │
│   ├─ selectPivot(nums, left, right)   ← 三数取中选基准, 交换到首位
│   │
│   ├─ partitionLomuto(nums, left, right, pivot)
│   │   ├─ 右指针左移跳过 ≥ pivot
│   │   ├─ 搬运到左空位
│   │   ├─ 左指针右移跳过 ≤ pivot
│   │   ├─ 搬运到右空位
│   │   └─ nums[left] = pivot 归位, 返回 mid
│   │
│   ├─ mid-left < right-mid?
│   │   ├─ YES: QuickSort(nums, left, mid-1) 递归短的左边
│   │   │        left = mid + 1               长的右边留给 for
│   │   └─ NO:  QuickSort(nums, mid+1, right) 递归短的右边
│   │            right = mid - 1               长的左边留给 for
│   │
│   └─ 继续 for 循环...直到 left >= right, 排序完成
```