package trappingrainwater

import "testing"

var trapCases = []struct {
	name   string
	height []int
	want   int
}{
	// LeetCode 官方示例
	{name: "示例1", height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, want: 6},
	{name: "示例2", height: []int{4, 2, 0, 3, 2, 5}, want: 9},

	// 边界：柱子太少接不到水
	{name: "空数组", height: []int{}, want: 0},
	{name: "单柱", height: []int{1}, want: 0},
	{name: "两柱", height: []int{1, 2}, want: 0},
	{name: "两柱等高", height: []int{3, 3}, want: 0},

	// 边界：平地
	{name: "全为0", height: []int{0, 0, 0, 0}, want: 0},
	{name: "全等高", height: []int{2, 2, 2, 2}, want: 0},

	// 边界：单调数组（无凹槽）
	{name: "严格递增", height: []int{0, 1, 2, 3}, want: 0},
	{name: "严格递减", height: []int{3, 2, 1, 0}, want: 0},

	// 边界：单个凹槽
	{name: "单凹槽", height: []int{2, 0, 2}, want: 2},
	{name: "单凹槽更高", height: []int{3, 0, 0, 3}, want: 6},
	{name: "深槽", height: []int{5, 0, 0, 0, 5}, want: 15},

	// 边界：凹槽不对称
	{name: "左高右低", height: []int{5, 2, 1, 2, 1, 5}, want: 14},
	{name: "左低右高", height: []int{1, 4, 2, 5, 6, 3}, want: 2},

	// 边界：多凹槽
	{name: "多凹槽", height: []int{3, 0, 1, 0, 2}, want: 5},
}

func TestTrap(t *testing.T) {
	for _, tt := range trapCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trap(tt.height); got != tt.want {
				t.Errorf("Trap(%v) = %d, want %d", tt.height, got, tt.want)
			}
		})
	}
}

func TestTrapAlternative(t *testing.T) {
	for _, tt := range trapCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrapAlternative(tt.height); got != tt.want {
				t.Errorf("TrapAlternative(%v) = %d, want %d", tt.height, got, tt.want)
			}
		})
	}
}

var benchHeights = []struct {
	name   string
	height []int
}{
	{"len=12", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
	{"len=100", generateHeights(100)},
	{"len=1000", generateHeights(1000)},
	{"len=10000", generateHeights(10000)},
}

func BenchmarkTrap(b *testing.B) {
	for _, bm := range benchHeights {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Trap(bm.height)
			}
		})
	}
}

func BenchmarkTrapAlternative(b *testing.B) {
	for _, bm := range benchHeights {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				TrapAlternative(bm.height)
			}
		})
	}
}

// generateHeights 生成长度为 n 的伪随机柱状图高度
func generateHeights(n int) []int {
	height := make([]int, n)
	for i := 0; i < n; i++ {
		height[i] = (i*31 + 7) % 20
	}
	return height
}
