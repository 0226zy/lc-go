package findthecelebrity

import "testing"

// makeKnows 根据邻接矩阵构造 knows 接口的 mock 实现
// graph[i][j] == 1 表示 i 认识 j
func makeKnows(graph [][]int) KnowsFunc {
	return func(a, b int) bool {
		return graph[a][b] == 1
	}
}

func TestFindCelebrity(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		graph [][]int
		want  int
	}{
		// LeetCode 官方示例
		{
			"示例1: 1号是名人",
			3,
			[][]int{{1, 1, 0}, {0, 1, 0}, {1, 1, 1}},
			1,
		},
		{
			"示例2: 没有名人",
			3,
			[][]int{{1, 0, 1}, {1, 1, 0}, {0, 1, 1}},
			-1,
		},

		// 边界：两人聚会
		{
			"两人互相认识无名人",
			2,
			[][]int{{1, 1}, {1, 1}},
			-1,
		},
		{
			"两人互不认识无名人",
			2,
			[][]int{{1, 0}, {0, 1}},
			-1,
		},
		{
			"两人中1号是名人",
			2,
			[][]int{{1, 1}, {0, 1}},
			1,
		},

		// 典型场景
		{
			"名人是0号",
			4,
			[][]int{{1, 0, 0, 0}, {1, 1, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 1}},
			0,
		},
		{
			"名人是最后一位",
			4,
			[][]int{{1, 1, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 1}, {0, 0, 0, 1}},
			3,
		},
		{
			"候选人认识他人验证失败",
			3,
			[][]int{{1, 0, 1}, {0, 1, 0}, {0, 1, 1}},
			-1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			knows := makeKnows(tt.graph)
			if got := FindCelebrity(tt.n, knows); got != tt.want {
				t.Errorf("FindCelebrity(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkFindCelebrity(b *testing.B) {
	// 构造 100 人的聚会（题目数据上限），名人编号为 99
	n := 100
	graph := make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, n)
		graph[i][i] = 1
		graph[i][n-1] = 1
	}
	graph[n-1][n-1] = 1
	knows := makeKnows(graph)

	b.Run("百人聚会", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FindCelebrity(n, knows)
		}
	})
}
