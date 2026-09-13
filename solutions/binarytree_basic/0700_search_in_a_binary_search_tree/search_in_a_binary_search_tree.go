package searchbst

import (
	"github.com/0226zy/lc-go/pkg/datastructures"
)

// SearchBST 二叉搜索树中的搜索
// 给定二叉搜索树的根节点 root 和一个值 val，
// 返回树中值等于 val 的节点（以其为根的子树），不存在则返回 nil。
// 时间复杂度: O(h) h 为树高，平衡时 O(log n)  空间复杂度: O(h) 递归栈深度
func SearchBST(root *datastructures.TreeNode, val int) *datastructures.TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val < root.Val {
		return SearchBST(root.Left, val)
	}
	return SearchBST(root.Right, val)
}
