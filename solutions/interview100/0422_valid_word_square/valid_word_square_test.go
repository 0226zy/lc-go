package validwordsquare

import (
	"strings"
	"testing"
)

func TestValidWordSquare(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  bool
	}{
		// LeetCode 官方示例
		{"示例1: 4x4 有效方块", []string{"abcd", "bnrt", "crmy", "dtye"}, true},
		{"示例2: 行长度不一但有效", []string{"abcd", "bnrt", "crm", "dt"}, true},
		{"示例3: ball 与 barl 不对称", []string{"ball", "area", "read", "lady"}, false},

		// 边界：空数组（没有第 0 行，也没有第 0 列，视为有效）
		{"空数组", []string{}, true},

		// 边界：单个单词
		{"单字符单词", []string{"a"}, true},
		{"单个多字符单词", []string{"abc"}, false}, // 第 1 列只有 'a'，与第 1 行 "abc" 不等

		// 边界：行数多于列数（某一列不存在）
		{"行数多于列长度", []string{"ab", "cd", "e"}, false},

		// 边界：某行比对应列短
		{"短行配长列", []string{"abc", "b", "c"}, true},
		{"短行配长列无效", []string{"abc", "b", "d"}, false}, // 第 1 列 "abd" != 第 1 行 "abc"

		// 典型场景：最小有效方阵
		{"2x2 有效", []string{"ab", "bc"}, true},
		{"2x2 无效", []string{"ab", "ac"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidWordSquare(tt.words); got != tt.want {
				t.Errorf("ValidWordSquare(%v) = %v, want %v", tt.words, got, tt.want)
			}
		})
	}
}

func BenchmarkValidWordSquare(b *testing.B) {
	// 构造最大规模的有效单词方块：500 个长度为 500 的单词，words[i][j] 由 (i+j) 决定保证对称
	// 简单起见用同一字母填充的方阵（显然有效）
	makeSquare := func(n, m int) []string {
		words := make([]string, n)
		row := strings.Repeat("a", m)
		for i := range words {
			words[i] = row
		}
		return words
	}

	benchmarks := []struct {
		name  string
		words []string
	}{
		{"4x4", makeSquare(4, 4)},
		{"50x50", makeSquare(50, 50)},
		{"500x500", makeSquare(500, 500)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ValidWordSquare(bm.words)
			}
		})
	}
}
