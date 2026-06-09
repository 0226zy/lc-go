package convertsortedarraytobinarysearchtree

import (
	"github.com/0226zy/lc-go/pkg/datastructures"
)

// SortedArrayToBST 将有序数组转换为二叉搜索树（递归分治）
// 选择中间元素作为根节点，递归构建左右子树，保证平衡。
// 时间复杂度: O(n)  空间复杂度: O(log n)
func SortedArrayToBST(nums []int) *datastructures.TreeNode {
	return buildBST(nums, 0, len(nums)-1)
}

// buildBST 递归构建平衡 BST
func buildBST(nums []int, left, right int) *datastructures.TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	root := &datastructures.TreeNode{Val: nums[mid]}
	root.Left = buildBST(nums, left, mid-1)
	root.Right = buildBST(nums, mid+1, right)
	return root
}

// SortedArrayToBSTIterative 将有序数组转换为二叉搜索树（迭代）
// 使用队列模拟递归过程。
// 时间复杂度: O(n)  空间复杂度: O(n)
func SortedArrayToBSTIterative(nums []int) *datastructures.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	type item struct {
		parent **datastructures.TreeNode
		left   int
		right  int
	}
	root := &datastructures.TreeNode{Val: 0}
	queue := []item{{parent: &root, left: 0, right: len(nums) - 1}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.left > cur.right {
			*cur.parent = nil
			continue
		}
		mid := cur.left + (cur.right-cur.left)/2
		node := &datastructures.TreeNode{Val: nums[mid]}
		*cur.parent = node
		queue = append(queue,
			item{parent: &node.Left, left: cur.left, right: mid - 1},
			item{parent: &node.Right, left: mid + 1, right: cur.right},
		)
	}
	return root
}
