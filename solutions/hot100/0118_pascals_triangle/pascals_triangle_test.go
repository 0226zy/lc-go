package pascalstriangle

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name    string
		numRows int
		want    [][]int
	}{
		{"官方示例1 numRows=5", 5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{"官方示例2 numRows=1", 1, [][]int{{1}}},
		{"下界 numRows=1 单行", 1, [][]int{{1}}},
		{"numRows=2 两行", 2, [][]int{{1}, {1, 1}}},
		{"numRows=3 三行", 3, [][]int{{1}, {1, 1}, {1, 2, 1}}},
		{"numRows=10 较大值", 10, [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
			{1, 5, 10, 10, 5, 1},
			{1, 6, 15, 20, 15, 6, 1},
			{1, 7, 21, 35, 35, 21, 7, 1},
			{1, 8, 28, 56, 70, 56, 28, 8, 1},
			{1, 9, 36, 84, 126, 126, 84, 36, 9, 1},
		}},
		{"上界 numRows=30", 30, nil}, // 仅验证行数与首尾，不硬编码完整结果
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Generate(tt.numRows)
			if tt.numRows == 30 {
				// 上界仅做结构校验
				if len(got) != 30 {
					t.Fatalf("Generate(30) 行数 = %d, want 30", len(got))
				}
				last := got[29]
				if len(last) != 30 {
					t.Fatalf("第 30 行长度 = %d, want 30", len(last))
				}
				if last[0] != 1 || last[29] != 1 {
					t.Fatalf("第 30 行首尾应为 1, got %d, %d", last[0], last[29])
				}
				// 验证对称性
				for j := 0; j < 15; j++ {
					if last[j] != last[29-j] {
						t.Fatalf("第 30 行不对称: last[%d]=%d, last[%d]=%d", j, last[j], 29-j, last[29-j])
					}
				}
				return
			}
			if !utils.Equal2DIntSlice(got, tt.want) {
				t.Errorf("Generate(%d) = %v, want %v", tt.numRows, got, tt.want)
			}
		})
	}
}

func TestGenerateInPlace(t *testing.T) {
	tests := []struct {
		name    string
		numRows int
		want    [][]int
	}{
		{"官方示例1 numRows=5", 5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{"官方示例2 numRows=1", 1, [][]int{{1}}},
		{"numRows=2 两行", 2, [][]int{{1}, {1, 1}}},
		{"numRows=3 三行", 3, [][]int{{1}, {1, 1}, {1, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateInPlace(tt.numRows)
			if !utils.Equal2DIntSlice(got, tt.want) {
				t.Errorf("GenerateInPlace(%d) = %v, want %v", tt.numRows, got, tt.want)
			}
		})
	}
}

func TestGenerateConsistency(t *testing.T) {
	// 验证两种实现结果一致
	for n := 1; n <= 30; n++ {
		a := Generate(n)
		b := GenerateInPlace(n)
		if !utils.Equal2DIntSlice(a, b) {
			t.Errorf("两种实现结果不一致: numRows=%d", n)
		}
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate(30)
	}
}

func BenchmarkGenerateInPlace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateInPlace(30)
	}
}
