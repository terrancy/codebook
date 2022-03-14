package challenge

//
//  CountHighestScoreNodes
//  @Description: 2049. 统计最高分的节点数目
//  @param parents
//  @return int
//
func CountHighestScoreNodes(parents []int) (ans int) {
    maxScore, n, graph := int64(0), len(parents), map[int][]int{}
    for i := 1; i < n; i++ {
        graph[parents[i]] = append(graph[parents[i]], i)
    }
    
    var dfs func(node int) int
    dfs = func(node int) int {
        left, right := 0, 0
        children := graph[node]
        if len(children) > 0 {
            left = dfs(children[0])
            if len(children) > 1 {
                right = dfs(children[1])
            }
        }
        if score := maxCountHighestScoreNodes(left, 1) * maxCountHighestScoreNodes(right, 1) * maxCountHighestScoreNodes(n-1-left-right, 1); score > maxScore {
            maxScore, ans = score, 1
        } else if score == maxScore {
            ans++
        }
        return left + right + 1
    }
    dfs(0)
    return ans
}

func maxCountHighestScoreNodes(a, b int) int64 {
    if a > b {
        return int64(a)
    }
    return int64(b)
}
