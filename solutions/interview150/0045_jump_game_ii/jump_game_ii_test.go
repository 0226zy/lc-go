package jumpgameii

import "testing"

var jumpCases = []struct {
	name string
	nums []int
	want int
}{
	// LeetCode 官方示例
	{name: "示例1: [2,3,1,1,4]", nums: []int{2, 3, 1, 1, 4}, want: 2},
	{name: "示例2: [2,3,0,1,4]", nums: []int{2, 3, 0, 1, 4}, want: 2},

	// 边界：单元素，无需跳跃
	{name: "单元素", nums: []int{0}, want: 0},
	{name: "单元素非0", nums: []int{5}, want: 0},

	// 边界：一步或两步恰达
	{name: "两步到达", nums: []int{1, 1}, want: 1},
	{name: "一次大跳直达", nums: []int{4, 0, 0, 0, 0}, want: 1},
	{name: "全是大步", nums: []int{9, 8, 7, 6, 5}, want: 1},

	// 边界：全为1，只能一步一步跳
	{name: "全为1", nums: []int{1, 1, 1, 1, 1}, want: 4},

	// 边界：中间有0需要绕行（原用例 [3,2,1,0,4] 实际不可达，违反题目保证，已移除）
	{name: "含零但不挡路", nums: []int{3, 0, 0, 2, 1}, want: 2},

	// 边界：递减步数
	{name: "递减步数", nums: []int{4, 3, 2, 1, 0}, want: 1},

	// 边界：先小后大
	{name: "先小后大", nums: []int{1, 2, 1, 1, 1}, want: 3},
	{name: "需要先小跳再大跳", nums: []int{1, 3, 2}, want: 2},
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

var benchNums = []struct {
	name string
	nums []int
}{
	{"len=10", []int{2, 3, 1, 1, 4, 2, 1, 3, 1, 1}},
	{"len=1000", generateNums(1000)},
	{"len=10000", generateNums(10000)},
}

func BenchmarkJump(b *testing.B) {
	for _, bm := range benchNums {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Jump(bm.nums)
			}
		})
	}
}

func BenchmarkJumpAlternative(b *testing.B) {
	for _, bm := range benchNums {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				JumpAlternative(bm.nums)
			}
		})
	}
}

// generateNums 生成步数递减的跳跃数组，保证可到达终点
func generateNums(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = n - i
	}
	return nums
}
