package removeinterval

// RemoveInterval 删除区间
// 给定有序的不相交区间列表 intervals 和待删除区间 toBeRemoved，
// 删除所有与 toBeRemoved 有交集的部分，返回剩余区间的有序列表。
// 时间复杂度: O(n)  空间复杂度: O(1)（不计返回结果）
func RemoveInterval(intervals [][]int, toBeRemoved []int) [][]int {
	res := make([][]int, 0)
	for _, iv := range intervals {
		a, b := iv[0], iv[1]
		// 与待删除区间不相交，整个保留
		if b <= toBeRemoved[0] || a >= toBeRemoved[1] {
			res = append(res, []int{a, b})
			continue
		}
		// 有交集：左半段非空则保留
		if a < toBeRemoved[0] {
			res = append(res, []int{a, toBeRemoved[0]})
		}
		// 右半段非空则保留
		if b > toBeRemoved[1] {
			res = append(res, []int{toBeRemoved[1], b})
		}
	}
	return res
}
