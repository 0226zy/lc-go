package wordbreak

// WordBreak 单词拆分（标准 DP 数组版）
// 判断字符串 s 能否用字典 wordDict 中的单词（可重复使用）拼接而成。
// dp[i] 表示前 i 个字符 s[:i] 能否被拆分，
// dp[i] = true 当且仅当存在 j < i 使得 dp[j] == true 且 s[j:i] 在字典中。
// 时间复杂度: O(n^2)  空间复杂度: O(n + 字典总字符数)
func WordBreak(s string, wordDict []string) bool {
	n := len(s)
	// 字典放入哈希表，O(1) 查询某个子串是否是合法单词
	wordSet := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		wordSet[w] = true
	}
	dp := make([]bool, n+1)
	dp[0] = true // base case：空前缀不需要任何单词
	for i := 1; i <= n; i++ {
		// 枚举最后一个单词的起点 j：s[:j] 能拼出，且 s[j:i] 本身是单词
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break // 找到一种拼法即可
			}
		}
	}
	return dp[n]
}
