package perfectsquares

import "testing"

func TestNumSquares(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"官方示例1 n=12", 12, 3},
		{"官方示例2 n=13", 13, 2},
		{"下界 n=1", 1, 1},
		{"完全平方数 n=4", 4, 1},
		{"完全平方数 n=9", 9, 1},
		{"完全平方数 n=16", 16, 1},
		{"n=2 两个1", 2, 2},
		{"n=3 三个1", 3, 3},
		{"n=7 需4个", 7, 4},    // 7=4^0*(8*0+7) → 答案4
		{"n=15 需4个", 15, 4},  // 15=4^0*(8*1+7) → 答案4
		{"n=23 需4个", 23, 4},  // 23=4^0*(8*2+7) → 答案4
		{"n=28 需4个", 28, 4},  // 28=4^1*(8*0+7) → 答案4
		{"n=43 需3个", 43, 3},  // 43=25+9+9
		{"n=50 需2个", 50, 2},  // 50=25+25
		{"n=100 完全平方", 100, 1},
		{"上界 n=9999", 9999, 4}, // 9999=4^0*(8*1249+7)
		{"上界 n=10000", 10000, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumSquares(tt.n); got != tt.want {
				t.Errorf("NumSquares(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestNumSquaresMath(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"官方示例1 n=12", 12, 3},
		{"官方示例2 n=13", 13, 2},
		{"下界 n=1", 1, 1},
		{"完全平方数 n=4", 4, 1},
		{"n=2 两个1", 2, 2},
		{"n=7 需4个", 7, 4},
		{"n=15 需4个", 15, 4},
		{"n=28 需4个", 28, 4},
		{"n=43 需3个", 43, 3},
		{"n=50 需2个", 50, 2},
		{"上界 n=9999", 9999, 4},
		{"上界 n=10000", 10000, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumSquaresMath(tt.n); got != tt.want {
				t.Errorf("NumSquaresMath(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestNumSquaresBFS(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"官方示例1 n=12", 12, 3},
		{"官方示例2 n=13", 13, 2},
		{"下界 n=1", 1, 1},
		{"完全平方数 n=4", 4, 1},
		{"n=2 两个1", 2, 2},
		{"n=7 需4个", 7, 4},
		{"n=15 需4个", 15, 4},
		{"n=43 需3个", 43, 3},
		{"n=50 需2个", 50, 2},
		{"上界 n=9999", 9999, 4},
		{"上界 n=10000", 10000, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumSquaresBFS(tt.n); got != tt.want {
				t.Errorf("NumSquaresBFS(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

// TestConsistency 验证三种实现结果一致
func TestConsistency(t *testing.T) {
	for n := 1; n <= 1000; n++ {
		a := NumSquares(n)
		b := NumSquaresMath(n)
		c := NumSquaresBFS(n)
		if a != b || b != c {
			t.Errorf("三种实现结果不一致: n=%d, DP=%d, Math=%d, BFS=%d", n, a, b, c)
		}
	}
}

func BenchmarkNumSquares(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumSquares(10000)
	}
}

func BenchmarkNumSquaresMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumSquaresMath(10000)
	}
}

func BenchmarkNumSquaresBFS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumSquaresBFS(10000)
	}
}
