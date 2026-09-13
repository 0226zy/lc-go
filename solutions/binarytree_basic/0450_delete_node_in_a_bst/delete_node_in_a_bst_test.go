package deletenodeinabst

import (
	"math"
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// 空节点占位符，配合 datastructures.NewTreeFromSlice 使用
const nilNode = math.MinInt32

func TestDeleteNode(t *testing.T) {
	tests := []struct {
		name string
		vals []int // 层序构造原树，nilNode 表示空节点
		key  int
		want []int // 层序构造期望树
	}{
		// LeetCode 官方示例
		{"示例1: 删除双子节点3（用后继4替换）",
			[]int{5, 3, 6, 2, 4, nilNode, 7}, 3,
			[]int{5, 4, 6, 2, nilNode, nilNode, 7}},
		{"示例2: key不存在，树不变",
			[]int{5, 3, 6, 2, 4, nilNode, 7}, 0,
			[]int{5, 3, 6, 2, 4, nilNode, 7}},

		// 边界情况
		{"空树", nil, 1, nil},
		{"单节点树删除根", []int{1}, 1, nil},
		{"单节点树key不存在", []int{1}, 2, []int{1}},
		{"删除叶节点",
			[]int{2, 1, 3}, 1,
			[]int{2, nilNode, 3}},
		{"删除只有右孩子的节点",
			[]int{5, 3, 6, 2, 4, nilNode, 7}, 6,
			[]int{5, 3, 7, 2, 4}},
		{"删除只有左孩子的节点",
			[]int{5, 3, nilNode, 2, 4}, 3,
			[]int{5, 2, nilNode, nilNode, nilNode, nilNode, 4}},
		{"删除根节点（双子节点，后继为右孩子）",
			[]int{5, 3, 6, 2, 4, nilNode, 7}, 5,
			[]int{6, 3, 7, 2, 4}},
		{"删除根节点（后继深埋在右子树左侧）",
			[]int{5, 3, 8, 2, 4, 6, 9}, 5,
			[]int{6, 3, 8, 2, 4, nilNode, 9}},

		// 链式树（退化 BST）
		{"右斜链删除中间节点",
			[]int{1, nilNode, 2, nilNode, 3}, 2,
			[]int{1, nilNode, 3}},
		{"右斜链删除根",
			[]int{1, nilNode, 2, nilNode, 3}, 1,
			[]int{2, nilNode, 3}},

		// 含负数
		{"含负数删除叶子",
			[]int{-10, -20, 5}, -20,
			[]int{-10, nilNode, 5}},
		{"含负数删除根",
			[]int{-10, -20, 5}, -10,
			[]int{5, -20, nilNode}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			got := DeleteNode(root, tt.key)
			want := datastructures.NewTreeFromSlice(tt.want)
			gotLevel, wantLevel := got.LevelOrder(), want.LevelOrder()
			if !reflect.DeepEqual(gotLevel, wantLevel) {
				t.Errorf("DeleteNode(%v, %d) 层序 = %v, 期望 %v",
					tt.vals, tt.key, gotLevel, wantLevel)
			}
		})
	}
}

// buildRightChain 构造值为 1..n 的右斜链 BST，用于压测退化场景
func buildRightChain(n int) *datastructures.TreeNode {
	if n <= 0 {
		return nil
	}
	root := &datastructures.TreeNode{Val: 1}
	cur := root
	for i := 2; i <= n; i++ {
		cur.Right = &datastructures.TreeNode{Val: i}
		cur = cur.Right
	}
	return root
}

func BenchmarkDeleteNode(b *testing.B) {
	b.Run("平衡树删除根节点", func(b *testing.B) {
		vals := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
		for i := 0; i < b.N; i++ {
			root := datastructures.NewTreeFromSlice(vals)
			DeleteNode(root, 8)
		}
	})
	b.Run("链式树删除最深节点", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			root := buildRightChain(1000)
			DeleteNode(root, 1000)
		}
	})
}
