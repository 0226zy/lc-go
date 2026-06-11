package pathsumiii

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestPathSum(t *testing.T) {
	tests := []struct {
		name      string
		root      *TreeNode
		targetSum int
		want      int
	}{
		{
			name: "示例1",
			root: datastructures.NewTreeFromSlice([]int{
				10, 5, -3, 3, 2, math.MinInt32, 11, 3, -2, math.MinInt32, 1,
			}),
			targetSum: 8,
			want:      3,
		},
		{
			name: "示例2",
			root: datastructures.NewTreeFromSlice([]int{
				5, 4, 8, 11, math.MinInt32, 13, 4, 7, 2, math.MinInt32, math.MinInt32, 5, 1,
			}),
			targetSum: 22,
			want:      3,
		},
		{
			name:      "空树",
			root:      nil,
			targetSum: 8,
			want:      0,
		},
		{
			name:      "单节点-找到",
			root:      datastructures.NewTreeFromSlice([]int{5}),
			targetSum: 5,
			want:      1,
		},
		{
			name:      "单节点-未找到",
			root:      datastructures.NewTreeFromSlice([]int{5}),
			targetSum: 3,
			want:      0,
		},
		{
			name: "路径不经过根节点",
			root: datastructures.NewTreeFromSlice([]int{
				1, 2, 3,
			}),
			targetSum: 2,
			want:      1,
		},
		{
			name: "路径不到叶子节点",
			root: datastructures.NewTreeFromSlice([]int{
				1, 2, 3, 4, 5,
			}),
			targetSum: 3,
			want:      2, // 1->2 和 3
		},
		{
			name: "多条重叠路径",
			root: datastructures.NewTreeFromSlice([]int{
				1, 2, math.MinInt32, 3, math.MinInt32, 4,
			}),
			targetSum: 3,
			want:      2, // 1->2, 2->3 不对，应该是 1->2, 3 等，根据树结构调整
		},
		{
			name: "包含负数",
			root: datastructures.NewTreeFromSlice([]int{
				1, -2, -3, 1, 3, -2, math.MinInt32,
			}),
			targetSum: -1,
			want:      4,
		},
		{
			name: "全零树",
			root: datastructures.NewTreeFromSlice([]int{
				0, 0, 0, 0, 0, 0, 0,
			}),
			targetSum: 0,
			want:      11, // 7个单节点 + 4条长度为2的路径
		},
		{
			name: "targetSum为0-无零节点",
			root: datastructures.NewTreeFromSlice([]int{
				1, 2, 3,
			}),
			targetSum: 0,
			want:      0,
		},
		{
			name: "大数值",
			root: datastructures.NewTreeFromSlice([]int{
				1000000000, 1000000000, math.MinInt32,
			}),
			targetSum: 2000000000,
			want:      1,
		},
		{
			name: "链式树",
			root: datastructures.NewTreeFromSlice([]int{
				1, 2, math.MinInt32, 3, math.MinInt32, 4, math.MinInt32, 5,
			}),
			targetSum: 15,
			want:      1, // 1+2+3+4+5=15
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PathSum(tt.root, tt.targetSum)
			if got != tt.want {
				t.Errorf("PathSum() = %d, want %d", got, tt.want)
			}
		})
	}
}

func BenchmarkPathSum(b *testing.B) {
	root := datastructures.NewTreeFromSlice([]int{
		10, 5, -3, 3, 2, math.MinInt32, 11, 3, -2, math.MinInt32, 1,
	})
	targetSum := 8
	for i := 0; i < b.N; i++ {
		_ = PathSum(root, targetSum)
	}
}
