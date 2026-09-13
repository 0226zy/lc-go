package searchbst

import (
	"math"
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestSearchBST(t *testing.T) {
	min := math.MinInt32
	tests := []struct {
		name string
		vals []int // 层序构造，math.MinInt32 表示 nil 节点
		val  int
		want [][]int // 期望子树的层序遍历结果，nil 表示未找到
	}{
		// LeetCode 官方示例
		{"示例1: 在[4,2,7,1,3]中查找2", []int{4, 2, 7, 1, 3}, 2, [][]int{{2}, {1, 3}}},
		{"示例2: 在[4,2,7,1,3]中查找5", []int{4, 2, 7, 1, 3}, 5, nil},

		// 边界情况
		{"空树", nil, 1, nil},
		{"单节点命中", []int{7}, 7, [][]int{{7}}},
		{"单节点未命中", []int{7}, 3, nil},
		{"查找根节点", []int{4, 2, 7, 1, 3}, 4, [][]int{{4}, {2, 7}, {1, 3}}},
		{"查找叶子节点", []int{4, 2, 7, 1, 3}, 3, [][]int{{3}}},

		// 链式树（退化情况）
		{"左链式树查找最深节点", []int{5, 3, min, 1}, 1, [][]int{{1}}},
		{"右链式树查找最深节点", []int{1, min, 3, min, 5}, 5, [][]int{{5}}},
		{"左链式树未命中", []int{5, 3, min, 1}, 4, nil},

		// 含负数
		{"含负数查找负值", []int{0, -10, 10, -20, -5}, -5, [][]int{{-5}}},
		{"含负数未命中", []int{0, -10, 10}, 7, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			got := SearchBST(root, tt.val)
			if got == nil {
				if tt.want != nil {
					t.Errorf("SearchBST(root, %d) = nil, want %v", tt.val, tt.want)
				}
				return
			}
			if tt.want == nil {
				t.Errorf("SearchBST(root, %d) = %v, want nil", tt.val, got.LevelOrder())
				return
			}
			if !reflect.DeepEqual(got.LevelOrder(), tt.want) {
				t.Errorf("SearchBST(root, %d) = %v, want %v", tt.val, got.LevelOrder(), tt.want)
			}
		})
	}
}

// buildBalancedBST 从升序切片构造平衡 BST，用于基准测试
func buildBalancedBST(sorted []int) *datastructures.TreeNode {
	if len(sorted) == 0 {
		return nil
	}
	mid := len(sorted) / 2
	return &datastructures.TreeNode{
		Val:   sorted[mid],
		Left:  buildBalancedBST(sorted[:mid]),
		Right: buildBalancedBST(sorted[mid+1:]),
	}
}

func BenchmarkSearchBST(b *testing.B) {
	// 构造一棵含 16383 个节点的满二叉搜索树（高度 14）
	n := 1<<14 - 1
	sorted := make([]int, n)
	for i := range sorted {
		sorted[i] = i
	}
	root := buildBalancedBST(sorted)

	b.Run("查找最深叶子", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SearchBST(root, n-1)
		}
	})
	b.Run("查找不存在元素", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SearchBST(root, -1)
		}
	})
}
