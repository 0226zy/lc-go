package diameterofbinarytree

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	N := math.MinInt32 // 表示空节点
	tests := []struct {
		name string
		vals []int // 层序遍历切片，N 表示 nil 空节点
		want int
	}{
		// LeetCode 官方示例
		{"示例1: 树[1,2,3,4,5]直径为3(路径4-2-1-3或5-2-1-3)",
			[]int{1, 2, 3, 4, 5}, 3},
		{"示例2: 树[1,2]直径为1",
			[]int{1, 2}, 1},

		// 边界情况
		{"空树直径为0", nil, 0},
		{"单节点树直径为0", []int{1}, 0},
		{"左斜链式树直径为3", []int{1, 2, N, 3, N, 4}, 3},
		{"右斜链式树直径为3", []int{1, N, 2, N, 3, N, 4}, 3},
		{"完全二叉树[1..7]直径为4(路径4-2-1-3-7)",
			[]int{1, 2, 3, 4, 5, 6, 7}, 4},
		{"含负数节点的树", []int{-1, -2, -3, -4, -5}, 3},

		// 直径不经过根节点的情况
		{"直径不经过根节点(路径6-5-4-2-3)",
			[]int{1, 2, N, 3, 4, N, N, N, 5, 6}, 4},
		// 左右子树严重不对称
		{"左子树很深的树", []int{1, 2, 3, 4, N, N, N, 5, N, 6}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			if got := DiameterOfBinaryTree(root); got != tt.want {
				t.Errorf("DiameterOfBinaryTree(%v) = %d, want %d", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkDiameterOfBinaryTree(b *testing.B) {
	benchmarks := []struct {
		name string
		vals []int
	}{
		{"深度15的满二叉树", buildFullTree(15)},
		{"深度20的满二叉树", buildFullTree(20)},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			root := datastructures.NewTreeFromSlice(bm.vals)
			for i := 0; i < b.N; i++ {
				DiameterOfBinaryTree(root)
			}
		})
	}
}

// buildFullTree 生成深度为 depth 的满二叉树的层序切片（节点数为 2^depth - 1）
func buildFullTree(depth int) []int {
	size := 1<<depth - 1
	vals := make([]int, size)
	for i := range vals {
		vals[i] = i + 1
	}
	return vals
}
