package jumpgame

import "testing"

var canJumpCases = []struct {
	name string
	nums []int
	want bool
}{
	// LeetCode 官方示例
	{name: "示例1: 可以到达", nums: []int{2, 3, 1, 1, 4}, want: true},
	{name: "示例2: 被0挡住", nums: []int{3, 2, 1, 0, 4}, want: false},

	// 边界：单元素（已在终点）
	{name: "单元素", nums: []int{0}, want: true},

	// 边界：第一个就是0且长度大于1
	{name: "起点就是0", nums: []int{0, 1}, want: false},

	// 边界：恰好一步到达
	{name: "恰好一步到终点", nums: []int{1}, want: true},
	{name: "刚好够到终点", nums: []int{1, 1, 1, 1}, want: true},
	{name: "一步直达", nums: []int{5, 0, 0, 0}, want: true},

	// 边界：大跳数越过后面的0
	{name: "大跳越过0", nums: []int{4, 2, 0, 0, 1, 1}, want: true},

	// 边界：中间有0但整体可达
	{name: "中间含0仍可达", nums: []int{2, 0, 2, 0, 1}, want: true},

	// 边界：结尾是0不影响
	{name: "结尾是0", nums: []int{2, 3, 1, 0}, want: true},

	// 边界：恰好卡在最后一个零
	{name: "恰好卡在最后一个零", nums: []int{1, 0, 1, 0}, want: false},

	// 边界：全为0且长度大于1
	{name: "全为0长度2", nums: []int{0, 0}, want: false},

	// 边界：大数值一步直达
	{name: "极大跳跃值", nums: []int{10, 0, 0, 0, 0, 0}, want: true},
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

// benchNums 生成 1000 长度的跳跃数组用于基准测试
var benchNums = func() []int {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = (i*3 + 1) % 5 // 步长 0~4
	}
	return nums
}()

func BenchmarkCanJump(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CanJump(benchNums)
	}
}

func BenchmarkCanJumpAlternative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CanJumpAlternative(benchNums)
	}
}
