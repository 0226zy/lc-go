package maximumdistanceinarrays

import "testing"

func TestMaxDistance(t *testing.T) {
	tests := []struct {
		name   string
		arrays [][]int
		want   int
	}{
		// LeetCode 官方示例
		{"示例1: 三个数组", [][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}, 4},
		{"示例2: 两个相同的单元素数组", [][]int{{1}, {1}}, 0},

		// 边界：最小规模的两个数组
		{"两个单元素数组", [][]int{{1}, {5}}, 4},

		// 边界：全局最小值与全局最大值在同一个数组中，正确答案是 max(|1-50|, |100-50|) = 50
		{"最值同数组", [][]int{{1, 100}, {50}}, 50},
		{"最值同数组-反向", [][]int{{50}, {1, 100}}, 50},

		// 边界：包含负数
		{"含负数", [][]int{{-5, -3}, {0, 2}}, 7},
		{"极端值", [][]int{{-10000}, {10000}}, 20000},

		// 典型场景：递减分布
		{"首数组大末数组小", [][]int{{10, 20}, {1, 2}}, 19},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxDistance(tt.arrays); got != tt.want {
				t.Errorf("MaxDistance(%v) = %d, want %d", tt.arrays, got, tt.want)
			}
		})
	}
}

func BenchmarkMaxDistance(b *testing.B) {
	// 构造 10^4 个数组，每个数组 10 个元素
	arrays := make([][]int, 10000)
	for i := range arrays {
		arr := make([]int, 10)
		for j := range arr {
			arr[j] = i*10 + j - 50000
		}
		arrays[i] = arr
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxDistance(arrays)
	}
}
