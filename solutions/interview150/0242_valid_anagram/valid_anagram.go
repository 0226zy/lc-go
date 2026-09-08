package validanagram

// IsAnagram 有效的字母异位词
// 判断字符串 t 是否是字符串 s 的字母异位词（由 s 的字符重新排列而成，顺序可以不同）。
// 时间复杂度: O(n)  空间复杂度: O(1)（固定 26 个小写字母的计数数组）
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// 先对 s 的字符计数，再用 t 的字符抵消，全部归零则说明是异位词
	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		count[t[i]-'a']--
		if count[t[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
