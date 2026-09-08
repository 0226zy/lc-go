package wordbreak

// WordBreak 单词拆分
// 判断字符串 s 能否被空格拆分为一个或多个字典 wordDict 中的单词，字典单词可重复使用。
// 时间复杂度: O(n*m*k)，n 为 s 长度，m 为最长字典单词长度，k 为子串哈希/比较开销  空间复杂度: O(n)
func WordBreak(s string, wordDict []string) bool {
	// 字典放入哈希集合加速查询，同时记录最长单词长度用于剪枝
	dict := make(map[string]struct{}, len(wordDict))
	maxLen := 0
	for _, w := range wordDict {
		dict[w] = struct{}{}
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	n := len(s)
	// dp[i] 表示 s[:i] 能否用字典单词拼出，空串视为可以
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		// 只枚举最后一个单词的起点 j，长度超过 maxLen 的子串一定不在字典中
		start := i - maxLen
		if start < 0 {
			start = 0
		}
		for j := start; j < i; j++ {
			if !dp[j] {
				continue
			}
			if _, ok := dict[s[j:i]]; ok {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
