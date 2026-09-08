package minimumgeneticmutation

import "testing"

func TestMinMutation(t *testing.T) {
	tests := []struct {
		name      string
		startGene string
		endGene   string
		bank      []string
		want      int
	}{
		// LeetCode 官方示例
		{"示例1: 一步直达", "AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}, 1},
		{"示例2: 两步变化", "AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}, 2},

		// 边界：起点即终点
		{"起点等于终点", "AACCGGTT", "AACCGGTT", []string{"AACCGGTA"}, 0},

		// 边界：终点不在基因库
		{"终点不在基因库", "AACCGGTT", "AACCGGTA", []string{"AACCGGTC"}, -1},

		// 边界：基因库为空
		{"基因库为空", "AACCGGTT", "AACCGGTA", nil, -1},

		// 边界：多个分支，需找最短路径
		{"多分支取最短", "AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}, 3},

		// 边界：存在环，避免重复访问
		{"存在环", "AACCGGTT", "AACCGGTA", []string{"AACCGGTA", "AACCGGTC", "AACCGGCC"}, 1},

		// 边界：只能单向变化（差多个位且库中无中间态）
		{"无中间态不可达", "AAAAAAAA", "CCCCCCCC", []string{"AAAAAAAC"}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinMutation(tt.startGene, tt.endGene, tt.bank); got != tt.want {
				t.Errorf("MinMutation(%q, %q, %v) = %d, want %d",
					tt.startGene, tt.endGene, tt.bank, got, tt.want)
			}
		})
	}
}

func BenchmarkMinMutation(b *testing.B) {
	bank := []string{
		"AACCGGTA", "AACCGCTA", "AAACGGTA", "AAACGGTT", "AACCGGTC",
		"AACCGGCC", "AACCGGCT", "AAACGGTC", "AAACGGCT", "AAAAGGTT",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinMutation("AACCGGTT", "AAACGGTA", bank)
	}
}
