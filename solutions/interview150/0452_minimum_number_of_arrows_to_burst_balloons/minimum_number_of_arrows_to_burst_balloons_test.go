package minimumnumberofarrowstoburstballoons

import "testing"

func TestFindMinArrowShots(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		// LeetCode 官方示例
		{"示例1: 两箭引爆四个气球", [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2},
		{"示例2: 互不相交各需一箭", [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 4},
		{"示例3: 端点相接两箭", [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2},

		// 边界：单个气球
		{"单个气球", [][]int{{1, 2}}, 1},
		{"单个气球单点", [][]int{{5, 5 + 1}}, 1},

		// 边界：完全重叠
		{"完全相同的区间", [][]int{{1, 6}, {1, 6}, {1, 6}}, 1},
		{"大区间与相邻小区间共一箭", [][]int{{1, 10}, {2, 3}, {3, 4}}, 1},

		// 边界：包含负坐标
		{"含负坐标", [][]int{{-10, -5}, {-8, -6}, {1, 3}}, 2},
		{"全为负坐标", [][]int{{-6, -5}, {-4, -3}, {-2, -1}}, 3},

		// 边界：交错重叠
		{"交错重叠需两箭", [][]int{{1, 5}, {2, 6}, {7, 10}, {8, 12}}, 2},
		{"前两个重叠后两个重叠", [][]int{{1, 4}, {2, 5}, {8, 9}, {9, 10}}, 2},

		// 边界：大量区间依次相接成链
		{"链条状相接", [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 深拷贝输入，防止实现内部排序修改原切片干扰断言
			points := make([][]int, len(tt.points))
			for i, p := range tt.points {
				points[i] = append([]int(nil), p...)
			}
			got := FindMinArrowShots(points)
			if got != tt.want {
				t.Errorf("FindMinArrowShots(%v) = %d, want %d", tt.points, got, tt.want)
			}
		})
	}
}

func BenchmarkFindMinArrowShots(b *testing.B) {
	// 构造 n 个随机宽度的区间
	generate := func(n int) [][]int {
		points := make([][]int, n)
		for i := 0; i < n; i++ {
			start := (i * 7919) % 1000000
			points[i] = []int{start, start + 1 + (i*104729)%100}
		}
		return points
	}

	benchmarks := []struct {
		name   string
		points [][]int
	}{
		{"n=10", generate(10)},
		{"n=100", generate(100)},
		{"n=1000", generate(1000)},
		{"n=100000", generate(100000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				points := make([][]int, len(bm.points))
				for j, p := range bm.points {
					points[j] = append([]int(nil), p...)
				}
				FindMinArrowShots(points)
			}
		})
	}
}
