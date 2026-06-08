package diameterofbinarytree

import (
	"github.com/0226zy/lc-go/pkg/datastructures"
)

// DiameterOfBinaryTree 二叉树的直径
// 给你一棵二叉树的根节点，返回该树的直径。
// 直径 = 任意两个节点之间最长路径的边数。
// 时间复杂度: O(n)  空间复杂度: O(h)
func DiameterOfBinaryTree(root *datastructures.TreeNode) int {
	diameter := 0
	var depth func(node *datastructures.TreeNode) int
	depth = func(node *datastructures.TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := depth(node.Left)
		rightDepth := depth(node.Right)
		// 以当前节点为最高点的路径长度（边数）
		if leftDepth+rightDepth > diameter {
			diameter = leftDepth + rightDepth
		}
		if leftDepth > rightDepth {
			return leftDepth + 1
		}
		return rightDepth + 1
	}
	depth(root)
	return diameter
}
