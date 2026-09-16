package allpathsfromsourceleadtodestination

import (
	"fmt"
	"testing"
)

func TestLeadsToDestination(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		edges       [][]int
		source      int
		destination int
		want        bool
	}{
		// LeetCode 官方示例
		{"示例1: 存在非终点的死胡同", 3, [][]int{{0, 1}, {0, 2}}, 0, 2, false},
		{"示例2: 存在环", 4, [][]int{{0, 1}, {0, 3}, {1, 2}, {2, 1}}, 0, 3, false},
		{"示例3: 菱形两条路径都到终点", 4, [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}, 0, 3, true},

		// 边界：起点即终点
		{"起点即终点且终点无出边", 1, [][]int{}, 0, 0, true},

		// 边界：无边
		{"无边且起点不是终点", 2, [][]int{}, 0, 1, false},
		{"无边且起点即终点", 3, [][]int{}, 1, 1, true},

		// 边界：终点有出边
		{"终点有出边指向死胡同", 3, [][]int{{0, 1}, {1, 2}}, 0, 1, false},
		{"终点在环上", 3, [][]int{{0, 1}, {1, 2}, {2, 1}}, 0, 1, false},

		// 边界：自环不在输入中，但链式单一路径
		{"单一路径直达终点", 4, [][]int{{0, 1}, {1, 2}, {2, 3}}, 0, 3, true},
		{"链式路径终点错误", 4, [][]int{{0, 1}, {1, 2}, {2, 3}}, 0, 2, false},

		// 边界：重复边
		{"重复边不影响结果", 3, [][]int{{0, 1}, {0, 1}, {1, 2}, {1, 2}}, 0, 2, true},

		// 典型场景：部分分支成环
		{"一条分支到终点一条分支成环", 5, [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 4}, {4, 2}}, 0, 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeadsToDestination(tt.n, tt.edges, tt.source, tt.destination); got != tt.want {
				t.Errorf("LeadsToDestination(n=%d, edges=%v, source=%d, destination=%d) = %v, want %v",
					tt.n, tt.edges, tt.source, tt.destination, got, tt.want)
			}
		})
	}
}

func BenchmarkLeadsToDestination(b *testing.B) {
	// 构造一个有 1000 个节点、约 2000 条边的菱形链：i -> i+1, i -> i+2
	n := 1000
	edges := make([][]int, 0, 2*n)
	for i := 0; i+1 < n; i++ {
		edges = append(edges, []int{i, i + 1})
	}
	for i := 0; i+2 < n; i++ {
		edges = append(edges, []int{i, i + 2})
	}

	b.Run(fmt.Sprintf("n=%d_e=%d", n, len(edges)), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			LeadsToDestination(n, edges, 0, n-1)
		}
	})
}
