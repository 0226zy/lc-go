package wigglesubsequence

// WiggleMaxLength 摆动序列（标准 DP 数组版）
// 差值严格正负交替的序列称为摆动序列，求 nums 的最长摆动子序列长度。
// up[i]/down[i] 分别表示以 nums[i] 结尾、最后一段上升/下降的摆动序列最长长度；
// nums[i] > nums[i-1] 时 up[i] = down[i-1] + 1，小于时 down[i] = up[i-1] + 1，相等时原样继承。
// 时间复杂度: O(n)  空间复杂度: O(n)
func WiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n // 空数组或单元素，自身即答案
	}
	up := make([]int, n)   // 以 i 结尾、最后一段上升的摆动序列最长长度
	down := make([]int, n) // 以 i 结尾、最后一段下降的摆动序列最长长度
	up[0], down[0] = 1, 1  // base case：单元素序列长度为 1
	for i := 1; i < n; i++ {
		switch {
		case nums[i] > nums[i-1]:
			up[i] = down[i-1] + 1 // 接在下降结尾的序列之后
			down[i] = down[i-1]
		case nums[i] < nums[i-1]:
			down[i] = up[i-1] + 1 // 接在上升结尾的序列之后
			up[i] = up[i-1]
		default:
			// 相等差值为 0，不构成摆动，两个状态都原样继承
			up[i] = up[i-1]
			down[i] = down[i-1]
		}
	}
	return max(up[n-1], down[n-1])
}

// WiggleMaxLengthOptimized 摆动序列（滚动变量空间优化版）
// up/down 只依赖前一位置，用两个滚动变量代替数组；等价于统计走向翻转的次数（贪心）。
// 时间复杂度: O(n)  空间复杂度: O(1)
func WiggleMaxLengthOptimized(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	up, down := 1, 1 // 分别代表 up[i-1]、down[i-1]
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
		// 相等时两个变量保持不变
	}
	return max(up, down)
}
