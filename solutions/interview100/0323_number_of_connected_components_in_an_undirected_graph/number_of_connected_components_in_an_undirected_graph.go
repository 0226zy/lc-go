package numberofconnectedcomponentsinanundirectedgraph

import "github.com/0226zy/lc-go/pkg/datastructures"

// CountComponents 无向图中连通分量的数目
// 给定 n 个节点（编号 0 到 n-1）和无向边列表 edges，求连通分量的数目。
// 使用并查集：初始 n 个分量，每成功合并一次分量数减一。
// 时间复杂度: O(n + e·α(n)) α 为反阿克曼函数  空间复杂度: O(n)
func CountComponents(n int, edges [][]int) int {
	uf := datastructures.NewUnionFind(n)
	for _, e := range edges {
		// Union 返回 true 表示两个节点原本不连通，合并后分量数减一
		uf.Union(e[0], e[1])
	}
	return uf.Count
}
