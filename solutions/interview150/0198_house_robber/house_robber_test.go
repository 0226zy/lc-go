package houserobber

import "testing"

var robCases = []struct {
	name string
	nums []int
	want int
}{
	// LeetCode 官方示例
	{name: "示例1：[1,2,3,1]", nums: []int{1, 2, 3, 1}, want: 4},
	{name: "示例2：[2,7,9,3,1]", nums: []int{2, 7, 9, 3, 1}, want: 12},

	// 边界
	{name: "空数组", nums: []int{}, want: 0},
	{name: "单房屋", nums: []int{5}, want: 5},
	{name: "两房屋取大", nums: []int{1, 3}, want: 3},
	{name: "两房屋取左", nums: []int{4, 2}, want: 4},

	// 其他场景
	{name: "全相同", nums: []int{2, 2, 2, 2}, want: 4},
	{name: "递增", nums: []int{1, 2, 3, 4, 5}, want: 9},
	{name: "递减", nums: []int{5, 4, 3, 2, 1}, want: 9},
	{name: "隔一偷更优", nums: []int{100, 1, 100, 1, 100}, want: 300},
	{name: "连续大额中间小", nums: []int{2, 1, 1, 2}, want: 4},
	{name: "含零", nums: []int{0, 0, 0}, want: 0},
	{name: "首尾大中间小", nums: []int{10, 1, 1, 10}, want: 20},
}

func TestRob(t *testing.T) {
	for _, tt := range robCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rob(tt.nums); got != tt.want {
				t.Errorf("Rob(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestRobOptimized(t *testing.T) {
	for _, tt := range robCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := RobOptimized(tt.nums); got != tt.want {
				t.Errorf("RobOptimized(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func generateNums(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = (i*37)%400 + 1
	}
	return nums
}

func BenchmarkRob(b *testing.B) {
	nums := generateNums(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Rob(nums)
	}
}

func BenchmarkRobOptimized(b *testing.B) {
	nums := generateNums(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RobOptimized(nums)
	}
}
