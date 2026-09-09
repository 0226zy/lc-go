package maximalsquare

import "testing"

var maximalSquareCases = []struct {
	name   string
	matrix [][]byte
	want   int
}{
	{
		name: "示例1：边长为2的正方形",
		matrix: [][]byte{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		},
		want: 4,
	},
	{
		name: "示例2：只有单个1",
		matrix: [][]byte{
			{'0', '1'},
			{'1', '0'},
		},
		want: 1,
	},
	{
		name:   "示例3：单个0",
		matrix: [][]byte{{'0'}},
		want:   0,
	},
	{
		name:   "单个1",
		matrix: [][]byte{{'1'}},
		want:   1,
	},
	{
		name: "全1矩阵",
		matrix: [][]byte{
			{'1', '1', '1'},
			{'1', '1', '1'},
			{'1', '1', '1'},
		},
		want: 9,
	},
	{
		name: "全0矩阵",
		matrix: [][]byte{
			{'0', '0'},
			{'0', '0'},
		},
		want: 0,
	},
	{
		name: "正方形在右下角",
		matrix: [][]byte{
			{'1', '0', '1'},
			{'1', '1', '1'},
			{'0', '1', '1'},
		},
		want: 4,
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
			if got := MaximalSquareOptimized(tt.matrix); got != tt.want {
				t.Errorf("MaximalSquareOptimized() = %d, want %d", got, tt.want)
			}
		})
	}
}

var benchmarkMatrix = func() [][]byte {
	// 构造 300x300 全 1 矩阵作为压力场景
	m := make([][]byte, 300)
	for i := range m {
		m[i] = make([]byte, 300)
		for j := range m[i] {
			m[i][j] = '1'
		}
	}
	return m
}()

func BenchmarkMaximalSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaximalSquare(benchmarkMatrix)
	}
}

func BenchmarkMaximalSquareOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaximalSquareOptimized(benchmarkMatrix)
	}
}
