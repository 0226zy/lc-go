package perfectsquares

import "testing"

var numSquaresCases = []struct {
	name string
	n    int
	want int
}{
	{name: "n=1（自身是平方数）", n: 1, want: 1},
	{name: "n=2", n: 2, want: 2},
	{name: "n=3", n: 3, want: 3},
	{name: "n=4（自身是平方数）", n: 4, want: 1},
	{name: "n=7（答案为4的特例）", n: 7, want: 4},
	{name: "示例1：n=12", n: 12, want: 3},
	{name: "示例2：n=13", n: 13, want: 2},
	{name: "n=43（答案为3）", n: 43, want: 3},
	{name: "n=10000（约束上限，100的平方）", n: 10000, want: 1},
	{name: "n=9999", n: 9999, want: 4},
}

func TestNumSquares(t *testing.T) {
	for _, tt := range numSquaresCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumSquares(tt.n); got != tt.want {
				t.Errorf("NumSquares(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestNumSquaresAlternative(t *testing.T) {
	for _, tt := range numSquaresCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumSquaresAlternative(tt.n); got != tt.want {
				t.Errorf("NumSquaresAlternative(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkNumSquares(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumSquares(10000)
	}
}

func BenchmarkNumSquaresAlternative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumSquaresAlternative(10000)
	}
}
