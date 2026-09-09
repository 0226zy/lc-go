package longestcommonsubsequence

// LongestCommonSubsequence 最长公共子序列（标准 DP 数组版）
// dp[i][j] 表示 text1 前 i 个字符与 text2 前 j 个字符的最长公共子序列长度。
// 若 text1[i-1] == text2[j-1]，dp[i][j] = dp[i-1][j-1] + 1；
// 否则 dp[i][j] = max(dp[i-1][j], dp[i][j-1])。
// 时间复杂度: O(m*n)  空间复杂度: O(m*n)
func LongestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1) // 第 0 行/第 0 列天然为 0，即 base case
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				// 末尾字符相等：放入公共子序列，累加 1
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 末尾字符不等：去掉 text1 末尾或 text2 末尾，取较优
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

// LongestCommonSubsequenceOptimized 最长公共子序列（一维滚动数组空间优化版）
// 第 i 行只依赖第 i-1 行，压缩为一维数组；用 prev 保存被覆盖前的左上角 dp[i-1][j-1]。
// 时间复杂度: O(m*n)  空间复杂度: O(n)
func LongestCommonSubsequenceOptimized(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([]int, n+1) // 滚动数组，全 0 即 base case
	for i := 1; i <= m; i++ {
		prev := 0 // 保存 dp[i-1][j-1]（左上角）
		for j := 1; j <= n; j++ {
			old := dp[j] // 更新前的 dp[j] 即 dp[i-1][j]，先存下来
			if text1[i-1] == text2[j-1] {
				dp[j] = prev + 1
			} else {
				// dp[j] 旧值 = 上一行 dp[i-1][j]，dp[j-1] 新值 = 本行 dp[i][j-1]
				dp[j] = max(dp[j], dp[j-1])
			}
			prev = old
		}
	}
	return dp[n]
}
