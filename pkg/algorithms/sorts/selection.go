package sorts

// SelectionSort 选择排序（升序）
// 对切片 nums 进行原地升序排序。
// 时间复杂度: O(n²)  空间复杂度: O(1)
func SelectionSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			nums[i], nums[minIdx] = nums[minIdx], nums[i]
		}
	}
}

// SelectionSortDesc 选择排序（降序）
// 对切片 nums 进行原地降序排序。
// 时间复杂度: O(n²)  空间复杂度: O(1)
func SelectionSortDesc(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[maxIdx] {
				maxIdx = j
			}
		}
		if maxIdx != i {
			nums[i], nums[maxIdx] = nums[maxIdx], nums[i]
		}
	}
}
