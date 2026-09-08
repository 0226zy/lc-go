package insertinterval

// Insert 插入区间
// 给一个无重叠、按起始端点升序排列的区间列表 intervals，
// 在其中插入新区间 newInterval，使结果仍有序且不重叠（必要时合并）。
// 时间复杂度: O(n)  空间复杂度: O(n) 输出数组
func Insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	i, n := 0, len(intervals)

	// 左段：右端点小于新区间左端点，不可能重叠，直接放入
	for i < n && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}
	// 中段：与新区间有重叠（端点相接也算），逐个合并进 newInterval
	for i < n && intervals[i][0] <= newInterval[1] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}
	res = append(res, newInterval) // 合并后的新区间
	// 右段：剩余区间不可能重叠，直接放入
	for i < n {
		res = append(res, intervals[i])
		i++
	}
	return res
}
