package mergeintervals

import "sort"

// Merge 合并区间
// 排序 + 贪心合并
// 时间复杂度: O(n log n) 排序占主导  空间复杂度: O(log n) 排序额外空间
func Merge(intervals [][]int) [][]int {
	// 按起点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}
	for _, interval := range intervals[1:] {
		last := merged[len(merged)-1]
		if interval[0] <= last[1] {
			// 重叠，合并
			if interval[1] > last[1] {
				last[1] = interval[1]
			}
		} else {
			// 不重叠，加入
			merged = append(merged, interval)
		}
	}
	return merged
}
