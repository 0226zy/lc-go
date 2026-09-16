package binarytreelongestconsecutivesequence

import "github.com/0226zy/lc-go/pkg/datastructures"

// LongestConsecutive 二叉树最长连续序列
// 返回二叉树中最长的连续序列路径长度：路径自上而下，每个节点的值比父节点恰好大 1。
// 思路：自顶向下 DFS，携带父节点值与当前连续长度，连续则 +1，断开则重置为 1。
// 时间复杂度: O(n)  空间复杂度: O(h)，h 为树高（递归栈深度）
func LongestConsecutive(root *datastructures.TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	// 父值初始取 root.Val，使根节点必然「断开」从长度 1 开始计数
	dfs(root, root.Val, 0, &ans)
	return ans
}

// dfs 遍历 node，parentVal 为父节点的值，length 为从上方延续下来的连续序列长度
func dfs(node *datastructures.TreeNode, parentVal, length int, ans *int) {
	if node == nil {
		return
	}
	if node.Val == parentVal+1 {
		length++ // 连续：长度加一
	} else {
		length = 1 // 断开：从当前节点重新开始
	}
	if length > *ans {
		*ans = length
	}
	dfs(node.Left, node.Val, length, ans)
	dfs(node.Right, node.Val, length, ans)
}
