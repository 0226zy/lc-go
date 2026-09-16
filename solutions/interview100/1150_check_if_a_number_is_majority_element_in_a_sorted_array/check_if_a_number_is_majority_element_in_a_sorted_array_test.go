package checkifanumberismajorityelementinasortedarray

import "testing"

func TestIsMajorityElement(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		// LeetCode 官方示例
		{"示例1: 5出现5次占多数", []int{2, 4, 5, 5, 5, 5, 5, 6, 6}, 5, true},
		{"示例2: 101出现2次恰好一半", []int{10, 100, 101, 101}, 101, false},

		// 边界：单元素数组
		{"单元素命中", []int{7}, 7, true},
		{"单元素未命中", []int{7}, 8, false},

		// 边界：target 不存在
		{"target小于所有元素", []int{2, 2, 3}, 1, false},
		{"target大于所有元素", []int{2, 2, 3}, 4, false},
		{"target落在元素间隙", []int{1, 3, 5}, 4, false},

		// 恰好一半 vs 超过一半
		{"恰好一半不算多数", []int{1, 1, 2, 2}, 1, false},
		{"超过一半是多数", []int{1, 1, 1, 2, 2}, 1, true},

		// target 在数组两端
		{"多数元素在头部", []int{1, 1, 1, 1, 2, 3, 4}, 1, true},
		{"多数元素在尾部", []int{1, 2, 3, 4, 4, 4, 4}, 4, true},
		{"尾部元素未过半", []int{1, 2, 3, 4}, 4, false},

		// 全部元素相同
		{"全部元素相同", []int{9, 9, 9, 9, 9}, 9, true},

		// 偶数长度需严格大于一半
		{"偶数长度刚好多一个", []int{1, 2, 2, 2}, 2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMajorityElement(tt.nums, tt.target); got != tt.want {
				t.Errorf("IsMajorityElement(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func BenchmarkIsMajorityElement(b *testing.B) {
	// 构造长度 1000 的排序数组，target 为多数元素
	nums := make([]int, 1000)
	for i := range nums {
		if i < 600 {
			nums[i] = 5
		} else {
			nums[i] = 100 + i
		}
	}

	b.Run("n=1000多数元素命中", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			IsMajorityElement(nums, 5)
		}
	})
	b.Run("n=1000未命中", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			IsMajorityElement(nums, 3)
		}
	})
}
