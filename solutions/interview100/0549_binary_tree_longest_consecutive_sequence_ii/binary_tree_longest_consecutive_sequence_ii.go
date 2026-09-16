package binarytreelongestconsecutivesequenceii

import "github.com/0226zy/lc-go/pkg/datastructures"

// LongestConsecutive 二叉树最长连续序列 II
// 找出二叉树中最长的连续序列路径长度（按节点数计），路径上相邻节点值恰好相差 1，
// 且路径可以是“子节点 -> 父节点 -> 子节点”的形式（先向上再向下）。
// 时间复杂度: O(n) 每个节点访问一次  空间复杂度: O(h) 递归栈深度，h 为树高
func LongestConsecutive(root *datastructures.TreeNode) int {
	ans := 0
	// dfs 返回 (inc, dec)：以 node 为端点向下延伸的最长严格递增 / 严格递减连续序列长度（含 node）
	var dfs func(node *datastructures.TreeNode) (int, int)
	dfs = func(node *datastructures.TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		inc, dec := 1, 1
		if node.Left != nil {
			lInc, lDec := dfs(node.Left)
			if node.Left.Val == node.Val+1 {
				inc = max(inc, lInc+1)
			}
			if node.Left.Val == node.Val-1 {
				dec = max(dec, lDec+1)
			}
		}
		if node.Right != nil {
			rInc, rDec := dfs(node.Right)
			if node.Right.Val == node.Val+1 {
				inc = max(inc, rInc+1)
			}
			if node.Right.Val == node.Val-1 {
				dec = max(dec, rDec+1)
			}
		}
		// 以 node 为拐点，一臂递增一臂递减，拼成经过 node 的最长连续序列
		ans = max(ans, inc+dec-1)
		return inc, dec
	}
	dfs(root)
	return ans
}
