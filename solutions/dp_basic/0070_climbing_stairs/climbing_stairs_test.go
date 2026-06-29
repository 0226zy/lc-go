package climbingstairs

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "1阶",
			n:    1,
			want: 1,
		},
		{
			name: "2阶",
			n:    2,
			want: 2,
		},
		{
			name: "3阶",
			n:    3,
			want: 3,
		},
		{
			name: "4阶",
			n:    4,
			want: 5,
		},
		{
			name: "5阶",
			n:    5,
			want: 8,
		},
		{
			name: "10阶",
			n:    10,
			want: 89,
		},
		{
			name: "20阶",
			n:    20,
			want: 10946,
		},
		{
			name: "45阶（约束上限）",
			n:    45,
			want: 1836311903,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClimbStairs(tt.n)
			if got != tt.want {
				t.Errorf("ClimbStairs(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkClimbStairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClimbStairs(45)
	}
}
