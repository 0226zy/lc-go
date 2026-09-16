package braceexpansion

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestExpand(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		// LeetCode 官方示例
		{"示例1: 两组花括号", "{a,b}c{d,e}f", []string{"acdf", "acef", "bcdf", "bcef"}},
		{"示例2: 无花括号", "abcd", []string{"abcd"}},

		// 边界：单字符
		{"单字符", "a", []string{"a"}},

		// 花括号内候选无序输入，结果仍须字典序
		{"候选无序输入", "{b,a}c", []string{"ac", "bc"}},
		{"三个候选单组", "{c,a,b}", []string{"a", "b", "c"}},

		// 三组花括号组合，2×2×2 = 8 种
		{"三组花括号", "{a,b}{c,d}{e,f}", []string{
			"ace", "acf", "ade", "adf",
			"bce", "bcf", "bde", "bdf",
		}},

		// 花括号与普通字符交错
		{"花括号开头", "{a,b}c", []string{"ac", "bc"}},
		{"花括号结尾", "ab{c,d}", []string{"abc", "abd"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Expand(tt.s); !utils.EqualStringSlice(got, tt.want) {
				t.Errorf("Expand(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkExpand(b *testing.B) {
	benchmarks := []struct {
		name string
		s    string
	}{
		{"结果数=4", "{a,b}c{d,e}f"},
		{"结果数=64", "{a,b,c,d}{e,f,g,h}{i,j,k,l}"},
		{"结果数=256", "{a,b,c,d}{e,f,g,h}{i,j,k,l}{m,n,o,p}"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Expand(bm.s)
			}
		})
	}
}
