package jumpgameii

import "testing"

var jumpCases = []struct {
	name string
	nums []int
	want int
}{
	{name: "示例1", nums: []int{2, 3, 1, 1, 4}, want: 2},
	{name: "示例2", nums: []int{2, 3, 0, 1, 4}, want: 2},
	{name: "单元素", nums: []int{0}, want: 0},
	{name: "一步直达", nums: []int{5, 1, 1, 1, 1}, want: 1},
	{name: "每步只能走一格", nums: []int{1, 1, 1, 1}, want: 3},
	{name: "全是大步", nums: []int{9, 8, 7, 6, 5}, want: 1},
	{name: "需要先小跳再大跳", nums: []int{1, 3, 2}, want: 2},
	{name: "含零但不挡路", nums: []int{3, 0, 0, 2, 1}, want: 2},
}

func TestJump(t *testing.T) {
	for _, tt := range jumpCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := Jump(tt.nums); got != tt.want {
				t.Errorf("Jump(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestJumpAlternative(t *testing.T) {
	for _, tt := range jumpCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := JumpAlternative(tt.nums); got != tt.want {
				t.Errorf("JumpAlternative(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

var benchNums = func() []int {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = (i*3+5)%7 + 1 // 步长 1~7，保证可达
	}
	return nums
}()

func BenchmarkJump(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Jump(benchNums)
	}
}

func BenchmarkJumpAlternative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JumpAlternative(benchNums)
	}
}
