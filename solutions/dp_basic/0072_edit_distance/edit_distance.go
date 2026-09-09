package editdistance

// MinDistance 编辑距离（标准 DP 数组版）
// 将 word1 转换成 word2 的最少操作数，支持插入、删除、替换三种操作。
// dp[i][j] 表示 word1 前 i 个字符转成 word2 前 j 个字符的最少操作数。
// 末尾字符相同则 dp[i][j] = dp[i-1][j-1]，否则取 删除/插入/替换 三者最小值加 1。
// 时间复杂度: O(m*n)  空间复杂度: O(m*n)
func MinDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i // base case：word2 为空，i 次删除
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j // base case：word1 为空，j 次插入
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] // 末尾字符相同，无需操作
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) // 删除/插入/替换
			}
		}
	}
	return dp[m][n]
}

// MinDistanceOptimized 编辑距离（滚动一维数组空间优化版）
// dp[j] 就地更新：更新前是正上方 dp[i-1][j]，用 prev 记录左上 dp[i-1][j-1]。
// 时间复杂度: O(m*n)  空间复杂度: O(min(m,n))
func MinDistanceOptimized(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m < n { // 让 word2 是较短的那个，压缩空间
		word1, word2 = word2, word1
		m, n = n, m
	}
	dp := make([]int, n+1)
	for j := 0; j <= n; j++ {
		dp[j] = j // 第 0 行：空串转化需 j 次插入
	}
	for i := 1; i <= m; i++ {
		prev := dp[0] // prev 记录左上 dp[i-1][j-1]
		dp[0] = i     // 第 i 行第 0 列：i 次删除
		for j := 1; j <= n; j++ {
			tmp := dp[j] // 暂存正上方 dp[i-1][j]，作为下一轮的左上
			if word1[i-1] == word2[j-1] {
				dp[j] = prev // 末尾字符相同，无需操作
			} else {
				dp[j] = 1 + min(dp[j], min(dp[j-1], prev)) // 删除/插入/替换
			}
			prev = tmp
		}
	}
	return dp[n]
}
