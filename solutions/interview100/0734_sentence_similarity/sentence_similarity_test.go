package sentencesimilarity

import "testing"

func TestAreSentencesSimilar(t *testing.T) {
	tests := []struct {
		name         string
		sentence1    []string
		sentence2    []string
		similarPairs [][]string
		want         bool
	}{
		// LeetCode 官方示例
		{
			"示例1: 每个位置的单词都相似",
			[]string{"great", "acting", "skills"},
			[]string{"fine", "drama", "talent"},
			[][]string{{"great", "fine"}, {"drama", "acting"}, {"skills", "talent"}},
			true,
		},
		{
			"示例2: 相同单词空词对",
			[]string{"great"},
			[]string{"great"},
			nil,
			true,
		},
		{
			"示例3: 长度不同",
			[]string{"great"},
			[]string{"doubleplus", "good"},
			[][]string{{"great", "doubleplus"}},
			false,
		},

		// 边界：单单词句子
		{"边界: 单单词相似", []string{"a"}, []string{"b"}, [][]string{{"a", "b"}}, true},
		{"边界: 单单词不相似", []string{"a"}, []string{"b"}, nil, false},

		// 相似关系是双向的：查询方向与词对给出方向相反
		{
			"相似关系双向",
			[]string{"fine", "acting"},
			[]string{"great", "drama"},
			[][]string{{"great", "fine"}, {"drama", "acting"}},
			true,
		},

		// 相似关系不可传递：great~fine、fine~good 不能推出 great~good
		{
			"相似关系不可传递",
			[]string{"great"},
			[]string{"good"},
			[][]string{{"great", "fine"}, {"fine", "good"}},
			false,
		},

		// 中间某个位置不匹配
		{
			"中间位置不匹配",
			[]string{"an", "extraordinary", "meal"},
			[]string{"one", "good", "dinner"},
			[][]string{{"an", "one"}, {"extraordinary", "good"}},
			false,
		},

		// 单词大小写敏感
		{
			"大小写敏感",
			[]string{"Great"},
			[]string{"great"},
			[][]string{{"great", "fine"}},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AreSentencesSimilar(tt.sentence1, tt.sentence2, tt.similarPairs); got != tt.want {
				t.Errorf("AreSentencesSimilar(%v, %v, %v) = %v, want %v",
					tt.sentence1, tt.sentence2, tt.similarPairs, got, tt.want)
			}
		})
	}
}

func BenchmarkAreSentencesSimilar(b *testing.B) {
	// 构造 1000 个单词的句子，2000 个相似词对（最大数据规模）
	s1 := make([]string, 1000)
	s2 := make([]string, 1000)
	pairs := make([][]string, 0, 2000)
	for i := 0; i < 1000; i++ {
		w1 := string(rune('a'+i%26)) + string(rune('a'+(i/26)%26)) + string(rune('a'+i/676))
		w2 := w1 + "x"
		s1[i] = w1
		s2[i] = w2
		pairs = append(pairs, []string{w1, w2})
	}
	// 补齐到 2000 个词对（冗余词对不影响结果）
	for i := 0; i < 1000; i++ {
		pairs = append(pairs, []string{s1[i], s1[(i+1)%1000]})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AreSentencesSimilar(s1, s2, pairs)
	}
}
