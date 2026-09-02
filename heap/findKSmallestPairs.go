package heap

import "container/heap"

type pairSum struct {
	sum  int
	i    int
	j    int
	nums []int
}

type pairHeap []*pairSum

func (h pairHeap) Len() int           { return len(h) }
func (h pairHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h pairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *pairHeap) Push(x any) {
	*h = append(*h, x.(*pairSum))
}

func (h *pairHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// FindKSmallestPairs
// @Title: LC373. 查找和最小的 K 对数字
// @Description: 给定两个以升序排列的整数数组 nums1 和 nums2 , 以及一个整数 k
// @Description: 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2
// @Description: 找到和最小的 k 个数字对 (u1,v1), (u2,v2) ... (uk,vk)
// @Link: https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/
// @param nums1 整数数组
// @param nums2 整数数组
// @param k 数字对数量
// @return 和最小的 k 个数字对
func FindKSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 || k <= 0 {
		return nil
	}

	var h pairHeap
	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(&h, &pairSum{
			sum: nums1[i] + nums2[0],
			i:   i,
			j:   0,
		})
	}

	res := make([][]int, 0, k)
	for h.Len() > 0 && len(res) < k {
		p := heap.Pop(&h).(*pairSum)
		res = append(res, []int{nums1[p.i], nums2[p.j]})
		if p.j+1 < len(nums2) {
			heap.Push(&h, &pairSum{
				sum: nums1[p.i] + nums2[p.j+1],
				i:   p.i,
				j:   p.j + 1,
			})
		}
	}

	return res
}