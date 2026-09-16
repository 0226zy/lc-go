package longestsubarray

import "testing"

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		// LeetCode 官方示例
		{
			"示例1: 删掉中间的0",
			[]int{1, 1, 0, 1},
			3,
		},
		{
			"示例2: 多个0间隔",
			[]int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			5,
		},
		{
			"示例3: 全1必须删一个",
			[]int{1, 1, 1},
			2,
		},

		// 边界：空数组
		{
			"空数组",
			[]int{},
			0,
		},

		// 边界：单元素
		{
			"单元素1: 删掉后为空",
			[]int{1},
			0,
		},
		{
			"单元素0: 删掉后为空",
			[]int{0},
			0,
		},

		// 边界：全 0 数组，删掉一个 0 也没有 1
		{
			"全0数组",
			[]int{0, 0, 0},
			0,
		},

		// 边界：0 在数组两端
		{
			"0在开头",
			[]int{0, 1, 1, 1},
			3,
		},
		{
			"0在末尾",
			[]int{1, 1, 1, 0},
			3,
		},

		// 0 交替出现，窗口最多跨一个 0
		{
			"01交替",
			[]int{0, 1, 0, 1, 0, 1},
			2,
		},

		// 长段 1 之间只有一个 0，拼接后最长
		{
			"两段长1拼一个0",
			[]int{1, 1, 1, 1, 0, 1, 1, 1},
			7,
		},

		// 只有一个 1，其余全 0
		{
			"只有一个1",
			[]int{0, 0, 1, 0, 0},
			1,
		},

		// 连续多个 0 相邻，无法连接两段 1
		{
			"连续两个0无法拼接",
			[]int{1, 1, 0, 0, 1, 1, 1},
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestSubarray(tt.nums); got != tt.want {
				t.Errorf("LongestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

// makeAllOnes 生成长度为 n、全为 1 的数组
func makeAllOnes(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = 1
	}
	return nums
}

// makeAlternatingNums 生成长度为 n、01 交替的数组
func makeAlternatingNums(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i % 2
	}
	return nums
}

func BenchmarkLongestSubarray(b *testing.B) {
	benchmarks := []struct {
		name string
		nums []int
	}{
		{"1000个全1", makeAllOnes(1000)},
		{"10万个全1", makeAllOnes(100000)},
		{"10万个01交替", makeAlternatingNums(100000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LongestSubarray(bm.nums)
			}
		})
	}
}
