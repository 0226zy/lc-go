package singlerowkeyboard

// CalculateTime 单行键盘
// 键盘只有一行，由 keyboard 给出 26 个小写字母的排列；手指初始在下标 0，
// 输入一个字符的耗时等于手指移动前后下标差的绝对值，返回输入整个 word 的总时间。
// 时间复杂度: O(26 + m)，m 为 word 长度  空间复杂度: O(26)
func CalculateTime(keyboard string, word string) int {
	// pos 记录每个字母在键盘上的下标（字母仅 26 个，用数组代替哈希表）
	var pos [26]int
	for i := 0; i < len(keyboard); i++ {
		pos[keyboard[i]-'a'] = i
	}

	total, prev := 0, 0
	for i := 0; i < len(word); i++ {
		cur := pos[word[i]-'a']
		if diff := cur - prev; diff >= 0 {
			total += diff
		} else {
			total -= diff
		}
		prev = cur
	}
	return total
}
