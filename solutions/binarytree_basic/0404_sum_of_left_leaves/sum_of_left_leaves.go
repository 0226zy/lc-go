package sumofleftleaves

import "github.com/0226zy/lc-go/pkg/datastructures"

// TreeNode 别名，便于阅读
type TreeNode = datastructures.TreeNode

// SumOfLeftLeaves 左叶子之和
// 给定二叉树的根节点 root，返回所有左叶子之和。
// 左叶子指“是父节点的左孩子、且自身没有子节点”的节点。
// 时间复杂度: O(n)  空间复杂度: O(h) 递归栈深度，h 为树高
func SumOfLeftLeaves(root *TreeNode) int {
	var dfs func(node *TreeNode, isLeft bool) int
	dfs = func(node *TreeNode, isLeft bool) int {
		if node == nil {
			return 0
		}
		// 叶子节点：只有同时是左孩子才计入答案
		if node.Left == nil && node.Right == nil {
			if isLeft {
				return node.Val
			}
			return 0
		}
		// 非叶子节点：累加左右子树中的左叶子之和
		return dfs(node.Left, true) + dfs(node.Right, false)
	}
	return dfs(root, false) // 根节点没有父节点，不是左叶子
}
