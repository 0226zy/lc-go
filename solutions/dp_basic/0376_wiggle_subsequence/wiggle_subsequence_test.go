package wigglesubsequence

import "testing"

var wiggleMaxLengthCases = []struct {
	name string
	nums []int
	want int
}{
	{name: "空数组", nums: []int{}, want: 0},
	{name: "单元素", nums: []int{1}, want: 1},
	{name: "官方示例1", nums: []int{1, 7, 4, 9, 2, 5}, want: 6},
	{name: "官方示例2", nums: []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}, want: 7},
	{name: "官方示例3单调递增", nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, want: 2},
	{name: "全部相等", nums: []int{0, 0}, want: 1},
	{name: "含相等元素", nums: []int{3, 3, 1, 2}, want: 3},
	{name: "先平后摆", nums: []int{1, 1, 1, 2, 1}, want: 3},
	{name: "两个不等元素", nums: []int{1, 2}, want: 2},
}

func TestWiggleMaxLength(t *testing.T) {
	for _, tt := range wiggleMaxLengthCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := WiggleMaxLength(tt.nums); got != tt.want {
				t.Errorf("WiggleMaxLength(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestWiggleMaxLengthOptimized(t *testing.T) {
	for _, tt := range wiggleMaxLengthCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := WiggleMaxLengthOptimized(tt.nums); got != tt.want {
				t.Errorf("WiggleMaxLengthOptimized(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkWiggleMaxLength(b *testing.B) {
	nums := []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
	for i := 0; i < b.N; i++ {
		WiggleMaxLength(nums)
	}
}

func BenchmarkWiggleMaxLengthOptimized(b *testing.B) {
	nums := []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
	for i := 0; i < b.N; i++ {
		WiggleMaxLengthOptimized(nums)
	}
}
