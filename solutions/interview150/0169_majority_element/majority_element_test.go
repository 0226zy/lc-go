package majorityelement

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: [3,2,3]", []int{3, 2, 3}, 3},
		{"示例2: [2,2,1,1,1,2,2]", []int{2, 2, 1, 1, 1, 2, 2}, 2},

		// 边界：单元素
		{"单元素", []int{1}, 1},
		// 边界：两元素
		{"两元素", []int{2, 2}, 2},
		// 边界：多数元素在开头
		{"多数元素在开头", []int{5, 5, 5, 5, 1, 2}, 5},
		// 边界：多数元素在结尾
		{"多数元素在结尾", []int{1, 2, 3, 4, 4, 4, 4}, 4},
		// 边界：含负数
		{"含负数", []int{-1, -1, -1, 0, 1}, -1},
		// 边界：最大规模提示 1e5
		{"大规模多数在末尾", []int{6, 5, 5, 6, 6, 6, 6, 6}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MajorityElement(tt.nums); got != tt.want {
				t.Errorf("MajorityElement(%v) = %d, want %d", tt.nums, got, tt.want)
			}
			if got := MajorityElementHash(tt.nums); got != tt.want {
				t.Errorf("MajorityElementHash(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkMajorityElement(b *testing.B) {
	nums := make([]int, 100000)
	for i := range nums {
		nums[i] = i % 7 // 制造一个出现次数过半的元素
	}
	b.Run("投票法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			MajorityElement(nums)
		}
	})
	b.Run("哈希法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			MajorityElementHash(nums)
		}
	})
}
