package pascalstriangleii

// GetRow 杨辉三角 II（标准 DP 数组版）
// 返回杨辉三角的第 rowIndex 行（从 0 开始），每个数等于上一行左上方与右上方两数之和。
// 逐行生成：next[j] = row[j-1] + row[j]，每行首尾固定为 1。
// 时间复杂度: O(rowIndex²)  空间复杂度: O(rowIndex)
func GetRow(rowIndex int) []int {
	row := []int{1} // base case：第 0 行为 [1]
	for i := 1; i <= rowIndex; i++ {
		next := make([]int, i+1)
		next[0], next[i] = 1, 1 // 每行首尾固定为 1
		for j := 1; j < i; j++ {
			next[j] = row[j-1] + row[j] // 左上方 + 右上方
		}
		row = next
	}
	return row
}

// GetRowOptimized 杨辉三角 II（单数组滚动优化版）
// 只用一个数组原地滚动：每行倒序执行 row[j] += row[j-1]，
// 倒序保证用到的 row[j-1] 仍是上一行的旧值；每行末尾补 1。
// 时间复杂度: O(rowIndex²)  空间复杂度: O(rowIndex)（无重复分配）
func GetRowOptimized(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1 // base case：第 0 行为 [1]
	for i := 1; i <= rowIndex; i++ {
		row[i] = 1 // 新行的行尾补 1
		for j := i - 1; j >= 1; j-- {
			row[j] += row[j-1] // 倒序更新，避免覆盖旧值
		}
	}
	return row
}
