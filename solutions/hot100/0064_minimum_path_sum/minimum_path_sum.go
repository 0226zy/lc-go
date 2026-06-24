package minimumpathsum

// MinPathSum 最小路径和
// 从网格左上角到右下角，每次只能向右或向下，求路径上数字和的最小值。
// 时间复杂度: O(m*n)  空间复杂度: O(n)
func MinPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	prev := 0
	for i := 0; i < n; i++ {
		grid[0][i] += prev
		prev = grid[0][i]
	}
	prev = 0
	for i := 0; i < m; i++ {
		grid[i][0] += prev
		prev = grid[i][0]
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				continue
			}
			up := grid[i-1][j]
			left := grid[i][j-1]
			grid[i][j] += min(up, left)
		}
	}
	return grid[m-1][n-1]
}
