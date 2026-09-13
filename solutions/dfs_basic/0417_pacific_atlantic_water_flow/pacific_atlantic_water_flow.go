package pacificatlantic

// PacificAtlantic 太平洋大西洋水流问题
// 给定一个 m x n 的高度矩阵 heights，雨水只能从高（或等高）处流向低（或等高）处，
// 矩阵的上边和左边邻接太平洋，下边和右边邻接大西洋。
// 找出所有雨水既能流到太平洋又能流到大西洋的单元格坐标。
// 时间复杂度: O(m*n) 每个格子在每个方向最多访问一次  空间复杂度: O(m*n) 两个可达标记矩阵与递归栈
func PacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return nil
	}
	m, n := len(heights), len(heights[0])

	// pacific[i][j] / atlantic[i][j] 表示 (i,j) 的水是否能到达太平洋 / 大西洋
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := range pacific {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// dfs 逆向搜索：从海岸线出发，向海拔更高（或等高）的内陆格子扩散并打标记，
	// 已标记的格子不再重复访问，因此无需在回溯时还原标记
	var dfs func(i, j int, reachable [][]bool)
	dfs = func(i, j int, reachable [][]bool) {
		reachable[i][j] = true
		for _, d := range dirs {
			ni, nj := i+d[0], j+d[1]
			if ni < 0 || ni >= m || nj < 0 || nj >= n || reachable[ni][nj] {
				continue
			}
			if heights[ni][nj] >= heights[i][j] { // 逆向：只能爬向不低的位置
				dfs(ni, nj, reachable)
			}
		}
	}

	// 上边与左边邻接太平洋，下边与右边邻接大西洋
	for i := 0; i < m; i++ {
		dfs(i, 0, pacific)
		dfs(i, n-1, atlantic)
	}
	for j := 0; j < n; j++ {
		dfs(0, j, pacific)
		dfs(m-1, j, atlantic)
	}

	// 两张可达表的交集即为答案
	var result [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}
