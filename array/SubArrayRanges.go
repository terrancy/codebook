package array

//
//  SubArrayRanges
//  @Description: 2104. 子数组范围和
//  @param nums
//  @return int64
//
func SubArrayRanges(nums []int) int64 {
    n := len(nums)
    if n == 1 {
        return 0
    }
    //sort.Ints(nums)
    dataList := make([][]int, 0)
    for i := 2; i <= n; i++ {
        dfs(nums, []int{}, &dataList, i, 0)
    }
    var res int64 = 0
    for _, data := range dataList {
        res += maxWidth(data)
    }
    return res
}

func dfs(nums []int, path []int, res *[][]int, k int, start int) {
    if len(path) == k {
        temp := make([]int, len(path))
        copy(temp, path)
        *res = append(*res, temp)
        return
    }
    for i := start; i < len(nums); i++ {
        if i > start && nums[i] == nums[i-1] {
            continue
        }
        path = append(path, nums[i])
        dfs(nums, path, res, k, i+1)
        if len(path) == k {
            break
        } else {
            path = path[:len(path)-1]
        }
    }
}

func maxWidth(data []int) int64 {
    n := len(data)
    if n == 1 {
        return 0
    }
    min := data[0]
    max := data[0]
    for i := 1; i < n; i++ {
        if min > data[i] {
            min = data[i]
        }
        if max < data[i] {
            max = data[i]
        }
    }
    return int64(max - min)
}

func subArrayRanges(nums []int) int64 {
    n, max, stack := len(nums), int64(0), []int{}
    for i := 0; i <= n; i++ {
        for l := len(stack); l > 0 && (i == n || nums[stack[l-1]] < nums[i]); {
            j := stack[l-1]
            stack = stack[:l-1]
            if l = len(stack); l > 0 {
                max += int64(nums[j]) * int64(i-j) * int64(j-stack[l-1])
            } else {
                max += int64(nums[j]) * int64(i-j) * int64(j+1)
            }
        }
        stack = append(stack, i)
    }
    stack = []int{}
    min := int64(0)
    for i := 0; i <= n; i++ {
        for l := len(stack); l > 0 && (i == n || nums[stack[l-1]] > nums[i]); {
            j := stack[l-1]
            stack = stack[:l-1]
            if l = len(stack); l > 0 {
                min += int64(nums[j]) * int64(i-j) * int64(j-stack[l-1])
            } else {
                min += int64(nums[j]) * int64(i-j) * int64(j+1)
            }
        }
        stack = append(stack, i)
    }
    return max - min
}
