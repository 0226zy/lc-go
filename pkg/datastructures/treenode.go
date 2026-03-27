package datastructures

import (
	"fmt"
	"math"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeFromSlice 从层序遍历切片构建二叉树，math.MinInt32 表示 nil 节点
func NewTreeFromSlice(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != math.MinInt32 {
			node.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != math.MinInt32 {
			node.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// LevelOrder 层序遍历返回切片
func (root *TreeNode) LevelOrder() [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		var level []int
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}

// String 以可读格式打印树
func (root *TreeNode) String() string {
	if root == nil {
		return "nil"
	}
	var sb strings.Builder
	printTree(&sb, root, "", false)
	return sb.String()
}

func printTree(sb *strings.Builder, node *TreeNode, prefix string, isLeft bool) {
	if node == nil {
		return
	}
	if isLeft {
		sb.WriteString(prefix + "├── ")
	} else {
		sb.WriteString(prefix + "└── ")
	}
	sb.WriteString(fmt.Sprintf("%d\n", node.Val))

	newPrefix := prefix
	if isLeft {
		newPrefix += "│   "
	} else {
		newPrefix += "    "
	}
	printTree(sb, node.Left, newPrefix, true)
	printTree(sb, node.Right, newPrefix, false)
}
