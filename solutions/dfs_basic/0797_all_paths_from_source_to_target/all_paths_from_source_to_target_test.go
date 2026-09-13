package allpaths

import (
	"reflect"
	"sort"
	"testing"
)

// normalize 对二维切片排序，用于忽略路径返回顺序的差异
func normalize(paths [][]int) [][]int {
	out := make([][]int, len(paths))
	for i, p := range paths {
		c := make([]int, len(p))
		copy(c, p)
		out[i] = c
	}
	sort.Slice(out, func(i, j int) bool {
		for k := 0; k < len(out[i]) && k < len(out[j]); k++ {
			if out[i][k] != out[j][k] {
				return out[i][k] < out[j][k]
			}
		}
		return len(out[i]) < len(out[j])
	})
	return out
}

func TestAllPathsSourceTarget(t *testing.T) {
	tests := []struct {
		name  string
		graph [][]int
		want  [][]int
	}{
		// LeetCode 官方示例
		{"示例1: graph=[[1,2],[3],[3],[]]",
			[][]int{{1, 2}, {3}, {3}, {}},
			[][]int{{0, 1, 3}, {0, 2, 3}}},
		{"示例2: graph=[[4,3,1],[3,2,4],[3],[4],[]]",
			[][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}},
			[][]int{
				{0, 4},
				{0, 3, 4},
				{0, 1, 3, 4},
				{0, 1, 2, 3, 4},
				{0, 1, 4},
			}},

		// 边界：只有两个节点，直接相连
		{"两节点直连", [][]int{{1}, {}}, [][]int{{0, 1}}},
		// 边界：链式图，路径唯一
		{"链式图唯一路径", [][]int{{1}, {2}, {}}, [][]int{{0, 1, 2}}},
		// 边界：多路汇合
		{"多路汇合", [][]int{{1, 2}, {3}, {3}, {4}, {}},
			[][]int{{0, 1, 3, 4}, {0, 2, 3, 4}}},
		// 边界：完全二分式的菱形图
		{"菱形图", [][]int{{1, 2}, {3}, {3}, {}},
			[][]int{{0, 1, 3}, {0, 2, 3}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllPathsSourceTarget(tt.graph)
			got, want := normalize(got), normalize(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("AllPathsSourceTarget(%v) = %v, want %v", tt.graph, got, want)
			}
		})
	}
}

func BenchmarkAllPathsSourceTarget(b *testing.B) {
	// 构造一个有 15 个节点的分层 DAG：相邻两层之间全连接，路径数指数增长
	layers := [][]int{{0}, {1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}, {11, 12}, {13}, {14}}
	graph := make([][]int, 15)
	for li := 0; li < len(layers)-1; li++ {
		for _, u := range layers[li] {
			graph[u] = append(graph[u], layers[li+1]...)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AllPathsSourceTarget(graph)
	}
}
