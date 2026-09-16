package maximumpointsyoucanobtainfromcards

import "testing"

func TestMaxScore(t *testing.T) {
	tests := []struct {
		name       string
		cardPoints []int
		k          int
		want       int
	}{
		// LeetCode 官方示例
		{
			"示例1: 常规混合牌堆",
			[]int{1, 2, 3, 4, 5, 6, 1},
			3,
			12,
		},
		{
			"示例2: 所有牌点数相同",
			[]int{2, 2, 2},
			2,
			4,
		},
		{
			"示例3: k等于数组长度",
			[]int{9, 7, 7, 9, 7, 7, 9},
			7,
			55,
		},

		// 边界：空输入
		{
			"空数组",
			[]int{},
			1,
			0,
		},

		// 边界：单元素
		{
			"单元素只拿一张",
			[]int{5},
			1,
			5,
		},

		// 边界：k 为 0（题目约束中不出现，保证健壮性）
		{
			"k为0不拿牌",
			[]int{1, 2, 3},
			0,
			0,
		},

		// 边界：只拿一张，取首尾最大值
		{
			"只拿一张取首尾较大值",
			[]int{1, 100, 2},
			1,
			2,
		},

		// 最优解在左侧
		{
			"最优解全在左侧",
			[]int{100, 1, 1, 1, 1},
			2,
			101,
		},

		// 最优解在右侧
		{
			"最优解全在右侧",
			[]int{1, 1, 1, 1, 100},
			2,
			101,
		},

		// 最优解左右混合
		{
			"最优解左右混合",
			[]int{6, 2, 1, 1, 3, 5},
			4,
			16,
		},

		// 极端值：全部最大点数
		{
			"极端值全为最大点数",
			[]int{10000, 10000, 10000, 10000},
			3,
			30000,
		},

		// 极端值：全部最小点数
		{
			"极端值全为最小点数",
			[]int{1, 1, 1, 1, 1},
			4,
			4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxScore(tt.cardPoints, tt.k); got != tt.want {
				t.Errorf("MaxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

// makeCardPoints 生成长度为 n、每张牌点数均为 fill 的数组
func makeCardPoints(n, fill int) []int {
	cardPoints := make([]int, n)
	for i := range cardPoints {
		cardPoints[i] = fill
	}
	return cardPoints
}

func BenchmarkMaxScore(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
		k    int
	}{
		{"长度1000拿500张", 1000, 500},
		{"长度100000拿50000张", 100000, 50000},
		{"长度100000全部拿走", 100000, 100000},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			cardPoints := makeCardPoints(bm.n, 7)
			for i := 0; i < b.N; i++ {
				MaxScore(cardPoints, bm.k)
			}
		})
	}
}
