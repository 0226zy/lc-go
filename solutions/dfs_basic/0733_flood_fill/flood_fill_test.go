package floodfill

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

// copyImage 深拷贝二维图像，避免测试用例之间互相污染（FloodFill 会原地修改图像）
func copyImage(image [][]int) [][]int {
	dup := make([][]int, len(image))
	for i := range image {
		dup[i] = make([]int, len(image[i]))
		copy(dup[i], image[i])
	}
	return dup
}

func TestFloodFill(t *testing.T) {
	tests := []struct {
		name  string
		image [][]int
		sr    int
		sc    int
		color int
		want  [][]int
	}{
		// LeetCode 官方示例
		{
			"示例1: 中心扩散染色",
			[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			1, 1, 2,
			[][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
		{
			"示例2: 新颜色与原颜色相同",
			[][]int{{0, 0, 0}, {0, 0, 0}},
			0, 0, 0,
			[][]int{{0, 0, 0}, {0, 0, 0}},
		},

		// 边界：单格图像
		{
			"单格图像染色",
			[][]int{{1}},
			0, 0, 5,
			[][]int{{5}},
		},
		{
			"单格图像颜色相同",
			[][]int{{3}},
			0, 0, 3,
			[][]int{{3}},
		},

		// 边界：整幅图像同色，全部染色
		{
			"全图同色",
			[][]int{{1, 1}, {1, 1}},
			0, 0, 9,
			[][]int{{9, 9}, {9, 9}},
		},

		// 边界：对角相连不算连通，只染四连通区域
		{
			"对角不连通",
			[][]int{{1, 0}, {0, 1}},
			0, 0, 2,
			[][]int{{2, 0}, {0, 1}},
		},

		// 边界：起点在角落
		{
			"起点在右上角",
			[][]int{{1, 1}, {1, 0}},
			0, 1, 7,
			[][]int{{7, 7}, {7, 0}},
		},

		// 复杂形状：被 0 分隔的多个同色区域，只染起点所在区域
		{
			"只染连通区域",
			[][]int{{1, 0, 1}, {1, 0, 1}, {1, 0, 1}},
			1, 0, 4,
			[][]int{{4, 0, 1}, {4, 0, 1}, {4, 0, 1}},
		},

		// 新颜色恰好与其他区域颜色相同，不影响结果正确性
		{
			"新颜色与其他区域同色",
			[][]int{{1, 2}, {1, 2}},
			0, 0, 2,
			[][]int{{2, 2}, {2, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FloodFill(copyImage(tt.image), tt.sr, tt.sc, tt.color)
			if !utils.Equal2DIntSlice(got, tt.want) {
				t.Errorf("FloodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}

// makeImage 生成 rows x cols 的图像，全部填充 fill 颜色
func makeImage(rows, cols, fill int) [][]int {
	image := make([][]int, rows)
	for i := range image {
		image[i] = make([]int, cols)
		for j := range image[i] {
			image[i][j] = fill
		}
	}
	return image
}

func BenchmarkFloodFill(b *testing.B) {
	benchmarks := []struct {
		name string
		rows int
		cols int
	}{
		{"10x10全同色", 10, 10},
		{"50x50全同色", 50, 50},
		{"100x100全同色", 100, 100},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FloodFill(makeImage(bm.rows, bm.cols, 1), 0, 0, 2)
			}
		})
	}
}
