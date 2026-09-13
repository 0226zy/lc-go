package mindepth

import "github.com/0226zy/lc-go/pkg/datastructures"

// MinDepth 二叉树的最小深度
// 给定一个二叉树，找出其最小深度，即根节点到最近叶子节点的最短路径上的节点数。
// 注意：叶子节点是没有左右孩子的节点，单侧子树为空时不能把它当作深度 0。
// 时间复杂度: O(n)  n 为节点数，每个节点最多访问一次  空间复杂度: O(h) 递归栈深度，最坏 O(n)
func MinDepth(root *datastructures.TreeNode) int {
	if root == nil {
		return 0
	}
	// 左子树为空：最小深度只能到右子树中寻找，不能误取 0
	if root.Left == nil {
		return MinDepth(root.Right) + 1
	}
	// 右子树为空：同理只能走左子树
	if root.Right == nil {
		return MinDepth(root.Left) + 1
	}
	// 两侧子树都非空：取两者较小值
	return min(MinDepth(root.Left), MinDepth(root.Right)) + 1
}
