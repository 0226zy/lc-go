package bitwiseandofnumbersrange

import "testing"

var rangeBitwiseAndCases = []struct {
	name  string
	left  int
	right int
	want  int
}{
	// LeetCode 官方示例
	{name: "示例1：left=5,right=7", left: 5, right: 7, want: 4},
	{name: "示例2：left=0,right=0", left: 0, right: 0, want: 0},
	{name: "示例3：left=1,right=2147483647", left: 1, right: 2147483647, want: 0},

	// 边界
	{name: "单点区间", left: 7, right: 7, want: 7},
	{name: "相邻两数", left: 8, right: 9, want: 8},
	{name: "跨过2的幂", left: 7, right: 8, want: 0},
	{name: "同前缀区间", left: 12, right: 15, want: 12},
	{name: "left=0跨区间", left: 0, right: 1, want: 0},
	{name: "大数相邻", left: 2147483646, right: 2147483647, want: 2147483646},
	{name: "中等区间", left: 10, right: 20, want: 0},
}

func TestRangeBitwiseAnd(t *testing.T) {
	for _, tt := range rangeBitwiseAndCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := RangeBitwiseAnd(tt.left, tt.right); got != tt.want {
				t.Errorf("RangeBitwiseAnd(%d, %d) = %d, want %d", tt.left, tt.right, got, tt.want)
			}
		})
	}
}

func TestRangeBitwiseAndShift(t *testing.T) {
	for _, tt := range rangeBitwiseAndCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := RangeBitwiseAndShift(tt.left, tt.right); got != tt.want {
				t.Errorf("RangeBitwiseAndShift(%d, %d) = %d, want %d", tt.left, tt.right, got, tt.want)
			}
		})
	}
}

func BenchmarkRangeBitwiseAnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RangeBitwiseAnd(1, 2147483647)
	}
}

func BenchmarkRangeBitwiseAndShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RangeBitwiseAndShift(1, 2147483647)
	}
}
