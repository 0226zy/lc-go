package trappingrainwater

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		// LeetCode 官方示例
		{"示例1", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{"示例2", []int{4, 2, 0, 3, 2, 5}, 9},

		// 边界：柱子太少接不到水
		{"空数组", []int{}, 0},
		{"单柱", []int{1}, 0},
		{"两柱", []int{1, 2}, 0},
		{"两柱等高", []int{3, 3}, 0},

		// 边界：平地
		{"全为0", []int{0, 0, 0, 0}, 0},
		{"全等高", []int{2, 2, 2, 2}, 0},

		// 边界：单调数组（无凹槽）
		{"严格递增", []int{0, 1, 2, 3}, 0},
		{"严格递减", []int{3, 2, 1, 0}, 0},

		// 边界：单个凹槽
		{"单凹槽", []int{2, 0, 2}, 2},
		{"单凹槽更高", []int{3, 0, 0, 3}, 6},

		// 边界：凹槽不对称
		{"左高右低", []int{5, 2, 1, 2, 1, 5}, 14},
		{"左低右高", []int{1, 4, 2, 5, 6, 3}, 2},

		// 边界：多凹槽
		{"多凹槽", []int{3, 0, 1, 0, 2}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trap(tt.height); got != tt.want {
				t.Errorf("Trap(%v) = %d, want %d", tt.height, got, tt.want)
			}
		})
	}
}

// TestTrapAllMethods 校验三种解法结果一致
func TestTrapAllMethods(t *testing.T) {
	tests := []struct {
		name   string
		height []int
	}{
		{"官方示例1", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
		{"官方示例2", []int{4, 2, 0, 3, 2, 5}},
		{"随机数组", []int{4, 2, 0, 3, 2, 5, 2, 1, 3, 0, 1}},
		{"V形", []int{5, 1, 1, 1, 5}},
		{"W形", []int{5, 1, 4, 1, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, b, c := Trap(tt.height), TrapDP(tt.height), TrapStack(tt.height)
			if a != b || b != c {
				t.Errorf("解法结果不一致: 双指针=%d, DP=%d, 单调栈=%d", a, b, c)
			}
		})
	}
}

func BenchmarkTrap(b *testing.B) {
	benchmarks := []struct {
		name   string
		height []int
	}{
		{"len=12", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
		{"len=100", generateHeights(100)},
		{"len=1000", generateHeights(1000)},
		{"len=10000", generateHeights(10000)},
		{"len=100000", generateHeights(100000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.Run("双指针", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					Trap(bm.height)
				}
			})
			b.Run("动态规划", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					TrapDP(bm.height)
				}
			})
			b.Run("单调栈", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					TrapStack(bm.height)
				}
			})
		})
	}
}

// generateHeights 生成长度为 n 的随机柱状图高度
func generateHeights(n int) []int {
	height := make([]int, n)
	for i := 0; i < n; i++ {
		height[i] = (i*31 + 7) % 20
	}
	return height
}
