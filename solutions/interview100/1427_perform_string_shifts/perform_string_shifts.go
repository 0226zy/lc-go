package performstringshifts

// StringShift 字符串的左右移
// 按顺序对字符串 s 执行 shift 中的移动操作（direction 为 0 左移、1 右移），返回最终字符串。
// 时间复杂度: O(m + n)，m 为操作数，n 为字符串长度  空间复杂度: O(n)
func StringShift(s string, shift [][]int) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	// 累计净位移：左移记为正，右移记为负
	net := 0
	for _, op := range shift {
		if op[0] == 0 {
			net += op[1]
		} else {
			net -= op[1]
		}
	}
	// 归一化为 [0, n) 内的左移位数（移动 n 位等于不动）
	net = ((net % n) + n) % n
	// 左移 net 位 = 前 net 个字符搬到末尾
	return s[net:] + s[:net]
}
