package pathsum

import "github.com/0226zy/lc-go/pkg/datastructures"

// HasPathSum 路径总和
// 给定二叉树的根节点 root 和一个整数 targetSum，
// 判断树中是否存在「根节点到叶节点」的路径，使得这条路径上所有节点值之和等于 targetSum。
// 时间复杂度: O(n) 每个节点访问一次  空间复杂度: O(h) 递归栈深度，h 为树高，最坏 O(n)
func HasPathSum(root *datastructures.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 到达叶节点：检查剩余目标和是否恰好等于叶节点的值
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	// 内部节点：把问题分解为左右子树剩余和为 targetSum - root.Val 的子问题
	remain := targetSum - root.Val
	return HasPathSum(root.Left, remain) || HasPathSum(root.Right, remain)
}
