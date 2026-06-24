package longestcommonsubsequence

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name  string
		text1 string
		text2 string
		want  int
	}{
		{
			name:  "示例1：部分匹配",
			text1: "abcde",
			text2: "ace",
			want:  3,
		},
		{
			name:  "示例2：完全相同",
			text1: "abc",
			text2: "abc",
			want:  3,
		},
		{
			name:  "示例3：完全不同",
			text1: "abc",
			text2: "def",
			want:  0,
		},
		{
			name:  "边界：空字符串",
			text1: "",
			text2: "abc",
			want:  0,
		},
		{
			name:  "两个都为空",
			text1: "",
			text2: "",
			want:  0,
		},
		{
			name:  "单字符相同",
			text1: "a",
			text2: "a",
			want:  1,
		},
		{
			name:  "单字符不同",
			text1: "a",
			text2: "b",
			want:  0,
		},
		{
			name:  "经典重复子序列",
			text1: "abcdefabcdef",
			text2: "abcdef",
			want:  6,
		},
		{
			name:  "长公共子序列",
			text1: "abcba",
			text2: "abcbcba",
			want:  5,
		},
		{
			name:  "分散公共子序列",
			text1: "hofubmnylkra",
			text2: "pqhgxgdofcvmr",
			want:  5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestCommonSubsequence(tt.text1, tt.text2)
			if got != tt.want {
				t.Errorf("LongestCommonSubsequence(%q, %q) = %d, want %d", tt.text1, tt.text2, got, tt.want)
			}
		})
	}
}

func BenchmarkLongestCommonSubsequence(b *testing.B) {
	text1 := "abcde"
	text2 := "ace"
	for i := 0; i < b.N; i++ {
		LongestCommonSubsequence(text1, text2)
	}
}
