package numberofprovinces

import "testing"

func TestFindCircleNum(t *testing.T) {
	tests := []struct {
		name        string
		isConnected [][]int
		want        int
	}{
		// LeetCode 官方示例
		{
			"示例1: 两组相连城市",
			[][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			2,
		},
		{
			"示例2: 三座城市互不相连",
			[][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			3,
		},

		// 边界情况
		{"单个城市", [][]int{{1}}, 1},
		{
			"两城市相连",
			[][]int{
				{1, 1},
				{1, 1},
			},
			1,
		},
		{
			"两城市不相连",
			[][]int{
				{1, 0},
				{0, 1},
			},
			2,
		},
		{
			"全部相连构成一个大省",
			[][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			1,
		},

		// 间接相连：0 与 2 不直接相连，但通过 1 间接相连
		{
			"间接相连算同一个省",
			[][]int{
				{1, 1, 0},
				{1, 1, 1},
				{0, 1, 1},
			},
			1,
		},
		{
			"链式相连与孤立城市混合",
			[][]int{
				{1, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 1},
			},
			2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindCircleNum(tt.isConnected); got != tt.want {
				t.Errorf("FindCircleNum() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestFindCircleNumUF(t *testing.T) {
	tests := []struct {
		name        string
		isConnected [][]int
		want        int
	}{
		{
			"示例1: 两组相连城市",
			[][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			2,
		},
		{
			"示例2: 三座城市互不相连",
			[][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			3,
		},
		{"单个城市", [][]int{{1}}, 1},
		{
			"全部相连构成一个大省",
			[][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			1,
		},
		{
			"间接相连算同一个省",
			[][]int{
				{1, 1, 0},
				{1, 1, 1},
				{0, 1, 1},
			},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindCircleNumUF(tt.isConnected); got != tt.want {
				t.Errorf("FindCircleNumUF() = %d, want %d", got, tt.want)
			}
		})
	}
}

// makeMatrix 生成 n x n 的邻接矩阵：full 为 true 时所有城市两两相连，否则完全孤立
func makeMatrix(n int, full bool) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			if i == j || full {
				matrix[i][j] = 1
			}
		}
	}
	return matrix
}

func BenchmarkFindCircleNum(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
		full bool
	}{
		{"20城全孤立", 20, false},
		{"100城全相连", 100, true},
		{"200城全相连", 200, true},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			matrix := makeMatrix(bm.n, bm.full)
			for i := 0; i < b.N; i++ {
				FindCircleNum(matrix)
			}
		})
	}
}

func BenchmarkFindCircleNumUF(b *testing.B) {
	matrix := makeMatrix(200, true)
	for i := 0; i < b.N; i++ {
		FindCircleNumUF(matrix)
	}
}
