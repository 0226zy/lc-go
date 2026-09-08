package lowestcommonancestorofabinarytree

import "github.com/0226zy/lc-go/pkg/datastructures"

// LowestCommonAncestor 二叉树的最近公共祖先
// 返回以 root 为根的树中 p 和 q 的最近公共祖先。
// 递归语义：若 p、q 不同时出现在子树中，则返回子树中能找到的那个（p 或 q），都找不到返回 nil。
// 时间复杂度: O(n)  空间复杂度: O(h)（h 为树高，即递归栈深度）
func LowestCommonAncestor(root, p, q *datastructures.TreeNode) *datastructures.TreeNode {
	if root == nil || root == p || root == q {
		// 空树返回 nil；root 本身就是 p/q 时直接上报 root
		return root
	}
	left := LowestCommonAncestor(root.Left, p, q)   // 左子树的查找结果
	right := LowestCommonAncestor(root.Right, p, q) // 右子树的查找结果
	if left != nil && right != nil {
		// p、q 分居 root 两侧，root 就是最近公共祖先
		return root
	}
	if left != nil {
		return left // 结果在左子树
	}
	return right // 结果在右子树（或都不在，返回 nil）
}
