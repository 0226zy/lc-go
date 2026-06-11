package pathsumiii

import "github.com/0226zy/lc-go/pkg/datastructures"

type TreeNode = datastructures.TreeNode

// PathSum 路径总和 III
// 给定一个二叉树的根节点 root 和一个整数 targetSum，
// 返回路径和等于 targetSum 的路径数量。
// 路径不需要从根节点开始，也不需要在叶子节点结束，但必须向下（只能从父节点到子节点）。
// 时间复杂度: O(?)  空间复杂度: O(?)
func PathSum(root *TreeNode, targetSum int) int {
	dict := map[int]int{0: 1}
	var dfs func(root *TreeNode, sum int) int
	dfs = func(root *TreeNode, sum int) int {
		if root == nil {
			return 0
		}
		sum += root.Val
		res := dict[sum-targetSum]
		dict[sum]++
		res += dfs(root.Left, sum)
		res += dfs(root.Right, sum)
		dict[sum]--
		return res
	}
	return dfs(root, 0)
}
