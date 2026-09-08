package editdistance

// MinDistance 编辑距离
// 给定两个单词 word1 和 word2，返回将 word1 转换成 word2 所需的最少操作数。
// 允许三种操作：插入一个字符、删除一个字符、替换一个字符。
// 时间复杂度: O(m*n)  空间复杂度: O(n)（使用滚动数组，只保留上一行 DP 状态）
func MinDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	// dp[j] 表示当前已处理的 word1 前缀到 word2[:j] 的编辑距离
	// 初始状态对应 word1 为空串：把空串变成 word2[:j] 需要插入 j 次
	dp := make([]int, n+1)
	for j := 0; j <= n; j++ {
		dp[j] = j
	}

	for i := 1; i <= m; i++ {
		// prev 记录 dp[i-1][j-1]，进入本轮时它是 dp[i-1][0] = i-1
		prev := dp[0]
		// dp[i][0]：word1[:i] 变成空串需要删除 i 次
		dp[0] = i
		for j := 1; j <= n; j++ {
			old := dp[j] // 暂存 dp[i-1][j]，供下一轮作为 dp[i-1][j-1] 使用
			if word1[i-1] == word2[j-1] {
				// 末尾字符相同，无需操作，直接继承 dp[i-1][j-1]
				dp[j] = prev
			} else {
				// 末尾字符不同，取三种操作的最小值再加 1：
				//   删除 word1[i-1] → dp[i-1][j]（即更新前的 dp[j]）
				//   插入 word2[j-1] → dp[i][j-1]（即已更新的 dp[j-1]）
				//   替换 word1[i-1] 为 word2[j-1] → dp[i-1][j-1]（即 prev）
				dp[j] = 1 + min(old, min(dp[j-1], prev))
			}
			prev = old
		}
	}

	return dp[n]
}
