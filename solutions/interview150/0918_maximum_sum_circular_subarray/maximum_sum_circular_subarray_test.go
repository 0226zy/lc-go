package maximumsumcircularsubarray

import "testing"

var maxSubarraySumCircularCases = []struct {
	name string
	nums []int
	want int
}{
	// LeetCode 官方示例
	{name: "示例1：不跨界", nums: []int{1, -2, 3, -2}, want: 3},
	{name: "示例2：跨界", nums: []int{5, -3, 5}, want: 10},
	{name: "示例3：全负取最大单元素", nums: []int{-3, -2, -3}, want: -2},

	// 边界：单元素
	{name: "单元素为正", nums: []int{3}, want: 3},
	{name: "单元素为负", nums: []int{-5}, want: -5},

	// 边界：全负数（不能返回 0）
	{name: "全负数", nums: []int{-2, -3, -1}, want: -1},
	{name: "全负数含0", nums: []int{-1, -2, 0, -3}, want: 0},
	{name: "全相同负数", nums: []int{-5, -5, -5}, want: -5},

	// 边界：全正数（最优就是整个数组，等价于不跨环）
	{name: "全正数", nums: []int{1, 2, 3}, want: 6},

	// 边界：跨环取两段
	{name: "跨环两段", nums: []int{3, -1, 2, -1}, want: 4},
	{name: "跨环两端各一个", nums: []int{1, -1, 1}, want: 2},
	{name: "最优在环尾和环首", nums: []int{8, -4, 3, -8, 8}, want: 16},
	{name: "官方样例跨界", nums: []int{3, -2, 2, -3}, want: 3},

	// 边界：普通 Kadane 更优
	{name: "中间一段最优", nums: []int{-2, 4, -5, 4, -5, 9, -20}, want: 9},
	{name: "含极大值", nums: []int{-1, 30000, -2, -3}, want: 30000},
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
	benchmarks := []struct {
		name string
		nums []int
	}{
		{"len=10", []int{5, -3, 5, -2, 1, -2, 3, -2, 4, -1}},
		{"len=1000", generateCircularNums(1000)},
		{"len=30000", generateCircularNums(30000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MaxSubarraySumCircular(bm.nums)
			}
		})
	}
}

func BenchmarkMaxSubarraySumCircularOptimized(b *testing.B) {
	benchmarks := []struct {
		name string
		nums []int
	}{
		{"len=10", []int{5, -3, 5, -2, 1, -2, 3, -2, 4, -1}},
		{"len=1000", generateCircularNums(1000)},
		{"len=30000", generateCircularNums(30000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MaxSubarraySumCircularOptimized(bm.nums)
			}
		})
	}
}

// generateCircularNums 生成长度为 n 的正负交错数组
func generateCircularNums(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			arr[i] = i % 101
		} else {
			arr[i] = -(i % 83)
		}
	}
	return arr
}
