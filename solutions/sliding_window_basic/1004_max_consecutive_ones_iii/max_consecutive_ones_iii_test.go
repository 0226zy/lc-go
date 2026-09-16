package longestones

import "testing"

func TestLongestOnes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		// LeetCode 官方示例
		{
			"示例1: 翻转2个0得到长度6",
			[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			2,
			6,
		},
		{
			"示例2: 翻转3个0得到长度10",
			[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			3,
			10,
		},

		// 边界：空输入
		{
			"空数组",
			[]int{},
			2,
			0,
		},

		// 边界：单元素
		{
			"单元素为1",
			[]int{1},
			0,
			1,
		},
		{
			"单元素为0且k为0",
			[]int{0},
			0,
			0,
		},
		{
			"单元素为0且k为1",
			[]int{0},
			1,
			1,
		},

		// 边界：k 为 0，等价于求原数组最长连续 1
		{
			"k为0时求最长连续1",
			[]int{1, 1, 0, 1, 1, 1, 0, 1},
			0,
			3,
		},

		// 边界：全 1 数组，无需翻转
		{
			"全1数组",
			[]int{1, 1, 1, 1, 1},
			0,
			5,
		},

		// 边界：全 0 数组，答案取决于 k
		{
			"全0数组且k不足",
			[]int{0, 0, 0, 0},
			2,
			2,
		},
		{
			"全0数组且k足够",
			[]int{0, 0, 0},
			5,
			3,
		},

		// 极端：k 等于数组长度，整个数组都可翻转为 1
		{
			"k等于数组长度",
			[]int{0, 1, 0, 1},
			4,
			4,
		},

		// 最优窗口在数组末尾
		{
			"最优窗口在末尾",
			[]int{0, 0, 1, 1, 1, 1},
			2,
			6,
		},

		// 最优窗口在数组开头
		{
			"最优窗口在开头",
			[]int{1, 1, 0, 0, 0, 1},
			1,
			3,
		},

		// 1 和 0 交替出现
		{
			"01交替",
			[]int{0, 1, 0, 1, 0, 1},
			2,
			5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestOnes(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("LongestOnes(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

// makeAlternating 生成长度为 n 的 0/1 交替数组
func makeAlternating(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i % 2
	}
	return nums
}

func BenchmarkLongestOnes(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
		k    int
	}{
		{"长度1000交替数组k=10", 1000, 10},
		{"长度10000交替数组k=100", 10000, 100},
		{"长度100000交替数组k=1000", 100000, 1000},
	}

	for _, bm := range benchmarks {
		nums := makeAlternating(bm.n)
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LongestOnes(nums, bm.k)
			}
		})
	}
}
