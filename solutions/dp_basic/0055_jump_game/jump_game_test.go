package jumpgame

import "testing"

var canJumpCases = []struct {
	name string
	nums []int
	want bool
}{
	{name: "示例1_可达", nums: []int{2, 3, 1, 1, 4}, want: true},
	{name: "示例2_被零挡住", nums: []int{3, 2, 1, 0, 4}, want: false},
	{name: "单元素", nums: []int{0}, want: true},
	{name: "一步直达", nums: []int{5, 0, 0, 0}, want: true},
	{name: "每步一格刚好到", nums: []int{1, 1, 1, 1}, want: true},
	{name: "起点为零且非终点", nums: []int{0, 1}, want: false},
	{name: "末尾的零不影响", nums: []int{2, 0, 0}, want: true},
	{name: "恰好卡在最后一个零", nums: []int{1, 0, 1, 0}, want: false},
}

func TestCanJump(t *testing.T) {
	for _, tt := range canJumpCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanJump(tt.nums); got != tt.want {
				t.Errorf("CanJump(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestCanJumpAlternative(t *testing.T) {
	for _, tt := range canJumpCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanJumpAlternative(tt.nums); got != tt.want {
				t.Errorf("CanJumpAlternative(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

var benchNums55 = func() []int {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = (i*3 + 1) % 5 // 步长 0~4
	}
	return nums
}()

func BenchmarkCanJump(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CanJump(benchNums55)
	}
}

func BenchmarkCanJumpAlternative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CanJumpAlternative(benchNums55)
	}
}
