package maximumsumcircularsubarray

// MaxSubarraySumCircular 环形子数组的最大和（标准 DP 数组版）
// 环形子数组两种形态：不跨界 = 普通最大子数组和 maxSum；
// 跨界 = 数组总和 total 减去中间最小子数组和 minSum。答案取两者较大。
// maxDP[i] / minDP[i] 表示以 nums[i] 结尾的子数组最大/最小和。
// 时间复杂度: O(n)  空间复杂度: O(n)
func MaxSubarraySumCircular(nums []int) int {
	n := len(nums)
	maxDP := make([]int, n) // 以 i 结尾的子数组最大和
	minDP := make([]int, n) // 以 i 结尾的子数组最小和
	maxDP[0], minDP[0] = nums[0], nums[0]
	total, maxSum, minSum := nums[0], nums[0], nums[0]
	for i := 1; i < n; i++ {
		maxDP[i] = max(nums[i], maxDP[i-1]+nums[i]) // 另起一段，或接在前一段后
		minDP[i] = min(nums[i], minDP[i-1]+nums[i]) // 同理求最小
		total += nums[i]
		maxSum = max(maxSum, maxDP[i])
		minSum = min(minSum, minDP[i])
	}
	if maxSum < 0 {
		// 全负数：total-minSum 对应"挖掉全部"的空子数组，不合法，只能取最大单元素
		return maxSum
	}
	return max(maxSum, total-minSum)
}

// MaxSubarraySumCircularOptimized 环形子数组的最大和（滚动变量空间优化版）
// 状态只依赖前一个位置，用滚动变量代替 maxDP/minDP 数组。
// 时间复杂度: O(n)  空间复杂度: O(1)
func MaxSubarraySumCircularOptimized(nums []int) int {
	total, curMax, curMin := nums[0], nums[0], nums[0]
	maxSum, minSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		curMax = max(nums[i], curMax+nums[i]) // 以 i 结尾的最大子数组和
		curMin = min(nums[i], curMin+nums[i]) // 以 i 结尾的最小子数组和
		maxSum = max(maxSum, curMax)
		minSum = min(minSum, curMin)
		total += nums[i]
	}
	if maxSum < 0 {
		return maxSum // 全负数特判
	}
	return max(maxSum, total-minSum)
}
