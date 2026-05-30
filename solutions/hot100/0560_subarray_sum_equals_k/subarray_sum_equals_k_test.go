package subarraysumequalk

import "testing"

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		k     int
		want  int
	}{
		// LeetCode 官方示例
		{"示例1: [1,1,1], k=2", []int{1, 1, 1}, 2, 2},
		{"示例2: [1,2,3], k=3", []int{1, 2, 3}, 3, 2},

		// 边界：单元素
		{"单元素等于k", []int{5}, 5, 1},
		{"单元素不等于k", []int{3}, 5, 0},

		// 边界：全为0
		{"全为0, k=0", []int{0, 0, 0}, 0, 6},
		{"全为0, k=1", []int{0, 0, 0}, 1, 0},

		// 边界：全负数
		{"全负数", []int{-1, -1, -1}, -2, 2},

		// 正负数混合
		{"正负混合1", []int{1, -1, 1, -1, 1}, 0, 6},
		{"正负混合2", []int{1, 2, -3, 3, 1}, 3, 4},

		// 重复前缀和
		{"重复前缀和", []int{1, 1, 1, 1}, 2, 3},

		// 空数组（题目要求至少1个元素，这里测试边界）
		// {"空数组", []int{}, 0, 0}, // 不符合题目约束，移除

		// 较大数组
		{"连续递增", []int{1, 2, 3, 4, 5}, 9, 2},
		{"含重复元素", []int{3, 4, 7, 2, 3, 1, 4, 2, 1}, 7, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SubarraySum(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("SubarraySum(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func BenchmarkSubarraySum(b *testing.B) {
	benchmarks := []struct {
		name  string
		nums  []int
		k     int
	}{
		{"len=10", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 15},
		{"len=100", generateArray(100), 500},
		{"len=1000", generateArray(1000), 5000},
		{"len=10000", generateArray(10000), 50000},
		{"全0_len=1000", make([]int, 1000), 0},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SubarraySum(bm.nums, bm.k)
			}
		})
	}
}

// generateArray 生成长度为 n 的数组 [1, 2, ..., n]
func generateArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	return arr
}
