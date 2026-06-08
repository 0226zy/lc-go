package binarytree

import (
	"github.com/0226zy/lc-go/pkg/datastructures"
)

// TreeTraversal 封装二叉树常用遍历算法
type TreeTraversal struct{}

// NewTreeTraversal 创建一个新的 TreeTraversal 实例
func NewTreeTraversal() *TreeTraversal {
	return &TreeTraversal{}
}

// PreorderRecursive 前序遍历（递归）
// 遍历顺序: 根 -> 左 -> 右
func (tt *TreeTraversal) PreorderRecursive(root *datastructures.TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := []int{}
	var dfs func(*datastructures.TreeNode)
	dfs = func(node *datastructures.TreeNode) {
		if node == nil {
			return
		}
		ret = append(ret, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ret
}

// PreorderIterative 前序遍历（非递归）
// 使用栈模拟递归过程
func (tt *TreeTraversal) PreorderIterative(root *datastructures.TreeNode) []int {
	if root == nil {
		return nil
	}
	ret, stack := []int{}, []*datastructures.TreeNode{root}
	for len(stack) > 0 {
		n := len(stack)
		node := stack[n-1]
		stack = stack[:n-1]
		ret = append(ret, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return ret
}

// InorderRecursive 中序遍历（递归）
// 遍历顺序: 左 -> 根 -> 右
func (tt *TreeTraversal) InorderRecursive(root *datastructures.TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := []int{}
	var dfs func(*datastructures.TreeNode)
	dfs = func(node *datastructures.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		ret = append(ret, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return ret
}

// InorderIterative 中序遍历（非递归）
// 使用栈模拟递归过程
func (tt *TreeTraversal) InorderIterative(root *datastructures.TreeNode) []int {
	if root == nil {
		return nil
	}
	ret, stack := []int{}, []*datastructures.TreeNode{}
	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		n := len(stack)
		curr = stack[n-1]
		stack = stack[:n-1]
		ret = append(ret, curr.Val)

		curr = curr.Right
	}
	return ret
}

// PostorderRecursive 后序遍历（递归）
// 遍历顺序: 左 -> 右 -> 根
func (tt *TreeTraversal) PostorderRecursive(root *datastructures.TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := []int{}
	var dfs func(*datastructures.TreeNode)
	dfs = func(node *datastructures.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		ret = append(ret, node.Val)
	}
	dfs(root)
	return ret
}

// PostorderIterative 后序遍历（非递归）
// 使用栈模拟递归过程
func (tt *TreeTraversal) PostorderIterative(root *datastructures.TreeNode) []int {
	if root == nil {
		return nil
	}
	ret, stack := []int{}, []*datastructures.TreeNode{}
	stack = append(stack, root)

	for len(stack) > 0 {
		n := len(stack)
		node := stack[n-1]
		stack = stack[:n-1]
		ret = append([]int{node.Val}, ret...)

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return ret
}

// LevelOrder 层序遍历
// 从上到下、从左到右依次遍历每个节点
func (tt *TreeTraversal) LevelOrder(root *datastructures.TreeNode) []int {
	if root == nil {
		return nil
	}
	ret, queue := []int{}, []*datastructures.TreeNode{root}
	for len(queue) > 0 {
		levelNum := len(queue)
		for i := 0; i < levelNum; i++ {
			node := queue[0]
			queue = queue[1:]
			ret = append(ret, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return ret
}

// LevelOrderByLevel 层序遍历，按层返回结果
// 每一层的节点值放在一个单独的切片中
func (tt *TreeTraversal) LevelOrderByLevel(root *datastructures.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret := [][]int{}
	queue := []*datastructures.TreeNode{root}
	for len(queue) > 0 {
		levelNum := len(queue)
		tmp := []int{}
		for i := 0; i < levelNum; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ret = append(ret, tmp)
	}
	return ret
}

// LevelOrderZigzag 层序遍历之字形（锯齿形）
// 奇数层从左到右，偶数层从右到左
func (tt *TreeTraversal) LevelOrderZigzag(root *datastructures.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret := [][]int{}
	queue := []*datastructures.TreeNode{root}
	leftToRight := true
	for len(queue) > 0 {
		levelNum := len(queue)
		tmp := make([]int, levelNum)
		for i := 0; i < levelNum; i++ {
			node := queue[0]
			queue = queue[1:]
			idx := i
			if !leftToRight {
				idx = levelNum - 1 - i
			}
			tmp[idx] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		leftToRight = !leftToRight
		ret = append(ret, tmp)
	}
	return ret
}
