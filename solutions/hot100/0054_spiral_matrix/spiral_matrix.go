package spiralmatrix

// SpiralOrder 螺旋矩阵
// 给你一个 m 行 n 列的矩阵 matrix，请按照顺时针螺旋顺序，返回矩阵中的所有元素。
// 时间复杂度: O(mn) 遍历所有元素  空间复杂度: O(1) 常数空间（除输出数组外）
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)

	top, bottom := 0, m-1
	left, right := 0, n-1

	for top <= bottom && left <= right {
		// 从左到右遍历上边界
		for j := left; j <= right; j++ {
			result = append(result, matrix[top][j])
		}
		top++

		// 从上到下遍历右边界
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// 如果还有行，从右到左遍历下边界
		if top <= bottom {
			for j := right; j >= left; j-- {
				result = append(result, matrix[bottom][j])
			}
			bottom--
		}

		// 如果还有列，从下到上遍历左边界
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}
