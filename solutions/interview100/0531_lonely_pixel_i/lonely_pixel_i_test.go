package lonelypixeli

import "testing"

func TestFindLonelyPixel(t *testing.T) {
	tests := []struct {
		name    string
		picture [][]byte
		want    int
	}{
		// LeetCode 官方示例
		{"示例1: 对角线上三个B", [][]byte{
			{'W', 'W', 'B'},
			{'W', 'B', 'W'},
			{'B', 'W', 'W'},
		}, 3},
		{"示例2: 全是B", [][]byte{
			{'B', 'B', 'B'},
			{'B', 'B', 'W'},
			{'B', 'B', 'B'},
		}, 0},

		// 边界：空图片
		{"空图片", [][]byte{}, 0},

		// 边界：单像素
		{"单个黑色像素", [][]byte{{'B'}}, 1},
		{"单个白色像素", [][]byte{{'W'}}, 0},

		// 边界：全白
		{"全白图片", [][]byte{
			{'W', 'W'},
			{'W', 'W'},
		}, 0},

		// 同行有两个 B → 都不是孤独像素
		{"同行两个B", [][]byte{
			{'B', 'B'},
			{'W', 'W'},
		}, 0},

		// 同列有两个 B → 都不是孤独像素
		{"同列两个B", [][]byte{
			{'B', 'W'},
			{'B', 'W'},
		}, 0},

		// 混合：一个孤独、一个不孤独
		{"一个孤独一个同行", [][]byte{
			{'B', 'W', 'B'},
			{'W', 'W', 'W'},
			{'W', 'B', 'W'},
		}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindLonelyPixel(tt.picture); got != tt.want {
				t.Errorf("FindLonelyPixel(%v) = %d, want %d", tt.picture, got, tt.want)
			}
		})
	}
}

func BenchmarkFindLonelyPixel(b *testing.B) {
	// 构造 500×500 的图片：主对角线上为 'B'，其余为 'W'
	picture := make([][]byte, 500)
	for i := range picture {
		picture[i] = make([]byte, 500)
		for j := range picture[i] {
			picture[i][j] = 'W'
		}
		picture[i][i] = 'B'
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindLonelyPixel(picture)
	}
}
