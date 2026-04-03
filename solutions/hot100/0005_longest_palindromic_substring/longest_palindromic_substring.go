package longestpalindromicsubstring

// LongestPalindrome 最长回文子串
// 中心扩展：对每个位置尝试奇数长度（一个中心）与偶数长度（两个中心）的回文，
// 记录最长区间。时间复杂度 O(n²)，空间复杂度 O(1)。
func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, end := 0, 0
	for i := range s {
		if l, r := findLongest(s, i, i); r-l > end-start {
			start, end = l, r
		}
		if l, r := findLongest(s, i, i+1); r-l > end-start {
			start, end = l, r
		}
	}
	return s[start : end+1]
}

func findLongest(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
