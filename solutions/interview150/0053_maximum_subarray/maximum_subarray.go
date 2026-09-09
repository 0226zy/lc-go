package maximumsubarray

// MaxSubArray 最大子数组和（标准 DP 数组版）
// 给定一个整数数组 nums，找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// dp[i] 表示以下标 i 结尾的连续子数组的最大和。
// dp[i] = max(nums[i], dp[i-1] + nums[i])，答案是所有 dp[i] 的最大值。
// 时间复杂度: O(n)  空间复杂度: O(n)
func MaxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0] // base case：以第 0 个元素结尾只有它自己
	ans := dp[0]
	for i := 1; i < n; i++ {
		// 要么接上前面的子数组，要么从 nums[i] 另起炉灶
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		ans = max(ans, dp[i])
	}
	return ans
}

// MaxSubArrayOptimized 最大子数组和（滚动变量空间优化版，即 Kadane 算法）
// dp[i] 只依赖 dp[i-1]，用一个滚动变量代替 dp 数组。
// 时间复杂度: O(n)  空间复杂度: O(1)
func MaxSubArrayOptimized(nums []int) int {
	prev, ans := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		prev = max(nums[i], prev+nums[i]) // 前面累计和为负就另起炉灶
		ans = max(ans, prev)
	}
	return ans
}
