package flattenbinarytreetolinkedlist

import "github.com/0226zy/lc-go/pkg/datastructures"

// Flatten 二叉树展开为链表
// 给定二叉树的根节点 root，原地将它展开为一个单链表：顺序与先序遍历一致，
// 链表节点使用 Right 指针相连，所有 Left 指针置为 nil。
// 时间复杂度: O(n) 每个节点恰好访问一次  空间复杂度: O(h) 递归栈深度，h 为树高
func Flatten(root *datastructures.TreeNode) {
	// prev 指向「已经处理好的链表的头节点」，即当前节点在先序序列中的后继
	var prev *datastructures.TreeNode

	// 按照「右 → 左 → 根」的顺序遍历（先序遍历的逆序），
	// 这样处理到某个节点时，它的先序后继 prev 一定是已经整理好的链表
	var dfs func(node *datastructures.TreeNode)
	dfs = func(node *datastructures.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		dfs(node.Left)
		// 当前节点的 Right 指向先序后继，Left 置空
		node.Right = prev
		node.Left = nil
		// 当前节点成为新的链表头
		prev = node
	}
	dfs(root)
}
