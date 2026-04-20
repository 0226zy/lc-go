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

// LongestPalindromeDP 动态规划解法
// dp[i][j] 表示 s[i..j] 是否为回文。按子串长度从小到大填表，
// 记录最长回文的起止下标。时间 O(n²)，空间 O(n²)。
func LongestPalindromeDP(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true // 长度 1，单字符必回文
	}

	start, maxLen := 0, 1

	// 长度 2
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start, maxLen = i, 2
		}
	}

	// 长度 3 ~ n
	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				if length > maxLen {
					start, maxLen = i, length
				}
			}
		}
	}

	return s[start : start+maxLen]
}
