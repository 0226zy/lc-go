package houserobber

// Rob 打家劫舍（标准 DP 数组版）
// 你是一名专业小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，
// 相邻的房屋装有相互连通的防盗系统：如果两间相邻的房屋在同一晚上被闯入，系统会自动报警。
// 给定代表每个房屋存放金额的非负整数数组 nums，计算在不触动警报的情况下一夜之内能够偷窃到的最高金额。
// dp[i] = max(dp[i-1], dp[i-2]+nums[i])，表示偷到第 i 间房时的最大金额。
// 时间复杂度: O(n)  空间复杂度: O(n)
func Rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1]) // 前两间选金额更大的那间
	for i := 2; i < n; i++ {
		// 不偷第 i 间：沿用 dp[i-1]；偷第 i 间：dp[i-2]+nums[i]
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

// RobOptimized 打家劫舍（滚动变量空间优化版）
// dp[i] 只依赖前两个状态，用 prev2 / prev1 滚动即可。
// 时间复杂度: O(n)  空间复杂度: O(1)
func RobOptimized(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	prev2, prev1 := nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		prev2, prev1 = prev1, max(prev1, prev2+nums[i])
	}
	return prev1
}
