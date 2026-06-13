package generateparentheses

// GenerateParenthesis 括号生成
// 给定 n 对括号，生成所有有效的括号组合。
// 时间复杂度: O(4^n/√n)  空间复杂度: O(n)
func GenerateParenthesis(n int) []string {
	var res []string
	buf := make([]byte, 0, 2*n)

	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		if open == n && close == n {
			res = append(res, string(buf))
			return
		}
		if open < n {
			buf = append(buf, '(')
			backtrack(open+1, close)
			buf = buf[:len(buf)-1]
		}
		if close < open {
			buf = append(buf, ')')
			backtrack(open, close+1)
			buf = buf[:len(buf)-1]
		}
	}
	backtrack(0, 0)
	return res
}

// GenerateParenthesisDP 括号生成（动态规划）
// 利用卡特兰数递推关系生成所有有效组合。
// 时间复杂度: O(4^n/√n)  空间复杂度: O(4^n/√n)
func GenerateParenthesisDP(n int) []string {
	// dp[i] 表示 i 对括号的所有合法组合
	dp := make([][]string, n+1)
	dp[0] = []string{""}

	for i := 1; i <= n; i++ {
		var cur []string
		// 枚举第一个左括号内包含的括号对数 j，剩余 i-1-j 对在右括号外
		for j := 0; j < i; j++ {
			for _, left := range dp[j] {
				for _, right := range dp[i-1-j] {
					cur = append(cur, "("+left+")"+right)
				}
			}
		}
		dp[i] = cur
	}
	return dp[n]
}
