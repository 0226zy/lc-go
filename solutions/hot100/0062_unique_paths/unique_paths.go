package uniquepaths

// UniquePaths 不同路径
// 机器人从 m x n 网格左上角到右下角，每次只能向右或向下移动，求不同路径数。
// 时间复杂度: O(m*n)  空间复杂度: O(n)
func UniquePaths(m int, n int) int {
	f := make([][]int, m)

	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				f[i][j] = 1
				continue
			}
			f[i][j] = f[i-1][j] + f[i][j-1]
		}
	}
	return f[m-1][n-1]
}
