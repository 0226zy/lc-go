package maximumsubarray

// MaxSubArray 最大子数组和
// Kadane 算法，动态规划空间优化
// 时间复杂度: O(n) 单次遍历  空间复杂度: O(1) 常数空间
func MaxSubArray(nums []int) int {
	maxSum := nums[0]
	curSum := 0
	for _, num := range nums {
		if curSum > 0 {
			curSum += num
		} else {
			curSum = num
		}
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}
