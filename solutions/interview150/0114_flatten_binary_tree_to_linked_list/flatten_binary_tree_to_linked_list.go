package flattenbinarytreetolinkedlist

import "github.com/0226zy/lc-go/pkg/datastructures"

// Flatten 二叉树展开为链表
// 给定二叉树的根节点 root，原地将它展开为一个单链表：
// 顺序为「先序遍历的顺序」，链表的节点使用二叉树的 Right 指针相连，Left 指针全部置空。
// 时间复杂度: O(n) 每个节点最多被访问常数次  空间复杂度: O(1) 迭代实现，常数额外空间
func Flatten(root *datastructures.TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left == nil {
			// 没有左子树，直接走到右子树
			curr = curr.Right
			continue
		}
		// 找到左子树的最右节点 pre
		// pre 是先序序列中 curr 之后的下一个节点（curr 的左子树最后一个节点）
		pre := curr.Left
		for pre.Right != nil {
			pre = pre.Right
		}
		// 把 curr 的右子树整体挂到 pre 的右边
		pre.Right = curr.Right
		// 左子树移到右边，并置空左指针
		curr.Right = curr.Left
		curr.Left = nil
		// 继续处理下一个节点（即原来的左子树根）
		curr = curr.Right
	}
}
