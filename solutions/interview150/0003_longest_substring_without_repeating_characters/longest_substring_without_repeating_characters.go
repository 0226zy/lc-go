package longestsubstringwithoutrepeatingcharacters

// LengthOfLongestSubstring 无重复字符的最长子串
// 给定一个字符串 s，请你找出其中不含有重复字符的最长子串的长度。
// 时间复杂度: O(n) 每个字符最多进出窗口一次  空间复杂度: O(|Σ|) Σ 为字符集大小
func LengthOfLongestSubstring(s string) int {
	// last 记录每个字符最近一次出现的位置（下标）
	last := make(map[byte]int)
	left, ans := 0, 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		// 若字符 c 已出现在当前窗口 [left, right) 内，则把左边界直接跳到它后面
		if idx, ok := last[c]; ok && idx >= left {
			left = idx + 1
		}
		last[c] = right
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}
