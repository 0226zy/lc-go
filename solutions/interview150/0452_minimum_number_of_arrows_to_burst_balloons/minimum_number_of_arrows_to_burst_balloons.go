package minimumnumberofarrowstoburstballoons

import "sort"

// FindMinArrowShots 用最少数量的箭引爆气球
// points[i] = [x_start, x_end] 表示气球的水平直径范围，
// 一支垂直射出的箭在坐标 x 处可以引爆所有满足 x_start <= x <= x_end 的气球。
// 返回引爆所有气球必须射出的最小弓箭数。
// 时间复杂度: O(n log n) 排序  空间复杂度: O(log n) 排序递归栈
func FindMinArrowShots(points [][]int) int {
	// 按右端点升序排序，贪心地让箭射在当前重叠区间的最右端点
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	ans := 1            // 至少射一支箭
	pos := points[0][1] // 第一支箭射在右端点最小的气球的右端点
	for i := 1; i < len(points); i++ {
		if points[i][0] > pos {
			// 左端点在箭的位置右侧，当前箭够不到，补射一支新箭
			ans++
			pos = points[i][1]
		}
	}
	return ans
}
