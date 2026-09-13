package zigzaglevelorder

import (
	"math"
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestZigzagLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want [][]int
	}{
		// LeetCode 官方示例
		{"示例1: 三层锯齿形", []int{3, 9, 20, math.MinInt32, math.MinInt32, 15, 7}, [][]int{{3}, {20, 9}, {15, 7}}},
		{"示例2: 单节点树", []int{1}, [][]int{{1}}},
		{"示例3: 空树", nil, nil},

		// 边界情况
		{"全左链式树", []int{1, 2, math.MinInt32, 3, math.MinInt32, 4}, [][]int{{1}, {2}, {3}, {4}}},
		{"全右链式树", []int{1, math.MinInt32, 2, math.MinInt32, 3}, [][]int{{1}, {2}, {3}}},
		{"满二叉树三层", []int{1, 2, 3, 4, 5, 6, 7}, [][]int{{1}, {3, 2}, {4, 5, 6, 7}}},
		{"满二叉树四层", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			[][]int{{1}, {3, 2}, {4, 5, 6, 7}, {15, 14, 13, 12, 11, 10, 9, 8}}},
		{"含负数节点", []int{-1, -2, -3, -4, math.MinInt32, math.MinInt32, -5},
			[][]int{{-1}, {-3, -2}, {-4, -5}}},
		{"单侧缺孩子的不平衡树", []int{1, 2, 3, 4, math.MinInt32, math.MinInt32, 6},
			[][]int{{1}, {3, 2}, {4, 6}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			got := ZigzagLevelOrder(root)
			if tt.want == nil {
				// 空树允许返回 nil 或空切片，二者语义一致
				if len(got) != 0 {
					t.Errorf("ZigzagLevelOrder(%v) = %v, 期望空结果", tt.vals, got)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZigzagLevelOrder(%v) = %v, 期望 %v", tt.vals, got, tt.want)
			}
		})
	}
}

// buildChain 构造长度为 n 的左链式树，节点值依次为 1..n
func buildChain(n int) *datastructures.TreeNode {
	if n == 0 {
		return nil
	}
	root := &datastructures.TreeNode{Val: 1}
	cur := root
	for i := 2; i <= n; i++ {
		cur.Left = &datastructures.TreeNode{Val: i}
		cur = cur.Left
	}
	return root
}

func BenchmarkZigzagLevelOrder(b *testing.B) {
	root := buildChain(2000) // 题目允许的最大规模
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ZigzagLevelOrder(root)
	}
}
