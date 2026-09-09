package combinationsumiv

// CombinationSum4 组合总和 IV（标准 DP 数组版）
// 从互不相同的 nums 中选数（可重复选），求总和为 target 的排列个数（顺序不同算不同组合）。
// dp[i] 表示总和为 i 的排列数，dp[i] = Σ dp[i-num]（num <= i）。
// 必须外层遍历金额、内层遍历物品：对每个 i 尝试所有数作为最后一个数，
// (1,2) 与 (2,1) 会被各数一遍，因此统计的是排列数而非组合数。
// 时间复杂度: O(target × len(nums))  空间复杂度: O(target)
func CombinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1 // base case：一个数都不选，是凑出 0 的唯一方式
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				// 排列的最后一个数取 num，前面部分有 dp[i-num] 种
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
