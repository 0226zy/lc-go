package countnodes

import "github.com/0226zy/lc-go/pkg/datastructures"

// CountNodes 完全二叉树的节点个数
// 给你一棵完全二叉树的根节点 root，求该树的节点总数。
// 利用完全二叉树的性质：比较根的最左路径与最右路径长度，若相等则为满二叉树，用公式 2^h - 1 直接计算。
// 时间复杂度: O(log^2 n)  空间复杂度: O(log n)
func CountNodes(root *datastructures.TreeNode) int {
	if root == nil {
		return 0
	}
	// 最左路径深度：一直沿左孩子走到底
	leftDepth := 0
	for cur := root; cur != nil; cur = cur.Left {
		leftDepth++
	}
	// 最右路径深度：一直沿右孩子走到底
	rightDepth := 0
	for cur := root; cur != nil; cur = cur.Right {
		rightDepth++
	}
	// 两路径等长说明该子树为满二叉树，节点数可直接用公式得出
	if leftDepth == rightDepth {
		return (1 << leftDepth) - 1
	}
	// 否则左右子树各自仍是完全二叉树，递归求和
	return CountNodes(root.Left) + CountNodes(root.Right) + 1
}
