package painthouseii

import "testing"

func TestMinCostII(t *testing.T) {
	tests := []struct {
		name  string
		costs [][]int
		want  int
	}{
		// LeetCode 官方示例
		{"示例1: 三颜色两栋房子", [][]int{{1, 5, 3}, {2, 9, 4}}, 5},
		{"示例2: 两颜色两栋房子", [][]int{{1, 3}, {2, 4}}, 5},

		// 边界：空输入与单栋房子
		{"空输入", nil, 0},
		{"单栋房子单色", [][]int{{5}}, 5},
		{"单栋房子多色取最小", [][]int{{8, 3, 6}}, 3},
		{"仅一种颜色且多栋房子无解", [][]int{{2}, {3}}, -1},

		// 典型场景
		{"最便宜颜色恰好交错", [][]int{{20, 1, 20}, {20, 20, 1}, {1, 20, 20}}, 3},
		{"递增矩阵", [][]int{{1, 2}, {3, 4}, {5, 6}}, 10},
		{"交替选择同色以外的最小值", [][]int{{3, 1}, {1, 3}}, 2},
		{"最小值颜色连续冲突需用次小值", [][]int{{1, 100}, {1, 100}, {1, 100}}, 102},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinCostII(tt.costs); got != tt.want {
				t.Errorf("MinCostII(%v) = %d, want %d", tt.costs, got, tt.want)
			}
		})
	}
}

func BenchmarkMinCostII(b *testing.B) {
	// 构造 100 栋房子、20 种颜色的花费矩阵（题目数据规模上限）
	n, k := 100, 20
	costs := make([][]int, n)
	for i := range costs {
		costs[i] = make([]int, k)
		for j := range costs[i] {
			costs[i][j] = (i*7+j*13)%20 + 1
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinCostII(costs)
	}
}
