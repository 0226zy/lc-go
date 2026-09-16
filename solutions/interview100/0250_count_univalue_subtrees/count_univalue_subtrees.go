package countunivaluesubtrees

import "github.com/0226zy/lc-go/pkg/datastructures"

// CountUnivalSubtrees 统计同值子树
// 给定二叉树根节点 root，统计所有节点值都相同的子树（同值子树）的数量。
// 时间复杂度: O(n) 每个节点访问一次  空间复杂度: O(h) h 为树高（递归栈）
func CountUnivalSubtrees(root *datastructures.TreeNode) int {
	count := 0
	isUnival(root, &count)
	return count
}

// isUnival 判断以 node 为根的子树是否同值，并在递归过程中累加同值子树数量
func isUnival(node *datastructures.TreeNode, count *int) bool {
	if node == nil {
		return true
	}
	// 后序遍历：先判断左右子树，再判断当前子树
	leftOK := isUnival(node.Left, count)
	rightOK := isUnival(node.Right, count)
	if !leftOK || !rightOK {
		return false
	}
	// 存在的孩子必须与当前节点同值
	if node.Left != nil && node.Left.Val != node.Val {
		return false
	}
	if node.Right != nil && node.Right.Val != node.Val {
		return false
	}
	*count++
	return true
}
