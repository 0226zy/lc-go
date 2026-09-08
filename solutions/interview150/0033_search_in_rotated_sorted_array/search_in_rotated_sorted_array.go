package searchinrotatedsortedarray

// Search 搜索旋转排序数组
// 给定一个按升序排列、值互不相同、且在某下标处旋转过的整数数组 nums，
// 查找目标值 target，存在则返回其下标，否则返回 -1。要求 O(log n)。
// 核心思想：旋转数组从中间切一刀，必有一半完全有序，根据 target 是否
// 落在有序半区来决定丢弃哪一半。
// 时间复杂度: O(log n)  空间复杂度: O(1)
func Search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2 // 防溢出
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			// 左半 [left, mid] 有序
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 右半 (mid, right] 有序
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
