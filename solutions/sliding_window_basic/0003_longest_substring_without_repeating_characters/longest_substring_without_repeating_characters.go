package longestsubstringwithoutrepeatingcharacters

// LengthOfLongestSubstring 无重复字符的最长子串
// 给定一个字符串 s，请你找出其中不含有重复字符的最长子串的长度。
// 时间复杂度: O(n) 每个字符最多进出窗口一次  空间复杂度: O(|Σ|) Σ 为字符集大小
func LengthOfLongestSubstring(s string) int {
	ret := 0
	last := make(map[byte]int)
	left, right := 0, 0
	for right < len(s) {
		if idx, ok := last[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		last[s[right]] = right
		ret = max(ret, right-left+1)
		right++
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
