package findminimuminrotatedsortedarray

// FindMin 寻找旋转排序数组中的最小值
// 使用二分查找定位数组中的最小值（旋转点）。
// 时间复杂度: O(log n)  空间复杂度: O(1)
func FindMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			// 最小值在右半部分
			left = mid + 1
		} else {
			// nums[mid] < nums[right]，最小值在左半部分或就是 mid
			// 由于元素互不相同，不会相等
			right = mid
		}
	}

	return nums[left]
}

// FindMinLinear 线性扫描寻找最小值
// 遍历数组找到最小值，仅作为对比解法。
// 时间复杂度: O(n)  空间复杂度: O(1)
func FindMinLinear(nums []int) int {
	minVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < minVal {
			minVal = nums[i]
		}
	}
	return minVal
}
