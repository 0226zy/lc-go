package candy

import "testing"

func TestCandy(t *testing.T) {
	tests := []struct {
		name    string
		ratings []int
		want    int
	}{
		// LeetCode 官方示例
		{"示例1: [1,0,2]", []int{1, 0, 2}, 5},
		{"示例2: [1,2,2]", []int{1, 2, 2}, 4},

		// 边界：单人
		{"单人", []int{5}, 1},

		// 边界：评分相同
		{"全相同", []int{3, 3, 3}, 3},

		// 边界：严格递增
		{"严格递增", []int{1, 2, 3, 4, 5}, 15},

		// 边界：严格递减
		{"严格递减", []int{5, 4, 3, 2, 1}, 15},

		// 边界：先升后降（V 形）
		{"先升后降", []int{1, 3, 2, 1}, 7}, // [1,2,2,1]
		{"V形谷底", []int{2, 1, 3}, 5},    // [2,1,2]

		// 边界：山峰
		{"山峰", []int{1, 3, 5, 3, 1}, 9}, // [1,2,3,2,1]

		// 边界：含相等评分打断递增
		{"递增中含相等", []int{1, 2, 2, 3, 4}, 9}, // [1,2,1,2,3]

		// 边界：谷底带平台（平台处评分相等，互不递增）
		{"谷底带平台", []int{3, 2, 1, 1, 2, 3}, 12}, // [3,2,1,1,2,3]

		// 边界：极大值与极小值
		{"交替起伏", []int{1, 3, 2, 4, 3, 5}, 9}, // [1,2,1,2,1,2]
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Candy(tt.ratings); got != tt.want {
				t.Errorf("Candy(%v) = %d, want %d", tt.ratings, got, tt.want)
			}
		})
	}
}

func BenchmarkCandy(b *testing.B) {
	benchmarks := []struct {
		name    string
		ratings []int
	}{
		{"len=10", []int{1, 0, 2, 3, 2, 1, 4, 5, 3, 2}},
		{"len=100", generateRatings(100)},
		{"len=1000", generateRatings(1000)},
		{"len=10000", generateRatings(10000)},
		{"len=100000", generateRatings(100000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Candy(bm.ratings)
			}
		})
	}
}

// generateRatings 生成长度为 n 的锯齿形评分数组
func generateRatings(n int) []int {
	ratings := make([]int, n)
	for i := 0; i < n; i++ {
		ratings[i] = (i * 7) % 100
	}
	return ratings
}
