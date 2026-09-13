package minimumabsolutedifferenceinbst

import (
	"math"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// GetMinimumDifference 二叉搜索树的最小绝对差
// 给定二叉搜索树的根节点，返回树中任意两个不同节点值之差的最小绝对值。
// 时间复杂度: O(n)  空间复杂度: O(h)，h 为树高（递归栈深度）
func GetMinimumDifference(root *datastructures.TreeNode) int {
	minDiff := math.MaxInt32
	// prev 指向上一个中序访问节点的值；用指针的 nil 区分"尚未访问任何节点"，
	// 避免哨兵值与节点真实取值冲突
	var prev *int

	var inorder func(node *datastructures.TreeNode)
	inorder = func(node *datastructures.TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prev != nil {
			if diff := node.Val - *prev; diff < minDiff {
				minDiff = diff
			}
		}
		val := node.Val
		prev = &val
		inorder(node.Right)
	}
	inorder(root)

	return minDiff
}
