package arithmeticslices

// NumberOfArithmeticSlices 等差数列划分（标准 DP 数组版）
// 求数组中所有长度至少为 3 的等差子数组个数。
// dp[i] 表示以 nums[i] 结尾的等差子数组个数：
// 若 nums[i]-nums[i-1] == nums[i-1]-nums[i-2]，则 dp[i] = dp[i-1] + 1，否则为 0。
// 答案为所有 dp[i] 之和。
// 时间复杂度: O(n)  空间复杂度: O(n)
func NumberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	ans := 0
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1 // 延伸旧子数组 + 新增长度 3 的子数组
		}
		ans += dp[i]
	}
	return ans
}

// NumberOfArithmeticSlicesOptimized 等差数列划分（滚动变量空间优化版）
// dp[i] 只依赖 dp[i-1]，用滚动变量 cur 代替 dp 数组。
// 时间复杂度: O(n)  空间复杂度: O(1)
func NumberOfArithmeticSlicesOptimized(nums []int) int {
	cur, ans := 0, 0
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			cur++
		} else {
			cur = 0 // 等差被打断，重新计数
		}
		ans += cur
	}
	return ans
}
