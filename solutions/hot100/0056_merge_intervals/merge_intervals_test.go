package mergeintervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		intervals [][]int
		want     [][]int
	}{
		// LeetCode 官方示例
		{"示例1", [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{"示例2", [][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},

		// 边界：单区间
		{"单区间", [][]int{{1, 3}}, [][]int{{1, 3}}},

		// 边界：全重叠
		{"全重叠", [][]int{{1, 5}, {2, 6}, {3, 7}}, [][]int{{1, 7}}},

		// 边界：全不重叠
		{"全不重叠", [][]int{{1, 2}, {3, 4}, {5, 6}}, [][]int{{1, 2}, {3, 4}, {5, 6}}},

		// 边界：包含关系
		{"包含关系", [][]int{{1, 10}, {2, 3}, {4, 5}}, [][]int{{1, 10}}},

		// 边界：逆序输入
		{"逆序输入", [][]int{{8, 10}, {1, 3}, {2, 6}}, [][]int{{1, 6}, {8, 10}}},

		// 边界：相邻区间
		{"相邻区间", [][]int{{1, 2}, {2, 3}, {3, 4}}, [][]int{{1, 4}}},

		// 较大数组
		{"多个区间混合", [][]int{{1, 3}, {2, 4}, {5, 7}, {6, 8}, {9, 10}}, [][]int{{1, 4}, {5, 8}, {9, 10}}},
		{"复杂场景", [][]int{{1, 4}, {0, 2}, {3, 5}}, [][]int{{0, 5}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制输入，避免测试间相互影响
			input := make([][]int, len(tt.intervals))
			for i := range tt.intervals {
				input[i] = make([]int, len(tt.intervals[i]))
				copy(input[i], tt.intervals[i])
			}
			got := Merge(input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge(%v) = %v, want %v", tt.intervals, got, tt.want)
			}
		})
	}
}

func BenchmarkMerge(b *testing.B) {
	benchmarks := []struct {
		name     string
		intervals [][]int
	}{
		{"len=4", [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
		{"len=10", generateIntervals(10)},
		{"len=100", generateIntervals(100)},
		{"len=1000", generateIntervals(1000)},
		{"len=10000", generateIntervals(10000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// 复制输入，避免原地修改影响后续迭代
				input := make([][]int, len(bm.intervals))
				for j := range bm.intervals {
					input[j] = make([]int, len(bm.intervals[j]))
					copy(input[j], bm.intervals[j])
				}
				Merge(input)
			}
		})
	}
}

// generateIntervals 生成 n 个随机区间
func generateIntervals(n int) [][]int {
	intervals := make([][]int, n)
	for i := 0; i < n; i++ {
		start := i * 2
		end := start + 1 + (i % 3) // 有些重叠，有些不重叠
		intervals[i] = []int{start, end}
	}
	return intervals
}
