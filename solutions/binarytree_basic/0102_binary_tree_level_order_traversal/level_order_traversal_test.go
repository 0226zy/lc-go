package levelorder

import (
	"math"
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		vals []int // 层序切片，math.MinInt32 表示空节点
		want [][]int
	}{
		// LeetCode 官方示例
		{"示例1: 普通二叉树",
			[]int{3, 9, 20, math.MinInt32, math.MinInt32, 15, 7},
			[][]int{{3}, {9, 20}, {15, 7}}},
		{"示例2: 单节点树",
			[]int{1},
			[][]int{{1}}},
		{"示例3: 空树",
			nil,
			nil},

		// 边界情况
		{"链式树: 只有左孩子",
			[]int{1, 2, math.MinInt32, 3},
			[][]int{{1}, {2}, {3}}},
		{"链式树: 只有右孩子",
			[]int{1, math.MinInt32, 2, math.MinInt32, 3},
			[][]int{{1}, {2}, {3}}},
		{"含负数与零的节点值",
			[]int{-1, 0, -3},
			[][]int{{-1}, {0, -3}}},
		{"满二叉树",
			[]int{1, 2, 3, 4, 5, 6, 7},
			[][]int{{1}, {2, 3}, {4, 5, 6, 7}}},
		{"第二层只有右孩子",
			[]int{1, 2, math.MinInt32, math.MinInt32, 3},
			[][]int{{1}, {2}, {3}}},
		{"单层多个孩子缺口",
			[]int{5, math.MinInt32, 8, 4, math.MinInt32},
			[][]int{{5}, {8}, {4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			got := LevelOrder(root)
			// 空树时两种实现都应返回 nil，避免 nil 与空切片比较失败
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelOrder() = %v, want %v", got, tt.want)
			}
			// 与标准库 TreeNode.LevelOrder() 交叉验证
			if std := root.LevelOrder(); !reflect.DeepEqual(got, std) {
				t.Errorf("LevelOrder() = %v, 与 TreeNode.LevelOrder() = %v 不一致", got, std)
			}
		})
	}
}

func BenchmarkLevelOrder(b *testing.B) {
	benchmarks := []struct {
		name string
		vals []int
	}{
		{"7层满二叉树", fullTreeVals(7)},
		{"10层满二叉树", fullTreeVals(10)},
		{"14层满二叉树", fullTreeVals(14)},
	}
	for _, bm := range benchmarks {
		root := datastructures.NewTreeFromSlice(bm.vals)
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LevelOrder(root)
			}
		})
	}
}

// fullTreeVals 生成 depth 层满二叉树的层序切片（节点值依次为 1..2^depth-1）
func fullTreeVals(depth int) []int {
	n := (1 << depth) - 1
	vals := make([]int, n)
	for i := range vals {
		vals[i] = i + 1
	}
	return vals
}
