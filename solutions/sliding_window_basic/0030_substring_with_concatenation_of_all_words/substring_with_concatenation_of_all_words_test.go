package substringwithconcatenationofallwords

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		words []string
		want  []int
	}{
		// LeetCode 官方示例
		{"示例1: barfoothefoobarman", "barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		{"示例2: 不存在串联子串", "wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, []int{}},
		{"示例3: barfoofoobarthefoobarman", "barfoofoobarthefoobarman", []string{"bar", "foo", "the"}, []int{6, 9, 12}},

		// 边界：s 长度不足以容纳所有单词
		{"s过短", "ab", []string{"abc"}, nil},
		// 边界：单单词单字符
		{"单单词", "a", []string{"a"}, []int{0}},
		// 边界：words 中全是重复单词，需要全部用上
		{"重复单词全部使用", "aaaaaaaa", []string{"aa", "aa", "aa"}, []int{0, 1, 2}},
		// 边界：s 中单词有重叠候选
		{"重叠候选", "abababab", []string{"ab", "ab"}, []int{0, 2, 4}},
		// 边界：words 中的单词在 s 中出现次数超过所需，应跳过超标窗口
		{"单词次数超标", "barfoofoobarman", []string{"bar", "foo"}, []int{0, 6}},
		// 边界：单词顺序与 words 不同（顺序任意）
		{"顺序不同也算", "foobar", []string{"bar", "foo"}, []int{0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindSubstring(tt.s, tt.words)
			// 结果顺序应天然递增，这里排序后比较避免顺序问题
			sort.Ints(got)
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindSubstring(%q, %v) = %v, want %v", tt.s, tt.words, got, tt.want)
			}
		})
	}
}

func BenchmarkFindSubstring(b *testing.B) {
	// 构造长串：abc 三个单词反复拼接并混入干扰字符
	long := make([]byte, 0, 9990)
	for len(long) < 9900 {
		long = append(long, "abc"...)
		long = append(long, "x"...)
	}

	benchmarks := []struct {
		name  string
		s     string
		words []string
	}{
		{"len=18", "barfoothefoobarman", []string{"foo", "bar"}},
		{"len=1000", string(long[:1000]), []string{"abc", "bca", "cab"}},
		{"len=5000", string(long[:5000]), []string{"abc", "bca", "cab"}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindSubstring(bm.s, bm.words)
			}
		})
	}
}
