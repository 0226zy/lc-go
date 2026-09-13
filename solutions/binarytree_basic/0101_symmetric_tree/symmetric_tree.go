package symmetrictree

import "github.com/0226zy/lc-go/pkg/datastructures"

// IsSymmetric 对称二叉树
// 给定一棵二叉树的根节点 root，判断它是否关于根节点左右对称（左子树与右子树互为镜像）。
// 时间复杂度: O(n) 每个节点最多比较一次  空间复杂度: O(h) 递归栈深度为树高 h
func IsSymmetric(root *datastructures.TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

// isMirror 递归判断以 a、b 为根的两棵子树是否互为镜像：
// 要求 a、b 同时为空或同时非空且值相等，并且 a 的左子树对应 b 的右子树、a 的右子树对应 b 的左子树。
func isMirror(a, b *datastructures.TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return isMirror(a.Left, b.Right) && isMirror(a.Right, b.Left)
}
