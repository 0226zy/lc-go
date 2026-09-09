package houserobberii

// Rob 打家劫舍 II（标准 DP 数组版）
// 房屋围成一圈，首尾相邻不能同时偷，求能偷到的最高金额。
// 拆成两个线性子问题：不偷最后一间 robRange(nums, 0, n-2)，
// 不偷第一间 robRange(nums, 1, n-1)，取两者最大值。
// 时间复杂度: O(n)  空间复杂度: O(n)
func Rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0] // 只有一间房，既是首也是尾，直接偷
	}
	return max(robRange(nums, 0, n-2), robRange(nums, 1, n-1))
}

// robRange 对闭区间 [lo, hi] 内的房屋跑 0198 的线性打家劫舍 DP
// dp[i] 表示只考虑区间内前 i 间房屋时的最高金额
func robRange(nums []int, lo, hi int) int {
	m := hi - lo + 1
	dp := make([]int, m+1)
	dp[1] = nums[lo] // base case：dp[0] = 0，dp[1] = 区间第一间
	for i := 2; i <= m; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[lo+i-1]) // 不偷 or 偷区间内第 i 间
	}
	return dp[m]
}

// RobOptimized 打家劫舍 II（滚动变量空间优化版）
// 环形仍需拆成两遍线性 DP，但每遍内部只依赖前两个状态，可用滚动变量。
// 时间复杂度: O(n)  空间复杂度: O(1)
func RobOptimized(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(robRangeOptimized(nums, 0, n-2), robRangeOptimized(nums, 1, n-1))
}

// robRangeOptimized 线性区间的滚动变量版
func robRangeOptimized(nums []int, lo, hi int) int {
	prev2, prev1 := 0, 0 // 分别代表 dp[i-2]、dp[i-1]
	for i := lo; i <= hi; i++ {
		prev2, prev1 = prev1, max(prev1, prev2+nums[i])
	}
	return prev1
}
