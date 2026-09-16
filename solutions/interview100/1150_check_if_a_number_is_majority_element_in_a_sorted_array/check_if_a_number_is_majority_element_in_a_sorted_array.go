package checkifanumberismajorityelementinasortedarray

import "sort"

// IsMajorityElement 检查一个数是否在数组中占绝大多数
// 给定非递减排序数组 nums 和目标值 target，若 target 出现次数严格大于 len(nums)/2 则返回 true。
// 时间复杂度: O(log n)  空间复杂度: O(1)
func IsMajorityElement(nums []int, target int) bool {
	n := len(nums)
	// 二分找到 target 第一次出现的位置
	first := sort.SearchInts(nums, target)
	// target 不存在
	if first == n || nums[first] != target {
		return false
	}
	// 相同元素连续排列：出现次数 > n/2 等价于 first + n/2 处仍是 target
	next := first + n/2
	return next < n && nums[next] == target
}
