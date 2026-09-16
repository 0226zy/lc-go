package findleavesofbinarytree

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
)

// nilV 表示层序切片中的空节点
const nilV = math.MinInt32

func TestFindLeaves(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want [][]int
	}{
		// LeetCode 官方示例
		{"示例1: [1,2,3,4,5]", []int{1, 2, 3, 4, 5}, [][]int{{4, 5, 3}, {2}, {1}}},
		{"示例2: 单节点 [1]", []int{1}, [][]int{{1}}},

		// 边界：空树
		{"空树", nil, [][]int{}},

		// 边界：只有左子树链（树退化为链表）
		{"左斜链", []int{1, 2, nilV, 3, nilV, 4}, [][]int{{4}, {3}, {2}, {1}}},

		// 边界：完全二叉树
		{"满二叉树", []int{1, 2, 3, 4, 5, 6, 7}, [][]int{{4, 5, 6, 7}, {2, 3}, {1}}},

		// 典型场景：左右子树高度不同
		//      1
		//     / \
		//    2   3
		//       / \
		//      4   5
		// 第0轮收集 2、4、5；第1轮收集 3；第2轮收集 1
		{"左右高度不等", []int{1, 2, 3, nilV, nilV, 4, 5}, [][]int{{2, 4, 5}, {3}, {1}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			got := FindLeaves(root)
			if !utils.Equal2DIntSlice(got, tt.want) {
				t.Errorf("FindLeaves(%v) = %v, want %v", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkFindLeaves(b *testing.B) {
	// 构造一棵高为 20 的满二叉树（约百万节点太大，用高 15 约 32767 个节点）
	depth := 15
	vals := make([]int, (1<<depth)-1)
	for i := range vals {
		vals[i] = i + 1
	}

	benchmarks := []struct {
		name  string
		depth int
	}{
		{"depth=4", 4},
		{"depth=8", 8},
		{"depth=12", 12},
		{"depth=15", depth},
	}

	for _, bm := range benchmarks {
		root := datastructures.NewTreeFromSlice(vals[:(1<<bm.depth)-1])
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindLeaves(root)
			}
		})
	}
}
