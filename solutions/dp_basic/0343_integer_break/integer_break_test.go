package integerbreak

import "testing"

var integerBreakCases = []struct {
	name string
	n    int
	want int
}{
	{name: "官方示例1最小值", n: 2, want: 1},
	{name: "3拆成1加2", n: 3, want: 2},
	{name: "4拆成2加2", n: 4, want: 4},
	{name: "5拆成2加3", n: 5, want: 6},
	{name: "官方示例2", n: 10, want: 36},
	{name: "约束上限", n: 58, want: 1549681956},
}

func TestIntegerBreak(t *testing.T) {
	for _, tt := range integerBreakCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntegerBreak(tt.n); got != tt.want {
				t.Errorf("IntegerBreak(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestIntegerBreakAlternative(t *testing.T) {
	for _, tt := range integerBreakCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntegerBreakAlternative(tt.n); got != tt.want {
				t.Errorf("IntegerBreakAlternative(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkIntegerBreak(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntegerBreak(58)
	}
}

func BenchmarkIntegerBreakAlternative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntegerBreakAlternative(58)
	}
}
