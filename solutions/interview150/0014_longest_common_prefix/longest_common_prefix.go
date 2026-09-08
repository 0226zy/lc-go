package longestcommonprefix

// LongestCommonPrefix 最长公共前缀
// 给定字符串数组 strs，返回所有字符串的最长公共前缀；不存在则返回空字符串。
// 时间复杂度: O(n * m)  n 为字符串个数，m 为最短字符串长度  空间复杂度: O(1)
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 以第一个字符串为基准，逐字符与其他字符串比较
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		// 逐个字符缩短 prefix，直到它是 strs[i] 的前缀
		for !hasPrefix(strs[i], prefix) {
			if prefix == "" {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

// hasPrefix 判断 s 是否以 prefix 开头
func hasPrefix(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	return s[:len(prefix)] == prefix
}
