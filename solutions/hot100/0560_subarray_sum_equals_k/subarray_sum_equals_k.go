package subarraysumequalk

// SubarraySum 和为 K 的子数组
// 前缀和 + 哈希表，统计和为 k 的连续子数组个数
// 时间复杂度: O(n) 单次遍历  空间复杂度: O(n) 哈希表存储前缀和
func SubarraySum(nums []int, k int) int {
	count := make(map[int]int)
	count[0] = 1 // 空前缀的和为 0
	ans := 0
	preSum := 0
	for _, num := range nums {
		preSum += num
		if c, ok := count[preSum-k]; ok {
			ans += c
		}
		count[preSum]++
	}
	return ans
}
