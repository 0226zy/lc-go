package summaryranges

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []string
	}{
		// LeetCode 官方示例
		{"示例1: [0,1,2,4,5,7]", []int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{"示例2: [0,2,3,4,6,8,9]", []int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},

		// 边界：空数组
		{"空数组", []int{}, []string{}},

		// 边界：单元素
		{"单元素", []int{-1}, []string{"-1"}},

		// 边界：整个数组连续
		{"整体连续", []int{0, 1, 2, 3, 4}, []string{"0->4"}},
		{"整体不连续", []int{0, 2, 4, 6}, []string{"0", "2", "4", "6"}},

		// 边界：极值
		{"包含int32最小值", []int{-2147483648, -2147483647, 2147483647}, []string{"-2147483648->-2147483647", "2147483647"}},
		{"单元素int32最大值", []int{2147483647}, []string{"2147483647"}},

		// 边界：连续区间在开头/结尾
		{"开头连续", []int{1, 2, 3, 5}, []string{"1->3", "5"}},
		{"结尾连续", []int{1, 3, 4, 5}, []string{"1", "3->5"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SummaryRanges(tt.nums)
			if !utils.EqualStringSlice(got, tt.want) {
				t.Errorf("SummaryRanges(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkSummaryRanges(b *testing.B) {
	// 构造 n 个长度为 3 的连续区间
	generate := func(n int) []int {
		nums := make([]int, 0, 3*n)
		for i := 0; i < n; i++ {
			base := i * 4 // 区间之间留一个空档
			nums = append(nums, base, base+1, base+2)
		}
		return nums
	}

	benchmarks := []struct {
		name string
		nums []int
	}{
		{"len=10", generate(10)},
		{"len=100", generate(100)},
		{"len=1000", generate(1000)},
		{"len=10000", generate(10000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SummaryRanges(bm.nums)
			}
		})
	}
}
