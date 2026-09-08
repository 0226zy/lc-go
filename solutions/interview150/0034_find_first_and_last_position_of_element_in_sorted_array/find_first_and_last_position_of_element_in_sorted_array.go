package findfirstandlastpositionofelementinsortedarray

// lowerBound 二分查找下界
// 返回有序数组 nums 中第一个 >= target 的元素下标（可能等于 len(nums)）。
// 模板要点：right 初始为 len(nums)，循环条件 left < right，
// nums[mid] >= target 时 right = mid（保留 mid），否则 left = mid + 1。
// 时间复杂度: O(log n)  空间复杂度: O(1)
func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid // mid 本身可能就是答案，保留
		} else {
			left = mid + 1
		}
	}
	return left
}

// SearchRange 在排序数组中查找元素的第一个和最后一个位置
// 给定非递减数组 nums 和目标值 target，返回 target 在数组中的开始位置和结束位置；
// 不存在则返回 [-1, -1]。要求 O(log n)。
// 思路：第一个位置 = lowerBound(target)；最后一个位置 = lowerBound(target+1) - 1。
// 时间复杂度: O(log n) 两次二分  空间复杂度: O(1)
func SearchRange(nums []int, target int) []int {
	first := lowerBound(nums, target)
	if first == len(nums) || nums[first] != target {
		return []int{-1, -1}
	}
	last := lowerBound(nums, target+1) - 1
	return []int{first, last}
}
