package factorialtrailingzeroes

// TrailingZeroes 阶乘后的零
// 返回 n! 结果中尾随零的数量。尾随零由因子 2×5 产生，2 充足，故只需统计因子 5 的个数。
// 时间复杂度: O(log n)  空间复杂度: O(1)
func TrailingZeroes(n int) int {
	count := 0
	for n > 0 {
		n /= 5
		count += n
	}
	return count
}
