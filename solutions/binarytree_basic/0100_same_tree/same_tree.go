package sametree

import "github.com/0226zy/lc-go/pkg/datastructures"

// IsSameTree 相同的树
// 给定两棵二叉树的根节点 p 和 q，检验它们在结构和节点值上是否完全相同。
// 时间复杂度: O(min(n, m))  空间复杂度: O(min(h1, h2)) 递归栈深度
func IsSameTree(p, q *datastructures.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}
