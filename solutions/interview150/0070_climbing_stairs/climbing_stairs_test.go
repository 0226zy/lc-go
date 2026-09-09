package climbingstairs

import "testing"

var climbStairsCases = []struct {
	name string
	n    int
	want int
}{
	// LeetCode 官方示例
	{name: "示例1: n=2", n: 2, want: 2},
	{name: "示例2: n=3", n: 3, want: 3},
	{name: "补充示例: n=4", n: 4, want: 5},

	// 边界：最小输入
	{name: "边界: n=1", n: 1, want: 1},

	// 斐波那契数列递推验证
	{name: "递推: n=5", n: 5, want: 8},
	{name: "递推: n=6", n: 6, want: 13},
	{name: "递推: n=7", n: 7, want: 21},
	{name: "递推: n=10", n: 10, want: 89},
	{name: "递推: n=20", n: 20, want: 10946},

	// 压力场景：约束最大值 n=45，结果仍在 int 范围内
	{name: "压力: n=44", n: 44, want: 1134903170},
	{name: "压力: n=45（约束上限）", n: 45, want: 1836311903},
}

func TestClimbStairs(t *testing.T) {
	for _, tt := range climbStairsCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClimbStairs(tt.n); got != tt.want {
				t.Errorf("ClimbStairs(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestClimbStairsOptimized(t *testing.T) {
	for _, tt := range climbStairsCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClimbStairsOptimized(tt.n); got != tt.want {
				t.Errorf("ClimbStairsOptimized(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkClimbStairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClimbStairs(45)
	}
}

func BenchmarkClimbStairsOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClimbStairsOptimized(45)
	}
}
