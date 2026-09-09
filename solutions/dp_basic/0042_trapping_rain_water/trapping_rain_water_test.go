package trappingrainwater

import "testing"

var trapCases = []struct {
	name   string
	height []int
	want   int
}{
	{name: "示例1", height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, want: 6},
	{name: "示例2", height: []int{4, 2, 0, 3, 2, 5}, want: 9},
	{name: "空数组", height: []int{}, want: 0},
	{name: "单根柱子", height: []int{5}, want: 0},
	{name: "两根柱子", height: []int{3, 1}, want: 0},
	{name: "递增无积水", height: []int{1, 2, 3, 4}, want: 0},
	{name: "递减无积水", height: []int{4, 3, 2, 1}, want: 0},
	{name: "凹槽", height: []int{3, 0, 3}, want: 3},
	{name: "全相同高度", height: []int{2, 2, 2}, want: 0},
	{name: "深槽", height: []int{5, 0, 0, 0, 5}, want: 15},
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

var benchHeights = func() []int {
	h := make([]int, 2000)
	for i := range h {
		h[i] = (i*7 + 13) % 97
	}
	return h
}()

func BenchmarkTrap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Trap(benchHeights)
	}
}

func BenchmarkTrapAlternative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrapAlternative(benchHeights)
	}
}
