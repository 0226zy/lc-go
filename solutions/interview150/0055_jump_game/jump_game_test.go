package jumpgame

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		// LeetCode 官方示例
		{"示例1: 可以到达", []int{2, 3, 1, 1, 4}, true},
		{"示例2: 被0挡住", []int{3, 2, 1, 0, 4}, false},

		// 边界：单元素（已在终点）
		{"单元素", []int{0}, true},

		// 边界：第一个就是0且长度大于1
		{"起点就是0", []int{0, 1}, false},

		// 边界：恰好一步到达
		{"恰好一步到终点", []int{1}, true},
		{"刚好够到终点", []int{1, 1, 1, 1}, true},

		// 边界：大跳数越过后面的0
		{"大跳越过0", []int{4, 2, 0, 0, 1, 1}, true},

		// 边界：中间有0但整体可达
		{"中间含0仍可达", []int{2, 0, 2, 0, 1}, true},

		// 边界：结尾是0不影响
		{"结尾是0", []int{2, 3, 1, 0}, true},

		// 边界：全为0且长度大于1
		{"全为0长度2", []int{0, 0}, false},

		// 边界：大数值一步直达
		{"极大跳跃值", []int{10, 0, 0, 0, 0, 0}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanJump(tt.nums); got != tt.want {
				t.Errorf("CanJump(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkCanJump(b *testing.B) {
	benchmarks := []struct {
		name string
		nums []int
	}{
		{"len=10", []int{2, 3, 1, 1, 4, 2, 1, 3, 1, 1}},
		{"len=1000", generateNums(1000)},
		{"len=10000", generateNums(10000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CanJump(bm.nums)
			}
		})
	}
}

// generateNums 生成步数递减的跳跃数组
func generateNums(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = n - i
	}
	return nums
}
