package backtrack

// FindWords
// @Title: LC212.单词搜索 II
// @Description: 给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words， 返回所有二维网格上的单词。
// @Description: 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中相邻单元格是那些水平相邻或垂直相邻的单元格。
// @Description: 同一个单元格内的字母在一个单词中不允许被重复使用。
// @Link: https://leetcode.cn/problems/word-search-ii/
// @param board
// @param words
// @return []string
func FindWords(board [][]byte, words []string) []string {
	var (
		m, n   = len(board), len(board[0])
		root   = buildTrie(words)
		dirs   = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		dfs    func(int, int, *trieNode)
		result []string
	)

	dfs = func(i, j int, node *trieNode) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		ch := board[i][j]
		if ch == '#' {
			return
		}
		idx := ch - 'a'
		next := node.children[idx]
		if next == nil {
			return
		}
		if next.word != "" {
			result = append(result, next.word)
			next.word = ""
		}
		board[i][j] = '#'
		for _, dir := range dirs {
			dfs(i+dir[0], j+dir[1], next)
		}
		board[i][j] = ch

		// 剪枝
		if len(next.children) == 0 {
			node.children[idx] = nil
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, root)
		}
	}

	return result
}

type trieNode struct {
	children [26]*trieNode
	word     string
}

func buildTrie(words []string) *trieNode {
	root := &trieNode{}
	for _, w := range words {
		node := root
		for i := 0; i < len(w); i++ {
			idx := w[i] - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &trieNode{}
			}
			node = node.children[idx]
		}
		node.word = w
	}
	return root
}
