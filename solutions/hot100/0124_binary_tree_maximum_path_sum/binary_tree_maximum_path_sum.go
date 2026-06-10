package binarytreemaximumpathsum

import (
	"math"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// MaxPathSum 二叉树中的最大路径和（递归 DFS）
// 对于每个节点，计算以该节点为端点的最大贡献值，同时更新全局最大路径和。
// 时间复杂度: O(n)  空间复杂度: O(h)
func MaxPathSum(root *datastructures.TreeNode) int {
	maxSum := math.MinInt32
	dfs(root, &maxSum)
	return maxSum
}

func dfs(node *datastructures.TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}
	// 如果子树贡献为负，则不选该子树（取 0）
	leftGain := max(dfs(node.Left, maxSum), 0)
	rightGain := max(dfs(node.Right, maxSum), 0)

	// 以当前节点为转折点的路径和
	currentPathSum := node.Val + leftGain + rightGain
	if currentPathSum > *maxSum {
		*maxSum = currentPathSum
	}

	// 返回以当前节点为端点的最大路径和（供父节点使用）
	return node.Val + max(leftGain, rightGain)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
