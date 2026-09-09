package longestincreasingsubsequence

import (
	"math/rand"
	"testing"
)

var lengthOfLISCases = []struct {
	name string
	nums []int
	want int
}{
	{name: "示例1", nums: []int{10, 9, 2, 5, 3, 7, 101, 18}, want: 4},
	{name: "示例2", nums: []int{0, 1, 0, 3, 2, 3}, want: 4},
	{name: "示例3：全部相等", nums: []int{7, 7, 7, 7, 7, 7, 7}, want: 1},
	{name: "单元素", nums: []int{1}, want: 1},
	{name: "严格递增", nums: []int{1, 2, 3, 4, 5}, want: 5},
	{name: "严格递减", nums: []int{5, 4, 3, 2, 1}, want: 1},
	{name: "含负数", nums: []int{-2, -1, -3, 0, -1, 2}, want: 4},
	{name: "最长子序列不在末尾", nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6}, want: 6},
}

func TestLengthOfLIS(t *testing.T) {
	for _, tt := range lengthOfLISCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLIS(tt.nums); got != tt.want {
				t.Errorf("LengthOfLIS(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestLengthOfLISAlternative(t *testing.T) {
	for _, tt := range lengthOfLISCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLISAlternative(tt.nums); got != tt.want {
				t.Errorf("LengthOfLISAlternative(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

var benchmarkNums = func() []int {
	// 构造 2500 个随机数作为压力场景
	r := rand.New(rand.NewSource(42))
	nums := make([]int, 2500)
	for i := range nums {
		nums[i] = r.Intn(20001) - 10000
	}
	return nums
}()

func BenchmarkLengthOfLIS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LengthOfLIS(benchmarkNums)
	}
}

func BenchmarkLengthOfLISAlternative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LengthOfLISAlternative(benchmarkNums)
	}
}
