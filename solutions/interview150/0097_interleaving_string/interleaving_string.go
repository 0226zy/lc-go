package interleavingstring

// IsInterleave 交错字符串
// 给定三个字符串 s1、s2 和 s3，判断 s3 是否由 s1 和 s2 交错组成。
// 交错组成要求 s1、s2 中字符的相对顺序在 s3 中保持不变：
// 即 s3 可以拆分成若干子串，交替取自 s1 和 s2，拼接后恰好等于 s3。
// 时间复杂度: O(m*n)  空间复杂度: O(n)（使用一维滚动数组优化，n 为 s2 的长度）
func IsInterleave(s1 string, s2 string, s3 string) bool {
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
			// s3 的第 i+j 个字符要么来自 s1 的第 i 个字符（对应 dp[j] 上一行，即原来的 dp[j]）
			// 要么来自 s2 的第 j 个字符（对应本行已更新的 dp[j-1]）
			dp[j] = (dp[j] && s1[i-1] == s3[i+j-1]) ||
				(dp[j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return dp[n]
}
