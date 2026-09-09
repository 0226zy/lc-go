package generateparentheses

// GenerateParenthesis 括号生成
// 数字 n 代表生成括号的对数，返回所有可能的并且有效的括号组合。
// 时间复杂度: O(4^n / √n) 即卡特兰数 Cn  空间复杂度: O(n) 递归栈深度
func GenerateParenthesis(n int) []string {
	path := make([]byte, 2*n)
	result := []string{}
	var backtrace func(open, close int)
	backtrace = func(open, close int) {
		if len(path) == 2*n {
			result = append(result, string(path))
			return
		}
		if open < n {
			path = append(path, '(')
			backtrace(open+1, close)
			path = path[:len(path)-1]
		}
		if close < open {
			path = append(path, ')')
			backtrace(open, close+1)
			path = path[:len(path)-1]
		}
	}
	backtrace(0, 0)
	return result

}
