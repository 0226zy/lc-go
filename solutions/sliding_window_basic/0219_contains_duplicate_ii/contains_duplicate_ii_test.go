package containsduplicateii

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		// LeetCode 官方示例
		{"示例1: [1,2,3,1] k=3", []int{1, 2, 3, 1}, 3, true},
		{"示例2: [1,0,1,1] k=1", []int{1, 0, 1, 1}, 1, true},
		{"示例3: [1,2,3,1,2,3] k=2", []int{1, 2, 3, 1, 2, 3}, 2, false},

		// 边界：单元素 / k=0
		{"单元素无法形成一对", []int{1}, 1, false},
		{"k=0 要求下标差为0不可能", []int{1, 1}, 0, false},
		{"相邻相同且 k=1", []int{1, 1}, 1, true},

		// 边界：刚好等于 k / 刚好超过 k
		{"距离刚好等于 k", []int{1, 2, 1}, 2, true},
		{"距离刚好超过 k", []int{1, 2, 1}, 1, false},
		{"多个候选取最近下标仍满足", []int{1, 1, 1}, 1, true},

		// 负数与极值
		{"负数重复在窗口内", []int{-1, 2, -1}, 2, true},
		{"大数重复", []int{1000000000, 0, 1000000000}, 2, true},

		// 无重复
		{"全部互异", []int{1, 2, 3, 4}, 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsNearbyDuplicate(tt.nums, tt.k); got != tt.want {
				t.Errorf("ContainsNearbyDuplicate(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func TestContainsNearbyDuplicateStress(t *testing.T) {
	n := 100000
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	t.Run("十万互异元素", func(t *testing.T) {
		if ContainsNearbyDuplicate(nums, 1000) {
			t.Errorf("期望 false")
		}
	})
	nums[n-1] = 0
	t.Run("首尾相同但距离过大", func(t *testing.T) {
		if ContainsNearbyDuplicate(nums, 10) {
			t.Errorf("期望 false，距离远大于 k")
		}
	})
	t.Run("首尾相同且 k 足够大", func(t *testing.T) {
		if !ContainsNearbyDuplicate(nums, n-1) {
			t.Errorf("期望 true")
		}
	})
}

func BenchmarkContainsNearbyDuplicate(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i % 997
	}
	benchmarks := []struct {
		name string
		nums []int
		k    int
	}{
		{"短数组命中", []int{1, 2, 3, 1}, 3},
		{"短数组未命中", []int{1, 2, 3, 1, 2, 3}, 2},
		{"一万元素", nums, 100},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ContainsNearbyDuplicate(bm.nums, bm.k)
			}
		})
	}
}
