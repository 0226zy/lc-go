package nthtribonumber

import "testing"

var tribonacciCases = []struct {
	name string
	n    int
	want int
}{
	{name: "T(0)=0", n: 0, want: 0},
	{name: "T(1)=1", n: 1, want: 1},
	{name: "T(2)=1", n: 2, want: 1},
	{name: "T(3)=2", n: 3, want: 2},
	{name: "T(4)=4（示例1）", n: 4, want: 4},
	{name: "T(5)=7", n: 5, want: 7},
	{name: "T(10)=149", n: 10, want: 149},
	{name: "T(25)=1389537（示例2）", n: 25, want: 1389537},
	{name: "T(37)=2082876103（约束上限）", n: 37, want: 2082876103},
}

func TestTribonacci(t *testing.T) {
	for _, tt := range tribonacciCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tribonacci(tt.n); got != tt.want {
				t.Errorf("Tribonacci(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestTribonacciOptimized(t *testing.T) {
	for _, tt := range tribonacciCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := TribonacciOptimized(tt.n); got != tt.want {
				t.Errorf("TribonacciOptimized(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkTribonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Tribonacci(37)
	}
}

func BenchmarkTribonacciOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TribonacciOptimized(37)
	}
}
