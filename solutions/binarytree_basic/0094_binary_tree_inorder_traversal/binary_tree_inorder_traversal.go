package inordertraversal

import "github.com/0226zy/lc-go/pkg/datastructures"

type TreeNode = datastructures.TreeNode

// InorderTraversal 二叉树的中序遍历
// 给定二叉树的根节点 root，按「左子树 -> 根节点 -> 右子树」的顺序返回所有节点值。
// 时间复杂度: O(n)  空间复杂度: O(h)，h 为树高（递归调用栈）
func InorderTraversal(root *TreeNode) []int {
	var result []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)               // 先递归左子树
		result = append(result, node.Val) // 再访问根节点
		dfs(node.Right)              // 最后递归右子树
	}
	dfs(root)
	return result
}

// InorderTraversalIterative 迭代法实现二叉树中序遍历（显式栈）
// 用栈模拟递归过程：一路向左入栈，弹栈访问后再转向右子树。
// 时间复杂度: O(n)  空间复杂度: O(h)，h 为树高
func InorderTraversalIterative(root *TreeNode) []int {
	var result []int
	stack := []*TreeNode{}
	cur := root

	for cur != nil || len(stack) > 0 {
		for cur != nil { // 沿左链一路入栈，直到最左节点
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1] // 弹出栈顶并访问
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		cur = cur.Right // 转向右子树
	}

	return result
}
