package longestsubstring

// LengthOfLongestSubstring 无重复字符的最长子串
// 滑动窗口 + 固定数组，避免 map 哈希开销
// 时间复杂度: O(n) 单次遍历  空间复杂度: O(1) 固定 128 大小数组
func LengthOfLongestSubstring(s string) int {
	var pos [128]int // pos[c] 存储字符 c 最近出现位置 +1，0 表示未出现
	maxLen := 0
	left := 0
	for right := 0; right < len(s); right++ {
		if p := pos[s[right]]; p > left {
			left = p
		}
		pos[s[right]] = right + 1
		if l := right - left + 1; l > maxLen {
			maxLen = l
		}
	}
	return maxLen
}
