package totalwaviness

import "testing"

func TestTotalWaviness(t *testing.T) {
	tests := []struct {
		name string
		num1 int64
		num2 int64
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
			name: "数字0",
			num1: 0,
			num2: 0,
			want: 0,
		},
		{
			name: "仅峰",
			num1: 132,
			num2: 132,
			want: 1,
		},
		{
			name: "仅谷",
			num1: 213,
			num2: 213,
			want: 1,
		},
		{
			name: "相等数位",
			num1: 1223,
			num2: 1223,
			want: 0,
		},
		{
			name: "连续谷值",
			num1: 210,
			num2: 219,
			want: 8,
		},
		{
			name: "多个峰谷",
			num1: 121,
			num2: 121,
			want: 1,
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

func TestTotalWavinessCompareWithBruteForce(t *testing.T) {
	// 在小范围内对比数位DP解法与暴力解法的结果，确保正确性
	tests := []struct {
		name string
		num1 int64
		num2 int64
	}{
		{"0-100", 0, 100},
		{"100-200", 100, 200},
		{"500-600", 500, 600},
		{"999-1001", 999, 1001},
		{"1000-1100", 1000, 1100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dpResult := TotalWaviness(tt.num1, tt.num2)
			bruteResult := TotalWavinessBruteForce(tt.num1, tt.num2)
			if dpResult != bruteResult {
				t.Errorf("TotalWaviness(%d, %d) DP=%d vs BruteForce=%d, want equal", tt.num1, tt.num2, dpResult, bruteResult)
			}
		})
	}
}

func TestGetDigits(t *testing.T) {
	tests := []struct {
		name string
		num  int64
		want []int
	}{
		{
			name: "数字0",
			num:  0,
			want: []int{0},
		},
		{
			name: "数字5",
			num:  5,
			want: []int{5},
		},
		{
			name: "数字123",
			num:  123,
			want: []int{1, 2, 3},
		},
		{
			name: "数字4848",
			num:  4848,
			want: []int{4, 8, 4, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getDigits(tt.num)
			if len(got) != len(tt.want) {
				t.Errorf("getDigits(%d) = %v, want %v", tt.num, got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("getDigits(%d) = %v, want %v", tt.num, got, tt.want)
					return
				}
			}
		})
	}
}

func TestCalculateWaviness(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   int64
	}{
		{
			name:   "少于3位",
			digits: []int{1, 2},
			want:   0,
		},
		{
			name:   "峰",
			digits: []int{1, 3, 2},
			want:   1,
		},
		{
			name:   "谷",
			digits: []int{2, 1, 3},
			want:   1,
		},
		{
			name:   "相等数位无波动",
			digits: []int{1, 2, 2, 3},
			want:   0,
		},
		{
			name:   "多个峰谷",
			digits: []int{4, 8, 4, 8},
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateWaviness(tt.digits)
			if got != tt.want {
				t.Errorf("calculateWaviness(%v) = %d, want %d", tt.digits, got, tt.want)
			}
		})
	}
}

func BenchmarkTotalWaviness(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TotalWaviness(120, 130)
	}
}

func BenchmarkTotalWavinessBigRange(b *testing.B) {
	// 测试数位DP在大范围下的性能
	num1 := int64(0)
	num2 := int64(1e8) // 暴力解法无法处理这么大的范围，但DP可以快速计算
	for i := 0; i < b.N; i++ {
		TotalWaviness(num1, num2)
	}
}
