package maximumlengthsubarraypositiveproduct

import "testing"

var getMaxLenCases = []struct {
	name string
	nums []int
	want int
}{
	{name: "示例1：[1,-2,-3,4]", nums: []int{1, -2, -3, 4}, want: 4},
	{name: "示例2：[0,1,-2,-3,-4]", nums: []int{0, 1, -2, -3, -4}, want: 3},
	{name: "示例3：[-1,-2,-3,0,1]", nums: []int{-1, -2, -3, 0, 1}, want: 2},
	{name: "单个正数", nums: []int{5}, want: 1},
	{name: "单个负数", nums: []int{-5}, want: 0},
	{name: "单个零", nums: []int{0}, want: 0},
	{name: "全部正数", nums: []int{1, 2, 3, 4, 5}, want: 5},
	{name: "奇数个负数", nums: []int{-1, -2, -3}, want: 2},
	{name: "偶数个负数", nums: []int{-1, -2, -3, -4}, want: 4},
	{name: "零分割多段", nums: []int{1, 2, 0, -1, -2, 0, 3}, want: 2},
	{name: "先负后正", nums: []int{-1, 2}, want: 1},
	{name: "负数在两端", nums: []int{-1, 1, 1, -1}, want: 4},
}

func TestGetMaxLen(t *testing.T) {
	for _, tt := range getMaxLenCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMaxLen(tt.nums); got != tt.want {
				t.Errorf("GetMaxLen(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestGetMaxLenOptimized(t *testing.T) {
	for _, tt := range getMaxLenCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMaxLenOptimized(tt.nums); got != tt.want {
				t.Errorf("GetMaxLenOptimized(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkGetMaxLen(b *testing.B) {
	nums := make([]int, 100000)
	for i := range nums {
		nums[i] = i%3 - 1 // -1, 0, 1 循环
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetMaxLen(nums)
	}
}

func BenchmarkGetMaxLenOptimized(b *testing.B) {
	nums := make([]int, 100000)
	for i := range nums {
		nums[i] = i%3 - 1 // -1, 0, 1 循环
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetMaxLenOptimized(nums)
	}
}
