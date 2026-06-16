package searcha2dmatrix

// SearchMatrix 搜索二维矩阵
// 将 m x n 的矩阵视为一个长度为 m*n 的有序数组，进行二分查找。
// 时间复杂度: O(log(m * n))  空间复杂度: O(1)
func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1

	for left <= right {
		mid := left + (right-left)/2
		row, col := mid/n, mid%n
		val := matrix[row][col]

		if val == target {
			return true
		} else if val < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

// SearchMatrixRowCol 搜索二维矩阵（按行再按列二分）
// 先二分查找目标值可能所在的行，再在该行内二分查找目标值。
// 时间复杂度: O(log m + log n) = O(log(m * n))  空间复杂度: O(1)
func SearchMatrixRowCol(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])

	// 第一步：二分定位行
	rowLeft, rowRight := 0, m-1
	for rowLeft <= rowRight {
		rowMid := rowLeft + (rowRight-rowLeft)/2
		if matrix[rowMid][0] == target {
			return true
		}
		if matrix[rowMid][n-1] == target {
			return true
		}
		if matrix[rowMid][0] < target && target < matrix[rowMid][n-1] {
			// 目标值只可能在该行中
			colLeft, colRight := 0, n-1
			for colLeft <= colRight {
				colMid := colLeft + (colRight-colLeft)/2
				if matrix[rowMid][colMid] == target {
					return true
				} else if matrix[rowMid][colMid] < target {
					colLeft = colMid + 1
				} else {
					colRight = colMid - 1
				}
			}
			return false
		} else if target < matrix[rowMid][0] {
			rowRight = rowMid - 1
		} else {
			rowLeft = rowMid + 1
		}
	}

	return false
}
