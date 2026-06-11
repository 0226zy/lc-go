package searcha2dmatrixii

// SearchMatrix 搜索二维矩阵 II
// 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的目标值 target。
// 该矩阵具有以下特性：
//   - 每行的元素从左到右升序排列。
//   - 每列的元素从上到下升序排列。
//
// 时间复杂度: O(?)  空间复杂度: O(?)
func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	row, col := 0, n-1
	for row < m && col >= 0 {
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}
