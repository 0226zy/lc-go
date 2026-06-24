package findtheduplicatenumber

// FindDuplicate 寻找重复数
// 给定 n+1 个值在 [1,n] 的整数，找出唯一的重复数。
// 时间复杂度: O(n)  空间复杂度: O(1)
func FindDuplicate(nums []int) int {
	// 快慢指针：快指针每次走两步，慢指针每次走一步
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	// 慢指针重回起点，两者同速前进，相遇处即环入口（重复数）
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

// FindDuplicateBinary 寻找重复数（二分查找）
// 在值域 [1,n] 上二分，统计 <= mid 的个数判断重复数所在区间。
// 时间复杂度: O(n log n)  空间复杂度: O(1)
func FindDuplicateBinary(nums []int) int {
	left, right := 1, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		count := 0
		for _, v := range nums {
			if v <= mid {
				count++
			}
		}
		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
