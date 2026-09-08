package zigzagconversion

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		numRows int
		want    string
	}{
		// LeetCode 官方示例
		{"示例1: PAYPALISHIRING 3行", "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"示例2: PAYPALISHIRING 4行", "PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"示例3: A 单行", "A", 1, "A"},

		// 边界：单字符多行
		{"单字符两行", "A", 2, "A"},

		// 边界：行数为 1
		{"单行不变", "ABCDEFG", 1, "ABCDEFG"},

		// 边界：行数 >= 字符串长度
		{"行数大于长度", "AB", 5, "AB"},
		{"行数等于长度", "ABC", 3, "ABC"},

		// 边界：两行
		{"两行交错", "ABCDEF", 2, "ACEBDF"},

		// 边界：三行
		{"三行奇数长度", "ABCDE", 3, "AEBDC"},

		// 边界：四行
		{"四行交错", "ABCDEFGHIJK", 4, "AGBFHCEIKDJ"},

		// 边界：含小写与符号
		{"大小写混合", "aBcDeFgH", 3, "aeBDFHcg"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert(tt.s, tt.numRows); got != tt.want {
				t.Errorf("Convert(%q, %d) = %q, want %q", tt.s, tt.numRows, got, tt.want)
			}
		})
	}
}

func BenchmarkConvert(b *testing.B) {
	s := "PAYPALISHIRING"
	benchmarks := []struct {
		name    string
		numRows int
	}{
		{"3行", 3},
		{"10行", 10},
		{"1000行", 1000},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Convert(s, bm.numRows)
			}
		})
	}
}
