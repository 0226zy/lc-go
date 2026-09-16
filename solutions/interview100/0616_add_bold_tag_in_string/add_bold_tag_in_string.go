package addboldtaginstring

import "strings"

// AddBoldTag 给字符串添加加粗标签
// 将 s 中所有在 words 中出现过的子串用 <b></b> 包裹，重叠或相邻的区间合并为一对标签。
// 时间复杂度: O(m·n·L)  空间复杂度: O(n)，m 为单词数，n 为 s 长度，L 为单词平均长度
func AddBoldTag(s string, words []string) string {
	n := len(s)
	bold := make([]bool, n) // bold[i] 表示 s[i] 是否需要加粗

	// 第一步：标记所有需要加粗的字符
	for _, word := range words {
		if word == "" {
			continue
		}
		// 从 from 处开始查找 word 的下一次出现，+1 步进以覆盖重叠出现
		for from := 0; from+len(word) <= n; {
			idx := strings.Index(s[from:], word)
			if idx < 0 {
				break
			}
			start := from + idx
			for i := start; i < start+len(word); i++ {
				bold[i] = true
			}
			from = start + 1
		}
	}

	// 第二步：扫描标记数组，在加粗区间边界处插入标签
	var sb strings.Builder
	for i := 0; i < n; i++ {
		if bold[i] && (i == 0 || !bold[i-1]) {
			sb.WriteString("<b>") // 进入一个加粗区间
		}
		sb.WriteByte(s[i])
		if bold[i] && (i == n-1 || !bold[i+1]) {
			sb.WriteString("</b>") // 离开一个加粗区间
		}
	}
	return sb.String()
}
