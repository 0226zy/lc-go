package sorts

// InsertionSort 插入排序（升序）
// 对切片 nums 进行原地升序排序。
// 时间复杂度: O(n²)  空间复杂度: O(1)
func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
}

// InsertionSortDesc 插入排序（降序）
// 对切片 nums 进行原地降序排序。
// 时间复杂度: O(n²)  空间复杂度: O(1)
func InsertionSortDesc(nums []int) {
	for i := 1; i < len(nums); i++ {
		currVal := nums[i]
		targetIdx := i - 1 // 先比较排序后的第一个
		// 从排序后的最后一个往前，如果当前值，比前面的大，则全部后移一位，给当前值腾位置
		for targetIdx >= 0 && nums[targetIdx] < currVal {
			nums[targetIdx+1] = nums[targetIdx]
			targetIdx--
		}
		// for 循环中 targetIdx 最后多减了一次
		nums[targetIdx+1] = currVal
	}
}
