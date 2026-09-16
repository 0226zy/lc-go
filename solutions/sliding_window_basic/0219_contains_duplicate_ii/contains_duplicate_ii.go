package containsduplicateii

// ContainsNearbyDuplicate 存在重复元素 II
// 判断数组中是否存在两个不同下标 i、j，使得 nums[i] == nums[j] 且 abs(i-j) <= k。
// 用哈希表记录每个值最近一次出现的下标。
// 时间复杂度: O(n)  空间复杂度: O(min(n, k))
func ContainsNearbyDuplicate(nums []int, k int) bool {
	last := make(map[int]int)
	for i, num := range nums {
		if j, ok := last[num]; ok && i-j <= k {
			return true
		}
		last[num] = i
	}
	return false
}
