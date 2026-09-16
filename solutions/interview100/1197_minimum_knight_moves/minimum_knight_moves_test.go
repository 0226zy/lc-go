package minimumknightmoves

import "testing"

func TestMinKnightMoves(t *testing.T) {
	tests := []struct {
		name string
		x    int
		y    int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: (2,1) 一步直达", 2, 1, 1},
		{"示例2: (5,5)", 5, 5, 4},

		// 边界：起点即终点
		{"原地不动", 0, 0, 0},

		// 经典绕路情形
		{"斜相邻需绕一步", 1, 1, 2},
		{"(1,0) 需要三步", 1, 0, 3},
		{"(2,2) 需要四步", 2, 2, 4},

		// 对称性：负坐标与交换坐标
		{"负坐标 (0,-1)", 0, -1, 3},
		{"负坐标 (-2,-1)", -2, -1, 1},
		{"交换坐标 (1,2)", 1, 2, 1},
		{"(3,3)", 3, 3, 2},
		{"(4,0)", 4, 0, 2},

		// 数据范围极限
		{"极限 (300,300)", 300, 300, 200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinKnightMoves(tt.x, tt.y); got != tt.want {
				t.Errorf("MinKnightMoves(%d, %d) = %d, want %d", tt.x, tt.y, got, tt.want)
			}
		})
	}
}

func BenchmarkMinKnightMoves(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinKnightMoves(300, 300)
	}
}
