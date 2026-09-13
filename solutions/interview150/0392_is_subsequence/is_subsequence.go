package issubsequence

// IsSubsequence 判断子序列
// 给定字符串 s 和 t，判断 s 是否为 t 的子序列（可删除 t 中若干字符，但不改变相对顺序）。
// 时间复杂度: O(n) n 为 t 的长度  空间复杂度: O(1)
func IsSubsequence(s, t string) bool {
	i := 0
	for j := 0; i < len(s) && j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == len(s)
}
