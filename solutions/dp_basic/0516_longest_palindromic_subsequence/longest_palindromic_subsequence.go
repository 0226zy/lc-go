package longestpalindromicsubsequence

// LongestPalindromeSubseq 最长回文子序列（标准二维 DP 数组版）
// 求字符串 s 中最长回文子序列的长度。
// dp[i][j] 表示子串 s[i..j] 中最长回文子序列的长度：
// s[i] == s[j] 时 dp[i][j] = dp[i+1][j-1] + 2，
// 否则 dp[i][j] = max(dp[i+1][j], dp[i][j-1])。
// i 倒序、j 正序遍历，保证短区间先于长区间求解。
// 时间复杂度: O(n^2)  空间复杂度: O(n^2)
func LongestPalindromeSubseq(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1 // base case：单个字符是长度为 1 的回文
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2 // 两端同时收入回文
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1]) // 舍弃左端或右端
			}
		}
	}
	return dp[0][n-1]
}

// LongestPalindromeSubseqOptimized 最长回文子序列（一维滚动数组空间优化版）
// dp[j] 滚动表示当前 i 行的 dp[i][j]，prev 在覆盖前记录 dp[i+1][j-1]。
// 区间 DP 依赖整段历史值，空间最多压到 O(n)，无法 O(1)。
// 时间复杂度: O(n^2)  空间复杂度: O(n)
func LongestPalindromeSubseqOptimized(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		dp[i] = 1 // base case
		prev := 0 // prev 保存 dp[i+1][j-1]（空区间时为 0）
		for j := i + 1; j < n; j++ {
			old := dp[j] // 旧的 dp[j] 即 dp[i+1][j]，留给下一轮做 prev
			if s[i] == s[j] {
				dp[j] = prev + 2
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			prev = old
		}
	}
	return dp[n-1]
}
