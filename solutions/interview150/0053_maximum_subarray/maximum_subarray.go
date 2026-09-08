package maximumsubarray

// MaxSubArray 最大子数组和
// 给定一个整数数组 nums，找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 使用 Kadane 算法：维护“以当前元素结尾的子数组最大和”，前面的和为负就果断抛弃、从当前元素重新开始。
// 时间复杂度: O(n) 一次遍历  空间复杂度: O(1) 常数变量
func MaxSubArray(nums []int) int {
	ans, cur := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		// cur + nums[i] 比 nums[i] 还小，说明前面的和是负资产，果断抛弃
		if cur+nums[i] > nums[i] {
			cur += nums[i]
		} else {
			cur = nums[i]
		}
		if cur > ans {
			ans = cur
		}
	}
	return ans
}
