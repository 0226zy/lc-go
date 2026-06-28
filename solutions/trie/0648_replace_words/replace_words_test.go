package replacewords

import "testing"

func TestReplaceWords(t *testing.T) {
	t.Skip("未实现")

	tests := []struct {
		name       string
		dictionary []string
		sentence   string
		want       string
	}{
		{
			name:       "示例1",
			dictionary: []string{"cat","bat","rat"},
			sentence:   "the cattle was rattled by the battery",
			want:       "the cat was rat by the bat",
		},
		{
			name:       "示例2",
			dictionary: []string{"a","b","c"},
			sentence:   "aadsfasf absbs bbab cadsfafs",
			want:       "a a b c",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceWords(tt.dictionary, tt.sentence); got != tt.want {
				t.Errorf("ReplaceWords() = %q, want %q", got, tt.want)
			}
		})
	}
}
