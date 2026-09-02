package others

// WordDictionary
// @Title: LC211.添加与搜索单词 - 数据结构设计
// @Description: 设计一个数据结构，支持 addWord(word) 和 search(word)，其中 word 可能包含 '.'
//               '.' 可以匹配任意一个字母
// @Link: https://leetcode.cn/problems/design-add-and-search-word-data-structure-design/
//
//	Example:
//		dict := WordDictionaryConstructor()
//		dict.AddWord("bad")
//		dict.AddWord("dad")
//		dict.AddWord("mad")
//		dict.Search("pad") // false
//		dict.Search("bad") // true
//		dict.Search(".ad") // true
//		dict.Search("b..") // true

type wordNode struct {
	children [26]*wordNode
	isEnd    bool
}

func newWordNode() *wordNode {
	return &wordNode{}
}

type WordDictionary struct {
	root *wordNode
}

func WordDictionaryConstructor() WordDictionary {
	return WordDictionary{root: newWordNode()}
}

func (wd *WordDictionary) AddWord(word string) {
	node := wd.root
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if node.children[idx] == nil {
			node.children[idx] = newWordNode()
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

func (wd *WordDictionary) Search(word string) bool {
	return wd.searchDFS(wd.root, word, 0)
}

func (wd *WordDictionary) searchDFS(node *wordNode, word string, pos int) bool {
	if pos == len(word) {
		return node.isEnd
	}
	ch := word[pos]
	if ch == '.' {
		// 通配符: 遍历所有 26 个子节点，任意一条路径匹配即可
		for _, child := range node.children {
			if child != nil && wd.searchDFS(child, word, pos+1) {
				return true
			}
		}
		return false
	}
	idx := ch - 'a'
	next := node.children[idx]
	if next == nil {
		return false
	}
	return wd.searchDFS(next, word, pos+1)
}