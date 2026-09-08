package searcha2dmatrix

// SearchMatrix 搜索二维矩阵
// 矩阵每行非递减、且每行首元素大于上一行尾元素，等价于一个全局有序的一维数组。
// 对一维下标 mid 做二分，用 row=mid/n, col=mid%n 映射回二维坐标。
// 时间复杂度: O(log(m*n)) 二分  空间复杂度: O(1) 常数变量
func SearchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := (left + right) / 2
		x := matrix[mid/n][mid%n] // 一维下标映射回二维坐标
		if x == target {
			return true
		} else if x < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
