package oneeditdistance

// IsOneEditDistance 相隔为 1 的编辑距离
// 给定两个字符串 s 和 t，判断它们的编辑距离是否恰好为 1。
// 一次编辑操作指插入、删除或替换一个字符。
// 时间复杂度: O(n) 最多扫描一遍字符串  空间复杂度: O(1)
func IsOneEditDistance(s string, t string) bool {
	// 保证 s 始终是较短的那个串
	if len(s) > len(t) {
		s, t = t, s
	}
	// 长度差大于 1，编辑距离至少为 2
	if len(t)-len(s) > 1 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			if len(s) == len(t) {
				// 长度相等：替换 s[i]，剩余部分必须完全相同
				return s[i+1:] == t[i+1:]
			}
			// 长度差 1：在 s 中插入 t[i]，剩余部分必须完全相同
			return s[i:] == t[i+1:]
		}
	}
	// 较短串是较长串的前缀：仅当较长串恰好多一个末尾字符时编辑距离为 1
	return len(t) == len(s)+1
}
