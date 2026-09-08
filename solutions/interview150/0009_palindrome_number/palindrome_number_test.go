package palindromenumber

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		// LeetCode 官方示例
		{"示例1: 121是回文", 121, true},
		{"示例2: -121不是回文", -121, false},
		{"示例3: 10不是回文", 10, false},

		// 边界：0 和个位数
		{"0是回文", 0, true},
		{"个位数7是回文", 7, true},

		// 边界：偶数位回文与非回文
		{"偶数位回文1221", 1221, true},
		{"偶数位非回文1234", 1234, false},

		// 边界：奇数位回文与非回文
		{"奇数位回文12321", 12321, true},
		{"奇数位非回文12345", 12345, false},

		// 边界：个位为0的数
		{"100不是回文", 100, false},
		{"1000不是回文", 1000, false},

		// 边界：负数
		{"-1不是回文", -1, false},
		{"-101不是回文", -101, false},

		// 边界：极大值
		{"int32最大值不是回文", 2147483647, false},
		{"接近回文的2147447412", 2147447412, true},

		// 易错：前半反转后可能看似相等但实际不等
		{"1112111是回文", 1112111, true},
		{"1123211是回文", 1123211, true},
		{"1000021不是回文", 1000021, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.x); got != tt.want {
				t.Errorf("IsPalindrome(%d) = %v, want %v", tt.x, got, tt.want)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	benchmarks := []struct {
		name string
		x    int
	}{
		{"回文123454321", 123454321},
		{"非回文123456789", 123456789},
		{"int32最大值", 2147483647},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsPalindrome(bm.x)
			}
		})
	}
}
