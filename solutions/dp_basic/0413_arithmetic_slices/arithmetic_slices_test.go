package arithmeticslices

import "testing"

var arithmeticSlicesCases = []struct {
	name string
	nums []int
	want int
}{
	{name: "官方示例1：[1,2,3,4]", nums: []int{1, 2, 3, 4}, want: 3},
	{name: "官方示例2：单元素", nums: []int{1}, want: 0},
	{name: "两个元素不足三个", nums: []int{1, 2}, want: 0},
	{name: "五个连续整数", nums: []int{1, 2, 3, 4, 5}, want: 6},
	{name: "公差为2的等差数组", nums: []int{2, 4, 6, 8, 10}, want: 6},
	{name: "中间被打断", nums: []int{1, 2, 3, 5}, want: 1},
	{name: "全部相等（公差为0）", nums: []int{7, 7, 7, 7}, want: 3},
	{name: "两段等差", nums: []int{1, 2, 3, 8, 9, 10}, want: 2},
	{name: "负数公差", nums: []int{3, -1, -5, -9}, want: 3},
	{name: "六个元素的等差数组", nums: []int{1, 3, 5, 7, 9, 11}, want: 10},
}

func TestNumberOfArithmeticSlices(t *testing.T) {
	for _, tt := range arithmeticSlicesCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberOfArithmeticSlices(tt.nums); got != tt.want {
				t.Errorf("NumberOfArithmeticSlices(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestNumberOfArithmeticSlicesOptimized(t *testing.T) {
	for _, tt := range arithmeticSlicesCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberOfArithmeticSlicesOptimized(tt.nums); got != tt.want {
				t.Errorf("NumberOfArithmeticSlicesOptimized(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkNumberOfArithmeticSlices(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		NumberOfArithmeticSlices(nums)
	}
}

func BenchmarkNumberOfArithmeticSlicesOptimized(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		NumberOfArithmeticSlicesOptimized(nums)
	}
}
