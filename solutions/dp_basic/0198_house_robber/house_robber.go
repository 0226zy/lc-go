package houserobber

// Rob 打家劫舍（标准 DP 数组版）
// 沿街偷窃，相邻两间不能同时偷，求能偷到的最高金额。
// dp[i] 表示只考虑前 i 间房屋时的最高金额，
// dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])（不偷第 i 间，或偷第 i 间）。
// 时间复杂度: O(n)  空间复杂度: O(n)
func Rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[1] = nums[0] // base case：dp[0] = 0（无房可偷），dp[1] = nums[0]
	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1]) // 不偷 or 偷第 i 间
	}
	return dp[n]
}

// RobOptimized 打家劫舍（滚动变量空间优化版）
// dp[i] 只依赖前两个状态，用两个滚动变量代替 dp 数组。
// 时间复杂度: O(n)  空间复杂度: O(1)
func RobOptimized(nums []int) int {
	prev2, prev1 := 0, 0 // 分别代表 dp[i-2]、dp[i-1]，初始对应「0 间房」
	for _, money := range nums {
		prev2, prev1 = prev1, max(prev1, prev2+money)
	}
	return prev1
}
