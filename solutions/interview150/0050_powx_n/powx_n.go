package powxn

// MyPow Pow(x, n)
// 实现 pow(x, n)：计算 x 的 n 次幂函数，n 是整数（可为负数），不得使用标准库 math.Pow。
// 时间复杂度: O(log n)  空间复杂度: O(1)
func MyPow(x float64, n int) float64 {
	// 指数转为 int64 处理：n 可能等于 -2^31，直接对 int 取负会溢出
	m := int64(n)
	if m < 0 {
		x = 1 / x
		m = -m
	}

	// 快速幂（二进制分解）：把指数按二进制位拆开，遇到 1 的位就把对应的 x^(2^k) 乘进结果
	ans := 1.0
	for m > 0 {
		if m&1 == 1 {
			ans *= x
		}
		x *= x
		m >>= 1
	}
	return ans
}
