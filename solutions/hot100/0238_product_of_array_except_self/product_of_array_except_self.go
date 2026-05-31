package productofarrayexceptself

// ProductExceptSelf 除自身以外数组的乘积
// 给你一个整数数组 nums，返回一个数组 answer，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。
// 不要使用除法，且在 O(n) 时间复杂度内完成此题。
// 时间复杂度: O(n) 两次遍历  空间复杂度: O(1) 常数空间（除输出数组外）
func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// 第一步：计算每个元素左侧的乘积
	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	// 第二步：计算每个元素右侧的乘积，并与左侧乘积相乘
	right := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= right
		right *= nums[i]
	}

	return answer
}
