package integertoroman

import "testing"

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		// LeetCode 官方示例
		{"示例1: 3", 3, "III"},
		{"示例2: 58", 58, "LVIII"},
		{"示例3: 1994", 1994, "MCMXCIV"},

		// 边界：最小值与最大值
		{"最小值 1", 1, "I"},
		{"最大值 3999", 3999, "MMMCMXCIX"},

		// 边界：需要减法表示的数字
		{"减法 4", 4, "IV"},
		{"减法 9", 9, "IX"},
		{"减法 40", 40, "XL"},
		{"减法 90", 90, "XC"},
		{"减法 400", 400, "CD"},
		{"减法 900", 900, "CM"},

		// 边界：进位附近的数字
		{"10", 10, "X"},
		{"14", 14, "XIV"},
		{"20", 20, "XX"},
		{"50", 50, "L"},
		{"100", 100, "C"},
		{"500", 500, "D"},
		{"1000", 1000, "M"},

		// 边界：混合数字
		{"58 已含", 58, "LVIII"},
		{"621", 621, "DCXXI"},
		{"1776", 1776, "MDCCLXXVI"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToRoman(tt.num); got != tt.want {
				t.Errorf("IntToRoman(%d) = %q, want %q", tt.num, got, tt.want)
			}
		})
	}
}

func BenchmarkIntToRoman(b *testing.B) {
	benchmarks := []struct {
		name string
		num  int
	}{
		{"小数字 3", 3},
		{"中数字 58", 58},
		{"大数字 1994", 1994},
		{"最大值 3999", 3999},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IntToRoman(bm.num)
			}
		})
	}
}
