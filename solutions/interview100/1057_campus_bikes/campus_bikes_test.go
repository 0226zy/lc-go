package campusbikes

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestAssignBikes(t *testing.T) {
	tests := []struct {
		name    string
		workers [][]int
		bikes   [][]int
		want    []int
	}{
		// LeetCode 官方示例
		{
			"示例1: 两名工人两辆自行车",
			[][]int{{0, 0}, {2, 1}},
			[][]int{{1, 2}, {3, 3}},
			[]int{1, 0},
		},
		{
			"示例2: 三名工人三辆自行车",
			[][]int{{0, 0}, {1, 1}, {2, 0}},
			[][]int{{1, 0}, {2, 2}, {2, 1}},
			[]int{0, 2, 1},
		},

		// 边界：单工人单自行车
		{
			"单工人单自行车",
			[][]int{{0, 0}},
			[][]int{{5, 5}},
			[]int{0},
		},

		// 边界：自行车多于工人，需挑最近的
		{
			"自行车多于工人",
			[][]int{{0, 0}},
			[][]int{{9, 9}, {0, 3}, {1, 0}},
			[]int{2},
		},

		// 边界：距离相同时按工人下标优先
		{
			"同距离按工人下标分配",
			[][]int{{0, 0}, {2, 0}},
			[][]int{{1, 0}, {3, 0}},
			[]int{0, 1},
		},

		// 边界：工人与自行车位置重合（距离为 0）
		{
			"位置重合距离为0",
			[][]int{{1, 1}, {0, 0}},
			[][]int{{1, 1}, {0, 0}},
			[]int{0, 1},
		},

		// 典型场景：贪心顺序影响后续分配
		{
			"最近对被抢占后顺延",
			[][]int{{0, 0}, {1, 0}},
			[][]int{{0, 0}, {10, 0}},
			[]int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AssignBikes(tt.workers, tt.bikes); !utils.EqualIntSlice(got, tt.want) {
				t.Errorf("AssignBikes(%v, %v) = %v, want %v", tt.workers, tt.bikes, got, tt.want)
			}
		})
	}
}

func BenchmarkAssignBikes(b *testing.B) {
	// 构造 100 名工人、100 辆自行车（10000 个配对）
	workers := make([][]int, 100)
	bikes := make([][]int, 100)
	for i := range workers {
		workers[i] = []int{i * 10, (i * 37) % 1000}
		bikes[i] = []int{(i * 53) % 1000, i * 7}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AssignBikes(workers, bikes)
	}
}
