package minimumsizesubarraysum

// MinSubArrayLen 长度最小的子数组
// 给定一个含有 n 个正整数的数组和一个正整数 target，找出该数组中满足其总和大于等于 target
// 的长度最小的连续子数组 [nums[l], nums[l+1], ..., nums[r-1], nums[r]]，并返回其长度。
// 如果不存在符合条件的子数组，返回 0。
// 时间复杂度: O(n) 每个元素最多进出窗口一次  空间复杂度: O(1)
func MinSubArrayLen(target int, nums []int) int {
	n := len(nums)
	left, sum := 0, 0
	ans := n + 1 // 用 n+1 表示“未找到”
	for right := 0; right < n; right++ {
		sum += nums[right]
		// 窗口内和已经达标，尝试收缩左边界，取更短的窗口
		for sum >= target {
			if right-left+1 < ans {
				ans = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	if ans == n+1 {
		return 0
	}
	return ans
}
