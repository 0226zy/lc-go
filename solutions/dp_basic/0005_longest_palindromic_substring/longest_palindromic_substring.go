package longestpalindromicsubstring

// LongestPalindrome 最长回文子串（标准 DP 数组版）
// dp[i][j] 表示子串 s[i..j] 是否为回文串。
// dp[i][j] = (s[i] == s[j]) && (j-i < 2 || dp[i+1][j-1])，按子串长度递增枚举。
// 时间复杂度: O(n²)  空间复杂度: O(n²)
func LongestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	start, maxLen := 0, 1                    // 最长回文的起点和长度
	for length := 1; length <= n; length++ { // 按子串长度枚举，保证内部子串先求解
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			if s[i] == s[j] {
				// 两端相等：长度小于 3 必为回文，否则看内部子串
				dp[i][j] = length < 3 || dp[i+1][j-1]
			}
			if dp[i][j] && length > maxLen {
				start, maxLen = i, length
			}
		}
	}
	return s[start : start+maxLen]
}

// LongestPalindromeAlternative 最长回文子串（中心扩展法）
// 枚举每个回文中心（单字符 + 相邻字符间隙，共 2n-1 个），向两边扩展找最长回文。
// 时间复杂度: O(n²)  空间复杂度: O(1)
func LongestPalindromeAlternative(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	start, maxLen := 0, 1
	expand := func(l, r int) { // 从中心 l,r 向两边扩展
		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 > maxLen {
				start, maxLen = l, r-l+1
			}
			l--
			r++
		}
	}
	for i := 0; i < n; i++ {
		expand(i, i)   // 奇数长度，中心是单个字符
		expand(i, i+1) // 偶数长度，中心是两个字符之间
	}
	return s[start : start+maxLen]
}
