package zigzagconversion

import "strings"

// Convert Z 字形变换
// 将字符串 s 以从上到下、从左到右的 Z 字形排列成 numRows 行，
// 再按行从左到右读取，返回读取出的字符串。
// 时间复杂度: O(n)  空间复杂度: O(n)
func Convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	var sb strings.Builder
	sb.Grow(len(s))
	cycle := 2*numRows - 2 // 一个完整“向下 + 向上”周期包含的字符数

	for row := 0; row < numRows; row++ {
		for i := row; i < len(s); i += cycle {
			sb.WriteByte(s[i]) // 每个周期内，第 row 行向下的字符
			// 中间行（非首行非末行）还有一个向上斜线经过的字符
			if row > 0 && row < numRows-1 {
				j := i + cycle - 2*row
				if j < len(s) {
					sb.WriteByte(s[j])
				}
			}
		}
	}
	return sb.String()
}
