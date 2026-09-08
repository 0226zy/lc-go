package romantointeger

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		// LeetCode 官方示例
		{"示例1: III", "III", 3},
		{"示例2: LVIII", "LVIII", 58},
		{"示例3: MCMXCIV", "MCMXCIV", 1994},

		// 边界：最小值与最大值
		{"最小值 I", "I", 1},
		{"最大值 MMMCMXCIX", "MMMCMXCIX", 3999},

		// 边界：单个字符
		{"单个字符 V", "V", 5},
		{"单个字符 M", "M", 1000},

		// 边界：仅减法表示
		{"减法 IV", "IV", 4},
		{"减法 IX", "IX", 9},
		{"减法 XL", "XL", 40},
		{"减法 XC", "XC", 90},
		{"减法 CD", "CD", 400},
		{"减法 CM", "CM", 900},

		// 边界：减法与加法混合
		{"混合 XIV", "XIV", 14},
		{"混合 LIX", "LIX", 59},
		{"混合 DCCCXC", "DCCCXC", 890},
		{"无减法 MMDCCLXXVII", "MMDCCLXXVII", 2777},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RomanToInt(tt.s); got != tt.want {
				t.Errorf("RomanToInt(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkRomanToInt(b *testing.B) {
	benchmarks := []struct {
		name string
		s    string
	}{
		{"短字符串 III", "III"},
		{"中字符串 LVIII", "LVIII"},
		{"长字符串 MCMXCIV", "MCMXCIV"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RomanToInt(bm.s)
			}
		})
	}
}
