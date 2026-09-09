package maximumsubarray

import "testing"

var maxSubArrayCases = []struct {
	name string
	nums []int
	want int
}{
	// LeetCode 官方示例
	{name: "示例1: [-2,1,-3,4,-1,2,1,-5,4]", nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
	{name: "示例2: [1]", nums: []int{1}, want: 1},
	{name: "示例3: [5,4,-1,7,8]", nums: []int{5, 4, -1, 7, 8}, want: 23},

	// 边界：单元素
	{name: "单元素为负数", nums: []int{-1}, want: -1},
	{name: "单元素为最小值", nums: []int{-10000}, want: -10000},

	// 边界：全负数
	{name: "全负数", nums: []int{-2, -1, -3, -4}, want: -1},
	{name: "全负数且递减", nums: []int{-5, -4, -3, -2, -1}, want: -1},

	// 边界：全正数
	{name: "全正数", nums: []int{1, 2, 3, 4}, want: 10},

	// 边界：正负交错
	{name: "正负交错", nums: []int{1, -1, 1, -1, 1}, want: 1},
	{name: "先正后负拖尾", nums: []int{3, -2, -1, -10}, want: 3},
	{name: "中间低谷后反弹", nums: []int{-2, -3, 4, -1, -2, 1, 5, -3}, want: 7},
	{name: "首尾大中间负", nums: []int{8, -19, 5}, want: 8},
	{name: "含零", nums: []int{-1, 0, -2}, want: 0},

	// 边界：答案就是整个数组
	{name: "累加最大", nums: []int{1, 2, -1, 2, 3}, want: 7},
	{name: "极大值混入", nums: []int{-1, 0, -2, 10000, -10000}, want: 10000},
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

var benchNums = []struct {
	name string
	nums []int
}{
	{"len=10", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4, 3}},
	{"len=100", generateNums(100)},
	{"len=1000", generateNums(1000)},
	{"len=10000", generateNums(10000)},
	{"len=100000", generateNums(100000)},
}

func BenchmarkMaxSubArray(b *testing.B) {
	for _, bm := range benchNums {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MaxSubArray(bm.nums)
			}
		})
	}
}

func BenchmarkMaxSubArrayOptimized(b *testing.B) {
	for _, bm := range benchNums {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MaxSubArrayOptimized(bm.nums)
			}
		})
	}
}

// generateNums 生成长度为 n 的正负交错数组
func generateNums(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			arr[i] = i % 97
		} else {
			arr[i] = -(i % 89)
		}
	}
	return arr
}
