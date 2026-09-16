package dividechocolate

import "testing"

func TestMaximizeSweetness(t *testing.T) {
	tests := []struct {
		name      string
		sweetness []int
		k         int
		want      int
	}{
		// LeetCode 官方示例
		{"示例1: 递增序列分6块", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, 6},
		{"示例2: 块数等于长度", []int{5, 6, 7, 8, 9, 1, 2, 3, 4}, 8, 1},
		{"示例3: 均分成3块", []int{1, 2, 2, 1, 2, 2, 1, 2, 2}, 2, 5},

		// 边界：不切，自己吃整块
		{"不切刀拿走全部", []int{3, 1, 2}, 0, 6},

		// 边界：单元素数组
		{"单元素数组", []int{7}, 0, 7},

		// 边界：所有元素相同
		{"全部相同均分", []int{4, 4, 4, 4}, 3, 4},
		{"全部相同无法均分", []int{5, 5, 5, 5}, 2, 5},

		// 典型场景：存在极小元素拖低答案
		{"极小元素限制答案", []int{100, 1, 100}, 2, 1},
		{"极小元素但不分给它", []int{100, 1, 100}, 1, 100},

		// 典型场景：贪心切割需要跨元素合并
		{"需要合并多元素", []int{1, 1, 1, 1, 10}, 1, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaximizeSweetness(tt.sweetness, tt.k); got != tt.want {
				t.Errorf("MaximizeSweetness(%v, %d) = %d, want %d", tt.sweetness, tt.k, got, tt.want)
			}
		})
	}
}

func BenchmarkMaximizeSweetness(b *testing.B) {
	// 构造 10^4 长度的数组：1..100 循环
	sweetness := make([]int, 0, 10000)
	for len(sweetness) < 10000 {
		for v := 1; v <= 100; v++ {
			sweetness = append(sweetness, v)
		}
	}
	sweetness = sweetness[:10000]

	b.Run("n=10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			MaximizeSweetness(sweetness, 100)
		}
	})
}
