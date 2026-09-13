package invertbinarytree

import (
	"math"
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestInvertTree(t *testing.T) {
	nilNode := math.MinInt32 // NewTreeFromSlice 中用 math.MinInt32 表示空节点
	tests := []struct {
		name string
		vals []int // 原树的层序序列
		want []int // 翻转后树的层序序列
	}{
		// LeetCode 官方示例
		{"示例1: 三层满二叉树", []int{4, 2, 7, 1, 3, 6, 9}, []int{4, 7, 2, 9, 6, 3, 1}},
		{"示例2: 三节点小树", []int{2, 1, 3}, []int{2, 3, 1}},
		{"示例3: 空树", nil, nil},

		// 边界情况
		{"单节点树", []int{1}, []int{1}},
		{"左斜链式树", []int{1, 2, nilNode, 3}, []int{1, nilNode, 2, nilNode, 3}},
		{"右斜链式树", []int{1, nilNode, 2, nilNode, 3}, []int{1, 2, nilNode, 3}},
		{"含负数的树", []int{-1, -2, -3}, []int{-1, -3, -2}},
		{"只有左孩子的根节点", []int{1, 2}, []int{1, nilNode, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			got := InvertTree(root).LevelOrder()
			want := datastructures.NewTreeFromSlice(tt.want).LevelOrder()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("InvertTree(%v) 层序结果 = %v, 期望 %v", tt.vals, got, want)
			}
		})
	}
}

// TestInvertTreeTwice 连续翻转两次应还原原树，验证翻转是幂等逆操作
func TestInvertTreeTwice(t *testing.T) {
	vals := []int{4, 2, 7, 1, 3, 6, 9}
	root := datastructures.NewTreeFromSlice(vals)
	restored := InvertTree(InvertTree(root)).LevelOrder()
	original := datastructures.NewTreeFromSlice(vals).LevelOrder()
	if !reflect.DeepEqual(restored, original) {
		t.Errorf("翻转两次后 = %v, 期望还原为 %v", restored, original)
	}
}

func BenchmarkInvertTree(b *testing.B) {
	// 构造深度为 10 的满二叉树（1023 个节点）的层序序列
	n := 1<<10 - 1
	vals := make([]int, n)
	for i := range vals {
		vals[i] = i + 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		root := datastructures.NewTreeFromSlice(vals)
		InvertTree(root)
	}
}
