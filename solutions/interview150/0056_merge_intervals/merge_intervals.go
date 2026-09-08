package mergeintervals

import "sort"

// Merge 合并区间
// 以数组 intervals 表示若干个区间的集合，合并所有重叠的区间，
// 返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
// 时间复杂度: O(n log n) 排序  空间复杂度: O(log n) 排序递归栈
func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	// 按左端点升序排序，保证可能重叠的区间连续排列
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		if intervals[i][0] <= last[1] {
			// 有重叠（端点相接也算），合并右端点取最大值
			if intervals[i][1] > last[1] {
				last[1] = intervals[i][1]
			}
		} else {
			// 无重叠，追加为新的基准区间
			res = append(res, intervals[i])
		}
	}
	return res
}
