package maximumsubarray

import "testing"

var maxSubArrayCases = []struct {
	name string
	nums []int
	want int
}{
	{name: "示例1", nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
	{name: "示例2_单元素", nums: []int{1}, want: 1},
	{name: "示例3", nums: []int{5, 4, -1, 7, 8}, want: 23},
	{name: "全负数", nums: []int{-3, -1, -2}, want: -1},
	{name: "全正数", nums: []int{1, 2, 3, 4}, want: 10},
	{name: "单个负数", nums: []int{-7}, want: -7},
	{name: "首尾大中间负", nums: []int{8, -19, 5}, want: 8},
	{name: "含零", nums: []int{-1, 0, -2}, want: 0},
}

func TestMaxSubArray(t *testing.T) {
	for _, tt := range maxSubArrayCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubArray(tt.nums); got != tt.want {
				t.Errorf("MaxSubArray(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestMaxSubArrayOptimized(t *testing.T) {
	for _, tt := range maxSubArrayCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubArrayOptimized(tt.nums); got != tt.want {
				t.Errorf("MaxSubArrayOptimized(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

var benchNums53 = func() []int {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = (i*13+7)%201 - 100 // 分布在 [-100, 100]
	}
	return nums
}()

func BenchmarkMaxSubArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxSubArray(benchNums53)
	}
}

func BenchmarkMaxSubArrayOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxSubArrayOptimized(benchNums53)
	}
}
