package countcompletetreenodes

import "github.com/0226zy/lc-go/pkg/datastructures"

// CountNodes 完全二叉树的节点个数
// 给定一棵完全二叉树，返回其节点总数。
// 利用完全二叉树性质：若左右最外路径等长则为满二叉树，可直接用公式 2^h - 1。
// 时间复杂度: O(log^2 n)  空间复杂度: O(log n)（递归栈深度等于树高）
func CountNodes(root *datastructures.TreeNode) int {
	if root == nil {
		return 0
	}
	// 沿最左路径求左子树高度
	lh := 0
	for node := root; node != nil; node = node.Left {
		lh++
	}
	// 沿最右路径求右子树高度
	rh := 0
	for node := root; node != nil; node = node.Right {
		rh++
	}
	if lh == rh {
		// 满二叉树，直接套公式 2^h - 1
		return (1 << lh) - 1
	}
	return CountNodes(root.Left) + CountNodes(root.Right) + 1
}
