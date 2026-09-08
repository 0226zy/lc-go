package reversebits

// ReverseBits 颠倒二进制位
// 颠倒给定的 32 位无符号整数的二进制位。
// 逐位拆分：每次取 n 的最低位，追加到结果的最高位（等价于结果左移一位后加上该位）。
// 时间复杂度: O(1) 固定 32 次迭代  空间复杂度: O(1)
func ReverseBits(n uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = res<<1 | n&1 // 结果左移腾出最低位，放入 n 的当前最低位
		n >>= 1            // n 右移，下一位变成最低位
	}
	return res
}
