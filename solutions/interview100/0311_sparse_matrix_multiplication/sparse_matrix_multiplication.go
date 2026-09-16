package sparsematrixmultiplication

// Multiply 稀疏矩阵的乘法
// 给定两个稀疏矩阵 mat1（m x k）和 mat2（k x n），返回它们相乘的结果矩阵（m x n）。
// 利用稀疏性：mat1[i][k] 为 0 时该项对结果没有贡献，直接跳过。
// 时间复杂度: O(m·k·n) 最坏情况，稀疏时约为 O(非零元素个数 × n)  空间复杂度: O(m·k)（非零元素预处理）
func Multiply(mat1 [][]int, mat2 [][]int) [][]int {
	m, k, n := len(mat1), len(mat1[0]), len(mat2[0])

	// 预处理 mat1：收集每一行的非零元素（列下标，值）
	type pair struct{ col, val int }
	nonZero := make([][]pair, m)
	for i := 0; i < m; i++ {
		for j := 0; j < k; j++ {
			if mat1[i][j] != 0 {
				nonZero[i] = append(nonZero[i], pair{col: j, val: mat1[i][j]})
			}
		}
	}

	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}

	// 按行累加：mat1 第 i 行的每个非零元素 (k, v) 将 mat2 的第 k 行乘以 v 累加到结果第 i 行
	for i := 0; i < m; i++ {
		for _, p := range nonZero[i] {
			for j := 0; j < n; j++ {
				if mat2[p.col][j] != 0 {
					result[i][j] += p.val * mat2[p.col][j]
				}
			}
		}
	}
	return result
}
