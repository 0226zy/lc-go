package invertbinarytree

import "github.com/0226zy/lc-go/pkg/datastructures"

// InvertTree 翻转二叉树
// 给定二叉树根节点 root，交换每个节点的左右子树（镜像翻转），返回翻转后的根节点。
// 时间复杂度: O(n) 每个节点访问一次  空间复杂度: O(h) 递归栈深度，h 为树高
func InvertTree(root *datastructures.TreeNode) *datastructures.TreeNode {
	if root == nil {
		return nil
	}
	root.Left = InvertTree(root.Left)
	root.Right = InvertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
