package deleteandearn

// DeleteAndEarn 删除并获得点数（标准 DP 数组版）
// 删除 nums[i] 得 nums[i] 点，但所有 nums[i]-1 和 nums[i]+1 必须一并删除，求最大点数。
// 先统计每个数值的总点数 sum[x]，转化为打家劫舍：数值轴上相邻数值不能同选。
// dp[i] 表示只考虑数值不超过 i 时的最大点数，dp[i] = max(dp[i-1], dp[i-2]+sum[i])。
// 时间复杂度: O(n + M)  空间复杂度: O(M)（M 为数值范围上限）
func DeleteAndEarn(nums []int) int {
	maxVal := 0
	for _, v := range nums {
		maxVal = max(maxVal, v)
	}
	sum := make([]int, maxVal+1)
	for _, v := range nums {
		sum[v] += v // 聚合每个数值的总点数
	}
	dp := make([]int, maxVal+1)
	dp[0] = sum[0]
	dp[1] = max(sum[0], sum[1]) // base case：0 和 1 相邻，二选一
	for i := 2; i <= maxVal; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+sum[i]) // 不拿 i，或拿 i 并放弃 i-1
	}
	return dp[maxVal]
}

// DeleteAndEarnOptimized 删除并获得点数（滚动变量空间优化版）
// dp[i] 只依赖前两个状态，用两个滚动变量代替 dp 数组（sum 桶数组仍需保留）。
// 时间复杂度: O(n + M)  空间复杂度: O(M)（仅为 sum 桶）
func DeleteAndEarnOptimized(nums []int) int {
	maxVal := 0
	for _, v := range nums {
		maxVal = max(maxVal, v)
	}
	sum := make([]int, maxVal+1)
	for _, v := range nums {
		sum[v] += v
	}
	prev2, prev1 := 0, sum[0] // 分别代表 dp[i-2]、dp[i-1]（i 从 1 开始）
	for i := 1; i <= maxVal; i++ {
		prev2, prev1 = prev1, max(prev1, prev2+sum[i])
	}
	return prev1
}
