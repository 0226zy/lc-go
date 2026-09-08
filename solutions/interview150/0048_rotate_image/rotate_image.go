package rotateimage

// Rotate 旋转图像
// 将 n x n 矩阵顺时针旋转 90 度，必须原地修改。
// 两步法：先沿主对角线转置，再对每一行做水平翻转。
// 时间复杂度: O(n^2)  空间复杂度: O(1) 原地修改
func Rotate(matrix [][]int) {
	n := len(matrix)

	// 第一步：转置，交换主对角线两侧的元素 matrix[i][j] 与 matrix[j][i]
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 第二步：将每一行左右翻转（相当于垂直镜像）
	for i := 0; i < n; i++ {
		for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
			matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
		}
	}
}

// RotateCopy 旋转图像（辅助数组解法，供对比）
// 新位置 (j, n-1-i) 存放旧位置 (i, j) 的元素。
// 时间复杂度: O(n^2)  空间复杂度: O(n^2) 需要一个同样大小的辅助矩阵
func RotateCopy(matrix [][]int) {
	n := len(matrix)
	copyMatrix := make([][]int, n)
	for i := 0; i < n; i++ {
		copyMatrix[i] = make([]int, n)
		copy(copyMatrix[i], matrix[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[j][n-1-i] = copyMatrix[i][j]
		}
	}
}
