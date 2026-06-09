package sorts

// QuickSort 快速排序（升序）
// 对切片 nums 进行原地升序排序。
// 时间复杂度: 平均 O(n log n)，最坏 O(n²)  空间复杂度: O(log n)
func QuickSort(nums []int) {
	quickSort(nums, func(a, b int) bool {
		return a < b
	})
}

// QuickSortDesc 快速排序（降序）
// 对切片 nums 进行原地降序排序。
// 时间复杂度: 平均 O(n log n)，最坏 O(n²)  空间复杂度: O(log n)
func QuickSortDesc(nums []int) {
	quickSort(nums, func(a, b int) bool {
		return a > b
	})
}

func quickSort(nums []int, less func(a, b int) bool) {
	if len(nums) <= 1 {
		return
	}
	quickSortHelper(nums, 0, len(nums)-1, less)
}

func quickSortHelper(nums []int, left, right int, less func(a, b int) bool) {
	if left >= right {
		return
	}

	pivotIndex := partition(nums, left, right, less)
	quickSortHelper(nums, left, pivotIndex-1, less)
	quickSortHelper(nums, pivotIndex+1, right, less)
}

func partition(nums []int, left, right int, less func(a, b int) bool) int {
	pivot := nums[right]
	storeIndex := left

	for i := left; i < right; i++ {
		if less(nums[i], pivot) {
			nums[storeIndex], nums[i] = nums[i], nums[storeIndex]
			storeIndex++
		}
	}

	nums[storeIndex], nums[right] = nums[right], nums[storeIndex]
	return storeIndex
}
