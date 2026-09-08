package textjustification

import "strings"

// FullJustify 文本左右对齐
// 给定单词数组 words 和宽度 maxWidth，用贪心策略排版：每行放尽可能多的单词，
// 单词间空格尽量均匀（左侧多于右侧），最后一行与单行词左对齐、行尾补空格。
// 时间复杂度: O(L)  L 为输出字符总数  空间复杂度: O(L) 答案存储
func FullJustify(words []string, maxWidth int) []string {
	res := []string{}
	for i := 0; i < len(words); {
		// 1. 贪心取词：cnt 为行内单词数，lineLen 为单词字符总长（不含空格）
		cnt, lineLen := 1, len(words[i])
		for i+cnt < len(words) && lineLen+len(words[i+cnt])+cnt <= maxWidth {
			lineLen += len(words[i+cnt])
			cnt++
		}

		// 2. 组装这一行
		var sb strings.Builder
		if i+cnt == len(words) || cnt == 1 {
			// 最后一行或只有一个词：左对齐，词间 1 个空格，行尾补满空格
			for j := 0; j < cnt; j++ {
				if j > 0 {
					sb.WriteByte(' ')
				}
				sb.WriteString(words[i+j])
			}
			for sb.Len() < maxWidth {
				sb.WriteByte(' ')
			}
		} else {
			// 普通行：空格均匀分配到词间空隙，余数从左往右每个空隙多分 1 个
			space, rem := (maxWidth-lineLen)/(cnt-1), (maxWidth-lineLen)%(cnt-1)
			for j := 0; j < cnt-1; j++ {
				sb.WriteString(words[i+j])
				sb.WriteString(strings.Repeat(" ", space))
				if j < rem {
					sb.WriteByte(' ')
				}
			}
			sb.WriteString(words[i+cnt-1]) // 行尾不补空格
		}

		res = append(res, sb.String())
		i += cnt
	}
	return res
}
