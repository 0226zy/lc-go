package uglynumberii

import "testing"

var nthUglyNumberCases = []struct {
	name string
	n    int
	want int
}{
	{name: "第1个丑数", n: 1, want: 1},
	{name: "第2个丑数", n: 2, want: 2},
	{name: "第5个丑数", n: 5, want: 5},
	{name: "第6个丑数", n: 6, want: 6},
	{name: "示例：第10个丑数", n: 10, want: 12},
	{name: "第11个丑数", n: 11, want: 15},
	{name: "第15个丑数", n: 15, want: 24},
	{name: "第20个丑数", n: 20, want: 36},
	{name: "第100个丑数", n: 100, want: 1536},
	{name: "约束上限：第1690个丑数", n: 1690, want: 2123366400},
}

func TestNthUglyNumber(t *testing.T) {
	for _, tt := range nthUglyNumberCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := NthUglyNumber(tt.n); got != tt.want {
				t.Errorf("NthUglyNumber(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkNthUglyNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NthUglyNumber(1690)
	}
}
