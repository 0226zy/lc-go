package symmetrictree

import "github.com/0226zy/lc-go/pkg/datastructures"

// IsSymmetric 对称二叉树
// 给定二叉树根节点 root，检查树是否轴对称（左右子树互为镜像）。
// 时间复杂度: O(n) 每个节点访问一次  空间复杂度: O(h) 递归栈深度，h 为树高
func IsSymmetric(root *datastructures.TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

// check 判断以 a、b 为根的两棵子树是否互为镜像（a 的左对应 b 的右）
func check(a, b *datastructures.TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return check(a.Left, b.Right) && check(a.Right, b.Left)
}
