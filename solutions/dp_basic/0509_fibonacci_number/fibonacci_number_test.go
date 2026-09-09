package fibonaccinumber

import "testing"

var fibCases = []struct {
	name string
	n    int
	want int
}{
	{name: "F(0)=0", n: 0, want: 0},
	{name: "F(1)=1", n: 1, want: 1},
	{name: "F(2)=1", n: 2, want: 1},
	{name: "F(3)=2", n: 3, want: 2},
	{name: "F(4)=3", n: 4, want: 3},
	{name: "F(5)=5", n: 5, want: 5},
	{name: "F(10)=55", n: 10, want: 55},
	{name: "F(20)=6765", n: 20, want: 6765},
	{name: "F(30)=832040（约束上限）", n: 30, want: 832040},
}

func TestFib(t *testing.T) {
	for _, tt := range fibCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib(tt.n); got != tt.want {
				t.Errorf("Fib(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestFibOptimized(t *testing.T) {
	for _, tt := range fibCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibOptimized(tt.n); got != tt.want {
				t.Errorf("FibOptimized(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(30)
	}
}

func BenchmarkFibOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibOptimized(30)
	}
}
