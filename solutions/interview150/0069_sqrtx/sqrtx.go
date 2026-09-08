package sqrtx

// MySqrt x 的平方根
// 给定一个非负整数 x，计算并返回 x 的算术平方根的整数部分（向下取整），
// 不使用任何内置指数函数和算符（如 math.Sqrt）。
// 时间复杂度: O(log x) 二分查找  空间复杂度: O(1) 只使用常数变量
func MySqrt(x int) int {
	// 0 和 1 的平方根就是自身，直接返回
	if x < 2 {
		return x
	}

	// 在闭区间 [2, x/2] 上二分查找最大的 mid，满足 mid*mid <= x
	// x/2 是上界：x >= 2 时其平方根一定不超过 x/2
	lo, hi, ans := 2, x/2, 1
	for lo <= hi {
		mid := lo + (hi-lo)/2 // 等价 (lo+hi)/2，养成防溢出习惯
		if mid <= x/mid {     // 用除法比较等价于 mid*mid <= x，彻底规避乘法溢出
			ans = mid  // mid 是一个可行解，尝试向右找更大的
			lo = mid + 1
		} else {
			hi = mid - 1 // mid 太大，向左缩小范围
		}
	}
	return ans
}
