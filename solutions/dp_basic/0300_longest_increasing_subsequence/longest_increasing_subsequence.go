package longestincreasingsubsequence

// LengthOfLIS 最长递增子序列（标准 DP 数组版）
// 求数组中最长严格递增子序列的长度。
// dp[i] 表示以 nums[i] 结尾的最长递增子序列长度，
// dp[i] = max(dp[j]) + 1，其中 j < i 且 nums[j] < nums[i]；答案为 max(dp)。
// 时间复杂度: O(n^2)  空间复杂度: O(n)
func LengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		dp[i] = 1 // base case：任何元素自己组成长度 1 的递增子序列
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] { // nums[i] 可以接在以 nums[j] 结尾的子序列后
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

// LengthOfLISAlternative 最长递增子序列（贪心 + 二分版）
// tails[k] 表示长度为 k+1 的递增子序列的最小结尾元素，始终严格递增。
// 对每个 x：若大于 tails 末尾则追加；否则二分替换第一个 >= x 的元素。
// 时间复杂度: O(n*log n)  空间复杂度: O(n)
func LengthOfLISAlternative(nums []int) int {
	tails := make([]int, 0, len(nums))
	for _, x := range nums {
		// 二分查找 tails 中第一个 >= x 的位置
		lo, hi := 0, len(tails)
		for lo < hi {
			mid := lo + (hi-lo)/2
			if tails[mid] < x {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		if lo == len(tails) {
			tails = append(tails, x) // x 比所有结尾都大，最长子序列延长一格
		} else {
			tails[lo] = x // 用更小的结尾替换，为后续拼接留出更多空间
		}
	}
	return len(tails)
}
