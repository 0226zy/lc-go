package lettercombinations

import (
	"reflect"
	"sort"
	"testing"
)

// normalize 将字符串切片排序，用于忽略返回顺序的无序比较
func normalize(s []string) []string {
	cp := make([]string, len(s))
	copy(cp, s)
	sort.Strings(cp)
	return cp
}

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		// LeetCode 官方示例
		{"示例1: 两位数字23", "23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"示例2: 空字符串", "", []string{}},
		{"示例3: 单个数字2", "2", []string{"a", "b", "c"}},

		// 边界情况：含 7 和 9（映射 4 个字母）
		{"数字7映射pqrs", "7", []string{"p", "q", "r", "s"}},
		{"数字9映射wxyz", "9", []string{"w", "x", "y", "z"}},
		{"两个4字母数字79", "79", []string{
			"pw", "px", "py", "pz",
			"qw", "qx", "qy", "qz",
			"rw", "rx", "ry", "rz",
			"sw", "sx", "sy", "sz",
		}},
		{"混合数字234", "234", []string{
			"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi",
			"bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi",
			"cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LetterCombinations(tt.digits)
			if !reflect.DeepEqual(normalize(got), normalize(tt.want)) {
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
	// 4 位全为 4 字母数字时应产生 4^4 = 256 种组合
	got = LetterCombinations("7979")
	if len(got) != 256 {
		t.Errorf("LetterCombinations(\"7979\") 长度 = %d, want 256", len(got))
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
		{"两位数字23", "23"},
		{"三位数字234", "234"},
		{"四位数字2345", "2345"},
		{"四位全4字母7979", "7979"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LetterCombinations(bm.digits)
			}
		})
	}
}
