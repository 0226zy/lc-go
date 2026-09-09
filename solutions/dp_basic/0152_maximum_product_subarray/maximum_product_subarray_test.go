package maximumproductsubarray

import "testing"

var maxProductCases = []struct {
	name string
	nums []int
	want int
}{
	{name: "示例1：正数段", nums: []int{2, 3, -2, 4}, want: 6},
	{name: "示例2：含0断开", nums: []int{-2, 0, -1}, want: 0},
	{name: "单个负数", nums: []int{-2}, want: -2},
	{name: "单个0", nums: []int{0}, want: 0},
	{name: "负负得正", nums: []int{-2, 3, -4}, want: 24},
	{name: "0之后重新开始", nums: []int{0, 2}, want: 2},
	{name: "两段负数取其一", nums: []int{-3, -1, -1}, want: 3},
	{name: "先乘后遇到0再重启", nums: []int{3, -1, 0, 5, 2}, want: 10},
	{name: "全负奇数个取后缀", nums: []int{-1, -2, -3, -4, -5}, want: 120},
	{name: "混合长数组", nums: []int{2, -5, -2, -4, 3}, want: 24},
}

func TestMaxProduct(t *testing.T) {
	for _, tt := range maxProductCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProduct(tt.nums); got != tt.want {
				t.Errorf("MaxProduct(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestMaxProductOptimized(t *testing.T) {
	for _, tt := range maxProductCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProductOptimized(tt.nums); got != tt.want {
				t.Errorf("MaxProductOptimized(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkMaxProduct(b *testing.B) {
	nums := []int{2, 3, -2, 4, -1, 2, 3, -4, 2, 1}
	for i := 0; i < b.N; i++ {
		MaxProduct(nums)
	}
}

func BenchmarkMaxProductOptimized(b *testing.B) {
	nums := []int{2, 3, -2, 4, -1, 2, 3, -4, 2, 1}
	for i := 0; i < b.N; i++ {
		MaxProductOptimized(nums)
	}
}
