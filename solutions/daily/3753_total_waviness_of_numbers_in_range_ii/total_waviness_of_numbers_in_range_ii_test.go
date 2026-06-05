package totalwaviness

import "testing"

func TestTotalWaviness(t *testing.T) {
	tests := []struct {
		name string
		num1 int
		num2 int
		want int64
	}{
		{
			name: "官方示例1",
			num1: 120,
			num2: 130,
			want: 3,
		},
		{
			name: "官方示例2",
			num1: 198,
			num2: 202,
			want: 3,
		},
		{
			name: "官方示例3",
			num1: 4848,
			num2: 4848,
			want: 2,
		},
		{
			name: "单元素区间无波动",
			num1: 1,
			num2: 1,
			want: 0,
		},
		{
			name: "两位数区间",
			num1: 10,
			num2: 20,
			want: 0,
		},
		{
			name: "大范围含多个峰谷",
			num1: 100,
			num2: 200,
			want: 63,
		},
		{
			name: "连续谷值",
			num1: 210,
			num2: 219,
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TotalWaviness(tt.num1, tt.num2)
			if got != tt.want {
				t.Errorf("TotalWaviness(%d, %d) = %d, want %d", tt.num1, tt.num2, got, tt.want)
			}
		})
	}
}

func BenchmarkTotalWaviness(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TotalWaviness(120, 130)
	}
}
