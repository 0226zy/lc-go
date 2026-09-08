package gasstation

import (
	"strconv"
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		name string
		gas  []int
		cost []int
		want int
	}{
		// LeetCode 官方示例
		{"示例1", []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{"示例2", []int{2, 3, 4}, []int{3, 4, 3}, -1},

		// 边界：恰好走一圈
		{"恰好够", []int{5}, []int{5}, 0},
		{"单站不够", []int{1}, []int{2}, -1},

		// 起点是 0 号站
		{"从0出发", []int{4}, []int{3}, 0},
		{"从0出发多站", []int{3, 1, 1}, []int{1, 2, 1}, 0},

		// 边界：所有 gas 与 cost 相等
		{"全程平衡", []int{2, 2, 2}, []int{2, 2, 2}, 0},

		// 边界：总量够但中途断油
		{"总量够但中途断油", []int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}, 4},

		// 边界：最后一站作为起点
		{"末站出发", []int{2, 1, 1, 4}, []int{2, 2, 2, 2}, 3},

		// 边界：两个站互补
		{"两站互补", []int{2, 3}, []int{3, 2}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanCompleteCircuit(tt.gas, tt.cost); got != tt.want {
				t.Errorf("CanCompleteCircuit(%v, %v) = %d, want %d", tt.gas, tt.cost, got, tt.want)
			}
		})
	}
}

func BenchmarkCanCompleteCircuit(b *testing.B) {
	benchmarks := []struct {
		name string
		gas  []int
		cost []int
	}{
		{"len=10", []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, []int{15, 25, 35, 45, 5, 55, 65, 75, 85, 95}},
	}
	for _, n := range []int{100, 1000, 10000, 100000} {
		gas, cost := generateGasCost(n)
		benchmarks = append(benchmarks, struct {
			name string
			gas  []int
			cost []int
		}{"len=" + strconv.Itoa(n), gas, cost})
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CanCompleteCircuit(bm.gas, bm.cost)
			}
		})
	}
}

// generateGasCost 生成长度为 n 的加油站数据，保证存在可行解
func generateGasCost(n int) ([]int, []int) {
	gas := make([]int, n)
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		gas[i] = i%50 + 10
		cost[i] = i%30 + 5
	}
	return gas, cost
}
