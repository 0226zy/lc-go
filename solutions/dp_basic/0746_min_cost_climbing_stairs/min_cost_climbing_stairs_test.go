package mincostclimbingstairs

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		name string
		cost []int
		want int
	}{
		{"示例1", []int{10, 15, 20}, 15},
		{"示例2", []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
		{"两台阶选便宜的", []int{5, 1}, 1},
		{"两台阶", []int{3, 5}, 3},
		{"单台阶", []int{7}, 0},
		{"全部相同", []int{2, 2, 2, 2, 2}, 4},
		{"递增", []int{1, 2, 3, 4, 5}, 6},
		{"递减", []int{5, 4, 3, 2, 1}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinCostClimbingStairs(tt.cost)
			if got != tt.want {
				t.Errorf("MinCostClimbingStairs(%v) = %d, want %d", tt.cost, got, tt.want)
			}
		})
	}
}

func BenchmarkMinCostClimbingStairs(b *testing.B) {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	for i := 0; i < b.N; i++ {
		MinCostClimbingStairs(cost)
	}
}
