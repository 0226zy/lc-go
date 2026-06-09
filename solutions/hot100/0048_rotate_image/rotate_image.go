package rotateimage

// Rotate 旋转图像（转置 + 反转每行）
// 先沿主对角线转置矩阵，再反转每一行，等价于顺时针旋转 90 度。
// 时间复杂度: O(n²)  空间复杂度: O(1)
func Rotate(matrix [][]int) {
	n := len(matrix)
	// 沿主对角线转置
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 反转每一行
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}

// RotateLayer 旋转图像（逐层旋转）
// 从外到内逐层旋转，每次将四个对应位置的元素轮换。
// 时间复杂度: O(n²)  空间复杂度: O(1)
func RotateLayer(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			// (i,j) → (j,n-1-i) → (n-1-i,n-1-j) → (n-1-j,i) → (i,j)
			matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i] =
				matrix[n-1-j][i], matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j]
		}
	}
}
