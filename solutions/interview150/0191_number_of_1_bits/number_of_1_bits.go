package numberof1bits

// HammingWeight 位1的个数（汉明重量）
// 编写一个函数，统计无符号整数二进制表示中 1 的个数。
// 使用 n&(n-1) 技巧：该操作恰好清除 n 最低位的 1，循环次数等于 1 的个数。
// 时间复杂度: O(k) k 为 1 的个数  空间复杂度: O(1)
func HammingWeight(n uint32) int {
	count := 0
	for n != 0 {
		n &= n - 1 // 清除最低位的 1
		count++
	}
	return count
}
