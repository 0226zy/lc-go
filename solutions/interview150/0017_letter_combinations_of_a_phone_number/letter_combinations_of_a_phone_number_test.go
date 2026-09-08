package lettercombinations

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		// LeetCode 官方示例
		{"示例1: digits=23", "23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"示例2: digits为空", "", []string{}},
		{"示例3: digits=2", "2", []string{"a", "b", "c"}},

		// 边界：含 7 和 9（对应 4 个字母）
		{"数字7对应pqrs", "7", []string{"p", "q", "r", "s"}},
		{"数字9对应wxyz", "9", []string{"w", "x", "y", "z"}},
		{"组合7和8", "78", []string{"pt", "pu", "pv", "qt", "qu", "qv", "rt", "ru", "rv", "st", "su", "sv"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LetterCombinations(tt.digits)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LetterCombinations(%q) = %v, want %v", tt.digits, got, tt.want)
			}
		})
	}
}

func TestLetterCombinationsCount(t *testing.T) {
	// 4 位数字每位 3 个字母时应产生 3^4 = 81 种组合
	got := LetterCombinations("2345")
	if len(got) != 81 {
		t.Errorf("LetterCombinations(\"2345\") 长度 = %d, want 81", len(got))
	}
	// 组合之间不应有重复
	seen := make(map[string]bool, len(got))
	for _, s := range got {
		if seen[s] {
			t.Errorf("出现重复组合 %q", s)
		}
		seen[s] = true
	}
}

func BenchmarkLetterCombinations(b *testing.B) {
	benchmarks := []struct {
		name   string
		digits string
	}{
		{"digits=23", "23"},
		{"digits=234", "234"},
		{"digits=2345", "2345"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LetterCombinations(bm.digits)
			}
		})
	}
}
