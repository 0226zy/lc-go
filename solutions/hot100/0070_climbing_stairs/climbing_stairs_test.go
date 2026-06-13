package climbingstairs

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"n=1", 1, 1},
		{"n=2", 2, 2},
		{"n=3 官方示例2", 3, 3},
		{"n=4", 4, 5},
		{"n=5", 5, 8},
		{"n=10", 10, 89},
		{"n=45 上界", 45, 1836311903},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClimbStairs(tt.n); got != tt.want {
				t.Errorf("ClimbStairs(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestClimbStairsMatrix(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"n=1", 1, 1},
		{"n=2", 2, 2},
		{"n=3 官方示例2", 3, 3},
		{"n=4", 4, 5},
		{"n=5", 5, 8},
		{"n=10", 10, 89},
		{"n=45 上界", 45, 1836311903},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClimbStairsMatrix(tt.n); got != tt.want {
				t.Errorf("ClimbStairsMatrix(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkClimbStairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClimbStairs(45)
	}
}

func BenchmarkClimbStairsMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClimbStairsMatrix(45)
	}
}
