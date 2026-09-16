package countingelements

// CountElements 数元素
// 统计数组中满足「x + 1 也在数组中」的元素 x 的数量，重复元素单独计数。
// 时间复杂度: O(n)  空间复杂度: O(n)
func CountElements(arr []int) int {
	// 哈希集合记录数组中出现过哪些值
	set := make(map[int]bool, len(arr))
	for _, v := range arr {
		set[v] = true
	}
	count := 0
	for _, v := range arr {
		// 若 v+1 存在，则 v 满足条件（重复元素逐个计数）
		if set[v+1] {
			count++
		}
	}
	return count
}
