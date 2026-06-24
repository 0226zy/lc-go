package editdistance

import "testing"

func TestMinDistance(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  int
	}{
		{
			name:  "示例1：horse到ros",
			word1: "horse",
			word2: "ros",
			want:  3,
		},
		{
			name:  "示例2：intention到execution",
			word1: "intention",
			word2: "execution",
			want:  5,
		},
		{
			name:  "边界：相同单词",
			word1: "abc",
			word2: "abc",
			want:  0,
		},
		{
			name:  "边界：word1为空",
			word1: "",
			word2: "abc",
			want:  3,
		},
		{
			name:  "边界：word2为空",
			word1: "abc",
			word2: "",
			want:  3,
		},
		{
			name:  "边界：都为空",
			word1: "",
			word2: "",
			want:  0,
		},
		{
			name:  "单字符替换",
			word1: "a",
			word2: "b",
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinDistance(tt.word1, tt.word2)
			if got != tt.want {
				t.Errorf("MinDistance(%q, %q) = %d, want %d", tt.word1, tt.word2, got, tt.want)
			}
		})
	}
}

func BenchmarkMinDistance(b *testing.B) {
	word1 := "horse"
	word2 := "ros"
	for i := 0; i < b.N; i++ {
		MinDistance(word1, word2)
	}
}
