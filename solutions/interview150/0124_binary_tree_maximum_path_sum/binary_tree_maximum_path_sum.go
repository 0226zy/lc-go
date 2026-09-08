package binarytreemaximumpathsum

import (
	"math"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// MaxPathSum 二叉树中的最大路径和
// 路径定义为从树中任意节点出发、沿父子连接到达任意节点的序列（同一节点至多出现一次），
// 返回所有非空路径中节点值之和的最大值。路径不必经过根节点。
// 时间复杂度: O(n) 每个节点访问一次  空间复杂度: O(h) 递归栈深度，h 为树高，最坏 O(n)
func MaxPathSum(root *datastructures.TreeNode) int {
	maxSum := math.MinInt32

	// maxGain 返回「以 node 为一端、向下延伸到某个后代节点」的路径最大和（即 node 的贡献值）。
	// 只有贡献值为正时对父节点才有用，否则返回 0（相当于舍弃该子树）。
	var maxGain func(node *datastructures.TreeNode) int
	maxGain = func(node *datastructures.TreeNode) int {
		if node == nil {
			return 0
		}
		// 递归计算左右子树的贡献值，负贡献直接舍弃
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)
		// 经过 node 的「弧形」路径和：左链 + node + 右链
		// 这是一条完整路径的候选答案，但不能再向上延伸，因此只用于更新全局最大值
		pathSum := node.Val + leftGain + rightGain
		if pathSum > maxSum {
			maxSum = pathSum
		}
		// 返回 node 的贡献值：只能选左、右中较大的一边继续向上延伸
		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root)
	return maxSum
}
