package searchinsertposition

// SearchInsert 搜索插入位置
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
// 如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 要求算法的时间复杂度为 O(log n)。
// 时间复杂度: O(log n)  空间复杂度: O(1)
func SearchInsert(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
