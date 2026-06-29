package fibonaccinumber

import "testing"

func TestFib(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"F(0)=0", 0, 0},
		{"F(1)=1", 1, 1},
		{"F(2)=1", 2, 1},
		{"F(3)=2", 3, 2},
		{"F(4)=3", 4, 3},
		{"F(5)=5", 5, 5},
		{"F(10)=55", 10, 55},
		{"F(20)=6765", 20, 6765},
		{"F(30)=832040", 30, 832040},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Fib(tt.n)
			if got != tt.want {
				t.Errorf("Fib(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(30)
	}
}
