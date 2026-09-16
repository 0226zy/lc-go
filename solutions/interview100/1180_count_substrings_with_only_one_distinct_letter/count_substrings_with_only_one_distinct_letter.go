package countsubstringswithonlyonedistinctletter

// CountLetters 统计只含单一字母的子串
// 返回 s 中所有字符都相同的子串个数，不同位置的相同子串重复计数。
// 思路：长度为 L 的连续相同字符段贡献 L*(L+1)/2 个子串。
// 时间复杂度: O(n)  空间复杂度: O(1)
func CountLetters(s string) int {
	total, run := 0, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			// 连续相同字符段延长
			run++
		} else {
			// 当前段结束，累加该段贡献并重置
			total += run * (run + 1) / 2
			run = 1
		}
	}
	// 累加最后一段（空串时 run=1 但循环未执行，需特判）
	if len(s) > 0 {
		total += run * (run + 1) / 2
	}
	return total
}
