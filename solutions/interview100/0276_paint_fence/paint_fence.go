package paintfence

// NumWays 栅栏涂色
// 有 n 根木柱的栅栏和 k 种颜色，要求最多只有两根相邻木柱颜色相同，
// 返回有效的涂色方案总数。
// 思路：按最后两根木柱是否同色划分状态，滚动递推。
// 时间复杂度: O(n)  空间复杂度: O(1)
func NumWays(n int, k int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return k
	}
	// i = 2 时的状态：same 为两根同色的方案数，diff 为两根异色的方案数
	same, diff := k, k*(k-1)
	for i := 3; i <= n; i++ {
		// 第 i 根与前一根同色：前一对必须异色
		// 第 i 根与前一根异色：之前任意合法方案，颜色有 k-1 种选择
		same, diff = diff, (same+diff)*(k-1)
	}
	return same + diff
}
