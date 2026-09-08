package summaryranges

import "strconv"

// SummaryRanges 汇总区间
// 给定一个无重复元素的有序整数数组 nums，返回恰好覆盖所有数字的最小有序区间列表。
// 每个区间按 "a->b"（a != b）或 "a"（a == b）的格式输出。
// 时间复杂度: O(n) 一次遍历  空间复杂度: O(1) 不计输出
func SummaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	res := []string{}
	start := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 { // 区间在 nums[i-1] 处断开
			res = append(res, formatRange(start, nums[i-1]))
			start = nums[i]
		}
	}
	res = append(res, formatRange(start, nums[len(nums)-1])) // 处理最后一个区间
	return res
}

// formatRange 按 "a->b" 或 "a" 的格式输出区间
func formatRange(a, b int) string {
	if a == b {
		return strconv.Itoa(a)
	}
	return strconv.Itoa(a) + "->" + strconv.Itoa(b)
}
