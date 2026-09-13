package lowestcommonancestorofabinarytree

import "github.com/0226zy/lc-go/pkg/datastructures"

// LowestCommonAncestor 二叉树的最近公共祖先
// 给定一棵二叉树和两个节点 p、q，返回 p 与 q 的最近公共祖先（允许一个节点是另一个节点的祖先）。
// 采用后序遍历 DFS：先在左右子树中查找，再在当前节点汇合结果。
// 时间复杂度: O(n)  空间复杂度: O(h)（h 为树高，即递归栈深度）
func LowestCommonAncestor(root, p, q *datastructures.TreeNode) *datastructures.TreeNode {
	// 空节点说明该路径上没有目标；root 命中 p 或 q 时直接上报 root，
	// 因为此时最近公共祖先只可能是 root 或 root 的祖先。
	if root == nil || root == p || root == q {
		return root
	}

	left := LowestCommonAncestor(root.Left, p, q)   // 左子树中找到的节点（p、q 或它们的公共祖先）
	right := LowestCommonAncestor(root.Right, p, q) // 右子树中找到的节点

	if left != nil && right != nil {
		// p、q 分别落在 root 的左右两侧，root 即最近公共祖先
		return root
	}
	if left != nil {
		// 只在左子树找到，向上传递左子树的结果
		return left
	}
	// 只在右子树找到，或两侧都没找到（返回 nil）
	return right
}
