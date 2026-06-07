package binarytreeinordertraversal

import "github.com/0226zy/lc-go/pkg/datastructures"

type TreeNode = datastructures.TreeNode

// InorderTraversalRecursive 递归法实现二叉树中序遍历
// 时间复杂度: O(n)  空间复杂度: O(h)，h 为树高
func InorderTraversalRecursive(root *TreeNode) []int {
	var result []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		result = append(result, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return result
}

// InorderTraversalIterative 迭代法实现二叉树中序遍历（显式栈）
// 时间复杂度: O(n)  空间复杂度: O(h)，h 为树高
func InorderTraversalIterative(root *TreeNode) []int {
	var result []int
	stack := []*TreeNode{}
	cur := root

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		cur = cur.Right
	}

	return result
}

// InorderTraversalMorris Morris 遍历实现二叉树中序遍历
// 利用空闲指针建立临时线索，无需额外栈空间
// 时间复杂度: O(n)  空间复杂度: O(1)
func InorderTraversalMorris(root *TreeNode) []int {
	var result []int
	cur := root

	for cur != nil {
		if cur.Left == nil {
			result = append(result, cur.Val)
			cur = cur.Right
		} else {
			// 找到左子树的最右节点（前驱）
			predecessor := cur.Left
			for predecessor.Right != nil && predecessor.Right != cur {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				// 建立临时线索
				predecessor.Right = cur
				cur = cur.Left
			} else {
				// 线索已存在，左子树遍历完毕，恢复指针
				predecessor.Right = nil
				result = append(result, cur.Val)
				cur = cur.Right
			}
		}
	}

	return result
}
