package interleavingstring

// IsInterleave 交错字符串（标准 DP 数组版）
// 给定三个字符串 s1、s2 和 s3，判断 s3 是否由 s1 和 s2 交错组成。
// 交错组成要求 s1、s2 中字符的相对顺序在 s3 中保持不变。
// dp[i][j] 表示 s1 的前 i 个字符与 s2 的前 j 个字符能否交错组成 s3 的前 i+j 个字符。
// 转移：s3[i+j-1] 要么来自 s1[i-1]（看 dp[i-1][j]），要么来自 s2[j-1]（看 dp[i][j-1]）。
// 时间复杂度: O(m*n)  空间复杂度: O(m*n)
func IsInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)

	// 长度不匹配时一定无法交错组成
	if m+n != len(s3) {
		return false
	}

	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true // base case：两个空串交错出空串

	// base case：第 0 列，只用 s1 的前 i 个字符逐位匹配 s3
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	// base case：第 0 行，只用 s2 的前 j 个字符逐位匹配 s3
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// s3 的第 i+j 个字符要么来自 s1 的第 i 个字符，要么来自 s2 的第 j 个字符
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) ||
				(dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return dp[m][n]
}

// IsInterleaveOptimized 交错字符串（滚动一维数组空间优化版）
// dp[i][j] 只依赖正上方和本行左侧：内层从左到右遍历时，
// dp[j] 尚未更新即上一行的 dp[i-1][j]，dp[j-1] 已更新即本行的 dp[i][j-1]。
// 时间复杂度: O(m*n)  空间复杂度: O(n)（n 为 s2 的长度）
func IsInterleaveOptimized(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)

	// 长度不匹配时一定无法交错组成
	if m+n != len(s3) {
		return false
	}

	// dp[j] 表示 s1 的前 i 个字符与 s2 的前 j 个字符
	// 能否交错组成 s3 的前 i+j 个字符（i 由外层循环隐式表示）
	dp := make([]bool, n+1)
	dp[0] = true

	// 初始化 i = 0 的一行：只用 s2 的前 j 个字符匹配 s3 的前 j 个字符
	for j := 1; j <= n; j++ {
		dp[j] = dp[j-1] && s2[j-1] == s3[j-1]
	}

	for i := 1; i <= m; i++ {
		// j = 0 的边界：只用 s1 的前 i 个字符匹配 s3 的前 i 个字符
		dp[0] = dp[0] && s1[i-1] == s3[i-1]
		for j := 1; j <= n; j++ {
			// dp[j] 是上一行的 dp[i-1][j]（来自 s1），dp[j-1] 是本行的 dp[i][j-1]（来自 s2）
			dp[j] = (dp[j] && s1[i-1] == s3[i+j-1]) ||
				(dp[j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return dp[n]
}
