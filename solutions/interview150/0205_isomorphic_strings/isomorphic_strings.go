package isomorphicstrings

// IsIsomorphic 同构字符串
// 判断字符串 s 和 t 是否同构：s 中每个字符可以映射到 t 中唯一的一个字符，且 t 中每个字符也只能被 s 中唯一的一个字符映射。
// 时间复杂度: O(n)  空间复杂度: O(1)（字符集大小固定）
func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// s2t: s 中字符到 t 中字符的映射；t2s: t 到 s 的反向映射，保证“双射”
	s2t := make([]int, 128)
	t2s := make([]int, 128)
	for i := 0; i < len(s); i++ {
		cs, ct := s[i], t[i]
		// 映射不存在时记录；存在时必须与已有映射一致
		if s2t[cs] == 0 && t2s[ct] == 0 {
			s2t[cs] = int(ct) + 1 // +1 是为了用 0 表示“未映射”
			t2s[ct] = int(cs) + 1
		} else if s2t[cs] != int(ct)+1 || t2s[ct] != int(cs)+1 {
			return false
		}
	}
	return true
}
