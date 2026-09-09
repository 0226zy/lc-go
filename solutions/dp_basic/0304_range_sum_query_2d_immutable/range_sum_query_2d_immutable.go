package rangesumquery2dimmutable

// NumMatrix 二维区域和检索（二维前缀和）
// pre[i][j] 表示以 (0,0) 为左上角、(i-1,j-1) 为右下角的矩形元素和，
// pre[i][j] = pre[i-1][j] + pre[i][j-1] - pre[i-1][j-1] + matrix[i-1][j-1]。
type NumMatrix struct {
	pre [][]int // 前缀和表，含哨兵行/列
}

// Constructor 初始化并预处理二维前缀和表
// 时间复杂度: O(m*n)  空间复杂度: O(m*n)
func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	pre := make([][]int, m+1)
	for i := range pre {
		pre[i] = make([]int, n+1) // 第 0 行、第 0 列为哨兵，默认 0
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 容斥：上方矩形 + 左方矩形 - 左上重叠 + 当前格
			pre[i][j] = pre[i-1][j] + pre[i][j-1] - pre[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{pre: pre}
}

// SumRegion 返回 (row1,col1) 到 (row2,col2) 的区域和
// 容斥：大矩形 - 上方多余 - 左方多余 + 左上角被减两次的部分
// 时间复杂度: O(1)  空间复杂度: O(1)
func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return nm.pre[row2+1][col2+1] - nm.pre[row1][col2+1] - nm.pre[row2+1][col1] + nm.pre[row1][col1]
}
