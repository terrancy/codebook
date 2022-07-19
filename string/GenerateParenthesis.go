package strings

//
// GenerateParenthesis
// @Description: NC26 括号生成
// @param n
// @return []string
//
func GenerateParenthesis(n int) []string {
    res := make([]string, 0)
    BSGenerateParenthesis("", 0, 0, n, &res)
    return res
}

func BSGenerateParenthesis(str string, open int, close int, n int, res *[]string) {
    if len(str) == n<<1 {
        *res = append(*res, str)
        return
    }
    
    if open < n {
        BSGenerateParenthesis(str+"(", open+1, close, n, res)
    }
    
    if close < open {
        BSGenerateParenthesis(str+")", open, close+1, n, res)
    }
}
