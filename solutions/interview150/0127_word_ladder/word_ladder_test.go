package wordladder

import "testing"

func TestLadderLength(t *testing.T) {
	tests := []struct {
		name      string
		beginWord string
		endWord   string
		wordList  []string
		want      int
	}{
		// LeetCode 官方示例
		{"示例1: hit到cog", "hit", "cog",
			[]string{"hot", "dot", "dog", "lot", "log", "cog"}, 5},
		{"示例2: endWord不在字典", "hit", "cog",
			[]string{"hot", "dot", "dog", "lot", "log"}, 0},

		// 边界：beginWord 与 endWord 只差一位
		{"只差一位", "a", "c", []string{"c"}, 2},

		// 边界：字典为空
		{"字典为空", "hit", "cog", nil, 0},

		// 边界：字典只有一个词且就是 endWord
		{"字典仅含endWord", "hit", "hot", []string{"hot"}, 2},

		// 边界：存在环，需靠 visited 去重
		{"含环的字典", "hit", "cog",
			[]string{"hot", "dot", "dog", "lot", "log", "cog", "hig"}, 5},

		// 边界：endWord 可达（hit->hot->hog->cog）
		{"需绕过死胡同", "hit", "cog",
			[]string{"hot", "dot", "hog", "cog"}, 4},

		// 边界：beginWord 出现在字典中
		{"beginWord在字典中", "hit", "cog",
			[]string{"hit", "hot", "dot", "dog", "lot", "log", "cog"}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LadderLength(tt.beginWord, tt.endWord, tt.wordList); got != tt.want {
				t.Errorf("LadderLength(%q, %q, %v) = %d, want %d",
					tt.beginWord, tt.endWord, tt.wordList, got, tt.want)
			}
		})
	}
}

func BenchmarkLadderLength(b *testing.B) {
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog", "hig", "hog", "hag"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LadderLength("hit", "cog", wordList)
	}
}
