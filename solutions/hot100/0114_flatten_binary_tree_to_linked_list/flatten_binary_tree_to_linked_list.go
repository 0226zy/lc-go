package flattenbinarytreetolinkedlist

import (
	"github.com/0226zy/lc-go/pkg/datastructures"
)

// Flatten 二叉树展开为链表（后序遍历：右→左→根）
// 利用先序遍历的逆序，用 prev 记录上一个处理的节点，将当前节点右指针指向 prev。
// 时间复杂度: O(n)  空间复杂度: O(h)
func Flatten(root *datastructures.TreeNode) {
	var prev *datastructures.TreeNode
	postOrder(root, &prev)
}

func postOrder(node *datastructures.TreeNode, prev **datastructures.TreeNode) {
	if node == nil {
		return
	}
	postOrder(node.Right, prev)
	postOrder(node.Left, prev)
	node.Right = *prev
	node.Left = nil
	*prev = node
}

// FlattenIterative 二叉树展开为链表（迭代法，类似 Morris 遍历）
// 找到左子树的最右节点作为前驱，将右子树挂到前驱上，左子树移到右侧。
// 时间复杂度: O(n)  空间复杂度: O(1)
func FlattenIterative(root *datastructures.TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			// 找到左子树的最右节点（前驱）
			predecessor := curr.Left
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			// 将当前右子树挂到前驱的右指针上
			predecessor.Right = curr.Right
			// 左子树移到右侧
			curr.Right = curr.Left
			curr.Left = nil
		}
		curr = curr.Right
	}
}
