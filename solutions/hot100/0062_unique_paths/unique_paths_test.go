package uniquepaths

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{
			name: "示例1：3x7网格",
			m:    3,
			n:    7,
			want: 28,
		},
		{
			name: "示例2：3x2网格",
			m:    3,
			n:    2,
			want: 3,
		},
		{
			name: "边界：1x1网格",
			m:    1,
			n:    1,
			want: 1,
		},
		{
			name: "边界：1x10网格",
			m:    1,
			n:    10,
			want: 1,
		},
		{
			name: "边界：10x1网格",
			m:    10,
			n:    1,
			want: 1,
		},
		{
			name: "正方形小网格",
			m:    2,
			n:    2,
			want: 2,
		},
		{
			name: "7x3网格",
			m:    7,
			n:    3,
			want: 28,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UniquePaths(tt.m, tt.n)
			if got != tt.want {
				t.Errorf("UniquePaths(%d, %d) = %d, want %d", tt.m, tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkUniquePaths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UniquePaths(3, 7)
	}
}
