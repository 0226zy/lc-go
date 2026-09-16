package longestones

// LongestOnes 最大连续1的个数 III
// 给定一个二进制数组 nums 和一个整数 k，最多可以翻转 k 个 0，
// 返回数组中连续 1 的最长个数。
// 滑动窗口：维护窗口内 0 的个数不超过 k 的最长窗口。
// 时间复杂度: O(n)  空间复杂度: O(1)
func LongestOnes(nums []int, k int) int {
	left, zeros, ans := 0, 0, 0
	for right := 0; right < len(nums); right++ {
		// 右边界扩张：遇到 0 计入待翻转数量
		if nums[right] == 0 {
			zeros++
		}
		// 0 的数量超过 k，收缩左边界直到窗口重新合法
		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}
		// 当前窗口 [left, right] 内 0 不超过 k 个，更新最大长度
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}
