package sumroottoleafnumbers

import "github.com/0226zy/lc-go/pkg/datastructures"

// SumNumbers 求根节点到叶节点数字之和
// 二叉树中每条从根到叶的路径上的节点值可以拼接成一个十进制数字，
// 返回所有这样的数字之和。根节点的值不为 0 时顶层也不含前导零。
// 时间复杂度: O(n) 每个节点访问一次  空间复杂度: O(h) 递归栈深度，h 为树高，最坏 O(n)
func SumNumbers(root *datastructures.TreeNode) int {
	var dfs func(node *datastructures.TreeNode, acc int) int
	dfs = func(node *datastructures.TreeNode, acc int) int {
		if node == nil {
			return 0
		}
		// 走到当前节点，路径数字整体左移一位（乘 10）再加上当前位
		acc = acc*10 + node.Val
		// 叶节点：这条路径的数字已完整，返回它
		if node.Left == nil && node.Right == nil {
			return acc
		}
		// 内部节点：累加左右子树各自形成的数字之和
		return dfs(node.Left, acc) + dfs(node.Right, acc)
	}
	return dfs(root, 0)
}
