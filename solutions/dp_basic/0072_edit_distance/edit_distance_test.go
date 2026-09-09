package editdistance

import "testing"

var minDistanceCases = []struct {
	name  string
	word1 string
	word2 string
	want  int
}{
	{name: "示例1：horse转ros", word1: "horse", word2: "ros", want: 3},
	{name: "示例2：intention转execution", word1: "intention", word2: "execution", want: 5},
	{name: "两空串", word1: "", word2: "", want: 0},
	{name: "word1为空", word1: "", word2: "abc", want: 3},
	{name: "word2为空", word1: "abc", word2: "", want: 3},
	{name: "完全相同", word1: "abc", word2: "abc", want: 0},
	{name: "单字符替换", word1: "a", word2: "b", want: 1},
	{name: "单字符相同", word1: "a", word2: "a", want: 0},
	{name: "删除加插入", word1: "ab", word2: "bc", want: 2},
	{name: "一个转自身前缀", word1: "abc", word2: "ab", want: 1},
}

func TestMinDistance(t *testing.T) {
	for _, tt := range minDistanceCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinDistance(tt.word1, tt.word2); got != tt.want {
				t.Errorf("MinDistance(%q, %q) = %d, want %d", tt.word1, tt.word2, got, tt.want)
			}
		})
	}
}

func TestMinDistanceOptimized(t *testing.T) {
	for _, tt := range minDistanceCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinDistanceOptimized(tt.word1, tt.word2); got != tt.want {
				t.Errorf("MinDistanceOptimized(%q, %q) = %d, want %d", tt.word1, tt.word2, got, tt.want)
			}
		})
	}
}

func BenchmarkMinDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinDistance("intention", "execution")
	}
}

func BenchmarkMinDistanceOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinDistanceOptimized("intention", "execution")
	}
}
