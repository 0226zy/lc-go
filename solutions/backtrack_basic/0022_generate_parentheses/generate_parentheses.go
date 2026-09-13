package generateparentheses

// GenerateParenthesis 括号生成
// 数字 n 代表生成括号的对数，返回所有可能的并且有效的括号组合。
// 时间复杂度: O(4^n / √n) 即卡特兰数 Cn 量级  空间复杂度: O(n) 递归栈深度
func GenerateParenthesis(n int) []string {
	result := []string{}
	path := make([]byte, 0, 2*n)

	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		// 路径长度达到 2n，收集一个合法组合
		if len(path) == 2*n {
			result = append(result, string(path))
			return
		}
		// 左括号还有剩余，添加左括号永远不会破坏前缀合法性
		if open < n {
			path = append(path, '(')
			backtrack(open+1, close)
			path = path[:len(path)-1]
		}
		// 右括号数量不能超过左括号，否则前缀非法，直接剪枝
		if close < open {
			path = append(path, ')')
			backtrack(open, close+1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, 0)
	return result
}
