package groupshiftedstrings

// GroupStrings 移位字符串分组
// 给定一个字符串数组 strings，将属于同一个「移位序列」的字符串分到同一组。
// 移位序列：对每个字母整体后移若干位（'z' 的下一字母为 'a'）所能得到的全部字符串。
// 时间复杂度: O(n*m) n 为字符串个数、m 为平均长度  空间复杂度: O(n*m)
func GroupStrings(strings []string) [][]string {
	// key 为相邻字符循环差值的定宽字节编码，互为移位字符串的 key 必然相同
	groups := make(map[string][]string)
	for _, s := range strings {
		key := shiftKey(s)
		groups[key] = append(groups[key], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

// shiftKey 计算字符串的循环差值编码：diff[i] = (s[i+1] - s[i] + 26) % 26
// 差值取值范围 0..25，恰好各占一个字节，拼接后不会歧义
func shiftKey(s string) string {
	buf := make([]byte, 0, len(s))
	for i := 1; i < len(s); i++ {
		diff := (int(s[i]) - int(s[i-1]) + 26) % 26
		buf = append(buf, byte(diff))
	}
	return string(buf)
}
