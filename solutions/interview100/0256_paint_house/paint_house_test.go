package painthouse

import "testing"

func TestMinCost(t *testing.T) {
	tests := []struct {
		name  string
		costs [][]int
		want  int
	}{
		// LeetCode 官方示例
		{"示例1: 三栋房子", [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}, 10},
		{"示例2: 单栋房子取最便宜颜色", [][]int{{7, 6, 2}}, 2},

		// 边界：空输入
		{"空输入", nil, 0},

		// 典型场景
		{"花费全相同", [][]int{{5, 5, 5}, {5, 5, 5}, {5, 5, 5}}, 15},
		{"每栋最便宜颜色恰好交错", [][]int{{1, 100, 100}, {100, 1, 100}, {100, 100, 1}}, 3},
		{"交叉贪便宜", [][]int{{1, 2, 3}, {3, 2, 1}, {1, 2, 3}}, 3},
		{"两栋房子", [][]int{{3, 5, 7}, {7, 5, 3}}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinCost(tt.costs); got != tt.want {
				t.Errorf("MinCost(%v) = %d, want %d", tt.costs, got, tt.want)
			}
		})
	}
}

func BenchmarkMinCost(b *testing.B) {
	// 构造 100 栋房子的花费矩阵（题目数据规模上限）
	costs := make([][]int, 100)
	for i := range costs {
		costs[i] = []int{(i*7)%20 + 1, (i*11)%20 + 1, (i*13)%20 + 1}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinCost(costs)
	}
}
