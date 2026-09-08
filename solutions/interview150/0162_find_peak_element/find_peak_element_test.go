package findpeakelement

import "testing"

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		// LeetCode 官方示例
		{"示例1: [1,2,3,1]", []int{1, 2, 3, 1}},
		{"示例2: [1,2,1,3,5,6,4]", []int{1, 2, 1, 3, 5, 6, 4}},

		// 边界：单元素（边界外都是 -∞，必为峰值）
		{"单元素", []int{42}},
		{"单元素负数", []int{-100}},

		// 边界：单调递增，峰值在末尾
		{"严格递增", []int{1, 2, 3, 4, 5}},
		// 边界：单调递减，峰值在开头
		{"严格递减", []int{5, 4, 3, 2, 1}},

		// 边界：两个元素
		{"两个元素升序", []int{1, 2}},
		{"两个元素降序", []int{2, 1}},

		// 边界：多个峰值，任一即可
		{"多个峰值", []int{1, 5, 1, 5, 1}},
		{"正负交替", []int{-3, -1, -4, -1, -6}},
		{"谷底型", []int{5, 1, 5}},

		// 边界：极值
		{"含最大最小int", []int{-2147483648, 2147483647}},
		{"峰值在中间偏左", []int{1, 3, 2, 4, 6, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindPeakElement(tt.nums)
			if !isPeak(tt.nums, got) {
				t.Errorf("FindPeakElement(%v) = %d, 下标 %d 处值 %d 不是峰值",
					tt.nums, got, got, tt.nums[got])
			}
		})
	}
}

// isPeak 校验下标 i 是否为峰值（边界外视为 -∞）
func isPeak(nums []int, i int) bool {
	if i < 0 || i >= len(nums) {
		return false
	}
	if i > 0 && nums[i] <= nums[i-1] {
		return false
	}
	if i < len(nums)-1 && nums[i] <= nums[i+1] {
		return false
	}
	return true
}

func BenchmarkFindPeakElement(b *testing.B) {
	benchmarks := []struct {
		name string
		nums []int
	}{
		{"len=10-多个峰", []int{1, 5, 1, 5, 1, 5, 1, 5, 1, 5}},
		{"len=100", generateWaveArray(100)},
		{"len=1000", generateWaveArray(1000)},
		{"len=10000", generateWaveArray(10000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindPeakElement(bm.nums)
			}
		})
	}
}

// generateWaveArray 生成长度为 n 的波浪数组（含多个峰值）
func generateWaveArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if i%4 < 2 {
			arr[i] = i
		} else {
			arr[i] = n - i
		}
	}
	return arr
}
