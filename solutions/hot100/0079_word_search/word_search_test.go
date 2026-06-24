package wordsearch

import "testing"

func TestExist(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		word  string
		want  bool
	}{
		{
			name:  "官方示例1 ABCCED",
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:  "ABCCED",
			want:  true,
		},
		{
			name:  "官方示例2 SEE",
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:  "SEE",
			want:  true,
		},
		{
			name:  "官方示例3 ABCB",
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:  "ABCB",
			want:  false,
		},
		{
			name:  "单格匹配",
			board: [][]byte{{'A'}},
			word:  "A",
			want:  true,
		},
		{
			name:  "单格不匹配",
			board: [][]byte{{'A'}},
			word:  "B",
			want:  false,
		},
		{
			name:  "单词比网格长",
			board: [][]byte{{'A', 'B'}},
			word:  "ABC",
			want:  false,
		},
		{
			name:  "重复使用同一格子应失败",
			board: [][]byte{{'A', 'A'}},
			word:  "AAA",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Exist(tt.board, tt.word)
			if got != tt.want {
				t.Errorf("Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExistInPlace(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		word  string
		want  bool
	}{
		{
			name:  "官方示例1 ABCCED",
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:  "ABCCED",
			want:  true,
		},
		{
			name:  "官方示例2 SEE",
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:  "SEE",
			want:  true,
		},
		{
			name:  "官方示例3 ABCB",
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:  "ABCB",
			want:  false,
		},
		{
			name:  "单格匹配",
			board: [][]byte{{'A'}},
			word:  "A",
			want:  true,
		},
		{
			name:  "单格不匹配",
			board: [][]byte{{'A'}},
			word:  "B",
			want:  false,
		},
		{
			name:  "单词比网格长",
			board: [][]byte{{'A', 'B'}},
			word:  "ABC",
			want:  false,
		},
		{
			name:  "重复使用同一格子应失败",
			board: [][]byte{{'A', 'A'}},
			word:  "AAA",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExistInPlace(tt.board, tt.word)
			if got != tt.want {
				t.Errorf("ExistInPlace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkExist(b *testing.B) {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	for i := 0; i < b.N; i++ {
		Exist(board, "ABCCED")
	}
}

func BenchmarkExistInPlace(b *testing.B) {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	for i := 0; i < b.N; i++ {
		ExistInPlace(board, "ABCCED")
	}
}
