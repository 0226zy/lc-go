package findminimuminrotatedsortedarray

// FindMin 寻找旋转排序数组中的最小值
// 给定一个原本升序、元素互不相同、经若干次旋转后的数组 nums，
// 用 O(log n) 时间复杂度找出其中的最小元素。
// 核心思想：比较 nums[mid] 与 nums[right]——若 nums[mid] > nums[right]，
// 说明 mid 落在大数段，最小值在右侧；否则最小值在 mid 或左侧。
// 时间复杂度: O(log n)  空间复杂度: O(1)
func FindMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			// mid 在大数段，最小值一定在 mid 右边
			left = mid + 1
		} else {
			// mid 与右端点在同一递增段，最小值在 mid 或左边
			right = mid
		}
	}
	return nums[left]
}
