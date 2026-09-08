package hindex

import (
	"math/rand"
	"testing"
)

func TestHIndex(t *testing.T) {
	tests := []struct {
		name      string
		citations []int
		want      int
	}{
		// LeetCode 官方示例
		{"示例1: [3,0,6,1,5]", []int{3, 0, 6, 1, 5}, 3},
		{"示例2: [1,3,1]", []int{1, 3, 1}, 1},

		// 边界：单篇论文
		{"单篇无引用", []int{0}, 0},
		{"单篇有引用", []int{100}, 1},

		// 边界：全为0
		{"全为0", []int{0, 0, 0, 0}, 0},

		// 边界：全部高引用，h 等于论文数
		{"全部高引用", []int{5, 5, 5, 5, 5}, 5},
		{"引用数远超论文数", []int{100, 200, 300}, 3},

		// 边界：递增/递减序列
		{"递增序列", []int{0, 1, 2, 3, 4}, 2},
		{"递减序列", []int{10, 8, 5, 4, 3}, 4},

		// 边界：只有一篇达标
		{"仅一篇达标", []int{2, 0, 0}, 1},

		// 随机混合
		{"混合引用", []int{4, 4, 0, 0, 0}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// HIndex 会原地排序，先复制一份
			citations := make([]int, len(tt.citations))
			copy(citations, tt.citations)
			if got := HIndex(citations); got != tt.want {
				t.Errorf("HIndex(%v) = %d, want %d", tt.citations, got, tt.want)
			}
		})
	}
}

// TestHIndexBucket 验证计数桶解法与排序法结果一致
func TestHIndexBucket(t *testing.T) {
	tests := []struct {
		name      string
		citations []int
		want      int
	}{
		{"示例1: [3,0,6,1,5]", []int{3, 0, 6, 1, 5}, 3},
		{"示例2: [1,3,1]", []int{1, 3, 1}, 1},
		{"全为0", []int{0, 0, 0, 0}, 0},
		{"全部高引用", []int{5, 5, 5, 5, 5}, 5},
		{"递增序列", []int{0, 1, 2, 3, 4}, 2},
		{"随机5000篇", randomCitations(5000, 1000), -1}, // -1 表示与排序法对比
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			if want == -1 {
				// 与排序法对齐期望值
				sorted := make([]int, len(tt.citations))
				copy(sorted, tt.citations)
				want = HIndex(sorted)
			}
			if got := HIndexBucket(tt.citations); got != want {
				t.Errorf("HIndexBucket(%v) = %d, want %d", tt.citations, got, want)
			}
		})
	}
}

func BenchmarkHIndex(b *testing.B) {
	benchmarks := []struct {
		name      string
		citations []int
	}{
		{"len=10", randomCitations(10, 100)},
		{"len=1000", randomCitations(1000, 1000)},
		{"len=5000", randomCitations(5000, 1000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// HIndex 原地排序，每次需复制
				citations := make([]int, len(bm.citations))
				copy(citations, bm.citations)
				HIndex(citations)
			}
		})
	}
}

func BenchmarkHIndexBucket(b *testing.B) {
	citations := randomCitations(5000, 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HIndexBucket(citations)
	}
}

// randomCitations 生成 n 篇论文、引用次数在 [0, maxCite] 的随机数组
func randomCitations(n, maxCite int) []int {
	citations := make([]int, n)
	for i := 0; i < n; i++ {
		citations[i] = rand.Intn(maxCite + 1)
	}
	return citations
}
