package graph

// 图节点的逻辑结构
type Vertex struct {
    Id        int
    Neighbors []Vertex
}

// 多叉树
type TreeNode struct {
    Val      int
    Children []TreeNode
}

func BuildGraph() {

}
