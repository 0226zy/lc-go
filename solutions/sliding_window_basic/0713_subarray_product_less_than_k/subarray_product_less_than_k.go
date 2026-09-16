package subarrayproductlessthank

// NumSubarrayProductLessThanK 乘积小于 K 的子数组
// 给定一个正整数数组 nums 和一个整数 k，返回乘积严格小于 k 的连续子数组个数。
// 滑动窗口：维护窗口乘积 product，product >= k 时收缩左边界，
// 每次以 right 结尾的合法子数组个数为 right-left+1。
// 时间复杂度: O(n) 左右指针各自最多移动 n 步  空间复杂度: O(1)
func NumSubarrayProductLessThanK(nums []int, k int) int {
	// 所有元素为正整数，k <= 1 时不可能存在乘积严格小于 k 的子数组
	if k <= 1 {
		return 0
	}

	count, left, product := 0, 0, 1
	for right := 0; right < len(nums); right++ {
		// 右指针扩张窗口，乘入新元素
		product *= nums[right]
		// 乘积不小于 k，收缩左边界直到窗口恢复合法
		for product >= k {
			product /= nums[left]
			left++
		}
		// 以 right 结尾且乘积小于 k 的子数组共有 right-left+1 个
		count += right - left + 1
	}
	return count
}
