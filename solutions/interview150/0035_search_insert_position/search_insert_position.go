package searchinsertposition

// SearchInsert 搜索插入位置
// 给定升序无重复数组 nums 和目标值 target，返回 target 的下标；
// 若不存在，返回按顺序插入的位置。本质是求 lower_bound（第一个 >= target 的位置）。
// 时间复杂度: O(log n) 二分查找  空间复杂度: O(1) 常数变量
func SearchInsert(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid // mid 可能是答案，保留在区间内
		} else {
			left = mid + 1 // mid 及左侧都太小，答案在右侧
		}
	}
	return left
}
