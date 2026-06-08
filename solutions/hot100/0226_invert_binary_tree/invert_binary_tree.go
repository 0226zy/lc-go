package invertbinarytree

import (
	"github.com/0226zy/lc-go/pkg/datastructures"
)

// InvertTree 翻转二叉树（递归 DFS）
// 给你一棵二叉树的根节点 root，翻转这棵二叉树，并返回其根节点。
// 时间复杂度: O(n)  空间复杂度: O(h)
func InvertTree(root *datastructures.TreeNode) *datastructures.TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	InvertTree(root.Left)
	InvertTree(root.Right)
	return root
}

// InvertTreeBFS 翻转二叉树（BFS 层序遍历）
// 使用队列逐层遍历，对每个节点交换其左右子树。
// 时间复杂度: O(n)  空间复杂度: O(w)
func InvertTreeBFS(root *datastructures.TreeNode) *datastructures.TreeNode {
	if root == nil {
		return nil
	}
	queue := []*datastructures.TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return root
}
