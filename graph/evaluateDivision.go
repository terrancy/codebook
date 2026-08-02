package graph

// EvaluateDivision
// @Title: LC399.除法求值
// @Description: 给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，其中 equations[i] = [Ai, Bi] 和 values[i] 共同表示等式 Ai / Bi = values[i]。每个 Ai 或 Bi 是一个表示单个变量的字符串。
// @Description: 另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj = ? 的结果作为答案。
// @Description: 返回 所有问题的答案。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。
// @Description: 如果问题中出现了给定的已知条件中没有出现的字符串，也需要用 -1.0 替代这个答案。
// @Link: https://leetcode.cn/problems/evaluate-division/
// @param equations
// @param values
// @param queries
// @return []float64
func EvaluateDivision(equations [][]string, values []float64, queries [][]string) []float64 {
	id := make(map[string]int)
	for _, eq := range equations {
		if _, ok := id[eq[0]]; !ok {
			id[eq[0]] = len(id)
		}
		if _, ok := id[eq[1]]; !ok {
			id[eq[1]] = len(id)
		}
	}

	n := len(id)
	parent := make([]int, n)
	weight := make([]float64, n)
	for i := range parent {
		parent[i] = i
		weight[i] = 1.0
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			origin := parent[x]
			parent[x] = find(parent[x])
			weight[x] *= weight[origin]
		}
		return parent[x]
	}

	union := func(x, y int, val float64) {
		rootX, rootY := find(x), find(y)
		if rootX != rootY {
			parent[rootX] = rootY
			weight[rootX] = val * weight[y] / weight[x]
		}
	}

	for i, eq := range equations {
		union(id[eq[0]], id[eq[1]], values[i])
	}

	ans := make([]float64, len(queries))
	for i, q := range queries {
		idx1, ok1 := id[q[0]]
		idx2, ok2 := id[q[1]]
		if !ok1 || !ok2 {
			ans[i] = -1.0
			continue
		}
		if find(idx1) != find(idx2) {
			ans[i] = -1.0
			continue
		}
		ans[i] = weight[idx1] / weight[idx2]
	}
	return ans
}