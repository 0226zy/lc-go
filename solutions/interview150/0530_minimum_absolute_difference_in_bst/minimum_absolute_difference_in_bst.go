package minimumabsolutedifferenceinbst

import (
	"math"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// GetMinimumDifference 二叉搜索树的最小绝对差
// 返回 BST 中任意两节点值之差的最小绝对值。
// 利用 BST 中序遍历递增的性质，只需比较中序序列中相邻元素的差。
// 时间复杂度: O(n) 每个节点访问一次  空间复杂度: O(h) 递归栈深度为树高
func GetMinimumDifference(root *datastructures.TreeNode) int {
	minDiff := math.MaxInt32
	prev := math.MinInt32 // 上一个中序访问的节点值；初始为极小哨兵
	var inorder func(node *datastructures.TreeNode)
	inorder = func(node *datastructures.TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prev != math.MinInt32 { // 跳过第一个节点（没有前驱）
			if d := node.Val - prev; d < minDiff {
				minDiff = d
			}
		}
		prev = node.Val
		inorder(node.Right)
	}
	inorder(root)
	return minDiff
}
