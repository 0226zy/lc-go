package jumpgameii

import "testing"

func TestJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: [2,3,1,1,4]", []int{2, 3, 1, 1, 4}, 2},
		{"示例2: [2,3,0,1,4]", []int{2, 3, 0, 1, 4}, 2},

		// 边界：单元素，无需跳跃
		{"单元素", []int{0}, 0},
		{"单元素非0", []int{5}, 0},

		// 边界：两步恰达
		{"两步到达", []int{1, 1}, 1},
		{"一次大跳直达", []int{4, 0, 0, 0, 0}, 1},

		// 边界：全为1，只能一步一步跳
		{"全为1", []int{1, 1, 1, 1, 1}, 4},

		// 边界：中间有0需要绕行
		{"绕过中间的0", []int{3, 2, 1, 0, 4}, 2},
		{"连续0", []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}, 3},

		// 边界：递减步数
		{"递减步数", []int{4, 3, 2, 1, 0}, 1},

		// 边界：先小后大
		{"先小后大", []int{1, 2, 1, 1, 1}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Jump(tt.nums); got != tt.want {
				t.Errorf("Jump(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkJump(b *testing.B) {
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
				Jump(bm.nums)
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
