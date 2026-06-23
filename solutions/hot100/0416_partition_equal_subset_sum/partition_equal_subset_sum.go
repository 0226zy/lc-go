package partitionequalsubsetsum

// CanPartition 分割等和子集
// 给定一个只包含正整数的非空数组 nums，判断是否可以将它分割成两个和相等的子集。
// 时间复杂度: O(n * target)  空间复杂度: O(target)
func CanPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	// 总和为奇数时无法均分
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[target]
}
