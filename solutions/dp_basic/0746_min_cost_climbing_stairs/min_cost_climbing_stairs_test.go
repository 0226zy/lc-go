package mincostclimbingstairs

import "testing"

var minCostCases = []struct {
	name string
	cost []int
	want int
}{
	{name: "示例1", cost: []int{10, 15, 20}, want: 15},
	{name: "示例2", cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, want: 6},
	{name: "两台阶选便宜的", cost: []int{5, 1}, want: 1},
	{name: "两台阶", cost: []int{3, 5}, want: 3},
	{name: "单台阶", cost: []int{7}, want: 0},
	{name: "全部相同", cost: []int{2, 2, 2, 2, 2}, want: 4},
	{name: "递增", cost: []int{1, 2, 3, 4, 5}, want: 6},
	{name: "递减", cost: []int{5, 4, 3, 2, 1}, want: 6},
	{name: "全零花费", cost: []int{0, 0, 0, 0}, want: 0},
}

func TestMinCostClimbingStairs(t *testing.T) {
	for _, tt := range minCostCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinCostClimbingStairs(tt.cost); got != tt.want {
				t.Errorf("MinCostClimbingStairs(%v) = %d, want %d", tt.cost, got, tt.want)
			}
		})
	}
}

func TestMinCostClimbingStairsOptimized(t *testing.T) {
	for _, tt := range minCostCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinCostClimbingStairsOptimized(tt.cost); got != tt.want {
				t.Errorf("MinCostClimbingStairsOptimized(%v) = %d, want %d", tt.cost, got, tt.want)
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

func BenchmarkMinCostClimbingStairsOptimized(b *testing.B) {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	for i := 0; i < b.N; i++ {
		MinCostClimbingStairsOptimized(cost)
	}
}
