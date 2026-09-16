package numberofconnectedcomponentsinanundirectedgraph

import "testing"

func TestCountComponents(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  int
	}{
		// LeetCode 官方示例
		{
			"示例1: 两个分量",
			5,
			[][]int{{0, 1}, {1, 2}, {3, 4}},
			2,
		},
		{
			"示例2: 全部连通成一条链",
			5,
			[][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			1,
		},

		// 边界：没有边，每个节点自成一个分量
		{
			"无边图全是孤立节点",
			4,
			[][]int{},
			4,
		},

		// 边界：单节点
		{
			"单节点无边",
			1,
			[][]int{},
			1,
		},

		// 典型场景：边形成环，分量数不变
		{
			"含环的图",
			4,
			[][]int{{0, 1}, {1, 2}, {2, 0}, {2, 3}},
			1,
		},

		// 典型场景：完全图
		{
			"完全图",
			4,
			[][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}},
			1,
		},

		// 典型场景：多对孤立组合
		{
			"三对节点各自成对",
			6,
			[][]int{{0, 1}, {2, 3}, {4, 5}},
			3,
		},

		// 典型场景：部分连通部分孤立
		{
			"一条链加孤立节点",
			6,
			[][]int{{0, 1}, {1, 2}, {2, 3}},
			3, // {0,1,2,3}、{4}、{5}
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountComponents(tt.n, tt.edges); got != tt.want {
				t.Errorf("CountComponents(%d, %v) = %d, want %d", tt.n, tt.edges, got, tt.want)
			}
		})
	}
}

func BenchmarkCountComponents(b *testing.B) {
	// 构造 2000 个节点、5000 条随机边
	n := 2000
	edges := make([][]int, 0, 5000)
	for i := 0; i < 5000; i++ {
		edges = append(edges, []int{(i * 7) % n, (i*13 + 1) % n})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountComponents(n, edges)
	}
}
