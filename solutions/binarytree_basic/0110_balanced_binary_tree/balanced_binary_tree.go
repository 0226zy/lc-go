package isbalanced

import "github.com/0226zy/lc-go/pkg/datastructures"

// IsBalanced 平衡二叉树
// 给定一棵二叉树，判断它是否是高度平衡的二叉树：
// 即每个节点的左右两个子树的高度差的绝对值不超过 1。
// 时间复杂度: O(n) 每个节点只访问一次  空间复杂度: O(h) 递归栈深度，h 为树高
func IsBalanced(root *datastructures.TreeNode) bool {
	return height(root) != -1
}

// height 自底向上计算以 node 为根的子树高度；
// 一旦发现某个子树不平衡，直接返回 -1 并逐层向上传播，实现提前剪枝。
func height(node *datastructures.TreeNode) int {
	if node == nil {
		return 0
	}
	left := height(node.Left)
	if left == -1 {
		return -1 // 左子树不平衡，整棵树不必再算
	}
	right := height(node.Right)
	if right == -1 {
		return -1 // 右子树不平衡，整棵树不必再算
	}
	if abs(left-right) > 1 {
		return -1 // 当前节点左右子树高度差超过 1
	}
	return max(left, right) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
