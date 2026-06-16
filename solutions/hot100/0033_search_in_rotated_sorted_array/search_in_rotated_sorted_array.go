package searchinrotatedsortedarray

// SearchInRotatedSortedArray 搜索旋转排序数组
// 通过一次二分查找直接定位 target，利用旋转数组的半区有序性判断搜索方向。
// 时间复杂度: O(log n)  空间复杂度: O(1)
func SearchInRotatedSortedArray(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		// 判断哪一半是有序的
		if nums[left] <= nums[mid] {
			// 左半部分 [left, mid] 有序
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 右半部分 [mid, right] 有序
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

// SearchInRotatedSortedArrayFindPivot 先找旋转点再二分
// 先通过二分查找找到数组的最小值（旋转点），再根据 target 与 nums[0] 的大小关系决定在哪一半进行普通二分查找。
// 时间复杂度: O(log n)  空间复杂度: O(1)
func SearchInRotatedSortedArrayFindPivot(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	// 找最小值下标，即旋转点
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	pivot := left

	// 判断 target 落在哪一半有序区间
	var l, r int
	if target >= nums[0] && pivot > 0 {
		l, r = 0, pivot-1
	} else {
		l, r = pivot, n-1
	}

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}
