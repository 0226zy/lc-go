package numberofprovinces

import "github.com/0226zy/lc-go/pkg/datastructures"

// FindCircleNum 省份数量
// 有 n 个城市，isConnected 是 n x n 的邻接矩阵，isConnected[i][j] = 1 表示城市 i 与城市 j 直接相连。
// 省份是一组直接或间接相连的城市，求省份的数量（即图的连通分量个数）。
// 时间复杂度: O(n^2) 每个城市最多访问一次，每次扫描整行邻接矩阵  空间复杂度: O(n) visited 数组与递归栈深度
func FindCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	count := 0

	var dfs func(city int)
	dfs = func(city int) {
		// 标记当前城市已访问
		visited[city] = true
		// 扫描整行邻接矩阵，访问所有直接相连且未访问过的城市
		for next := 0; next < n; next++ {
			if isConnected[city][next] == 1 && !visited[next] {
				dfs(next)
			}
		}
	}

	for city := 0; city < n; city++ {
		if !visited[city] {
			count++   // 发现一个未访问的城市，说明找到了一个新省份
			dfs(city) // 把这个省份的所有城市都标记掉
		}
	}
	return count
}

// FindCircleNumUF 省份数量（并查集解法）
// 把所有直接相连的城市合并到同一个集合，最终统计根节点个数即为省份数量。
// 时间复杂度: O(n^2*α(n))  α 为阿克曼函数的反函数，近似常数  空间复杂度: O(n)
func FindCircleNumUF(isConnected [][]int) int {
	n := len(isConnected)
	uf := datastructures.NewUnionFind(n)
	// 只遍历上三角，避免重复合并
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}
	// UnionFind 内部维护了连通分量数 Count，初始为 n，每成功合并一次减一
	return uf.Count
}
