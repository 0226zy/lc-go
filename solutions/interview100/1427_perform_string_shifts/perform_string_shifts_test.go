package performstringshifts

import "testing"

func TestStringShift(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		shift [][]int
		want  string
	}{
		// LeetCode 官方示例
		{"示例1: abc 左移1右移2", "abc", [][]int{{0, 1}, {1, 2}}, "cab"},
		{"示例2: abcdefg 多次移动", "abcdefg", [][]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}}, "efgabcd"},

		// 边界：单字符怎么移都不变
		{"单字符", "a", [][]int{{0, 5}, {1, 7}}, "a"},

		// 边界：净位移为 0
		{"左右恰好抵消", "hello", [][]int{{0, 3}, {1, 3}}, "hello"},
		{"移动位数为长度整数倍", "abcd", [][]int{{0, 8}}, "abcd"},

		// 边界：大位移取模
		{"位移超过长度", "abcde", [][]int{{1, 12}}, "deabc"},

		// 典型场景：纯左移与纯右移
		{"纯左移", "abcdef", [][]int{{0, 2}}, "cdefab"},
		{"纯右移", "abcdef", [][]int{{1, 2}}, "efabcd"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringShift(tt.s, tt.shift); got != tt.want {
				t.Errorf("StringShift(%q, %v) = %q, want %q", tt.s, tt.shift, got, tt.want)
			}
		})
	}
}

func BenchmarkStringShift(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
	shift := make([][]int, 0, 100)
	for i := 0; i < 100; i++ {
		shift = append(shift, []int{i % 2, 100})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringShift(s, shift)
	}
}
