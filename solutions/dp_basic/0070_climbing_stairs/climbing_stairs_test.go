package climbingstairs

import "testing"

var climbStairsCases = []struct {
	name string
	n    int
	want int
}{
	{name: "1阶", n: 1, want: 1},
	{name: "2阶", n: 2, want: 2},
	{name: "3阶", n: 3, want: 3},
	{name: "4阶", n: 4, want: 5},
	{name: "5阶", n: 5, want: 8},
	{name: "10阶", n: 10, want: 89},
	{name: "20阶", n: 20, want: 10946},
	{name: "45阶（约束上限）", n: 45, want: 1836311903},
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
