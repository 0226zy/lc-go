package removeinterval

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestRemoveInterval(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		toBeRemoved []int
		want        [][]int
	}{
		// LeetCode 官方示例
		{
			"示例1: 三个区间与删除区间相交",
			[][]int{{0, 2}, {3, 4}, {5, 7}},
			[]int{1, 6},
			[][]int{{0, 1}, {6, 7}},
		},
		{
			"示例2: 单个区间被中间挖空",
			[][]int{{0, 5}},
			[]int{2, 3},
			[][]int{{0, 2}, {3, 5}},
		},
		{
			"示例3: 含负数区间",
			[][]int{{-5, -4}, {-3, -2}, {1, 2}, {3, 5}, {8, 9}},
			[]int{-1, 4},
			[][]int{{-5, -4}, {-3, -2}, {4, 5}, {8, 9}},
		},

		// 边界：完全不相交
		{
			"删除区间在所有区间之前",
			[][]int{{5, 7}, {8, 10}},
			[]int{0, 5},
			[][]int{{5, 7}, {8, 10}},
		},
		{
			"删除区间在所有区间之后",
			[][]int{{0, 2}, {3, 4}},
			[]int{4, 6},
			[][]int{{0, 2}, {3, 4}},
		},

		// 边界：全部被删除
		{
			"删除区间覆盖所有区间",
			[][]int{{1, 3}, {4, 6}},
			[]int{0, 10},
			[][]int{},
		},

		// 边界：删除区间恰好等于某个区间
		{
			"删除区间与某区间完全重合",
			[][]int{{0, 2}, {3, 5}, {6, 8}},
			[]int{3, 5},
			[][]int{{0, 2}, {6, 8}},
		},

		// 边界：仅端点相触不算相交
		{
			"端点相触不相交",
			[][]int{{0, 2}, {4, 6}},
			[]int{2, 4},
			[][]int{{0, 2}, {4, 6}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveInterval(tt.intervals, tt.toBeRemoved); !utils.Equal2DIntSlice(got, tt.want) {
				t.Errorf("RemoveInterval(%v, %v) = %v, want %v", tt.intervals, tt.toBeRemoved, got, tt.want)
			}
		})
	}
}

func BenchmarkRemoveInterval(b *testing.B) {
	// 构造 10000 个不相交区间
	intervals := make([][]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		intervals = append(intervals, []int{i * 10, i*10 + 5})
	}
	toBeRemoved := []int{25000, 75000}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveInterval(intervals, toBeRemoved)
	}
}
