package validatebinarysearchtree

import (
	"math"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// IsValidBST 验证二叉搜索树（递归 DFS，带上下界）
// 每个节点必须满足 lower < node.Val < upper。
// 时间复杂度: O(n)  空间复杂度: O(h)
func IsValidBST(root *datastructures.TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

// validate 递归验证节点值是否在 (lower, upper) 范围内
func validate(node *datastructures.TreeNode, lower, upper int64) bool {
	if node == nil {
		return true
	}
	val := int64(node.Val)
	if val <= lower || val >= upper {
		return false
	}
	return validate(node.Left, lower, val) && validate(node.Right, val, upper)
}

// IsValidBSTInorder 验证二叉搜索树（中序遍历）
// BST 的中序遍历结果必须严格递增。
// 时间复杂度: O(n)  空间复杂度: O(h)
func IsValidBSTInorder(root *datastructures.TreeNode) bool {
	var prev *int // 记录中序遍历的前一个值
	stack := []*datastructures.TreeNode{}
	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil && curr.Val <= *prev {
			return false
		}
		prev = &curr.Val
		curr = curr.Right
	}
	return true
}
