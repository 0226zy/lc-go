package rotatearray

// Rotate 旋转数组
// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
// 时间复杂度: O(n) 三次翻转  空间复杂度: O(1) 常数空间
func Rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	// 翻转整个数组
	reverse(nums, 0, n-1)
	// 翻转前 k 个元素
	reverse(nums, 0, k-1)
	// 翻转后 n-k 个元素
	reverse(nums, k, n-1)
}

// reverse 翻转数组 nums 的 [left, right] 区间
func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
