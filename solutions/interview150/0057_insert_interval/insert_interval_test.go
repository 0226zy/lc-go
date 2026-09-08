package insertinterval

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		// LeetCode 官方示例
		{"示例1: 中间插入合并一个", [][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
		{"示例2: 跨越合并三个区间", [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}}},

		// 边界：空列表
		{"插入到空列表", [][]int{}, []int{5, 7}, [][]int{{5, 7}}},

		// 边界：新区间在最前/最后
		{"插入到最前面", [][]int{{3, 5}, {6, 9}}, []int{1, 2}, [][]int{{1, 2}, {3, 5}, {6, 9}}},
		{"插入到最后面", [][]int{{1, 3}, {6, 9}}, []int{10, 12}, [][]int{{1, 3}, {6, 9}, {10, 12}}},

		// 边界：与首/尾区间合并
		{"与首区间合并", [][]int{{1, 3}, {6, 9}}, []int{0, 5}, [][]int{{0, 5}, {6, 9}}},
		{"与尾区间合并", [][]int{{1, 3}, {6, 9}}, []int{8, 10}, [][]int{{1, 3}, {6, 10}}},

		// 边界：新区间被完全包含/完全包含所有区间
		{"新区间被包含", [][]int{{1, 5}}, []int{2, 3}, [][]int{{1, 5}}},
		{"新区间包含所有", [][]int{{1, 3}, {6, 9}, {12, 15}}, []int{0, 20}, [][]int{{0, 20}}},

		// 边界：端点相接
		{"与首区间端点相接", [][]int{{1, 3}, {6, 9}}, []int{3, 4}, [][]int{{1, 4}, {6, 9}}},
		{"与尾区间端点相接", [][]int{{1, 3}, {6, 9}}, []int{4, 6}, [][]int{{1, 3}, {4, 9}}},

		// 边界：单点区间
		{"新区间为单点", [][]int{{1, 3}, {6, 9}}, []int{4, 4}, [][]int{{1, 3}, {4, 4}, {6, 9}}},
		{"原列表含单点区间", [][]int{{1, 2}, {4, 4}, {6, 9}}, []int{2, 4}, [][]int{{1, 4}, {6, 9}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 深拷贝输入，防止实现修改原切片干扰断言
			intervals := make([][]int, len(tt.intervals))
			for i, iv := range tt.intervals {
				intervals[i] = append([]int(nil), iv...)
			}
			newInterval := append([]int(nil), tt.newInterval...)
			got := Insert(intervals, newInterval)
			if !utils.Equal2DIntSlice(got, tt.want) {
				t.Errorf("Insert(%v, %v) = %v, want %v", tt.intervals, tt.newInterval, got, tt.want)
			}
		})
	}
}

func BenchmarkInsert(b *testing.B) {
	// 构造 n 个互不重叠的区间 [0,1], [3,4], [6,7], ...
	generate := func(n int) [][]int {
		intervals := make([][]int, n)
		for i := 0; i < n; i++ {
			base := i * 3
			intervals[i] = []int{base, base + 1}
		}
		return intervals
	}

	benchmarks := []struct {
		name      string
		intervals [][]int
	}{
		{"n=10", generate(10)},
		{"n=100", generate(100)},
		{"n=1000", generate(1000)},
		{"n=10000", generate(10000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				intervals := make([][]int, len(bm.intervals))
				for j, iv := range bm.intervals {
					intervals[j] = append([]int(nil), iv...)
				}
				Insert(intervals, []int{1, 3*len(bm.intervals) - 2})
			}
		})
	}
}
