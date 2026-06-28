package wordsearchii

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindWords(t *testing.T) {
	t.Skip("未实现")

	tests := []struct {
		name  string
		board [][]byte
		words []string
		want  []string
	}{
		{
			name:  "示例1",
			board: [][]byte{{'o','a','a','n'},{'e','t','a','e'},{'i','h','k','r'},{'i','f','l','v'}},
			words: []string{"oath","pea","eat","rain"},
			want:  []string{"eat","oath"},
		},
		{
			name:  "示例2",
			board: [][]byte{{'a','b'},{'c','d'}},
			words: []string{"abcb"},
			want:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindWords(tt.board, tt.words)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
