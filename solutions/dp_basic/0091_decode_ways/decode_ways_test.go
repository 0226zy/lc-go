package decodeways

import (
	"strings"
	"testing"
)

var numDecodingsCases = []struct {
	name string
	s    string
	want int
}{
	{name: "示例1：12", s: "12", want: 2},
	{name: "示例2：226", s: "226", want: 3},
	{name: "示例3：06前导零", s: "06", want: 0},
	{name: "单个0", s: "0", want: 0},
	{name: "单个数字", s: "7", want: 1},
	{name: "整10", s: "10", want: 1},
	{name: "27只能拆开", s: "27", want: 1},
	{name: "26可合可拆", s: "26", want: 2},
	{name: "含10的11106", s: "11106", want: 2},
	{name: "0后不可合并", s: "2101", want: 1},
	{name: "30无法解码", s: "30", want: 0},
	{name: "全1长串", s: "1111111111", want: 89}, // 同斐波那契 F(11)
}

func TestNumDecodings(t *testing.T) {
	for _, tt := range numDecodingsCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumDecodings(tt.s); got != tt.want {
				t.Errorf("NumDecodings(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestNumDecodingsOptimized(t *testing.T) {
	for _, tt := range numDecodingsCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumDecodingsOptimized(tt.s); got != tt.want {
				t.Errorf("NumDecodingsOptimized(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

var benchInput = strings.Repeat("1", 45) // 保证答案不溢出 32 位的较长输入

func BenchmarkNumDecodings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumDecodings(benchInput)
	}
}

func BenchmarkNumDecodingsOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumDecodingsOptimized(benchInput)
	}
}
