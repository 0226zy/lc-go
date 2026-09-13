package maximalsquare

import "testing"

func toByteMatrix(rows []string) [][]byte {
	matrix := make([][]byte, len(rows))
	for i, row := range rows {
		matrix[i] = []byte(row)
	}
	return matrix
}

var maximalSquareCases = []struct {
	name   string
	matrix [][]byte
	want   int
}{
	// LeetCode 官方示例
	{
		name: "示例1：边长为2，面积4",
		matrix: toByteMatrix([]string{
			"10100",
			"10111",
			"11111",
			"10010",
		}),
		want: 4,
	},
	{
		name:   "示例2：交叉各一个1",
		matrix: toByteMatrix([]string{"01", "10"}),
		want:   1,
	},
	{
		name:   "示例3：单格0",
		matrix: toByteMatrix([]string{"0"}),
		want:   0,
	},

	// 边界
	{name: "空矩阵", matrix: [][]byte{}, want: 0},
	{name: "全0", matrix: toByteMatrix([]string{"000", "000"}), want: 0},
	{name: "全1单格", matrix: toByteMatrix([]string{"1"}), want: 1},
	{name: "全1两行两列", matrix: toByteMatrix([]string{"11", "11"}), want: 4},
	{name: "仅一个1", matrix: toByteMatrix([]string{"01", "00"}), want: 1},
	{name: "对角线1", matrix: toByteMatrix([]string{"100", "010", "001"}), want: 1},
	{
		name: "右下角边长4，面积16",
		matrix: toByteMatrix([]string{
			"11111",
			"11111",
			"11111",
			"01111",
		}),
		want: 16,
	},
	{name: "单列", matrix: toByteMatrix([]string{"1", "1", "0", "1"}), want: 1},
	{
		name: "更大正方形边长3",
		matrix: toByteMatrix([]string{
			"0111",
			"1111",
			"1111",
			"1111",
		}),
		want: 9,
	},
}

func TestMaximalSquare(t *testing.T) {
	for _, tt := range maximalSquareCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaximalSquare(tt.matrix); got != tt.want {
				t.Errorf("MaximalSquare() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestMaximalSquareOptimized(t *testing.T) {
	for _, tt := range maximalSquareCases {
		t.Run(tt.name, func(t *testing.T) {
			var matrix [][]byte
			if len(tt.matrix) > 0 {
				matrix = make([][]byte, len(tt.matrix))
				for i := range tt.matrix {
					matrix[i] = append([]byte(nil), tt.matrix[i]...)
				}
			}
			if got := MaximalSquareOptimized(matrix); got != tt.want {
				t.Errorf("MaximalSquareOptimized() = %d, want %d", got, tt.want)
			}
		})
	}
}

func BenchmarkMaximalSquare(b *testing.B) {
	matrix := toByteMatrix([]string{
		"10100101",
		"11110111",
		"11111111",
		"01111110",
		"11101111",
		"10111101",
		"11111011",
		"01111111",
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaximalSquare(matrix)
	}
}

func BenchmarkMaximalSquareOptimized(b *testing.B) {
	matrix := toByteMatrix([]string{
		"10100101",
		"11110111",
		"11111111",
		"01111110",
		"11101111",
		"10111101",
		"11111011",
		"01111111",
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaximalSquareOptimized(matrix)
	}
}
