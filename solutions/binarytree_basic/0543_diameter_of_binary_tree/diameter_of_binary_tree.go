package diameterofbinarytree

import "github.com/0226zy/lc-go/pkg/datastructures"

// DiameterOfBinaryTree 二叉树的直径
// 给定一棵二叉树，计算它的直径长度：任意两个节点之间最长路径的边数，
// 这条路径可能穿过也可能不穿过根节点。
// 时间复杂度: O(n)  n 为节点数，每个节点恰好访问一次
// 空间复杂度: O(h)  h 为树高，递归栈深度，最坏退化为链式树时 O(n)
func DiameterOfBinaryTree(root *datastructures.TreeNode) int {
	diameter := 0

	// height 返回以 node 为根的子树高度（边数），同时更新全局最大直径
	var height func(node *datastructures.TreeNode) int
	height = func(node *datastructures.TreeNode) int {
		if node == nil {
			return -1 // 空节点高度定义为 -1，使叶子节点高度为 0（叶子没有向下的边）
		}
		leftH := height(node.Left)
		rightH := height(node.Right)
		// 经过当前节点的最长路径 = 左子树高度 + 右子树高度 + 2 条边
		if leftH+rightH+2 > diameter {
			diameter = leftH + rightH + 2
		}
		// 向上返回当前子树的高度，供父节点继续累加
		return max(leftH, rightH) + 1
	}
	height(root)
	return diameter
}
