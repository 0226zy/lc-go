package populatingnextright

import (
	"math"
	"reflect"
	"testing"
)

// newNodeTree 从层序遍历切片构建带 Next 指针的二叉树，math.MinInt32 表示 nil 节点
func newNodeTree(vals []int) *Node {
	if len(vals) == 0 {
		return nil
	}
	root := &Node{Val: vals[0]}
	queue := []*Node{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != math.MinInt32 {
			node.Left = &Node{Val: vals[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != math.MinInt32 {
			node.Right = &Node{Val: vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// collectNextPointers 收集每一层各节点 Next 指向的值（nil 用 -1 表示），用于验证连接结果
func collectNextPointers(root *Node) [][]int {
	var result [][]int
	for curr := root; curr != nil; {
		var nextStart *Node
		var nextVals []int
		for node := curr; node != nil; node = node.Next {
			// 记录本层各节点 next 指向的值，nil 用 -1 表示
			if node.Next == nil {
				nextVals = append(nextVals, -1)
			} else {
				nextVals = append(nextVals, node.Next.Val)
			}
			// 找到下一层最左的起点
			if nextStart == nil {
				if node.Left != nil {
					nextStart = node.Left
				} else if node.Right != nil {
					nextStart = node.Right
				}
			}
		}
		result = append(result, nextVals)
		curr = nextStart
	}
	return result
}

func TestConnect(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want [][]int
	}{
		// LeetCode 官方示例：非完美二叉树
		{
			"示例1: [1,2,3,4,5,null,7]",
			[]int{1, 2, 3, 4, 5, math.MinInt32, 7},
			[][]int{{-1}, {3, -1}, {5, 7, -1}},
		},
		// LeetCode 官方示例：空树
		{"示例2: 空树", []int{}, nil},
		// 边界：单节点
		{"单节点", []int{1}, [][]int{{-1}}},
		// 只有左子树的链
		{"只有左子树", []int{1, 2, math.MinInt32, 4}, [][]int{{-1}, {-1}, {-1}}},
		// 只有右子树的链
		{"只有右子树", []int{1, math.MinInt32, 3, math.MinInt32, 7}, [][]int{{-1}, {-1}, {-1}}},
		// 跨子树连接：左子树最右节点连到右子树最左节点
		{
			"跨子树连接",
			[]int{1, 2, 3, 4, 5, 6, 7},
			[][]int{{-1}, {3, -1}, {5, 6, 7, -1}},
		},
		// 左子树缺右孩子，需要找到右子树的左孩子作为 next
		{
			"左缺右跨到右左孩子",
			[]int{1, 2, 3, 4, math.MinInt32, math.MinInt32, 7},
			[][]int{{-1}, {3, -1}, {7, -1}},
		},
		// 某层前面的节点没有孩子，需要沿 next 链找到本层有孩子的节点作为下一层起点
		{
			"沿next寻找下一层起点",
			[]int{1, 2, 3, math.MinInt32, math.MinInt32, 4},
			[][]int{{-1}, {3, -1}, {-1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := newNodeTree(tt.vals)
			root = Connect(root)
			got := collectNextPointers(root)
			if tt.want == nil {
				if got != nil {
					t.Errorf("Connect(%v) 的 next 指针 = %v, want nil", tt.vals, got)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connect(%v) 的 next 指针 = %v, want %v", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkConnect(b *testing.B) {
	// 构建一个每层节点数翻倍的完全二叉树，规模约 4095 个节点
	var vals []int
	for size := 1; len(vals)+size <= 4095; size *= 2 {
		for i := 0; i < size; i++ {
			vals = append(vals, i+1)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		root := newNodeTree(vals)
		Connect(root)
	}
}
