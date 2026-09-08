package spiralmatrix

// SpiralOrder 螺旋矩阵
// 从左上角开始，按顺时针螺旋顺序遍历 m x n 矩阵，返回展开后的切片。
// 通过不断收缩上下左右四条边界来实现“一圈一圈向内”的遍历。
// 时间复杂度: O(m*n) 每个元素访问一次  空间复杂度: O(1) 不计返回值
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)

	top, bottom := 0, m-1 // 上、下边界
	left, right := 0, n-1 // 左、右边界
	for top <= bottom && left <= right {
		// 向右：遍历上边一行
		for j := left; j <= right; j++ {
			result = append(result, matrix[top][j])
		}
		top++

		// 向下：遍历右边一列
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// 收缩后可能只剩一行或一列，需要防止重复遍历
		if top <= bottom {
			// 向左：遍历下边一行
			for j := right; j >= left; j-- {
				result = append(result, matrix[bottom][j])
			}
			bottom--
		}
		if left <= right {
			// 向上：遍历左边一列
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}
	return result
}
