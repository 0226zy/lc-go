package kthsmallest

import "github.com/0226zy/lc-go/pkg/datastructures"

// KthSmallest 二叉搜索树中第 K 小的元素
// 给定二叉搜索树的根节点 root 和整数 k，返回其中第 k 小的元素（从 1 开始计数）。
// 利用 BST 中序遍历结果递增的性质，遍历时计数，数到第 k 个节点立即提前终止。
// 时间复杂度: O(h + k)  空间复杂度: O(h)
func KthSmallest(root *datastructures.TreeNode, k int) int {
	count := 0 // 已访问的节点个数
	ans := 0   // 第 k 小的节点值

	// inorder 中序遍历，返回值表示是否已经找到第 k 小（找到则向上层层返回，实现剪枝）
	var inorder func(node *datastructures.TreeNode) bool
	inorder = func(node *datastructures.TreeNode) bool {
		if node == nil {
			return false
		}
		// 左子树中已找到则直接返回，不再访问当前节点和右子树
		if inorder(node.Left) {
			return true
		}
		count++
		if count == k { // 中序第 k 个被访问的节点就是第 k 小
			ans = node.Val
			return true
		}
		return inorder(node.Right)
	}

	inorder(root)
	return ans
}
