package longestconsecutivesequence

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"示例1", []int{100, 4, 200, 1, 3, 2}, 4},
		{"示例2", []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{"空数组", []int{}, 0},
		{"单元素", []int{1}, 1},
		{"全部相同", []int{1, 1, 1}, 1},
		{"无连续", []int{1, 3, 5, 7}, 1},
		{"全连续", []int{1, 2, 3, 4, 5}, 5},
		{"负数连续", []int{-1, -2, -3, 0, 1}, 5},
		{"含重复", []int{1, 2, 0, 1}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestConsecutive(tt.nums)
			if got != tt.want {
				t.Errorf("LongestConsecutive(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkLongestConsecutive(b *testing.B) {
	nums := []int{100, 4, 200, 1, 3, 2}
	for i := 0; i < b.N; i++ {
		LongestConsecutive(nums)
	}
}
