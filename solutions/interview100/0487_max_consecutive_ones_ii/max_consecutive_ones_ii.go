package maxconsecutiveonesii

// FindMaxConsecutiveOnes 最大连续1的个数 II
// 给定一个二进制数组 nums，最多可以把一个 0 翻转成 1，返回连续 1 的最大个数。
// 等价于求「最多含一个 0」的最长子数组长度，使用滑动窗口。
// 时间复杂度: O(n) 左右指针均单调右移  空间复杂度: O(1)
func FindMaxConsecutiveOnes(nums []int) int {
	left, zeros, ans := 0, 0, 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeros++
		}
		// 窗口内 0 超过一个时收缩左边界，直到重新合法
		for zeros > 1 {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}
