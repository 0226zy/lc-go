package countnodes

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestCountNodes(t *testing.T) {
	tests := []struct {
		name string
		vals []int // 层序表示，math.MinInt32 表示空节点
		want int
	}{
		// LeetCode 官方示例
		{"示例1: [1,2,3,4,5,6]", []int{1, 2, 3, 4, 5, 6}, 6},
		{"示例2: 空树", nil, 0},
		{"示例3: 单节点 [1]", []int{1}, 1},

		// 边界：两层满二叉树
		{"两层满树", []int{1, 2, 3}, 3},

		// 边界：最后一层只有最左一个节点
		{"最后一层仅一个节点", []int{1, 2, 3, 4}, 4},

		// 边界：最后一层缺少最右一个节点
		{"最后一层缺末尾节点", []int{1, 2, 3, 4, 5, 6, math.MinInt32}, 6},

		// 边界：链式树（左链 1→2→3→4）
		{"左链式树", []int{1, 2, math.MinInt32, 3, math.MinInt32, 4}, 4},

		// 边界：含负数节点值，计数与节点取值无关
		{"含负数节点", []int{-5, -2, -3, 0, 4, -1}, 6},

		// 完全树：高度 4，最后一层只填了左半边
		{"高度4的完全树", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 12},

		// 满树：高度 4，共 15 个节点
		{"高度4的满树", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			if got := CountNodes(root); got != tt.want {
				t.Errorf("CountNodes() = %d, want %d", got, tt.want)
			}
		})
	}
}

func BenchmarkCountNodes(b *testing.B) {
	// 满树：高度 15，共 32767 个节点
	fullTree := datastructures.NewTreeFromSlice(buildFullTreeVals(15))
	// 完全但不满的树：高度 15，前 14 层满、最后一层仅最左一个节点，共 16384 个节点
	almostFullTree := datastructures.NewTreeFromSlice(buildAlmostFullTreeVals(15))

	benchmarks := []struct {
		name string
		root *datastructures.TreeNode
	}{
		{"满树32767节点", fullTree},
		{"完全树16384节点", almostFullTree},
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
	for i := range vals {
		vals[i] = i + 1
	}
	return vals
}

// buildAlmostFullTreeVals 生成高度为 h 的完全二叉树的层序切片，最后一层仅含最左一个节点
func buildAlmostFullTreeVals(h int) []int {
	fullPart := (1 << (h - 1)) - 1 // 前 h-1 层为满树
	vals := make([]int, 0, fullPart+1)
	for i := 0; i < fullPart; i++ {
		vals = append(vals, i+1)
	}
	vals = append(vals, fullPart+1)
	return vals
}
