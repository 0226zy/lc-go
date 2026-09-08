package singlenumber

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		// LeetCode 官方示例
		{"示例1", []int{2, 2, 1}, 1},
		{"示例2", []int{4, 1, 2, 1, 2}, 4},
		{"示例3", []int{1}, 1},

		// 边界：成对元素相邻或相隔
		{"答案在开头", []int{9, 3, 3}, 9},
		{"答案在结尾", []int{3, 3, 9}, 9},
		{"多对混合", []int{1, 2, 3, 1, 2}, 3},

		// 边界：极值
		{"含最小整数", []int{-2147483648, 5, -2147483648}, 5},
		{"含最大整数", []int{2147483647, 0, 0}, 2147483647},
		{"负数答案", []int{-1, 2, 2}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SingleNumber(tt.nums); got != tt.want {
				t.Errorf("SingleNumber(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkSingleNumber(b *testing.B) {
	b.Run("长度11", func(b *testing.B) {
		nums := []int{4, 1, 2, 1, 2, 7, 8, 7, 9, 9, 3}
		for i := 0; i < b.N; i++ {
			SingleNumber(nums)
		}
	})
	b.Run("长度100001", func(b *testing.B) {
		nums := make([]int, 100001)
		for i := 0; i < len(nums); i += 2 {
			nums[i], nums[i+1] = i, i
		}
		nums[100000] = -12345 // 最后一个元素只出现一次
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			SingleNumber(nums)
		}
	})
}
