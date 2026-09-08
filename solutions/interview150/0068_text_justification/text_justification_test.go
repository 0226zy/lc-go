package textjustification

import (
	"reflect"
	"testing"
)

func TestFullJustify(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		maxWidth int
		want     []string
	}{
		// LeetCode 官方示例
		{
			"示例1: This is an example",
			[]string{"This", "is", "an", "example", "of", "text", "justification."},
			16,
			[]string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},
		{
			"示例2: What must be（含长单词与最后一行）",
			[]string{"What", "must", "be", "acknowledgment", "shall", "be"},
			16,
			[]string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		},
		{
			"示例3: Science is what（宽度 20）",
			[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			20,
			[]string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},

		// 边界：只有一个词
		{"单个单词", []string{"a"}, 1, []string{"a"}},
		{"单个单词需补空格", []string{"a"}, 3, []string{"a  "}},
		{"单个单词恰好占满", []string{"abc"}, 3, []string{"abc"}},

		// 边界：单词长度恰好等于 maxWidth
		{"单词恰好占满一行后跟其他词", []string{"abc", "def", "ghi"}, 3, []string{"abc", "def", "ghi"}},
		{
			"长短词混合且长词恰好占满",
			[]string{"aa", "bb", "cc", "dd", "eeeeeeee"},
			8,
			[]string{"aa bb cc", "dd      ", "eeeeeeee"},
		},

		// 边界：所有词恰好在同一行（即第一行也是最后一行）
		{"所有词同一行", []string{"a", "b", "c"}, 5, []string{"a b c"}},

		// 边界：普通行空格分配——均匀分配与余数靠左
		{
			"空格均匀分配无余数",
			[]string{"a", "b", "c", "d", "e"},
			7,
			[]string{"a b c d", "e      "},
		},
		{
			"余一个空格分给最左空隙",
			[]string{"a", "b", "c", "d", "e"},
			8,
			[]string{"a  b c d", "e       "},
		},

		// 边界：最后一行左对齐，词间只有单空格
		{
			"最后一行不拉伸",
			[]string{"a", "b", "c", "d", "e"},
			3,
			[]string{"a b", "c d", "e  "},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FullJustify(tt.words, tt.maxWidth)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FullJustify(%v, %d) =\n%q\nwant:\n%q", tt.words, tt.maxWidth, got, tt.want)
			}
		})
	}
}

// TestFullJustifyInvariants 校验输出满足题面硬性要求：每行长度恰为 maxWidth
func TestFullJustifyInvariants(t *testing.T) {
	cases := []struct {
		words    []string
		maxWidth int
	}{
		{[]string{"This", "is", "an", "example", "of", "text", "justification."}, 16},
		{[]string{"a"}, 1},
		{[]string{"a", "b", "c", "d", "e", "f", "g"}, 7},
		{[]string{"aaaa", "bb", "cc", "dd", "eeeeeeeee", "ff", "g"}, 11},
	}
	for _, c := range cases {
		lines := FullJustify(c.words, c.maxWidth)
		for _, line := range lines {
			if len(line) != c.maxWidth {
				t.Errorf("行长度 %d != maxWidth %d，行内容: %q", len(line), c.maxWidth, line)
			}
		}
	}
}

func BenchmarkFullJustify(b *testing.B) {
	// 构造 300 个单词的大规模输入（题面约束上限）
	words := make([]string, 300)
	for i := range words {
		if i%10 == 0 {
			words[i] = "acknowledgment"
		} else {
			words[i] = "word"
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FullJustify(words, 20)
	}
}
