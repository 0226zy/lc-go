package trappingrainwater

import (
	"math/rand"
	"testing"
)

// 测试用例：函数名 -> 表驱动测试
var cases = []struct {
	name   string
	height []int
	want   int
}{
	{"示例1", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
	{"示例2", []int{4, 2, 0, 3, 2, 5}, 9},
	{"空数组", []int{}, 0},
	{"单根柱子", []int{5}, 0},
	{"两根柱子", []int{3, 7}, 0},
	{"单调递增", []int{1, 2, 3, 4, 5}, 0},
	{"单调递减", []int{5, 4, 3, 2, 1}, 0},
	{"全部等高", []int{3, 3, 3, 3}, 0},
	{"V字形", []int{3, 0, 3}, 3},
	{"深V字形", []int{5, 0, 0, 0, 5}, 15},
	{"波浪形", []int{2, 0, 2, 0, 2}, 4},
	{"全零", []int{0, 0, 0}, 0},
	{"中间凹陷", []int{4, 0, 2, 0, 4}, 10},
}

func TestTrap(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trap(tt.height); got != tt.want {
				t.Errorf("Trap(%v) = %d, 期望 %d", tt.height, got, tt.want)
			}
		})
	}
}

func TestTrapDP(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrapDP(tt.height); got != tt.want {
				t.Errorf("TrapDP(%v) = %d, 期望 %d", tt.height, got, tt.want)
			}
		})
	}
}

func TestTrapStack(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrapStack(tt.height); got != tt.want {
				t.Errorf("TrapStack(%v) = %d, 期望 %d", tt.height, got, tt.want)
			}
		})
	}
}

// TestConsistency 随机数据一致性测试：三种解法结果必须一致
func TestConsistency(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	for i := 0; i < 100; i++ {
		n := rng.Intn(200) + 1
		height := make([]int, n)
		for j := range height {
			height[j] = rng.Intn(100)
		}
		a, b, c := Trap(height), TrapDP(height), TrapStack(height)
		if a != b || b != c {
			t.Fatalf("结果不一致: Trap=%d, TrapDP=%d, TrapStack=%d, 输入=%v", a, b, c, height)
		}
	}
}

// 大规模随机数据，用于 Benchmark
func genLargeInput(n int) []int {
	height := make([]int, n)
	rng := rand.New(rand.NewSource(42))
	for i := range height {
		height[i] = rng.Intn(100000)
	}
	return height
}

func BenchmarkTrap(b *testing.B) {
	height := genLargeInput(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Trap(height)
	}
}

func BenchmarkTrapDP(b *testing.B) {
	height := genLargeInput(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TrapDP(height)
	}
}

func BenchmarkTrapStack(b *testing.B) {
	height := genLargeInput(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TrapStack(height)
	}
}
