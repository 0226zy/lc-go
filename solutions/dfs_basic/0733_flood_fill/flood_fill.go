package floodfill

// FloodFill 图像渲染
// 给定一幅二维整数数组表示的图像、起点坐标 (sr, sc) 和目标颜色 color，
// 从起点开始把所有与之四连通且颜色相同的格子都渲染成 color，返回渲染后的图像。
// 时间复杂度: O(m*n) 每个格子最多被访问一次  空间复杂度: O(m*n) 递归栈深度最坏情况
func FloodFill(image [][]int, sr int, sc int, color int) [][]int {
	m, n := len(image), len(image[0])
	oldColor := image[sr][sc]
	// 新颜色与原颜色相同，无需渲染，直接返回
	if oldColor == color {
		return image
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		// 越界或颜色不等于原颜色，直接返回
		if i < 0 || i >= m || j < 0 || j >= n || image[i][j] != oldColor {
			return
		}
		// 染色即标记，避免重复访问
		image[i][j] = color
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	dfs(sr, sc)
	return image
}
