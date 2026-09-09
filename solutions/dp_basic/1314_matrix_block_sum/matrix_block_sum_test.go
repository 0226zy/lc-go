package matrixblocksum

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

var matrixBlockSumCases = []struct {
	name string
	mat  [][]int
	k    int
	want [][]int
}{
	{
		name: "示例1：3x3矩阵k=1",
		mat:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		k:    1,
		want: [][]int{{12, 21, 16}, {27, 45, 33}, {24, 39, 28}},
	},
	{
		name: "示例2：3x3矩阵k=2",
		mat:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		k:    2,
		want: [][]int{{45, 45, 45}, {45, 45, 45}, {45, 45, 45}},
	},
	{
		name: "单元素矩阵",
		mat:  [][]int{{7}},
		k:    1,
		want: [][]int{{7}},
	},
	{
		name: "k=0只含自身",
		mat:  [][]int{{1, 2}, {3, 4}},
		k:    0,
		want: [][]int{{1, 2}, {3, 4}},
	},
	{
		name: "单行矩阵",
		mat:  [][]int{{1, 2, 3, 4}},
		k:    1,
		want: [][]int{{3, 6, 9, 7}},
	},
	{
		name: "单列矩阵",
		mat:  [][]int{{1}, {2}, {3}, {4}},
		k:    1,
		want: [][]int{{3}, {6}, {9}, {7}},
	},
}

func TestMatrixBlockSum(t *testing.T) {
	for _, tt := range matrixBlockSumCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatrixBlockSum(tt.mat, tt.k); !utils.Equal2DIntSlice(got, tt.want) {
				t.Errorf("MatrixBlockSum(%v, %d) = %v, want %v", tt.mat, tt.k, got, tt.want)
			}
		})
	}
}

func BenchmarkMatrixBlockSum(b *testing.B) {
	mat := make([][]int, 100)
	for i := range mat {
		mat[i] = make([]int, 100)
		for j := range mat[i] {
			mat[i][j] = (i*100+j)%100 + 1
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MatrixBlockSum(mat, 50)
	}
}
