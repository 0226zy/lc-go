package sorts

// BubbleSort 冒泡排序（升序）
// 对切片 nums 进行原地升序排序。
// 时间复杂度: O(n²)  空间复杂度: O(1)
func BubbleSort(nums []int) {
	n := len(nums)

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// BubbleSortDesc 冒泡排序（降序）
// 对切片 nums 进行原地降序排序。
// 时间复杂度: O(n²)  空间复杂度: O(1)
func BubbleSortDesc(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if nums[j] < nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
