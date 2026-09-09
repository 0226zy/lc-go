package pascalstriangleii

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

var getRowCases = []struct {
	name     string
	rowIndex int
	want     []int
}{
	{name: "第0行（示例2）", rowIndex: 0, want: []int{1}},
	{name: "第1行（示例3）", rowIndex: 1, want: []int{1, 1}},
	{name: "第2行", rowIndex: 2, want: []int{1, 2, 1}},
	{name: "第3行（示例1）", rowIndex: 3, want: []int{1, 3, 3, 1}},
	{name: "第4行", rowIndex: 4, want: []int{1, 4, 6, 4, 1}},
	{name: "第5行", rowIndex: 5, want: []int{1, 5, 10, 10, 5, 1}},
}

func TestGetRow(t *testing.T) {
	for _, tt := range getRowCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRow(tt.rowIndex); !utils.EqualIntSlice(got, tt.want) {
				t.Errorf("GetRow(%d) = %v, want %v", tt.rowIndex, got, tt.want)
			}
		})
	}
}

func TestGetRowOptimized(t *testing.T) {
	for _, tt := range getRowCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRowOptimized(tt.rowIndex); !utils.EqualIntSlice(got, tt.want) {
				t.Errorf("GetRowOptimized(%d) = %v, want %v", tt.rowIndex, got, tt.want)
			}
		})
	}
}

func BenchmarkGetRow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRow(33)
	}
}

func BenchmarkGetRowOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRowOptimized(33)
	}
}
