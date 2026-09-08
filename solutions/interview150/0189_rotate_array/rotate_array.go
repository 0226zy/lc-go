package rotatearray

// Rotate 轮转数组
// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置（k 为非负数）。
// 时间复杂度: O(n) 三次翻转共 2n 次交换  空间复杂度: O(1) 原地轮转
func Rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	k %= n // k 可能大于数组长度，先取模
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

// reverse 翻转 nums[left:right] 闭区间
func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

// RotateCopy 轮转数组（辅助数组对照解法）
// 用一个新数组按最终位置拷贝，再复制回原数组。思路直观但需要 O(n) 额外空间。
// 时间复杂度: O(n)  空间复杂度: O(n)
func RotateCopy(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	k %= n
	tmp := make([]int, n)
	copy(tmp, nums[n-k:])
	copy(tmp[k:], nums[:n-k])
	copy(nums, tmp)
}
