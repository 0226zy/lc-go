package longestpalindromicsubstring

// LongestPalindrome 最长回文子串
// 给你一个字符串 s，找到 s 中最长的回文子串；如果有多个最长答案，返回其中任意一个。
// 时间复杂度: O(n^2)  空间复杂度: O(1)
func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1
	// 以 i（奇数长度回文）或 i、i+1 之间（偶数长度回文）为中心向两侧扩展
	for i := 0; i < len(s); i++ {
		if l := expand(s, i, i); l > maxLen {
			start, maxLen = i-l/2, l
		}
		if l := expand(s, i, i+1); l > maxLen {
			start, maxLen = i+1-l/2, l
		}
	}
	return s[start : start+maxLen]
}

// expand 以 left、right 为中心向两侧扩展，返回得到的回文串长度
func expand(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
