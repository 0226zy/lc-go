package uniquepaths

import "testing"

var uniquePathsCases = []struct {
	name string
	m    int
	n    int
	want int
}{
	{name: "示例1：3x7", m: 3, n: 7, want: 28},
	{name: "示例2：3x2", m: 3, n: 2, want: 3},
	{name: "单格1x1", m: 1, n: 1, want: 1},
	{name: "单行1x5", m: 1, n: 5, want: 1},
	{name: "单列5x1", m: 5, n: 1, want: 1},
	{name: "方阵3x3", m: 3, n: 3, want: 6},
	{name: "对称7x3", m: 7, n: 3, want: 28},
	{name: "较大19x13", m: 19, n: 13, want: 86493225},
	{name: "较大23x12", m: 23, n: 12, want: 193536720},
}

func TestUniquePaths(t *testing.T) {
	for _, tt := range uniquePathsCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniquePaths(tt.m, tt.n); got != tt.want {
				t.Errorf("UniquePaths(%d, %d) = %d, want %d", tt.m, tt.n, got, tt.want)
			}
		})
	}
}

func TestUniquePathsAlternative(t *testing.T) {
	for _, tt := range uniquePathsCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniquePathsAlternative(tt.m, tt.n); got != tt.want {
				t.Errorf("UniquePathsAlternative(%d, %d) = %d, want %d", tt.m, tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkUniquePaths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UniquePaths(23, 12)
	}
}

func BenchmarkUniquePathsAlternative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UniquePathsAlternative(23, 12)
	}
}
