package mergeintervals

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		// LeetCode 官方示例
		{"示例1: 三个独立区间", [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{"示例2: 端点相接", [][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},

		// 边界：单元素
		{"单个区间", [][]int{{1, 4}}, [][]int{{1, 4}}},
		{"单个点区间", [][]int{{1, 1}}, [][]int{{1, 1}}},

		// 边界：完全重叠/包含
		{"区间互相包含", [][]int{{1, 10}, {2, 3}, {4, 5}}, [][]int{{1, 10}}},
		{"完全相同的区间", [][]int{{2, 3}, {2, 3}, {2, 3}}, [][]int{{2, 3}}},

		// 边界：首尾相连成一条
		{"依次首尾相接", [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, [][]int{{1, 5}}},
		{"不相邻的区间", [][]int{{1, 3}, {5, 7}, {9, 11}}, [][]int{{1, 3}, {5, 7}, {9, 11}}},

		// 边界：乱序输入
		{"乱序重叠", [][]int{{8, 10}, {1, 3}, {2, 6}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{"后段包含前段", [][]int{{3, 5}, {1, 10}}, [][]int{{1, 10}}},

		// 边界：左端点相同
		{"左端点相同", [][]int{{1, 4}, {1, 5}}, [][]int{{1, 5}}},
		{"左端点相同且递减右端点", [][]int{{1, 10}, {1, 2}}, [][]int{{1, 10}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 深拷贝一份输入，防止 Merge 内部排序修改原切片干扰断言
			intervals := make([][]int, len(tt.intervals))
			for i, iv := range tt.intervals {
				intervals[i] = append([]int(nil), iv...)
			}
			got := Merge(intervals)
			if !utils.Equal2DIntSlice(got, tt.want) {
				t.Errorf("Merge(%v) = %v, want %v", tt.intervals, got, tt.want)
			}
		})
	}
}

func BenchmarkMerge(b *testing.B) {
	// 构造 n 个随机区间
	generate := func(n int) [][]int {
		intervals := make([][]int, n)
		for i := 0; i < n; i++ {
			start := (i * 7919) % 100000
			intervals[i] = []int{start, start + (i*104729)%1000}
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
				// 每次深拷贝，避免已排序输入影响基准
				intervals := make([][]int, len(bm.intervals))
				for j, iv := range bm.intervals {
					intervals[j] = append([]int(nil), iv...)
				}
				Merge(intervals)
			}
		})
	}
}
