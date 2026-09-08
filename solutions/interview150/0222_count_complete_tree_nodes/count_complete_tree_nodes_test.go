package countcompletetreenodes

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// newTree 从层序切片构建二叉树，MinInt32 表示 nil
func newTree(vals []int) *datastructures.TreeNode {
	return datastructures.NewTreeFromSlice(vals)
}

func TestCountNodes(t *testing.T) {
	tests := []struct {
		name string
		vals []int // 层序表示，MinInt32 为 nil
		want int
	}{
		// LeetCode 官方示例
		{"示例1: [1,2,3,4,5,6]", []int{1, 2, 3, 4, 5, 6}, 6},
		{"示例2: 空树", nil, 0},
		{"示例3: [1]", []int{1}, 1},

		// 边界：两层满树
		{"两层满树", []int{1, 2, 3}, 3},

		// 边界：最后一层只有最左一个节点
		{"最后一层单节点", []int{1, 2, 3, 4}, 4},

		// 边界：最后一层缺最后一个节点
		{"最后一层缺末尾", []int{1, 2, 3, 4, 5, 6, math.MinInt32}, 6},

		// 边界：左斜树（完全二叉树的退化形态 1→2→3→4）
		{"左斜树", []int{1, 2, math.MinInt32, 3, math.MinInt32, 4}, 4},

		// 边界：右斜树（1→2→3）
		{"右斜树", []int{1, math.MinInt32, 2, math.MinInt32, 3}, 3},

		// 完全树：高度 4，最后一层左半满
		{"高度4完全树", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 12},

		// 满树：高度 4
		{"高度4满树", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountNodes(newTree(tt.vals)); got != tt.want {
				t.Errorf("CountNodes() = %d, want %d", got, tt.want)
			}
		})
	}
}

func BenchmarkCountNodes(b *testing.B) {
	// 满树高度 15：32767 个节点
	full15 := newTree(buildFullTreeVals(15))
	// 完全但不满的树：高度 15，最后一层只有最左一个节点
	lastLayerOne := newTree(buildAlmostFullTreeVals(15))

	benchmarks := []struct {
		name string
		root *datastructures.TreeNode
	}{
		{"满树32767节点", full15},
		{"完全树16384节点", lastLayerOne},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CountNodes(bm.root)
			}
		})
	}
}

// buildFullTreeVals 生成高度为 h 的满二叉树的层序切片
func buildFullTreeVals(h int) []int {
	n := (1 << h) - 1
	vals := make([]int, n)
	for i := 0; i < n; i++ {
		vals[i] = i + 1
	}
	return vals
}

// buildAlmostFullTreeVals 生成高度为 h 的完全二叉树的层序切片，最后一层仅有最左一个节点
func buildAlmostFullTreeVals(h int) []int {
	n := (1 << (h - 1)) // 前 h-1 层满，加 1 个节点
	vals := make([]int, 0, n)
	for i := 0; i < (1<<(h-1))-1; i++ {
		vals = append(vals, i+1)
	}
	vals = append(vals, (1 << (h - 1)))
	return vals
}
