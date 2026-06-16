package findfirstandlastpositionofelementinsortedarray

// SearchRange 在排序数组中查找元素的第一个和最后一个位置
// 使用两次二分查找分别定位 target 的左边界和右边界。
// 时间复杂度: O(log n)  空间复杂度: O(1)
func SearchRange(nums []int, target int) []int {
	first := findFirst(nums, target)
	if first == -1 {
		return []int{-1, -1}
	}
	last := findLast(nums, target)
	return []int{first, last}
}

// findFirst 二分查找 target 第一次出现的位置
func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ans := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			ans = mid
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

// findLast 二分查找 target 最后一次出现的位置
func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ans := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			ans = mid
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

// SearchRangeBinarySearchAndExpand 二分查找 + 线性扩展
// 先二分查找 target 的任意一个位置，再向两侧线性扩展得到完整区间。
// 时间复杂度: 最坏 O(n)，平均 O(log n)  空间复杂度: O(1)
func SearchRangeBinarySearchAndExpand(nums []int, target int) []int {
	idx := binarySearchAny(nums, target)
	if idx == -1 {
		return []int{-1, -1}
	}

	left, right := idx, idx
	for left > 0 && nums[left-1] == target {
		left--
	}
	for right < len(nums)-1 && nums[right+1] == target {
		right++
	}
	return []int{left, right}
}

// binarySearchAny 二分查找 target 的任意一个位置
func binarySearchAny(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
