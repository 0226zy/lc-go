package maximumsumcircularsubarray

import "testing"

var maxSubarraySumCircularCases = []struct {
	name string
	nums []int
	want int
}{
	{name: "示例1不跨界", nums: []int{1, -2, 3, -2}, want: 3},
	{name: "示例2跨界", nums: []int{5, -3, 5}, want: 10},
	{name: "示例3全负取最大单元素", nums: []int{-3, -2, -3}, want: -2},
	{name: "单元素", nums: []int{7}, want: 7},
	{name: "单元素负数", nums: []int{-7}, want: -7},
	{name: "全正取整个数组", nums: []int{1, 2, 3, 4}, want: 10},
	{name: "跨界挖掉中间负数", nums: []int{3, -1, 2, -1}, want: 4},
	{name: "官方样例跨界", nums: []int{3, -2, 2, -3}, want: 3},
	{name: "两端大中间小负", nums: []int{2, -1, -1, 2}, want: 4},
	{name: "全相同负数", nums: []int{-5, -5, -5}, want: -5},
}

func TestMaxSubarraySumCircular(t *testing.T) {
	for _, tt := range maxSubarraySumCircularCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubarraySumCircular(tt.nums); got != tt.want {
				t.Errorf("MaxSubarraySumCircular(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestMaxSubarraySumCircularOptimized(t *testing.T) {
	for _, tt := range maxSubarraySumCircularCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubarraySumCircularOptimized(tt.nums); got != tt.want {
				t.Errorf("MaxSubarraySumCircularOptimized(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkMaxSubarraySumCircular(b *testing.B) {
	nums := []int{5, -3, 5, -2, 8, -10, 6, 4, -1, 9}
	for i := 0; i < b.N; i++ {
		MaxSubarraySumCircular(nums)
	}
}

func BenchmarkMaxSubarraySumCircularOptimized(b *testing.B) {
	nums := []int{5, -3, 5, -2, 8, -10, 6, 4, -1, 9}
	for i := 0; i < b.N; i++ {
		MaxSubarraySumCircularOptimized(nums)
	}
}
