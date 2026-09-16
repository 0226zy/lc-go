package longestsubarray

// LongestSubarray 删掉一个元素以后全为 1 的最长子数组
// 给定一个二进制数组 nums，必须从中删掉恰好一个元素，
// 返回删掉后只包含 1 的最长非空子数组的长度；若不存在这样的子数组，返回 0。
// 滑动窗口：窗口内最多允许 1 个 0，超过则收缩左边界；因必须删一个元素，答案为窗口长度 - 1。
// 时间复杂度: O(n) 左右指针均单调右移，每个元素最多进出窗口一次  空间复杂度: O(1) 仅常数个变量
func LongestSubarray(nums []int) int {
	left, zeroCount, ans := 0, 0, 0
	for right := 0; right < len(nums); right++ {
		// 右端元素进窗口，统计窗口内 0 的个数
		if nums[right] == 0 {
			zeroCount++
		}
		// 窗口内 0 超过 1 个，收缩左边界直到合法
		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		// 必须删掉恰好一个元素（窗口内唯一的 0，或全 1 时任删一个 1），贡献为窗口长度 - 1
		if right-left > ans {
			ans = right - left
		}
	}
	return ans
}
