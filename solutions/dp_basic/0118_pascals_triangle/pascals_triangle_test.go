package pascalstriangle

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

var generateCases = []struct {
	name    string
	numRows int
	want    [][]int
}{
	{name: "1行（示例2）", numRows: 1, want: [][]int{{1}}},
	{name: "2行", numRows: 2, want: [][]int{{1}, {1, 1}}},
	{name: "3行", numRows: 3, want: [][]int{{1}, {1, 1}, {1, 2, 1}}},
	{name: "5行（示例1）", numRows: 5, want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
	{
		name:    "6行",
		numRows: 6,
		want:    [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}},
	},
}

func TestGenerate(t *testing.T) {
	for _, tt := range generateCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := Generate(tt.numRows); !utils.Equal2DIntSlice(got, tt.want) {
				t.Errorf("Generate(%d) = %v, want %v", tt.numRows, got, tt.want)
			}
		})
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate(30)
	}
}
