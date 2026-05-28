package longestincreasingsubsequence

// LengthOfLIS 最长递增子序列
// 给定一个整数数组 nums，找到其中最长严格递增子序列的长度。
// 时间复杂度: O(n^2)  空间复杂度: O(n)
func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	res := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
