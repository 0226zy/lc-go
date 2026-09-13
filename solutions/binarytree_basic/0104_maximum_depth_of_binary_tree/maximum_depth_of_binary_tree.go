package maximumdepthofbinarytree

import "github.com/0226zy/lc-go/pkg/datastructures"

// MaxDepth 二叉树的最大深度
// 给定一个二叉树的根节点 root，返回它的最大深度：即从根节点到最远叶子节点的最长路径上的节点个数。
// 时间复杂度: O(n)  每个节点恰好访问一次
// 空间复杂度: O(h)  递归调用栈的深度，h 为树高；最坏情况（链式树）退化为 O(n)
func MaxDepth(root *datastructures.TreeNode) int {
	if root == nil {
		return 0 // 空子树深度为 0
	}
	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)
	return max(left, right) + 1 // 当前节点深度 = 左右子树较大深度 + 1
}
