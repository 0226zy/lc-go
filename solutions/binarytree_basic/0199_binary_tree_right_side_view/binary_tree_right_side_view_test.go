package rightsiderview

import (
	"math"
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// 右视图测试用例集合，BFS 与 DFS 两种实现共用
var rightSideViewCases = []struct {
	name string
	vals []int
	want []int
}{
	// LeetCode 官方示例
	{"示例1: 根左右子树交错缺子", []int{1, 2, 3, math.MinInt32, 5, math.MinInt32, 4}, []int{1, 3, 4}},
	{"示例2: 只有右子树", []int{1, math.MinInt32, 3}, []int{1, 3}},
	{"示例3: 空树", nil, nil},

	// 边界：空树与单节点
	{"空树(nil 切片)", []int{}, nil},
	{"单节点", []int{1}, []int{1}},

	// 边界：最右节点来自左子树
	{"左子树更深, 底层看到左子树节点", []int{1, 2, 3, 4}, []int{1, 3, 4}},
	{"只有左子树的链式树", []int{1, 2, math.MinInt32, 3}, []int{1, 2, 3}},
	{"只有右子树的链式树", []int{1, math.MinInt32, 2, math.MinInt32, 3}, []int{1, 2, 3}},

	// 常规情况
	{"满二叉树", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 3, 7}},
	{"交错缺子, 底层最右是5", []int{1, 2, 3, math.MinInt32, 4, 5}, []int{1, 3, 5}},

	// 边界：含负数节点值
	{"含负数节点值", []int{-1, -2, -3, math.MinInt32, -5, math.MinInt32, -4}, []int{-1, -3, -4}},
}

func TestRightSideView(t *testing.T) {
	for _, tt := range rightSideViewCases {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			got := RightSideView(root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RightSideView(%v) = %v, want %v", tt.vals, got, tt.want)
			}
		})
	}
}

func TestRightSideViewDFS(t *testing.T) {
	for _, tt := range rightSideViewCases {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			got := RightSideViewDFS(root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RightSideViewDFS(%v) = %v, want %v", tt.vals, got, tt.want)
			}
		})
	}
}

// TestRightSideView一致性 验证 BFS 与 DFS 两种实现结果一致
func TestRightSideView一致性(t *testing.T) {
	for _, tt := range rightSideViewCases {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			gotBFS := RightSideView(root)
			gotDFS := RightSideViewDFS(root)
			if !reflect.DeepEqual(gotBFS, gotDFS) {
				t.Errorf("BFS 与 DFS 结果不一致: BFS=%v, DFS=%v", gotBFS, gotDFS)
			}
		})
	}
}

// buildCompleteTree 构造深度为 depth 的满二叉树，用于基准测试
func buildCompleteTree(depth int) *datastructures.TreeNode {
	val := 1
	var build func(d int) *datastructures.TreeNode
	build = func(d int) *datastructures.TreeNode {
		if d == 0 {
			return nil
		}
		node := &datastructures.TreeNode{Val: val}
		val++
		node.Left = build(d - 1)
		node.Right = build(d - 1)
		return node
	}
	return build(depth)
}

func BenchmarkRightSideView(b *testing.B) {
	benchmarks := []struct {
		name  string
		depth int
	}{
		{"BFS_depth=10(1023节点)", 10},
		{"DFS_depth=10(1023节点)", 10},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			root := buildCompleteTree(bm.depth)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if bm.name[0] == 'B' { // BFS 实现
					RightSideView(root)
				} else { // DFS 实现
					RightSideViewDFS(root)
				}
			}
		})
	}
}
