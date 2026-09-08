package generateparentheses

// GenerateParenthesis 括号生成
// 数字 n 代表生成括号的对数，返回所有可能的并且有效的括号组合。
// 时间复杂度: O(4^n / √n) 即卡特兰数 Cn  空间复杂度: O(n) 递归栈深度
func GenerateParenthesis(n int) []string {
	var result []string
	path := make([]byte, 0, 2*n)

	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		if len(path) == 2*n {
			result = append(result, string(path))
			return
		}
		// 左括号还没用完，随时可以加
		if open < n {
			path = append(path, '(')
			backtrack(open+1, close)
			path = path[:len(path)-1]
		}
		// 右括号不能超过左括号数量，否则前缀无效
		if close < open {
			path = append(path, ')')
			backtrack(open, close+1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, 0)
	return result
}
