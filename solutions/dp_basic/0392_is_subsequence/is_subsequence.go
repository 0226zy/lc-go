package issubsequence

// IsSubsequence 判断子序列（标准 DP 数组版）
// 判断 s 是否为 t 的子序列。
// dp[i][j] 表示 s 的前 i 个字符是否为 t 的前 j 个字符的子序列：
// s[i-1] == t[j-1] 时 dp[i][j] = dp[i-1][j-1]，否则 dp[i][j] = dp[i][j-1]。
// 时间复杂度: O(m*n)  空间复杂度: O(m*n)
func IsSubsequence(s string, t string) bool {
	m, n := len(s), len(t)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = true // base case：空串 s 是任何 t 前缀的子序列
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] // 匹配成功，两个前缀同时缩短
			} else {
				dp[i][j] = dp[i][j-1] // 跳过 t[j-1]
			}
		}
	}
	return dp[m][n]
}

// IsSubsequenceAlternative 判断子序列（双指针版）
// i 指向 s 待匹配的字符，j 扫描 t；相等则 i 前进，j 始终前进。
// 扫描结束 i == len(s) 说明 s 的字符按顺序全部在 t 中出现。
// 时间复杂度: O(m+n)  空间复杂度: O(1)
func IsSubsequenceAlternative(s string, t string) bool {
	i := 0
	for j := 0; j < len(t) && i < len(s); j++ {
		if s[i] == t[j] {
			i++ // 匹配成功，推进 s 的指针
		}
	}
	return i == len(s)
}
