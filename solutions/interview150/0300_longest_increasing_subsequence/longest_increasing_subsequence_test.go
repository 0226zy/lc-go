package longestincreasingsubsequence

import "testing"

var lisCases = []struct {
	name string
	nums []int
	want int
}{
	// LeetCode 官方示例
	{name: "官方示例1", nums: []int{10, 9, 2, 5, 3, 7, 101, 18}, want: 4},
	{name: "官方示例2", nums: []int{0, 1, 0, 3, 2, 3}, want: 4},
	{name: "官方示例3", nums: []int{7, 7, 7, 7, 7, 7, 7}, want: 1},

	// 边界
	{name: "空输入", nums: []int{}, want: 0},
	{name: "单元素", nums: []int{1}, want: 1},
	{name: "严格递增", nums: []int{1, 2, 3, 4, 5}, want: 5},
	{name: "严格递减", nums: []int{5, 4, 3, 2, 1}, want: 1},
	{name: "乱序", nums: []int{1, 2, 4, 3}, want: 3},
	{name: "含负数", nums: []int{-2, -1, 0, 1, 2}, want: 5},
	{name: "负数乱序", nums: []int{3, -2, 0, -1, 2}, want: 3},
	{name: "全相同", nums: []int{5, 5, 5, 5}, want: 1},
	{name: "交替升降", nums: []int{1, 3, 2, 4, 3, 5, 4, 6}, want: 5},
}

func TestLengthOfLIS(t *testing.T) {
	for _, tt := range lisCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLIS(tt.nums); got != tt.want {
				t.Errorf("LengthOfLIS(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestLengthOfLISBinary(t *testing.T) {
	for _, tt := range lisCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLISBinary(tt.nums); got != tt.want {
				t.Errorf("LengthOfLISBinary(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestConsistency(t *testing.T) {
	for _, tt := range lisCases {
		t.Run(tt.name, func(t *testing.T) {
			a := LengthOfLIS(tt.nums)
			b := LengthOfLISBinary(tt.nums)
			if a != b {
				t.Errorf("两种实现不一致: nums=%v, DP=%d, Binary=%d", tt.nums, a, b)
			}
		})
	}
}

func BenchmarkLengthOfLIS(b *testing.B) {
	nums := make([]int, 2500)
	for i := range nums {
		nums[i] = (i*7 + 3) % 10000
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LengthOfLIS(nums)
	}
}

func BenchmarkLengthOfLISBinary(b *testing.B) {
	nums := make([]int, 2500)
	for i := range nums {
		nums[i] = (i*7 + 3) % 10000
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LengthOfLISBinary(nums)
	}
}
