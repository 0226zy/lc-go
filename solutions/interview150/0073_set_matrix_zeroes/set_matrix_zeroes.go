package setmatrixzeroes

// SetZeroes 矩阵置零
// 若矩阵中某元素为 0，则将其所在行和所在列的所有元素都置为 0，原地修改。
// 用矩阵的首行和首列充当标记数组记录哪些行列需要置零；
// 首行、首列自身是否需要置零用两个布尔变量单独记录，最后统一处理。
// 时间复杂度: O(m*n) 两次遍历  空间复杂度: O(1)
func SetZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	// 首行、首列自身是否原本就含 0
	row0HasZero, col0HasZero := false, false
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			row0HasZero = true
			break
		}
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			col0HasZero = true
			break
		}
	}

	// 用首行首列做标记：matrix[i][0]=0 表示第 i 行要置零，matrix[0][j]=0 表示第 j 列要置零
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 根据标记将内部元素置零
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 最后处理首行和首列，避免标记阶段被提前污染
	if row0HasZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if col0HasZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
