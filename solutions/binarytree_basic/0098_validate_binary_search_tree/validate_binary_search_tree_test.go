package isvalidbst

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name string
		vals []int // 层序构造数组，math.MinInt32 表示空节点
		want bool
	}{
		// LeetCode 官方示例
		{"示例1: 合法BST [2,1,3]", []int{2, 1, 3}, true},
		{"示例2: 右子树藏小值 [5,1,4,空,空,3,6]", []int{5, 1, 4, math.MinInt32, math.MinInt32, 3, 6}, false},

		// 边界：空树与单节点
		{"空树视为合法", nil, true},
		{"单节点", []int{42}, true},

		// 边界：链式树（退化为链表）
		{"右链合法", []int{1, math.MinInt32, 2, math.MinInt32, 3}, true},
		{"左链非法", []int{3, 2, math.MinInt32, 4}, false},

		// 边界：相等值不合法（BST 要求严格小于/大于）
		{"根与左孩子相等", []int{1, 1}, false},
		{"根与右孩子相等", []int{2, math.MinInt32, 2}, false},

		// 边界：含负数
		{"含负数合法", []int{0, -3, 9, math.MinInt32, -1, 5}, true},
		{"含负数非法", []int{0, -3, 9, math.MinInt32, math.MinInt32, -5}, false},

		// 经典陷阱：深层违规——只看父子关系无法发现，必须携带上下界
		{"深层违规：右子树中藏有不大于根的值", []int{10, 5, 15, math.MinInt32, math.MinInt32, 6, 20}, false},

		// 边界：int32 极值（检验哨兵边界不会误判）
		{"int32最大值单节点", []int{math.MaxInt32}, true},
		{"int32最小值作根", []int{math.MinInt32}, true},
		{"右孩子为int32最大值合法", []int{math.MaxInt32 - 1, math.MinInt32, math.MaxInt32}, true},

		// 较大的合法 BST
		{"较大合法BST", []int{8, 3, 10, 1, 6, math.MinInt32, 14, math.MinInt32, math.MinInt32, 4, 7, 13}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			if got := IsValidBST(root); got != tt.want {
				t.Errorf("IsValidBST(%v) = %v, want %v", tt.vals, got, tt.want)
			}
		})
	}
}

// buildValidBST 构造深度为 depth 的满合法 BST（值域 1..2^depth-1），用于基准测试
func buildValidBST(depth int) *datastructures.TreeNode {
	var build func(lo, hi int) *datastructures.TreeNode
	build = func(lo, hi int) *datastructures.TreeNode {
		if lo > hi {
			return nil
		}
		mid := (lo + hi) / 2
		return &datastructures.TreeNode{
			Val:   mid,
			Left:  build(lo, mid-1),
			Right: build(mid+1, hi),
		}
	}
	return build(1, (1<<depth)-1)
}

func BenchmarkIsValidBST(b *testing.B) {
	benchmarks := []struct {
		name  string
		depth int
	}{
		{"深度10共1023节点", 10},
		{"深度14共16383节点", 14},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			root := buildValidBST(bm.depth)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				IsValidBST(root)
			}
		})
	}
}
