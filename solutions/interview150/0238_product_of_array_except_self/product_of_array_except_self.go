package productofarrayexceptself

// ProductExceptSelf 除了自身以外数组的乘积
// 给定一个整数数组 nums，返回数组 answer，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。
// 要求不使用除法，且在 O(n) 时间内完成；返回数组不计入额外空间。
// 时间复杂度: O(n) 两次遍历  空间复杂度: O(1) 除返回数组外只使用常数空间
func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// 第一遍：answer[i] 先存 nums[0..i-1] 的前缀乘积
	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	// 第二遍：用后缀乘积从右往左累乘到 answer 中
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= suffix
		suffix *= nums[i]
	}
	return answer
}
