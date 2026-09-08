package numberofislands

import "github.com/0226zy/lc-go/pkg/datastructures"

// NumIslands 岛屿数量
// 给你一个由 '1'（陆地）和 '0'（水）组成的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// 时间复杂度: O(m*n) 每个格子最多被访问一次  空间复杂度: O(m*n) 递归栈深度最坏情况
func NumIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	count := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		// 越界或遇到水，直接返回
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
			return
		}
		// 淹掉当前陆地，避免重复访问
		grid[i][j] = '0'
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j) // 把整座岛屿淹掉
			}
		}
	}
	return count
}

// NumIslandsUF 岛屿数量（并查集解法）
// 把所有相邻的陆地合并到同一个集合，最终集合个数即为岛屿数量。
// 时间复杂度: O(m*n*α(m*n))  α 为阿克曼函数的反函数，近似常数  空间复杂度: O(m*n)
func NumIslandsUF(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	uf := datastructures.NewUnionFind(m * n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			// 只向右、向下合并，避免重复
			if i+1 < m && grid[i+1][j] == '1' {
				uf.Union(i*n+j, (i+1)*n+j)
			}
			if j+1 < n && grid[i][j+1] == '1' {
				uf.Union(i*n+j, i*n+j+1)
			}
		}
	}
	// 统计根节点属于自己的陆地格子数
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && uf.Find(i*n+j) == i*n+j {
				count++
			}
		}
	}
	return count
}
