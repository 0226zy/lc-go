package findsmallestcommonelementinallrows

// SmallestCommonElement 找出所有行中最小公共元素
// 矩阵 mat 的每一行严格递增，返回所有行都包含的最小元素；不存在则返回 -1。
// 思路：统计每个元素出现的行数（行内无重复，直接计数），
// 从小到大找第一个出现次数等于行数的元素。
// 时间复杂度: O(m·n + V)，V = 10^4 为元素值域  空间复杂度: O(V)
func SmallestCommonElement(mat [][]int) int {
	const maxVal = 10000 // 元素值域 1..10^4
	m := len(mat)
	cnt := make([]int, maxVal+1)
	// 每行严格递增，行内无重复，直接计数即可
	for _, row := range mat {
		for _, v := range row {
			cnt[v]++
		}
	}
	// 从小到大找第一个出现在所有行中的元素
	for v := 1; v <= maxVal; v++ {
		if cnt[v] == m {
			return v
		}
	}
	return -1
}
