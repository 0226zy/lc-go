package kthsmallestelementinabst

import "github.com/0226zy/lc-go/pkg/datastructures"

// KthSmallest 二叉搜索树中第 K 小的元素
// 返回 BST 中按升序排列后的第 k 个节点值。
// 利用 BST 中序遍历递增的性质，遍历时计数，数到 k 即可提前返回。
// 时间复杂度: O(h + k) h 为树高，只需走到第 k 个节点  空间复杂度: O(h) 递归栈深度
func KthSmallest(root *datastructures.TreeNode, k int) int {
	count := 0
	result := -1
	var inorder func(node *datastructures.TreeNode) bool
	inorder = func(node *datastructures.TreeNode) bool {
		if node == nil {
			return false
		}
		// 左子树中找到则直接返回
		if inorder(node.Left) {
			return true
		}
		count++
		if count == k { // 中序第 k 个节点即第 k 小
			result = node.Val
			return true
		}
		return inorder(node.Right)
	}
	inorder(root)
	return result
}
