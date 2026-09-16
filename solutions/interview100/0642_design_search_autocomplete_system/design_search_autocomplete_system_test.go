package designsearchautocompletesystem

import (
	"strings"
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestAutocompleteSystem(t *testing.T) {
	tests := []struct {
		name      string
		sentences []string
		times     []int
		inputs    string     // 逐字符输入序列
		want      [][]string // 每次 input 的期望返回（'#' 对应空列表）
	}{
		{
			"LeetCode官方示例1",
			[]string{"i love you", "island", "iroman", "i love leetcode"},
			[]int{5, 3, 2, 2},
			"i a#",
			[][]string{
				{"i love you", "island", "i love leetcode"},
				{"i love you", "i love leetcode"},
				{},
				{},
			},
		},
		{
			"长度不同的同前缀句子",
			[]string{"abc", "ab", "a"},
			[]int{1, 2, 3},
			"a#a",
			[][]string{
				{"a", "ab", "abc"},
				{},
				{"a", "ab", "abc"},
			},
		},
		{
			"井号新增句子参与后续补全",
			[]string{"i love you", "island", "iroman", "i love leetcode"},
			[]int{5, 3, 2, 2},
			"i a#i a",
			[][]string{
				{"i love you", "island", "i love leetcode"},
				{"i love you", "i love leetcode"},
				{},
				{},
				{"i love you", "island", "i love leetcode"},
				{"i love you", "i love leetcode", "i a"},
				{"i a"},
			},
		},
		{
			"热度相同时按字典序升序",
			[]string{"ab", "aa", "ac"},
			[]int{2, 2, 1},
			"a",
			[][]string{{"aa", "ab", "ac"}},
		},
		{
			"超过3个候选只取前3",
			[]string{"a1", "a2", "a3", "a4", "a5"},
			[]int{5, 4, 3, 2, 1},
			"a",
			[][]string{{"a1", "a2", "a3"}},
		},
		{
			"无任何匹配",
			[]string{"hello"},
			[]int{1},
			"x",
			[][]string{{}},
		},
		{
			"边界：初始热度为0的句子也能补全",
			[]string{"abc"},
			[]int{0},
			"a",
			[][]string{{"abc"}},
		},
		{
			"边界：空历史库完全靠井号积累",
			[]string{},
			[]int{},
			"ab#a",
			[][]string{{}, {}, {}, {"ab"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := Constructor(tt.sentences, tt.times)
			if len(tt.inputs) != len(tt.want) {
				t.Fatalf("输入字符数与期望数不一致: %d vs %d", len(tt.inputs), len(tt.want))
			}
			for i := 0; i < len(tt.inputs); i++ {
				got := as.Input(tt.inputs[i])
				if !utils.EqualStringSlice(got, tt.want[i]) {
					t.Errorf("第 %d 次输入 %q: got %v, want %v", i, string(tt.inputs[i]), got, tt.want[i])
				}
			}
		})
	}
}

func BenchmarkAutocompleteSystem(b *testing.B) {
	// 构造 100 个历史句子
	sentences := make([]string, 0, 100)
	times := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		sentences = append(sentences, "sentence "+strings.Repeat(string(rune('a'+i%26)), i%10+1)+string(rune('a'+i/26)))
		times = append(times, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		as := Constructor(sentences, times)
		// 模拟输入一句话再保存
		for _, c := range "sentence aaa#" {
			as.Input(byte(c))
		}
	}
}
