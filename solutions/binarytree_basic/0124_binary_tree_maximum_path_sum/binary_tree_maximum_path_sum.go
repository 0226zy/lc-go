package binarytreemaximumpathsum

import (
	"math"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// MaxPathSum 二叉树中的最大路径和
// 给定二叉树的根节点 root，返回其最大路径和。路径从任意节点出发沿父子边到达任意节点，
// 同一节点至多出现一次，至少包含一个节点，且不一定经过根节点。
// 时间复杂度: O(n)  n 为节点数，每个节点恰好访问一次
// 空间复杂度: O(h)  h 为树高，递归栈深度，最坏退化为链式树时 O(n)
func MaxPathSum(root *datastructures.TreeNode) int {
	// 空树没有路径，约定返回 0
	if root == nil {
		return 0
	}

	// 节点值可能为负，全局最大值必须初始化为最小整数而非 0
	maxSum := math.MinInt32

	// gain 返回以 node 为上端点、向下延伸到某个后代节点的单边路径最大和（node 的单边贡献），
	// 同时在递归过程中用「经过 node 的完整路径」更新全局最大值 maxSum
	var gain func(node *datastructures.TreeNode) int
	gain = func(node *datastructures.TreeNode) int {
		if node == nil {
			return 0
		}
		// 后序递归：先拿到左右子树的单边贡献，贡献为负时直接舍弃（按 0 处理）
		left := max(gain(node.Left), 0)
		right := max(gain(node.Right), 0)
		// 以 node 为最高点的完整路径 = 左链 + node + 右链，
		// 该路径在 node 处拐弯，无法再接给父节点，因此只用来更新答案
		if sum := node.Val + left + right; sum > maxSum {
			maxSum = sum
		}
		// 返回给父节点的单边贡献只能在左右中选较大的一条链继续延伸
		return node.Val + max(left, right)
	}

	gain(root)
	return maxSum
}
