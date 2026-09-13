package pathsumii

import "github.com/0226zy/lc-go/pkg/datastructures"

// PathSum 路径总和 II
// 给定二叉树的根节点 root 和一个整数目标和 targetSum，
// 找出所有「从根节点到叶子节点」且路径上节点值之和等于 targetSum 的路径。
// 时间复杂度: O(n·h) 每个节点访问一次，收集答案时还需拷贝长度为 h 的路径  空间复杂度: O(h) 递归栈深度与当前路径长度，h 为树高，最坏 O(n)
func PathSum(root *datastructures.TreeNode, targetSum int) [][]int {
	var result [][]int
	var path []int

	var dfs func(node *datastructures.TreeNode, remain int)
	dfs = func(node *datastructures.TreeNode, remain int) {
		if node == nil {
			return
		}
		// 做选择：把当前节点加入路径，并更新剩余目标和
		path = append(path, node.Val)
		remain -= node.Val
		// 到达叶节点且剩余和恰好为 0：收集一条合法路径（必须拷贝，否则后续回溯会覆盖）
		if node.Left == nil && node.Right == nil && remain == 0 {
			p := make([]int, len(path))
			copy(p, path)
			result = append(result, p)
		} else {
			dfs(node.Left, remain)
			dfs(node.Right, remain)
		}
		// 撤销选择：弹出当前节点，恢复现场，继续探索其他分支
		path = path[:len(path)-1]
	}
	dfs(root, targetSum)
	return result
}
