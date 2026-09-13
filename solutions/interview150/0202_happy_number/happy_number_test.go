package happynumber

import "testing"

func TestIsHappy(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		// LeetCode 官方示例
		{"示例1: 19 是快乐数", 19, true},
		{"示例2: 2 不是快乐数", 2, false},

		// 边界：1 本身就是快乐数
		{"1 是快乐数", 1, true},
		{"7 是快乐数", 7, true},
		{"10 是快乐数", 10, true},

		// 已知非快乐数
		{"4 进入环不是快乐数", 4, false},
		{"3 不是快乐数", 3, false},
		{"11 不是快乐数", 11, false},

		// 较大输入
		{"100 是快乐数", 100, true},
		{"int32最大值附近", 2147483647, false},
		{"82 在 19 的路径上是快乐数", 82, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHappy(tt.n); got != tt.want {
				t.Errorf("IsHappy(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestGetNext(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{19, 82},
		{82, 68},
		{68, 100},
		{100, 1},
		{1, 1},
		{7, 49},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := getNext(tt.n); got != tt.want {
				t.Errorf("getNext(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkIsHappy(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"快乐数19", 19},
		{"非快乐数2", 2},
		{"快乐数1", 1},
		{"大数2147483646", 2147483646},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsHappy(bm.n)
			}
		})
	}
}
