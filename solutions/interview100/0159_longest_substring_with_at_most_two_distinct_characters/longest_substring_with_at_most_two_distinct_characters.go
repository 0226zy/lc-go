package longestsubstringwithatmosttwodistinctcharacters

// LengthOfLongestSubstringTwoDistinct 至多包含两个不同字符的最长子串
// 给定一个字符串 s，找出至多包含两个不同字符的最长子串的长度。
// 时间复杂度: O(n) 每个字符最多进出哈希表一次  空间复杂度: O(1) 哈希表最多存 3 个键
func LengthOfLongestSubstringTwoDistinct(s string) int {
	// last 记录当前窗口内每种字符最近一次出现的位置（下标）
	last := make(map[byte]int)
	left, ans := 0, 0
	for right := 0; right < len(s); right++ {
		last[s[right]] = right
		// 窗口内出现了第 3 种字符，移出最近出现位置最靠左的那种
		if len(last) > 2 {
			minIdx := len(s)
			var minChar byte
			for c, idx := range last {
				if idx < minIdx {
					minIdx = idx
					minChar = c
				}
			}
			delete(last, minChar)
			left = minIdx + 1
		}
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}
