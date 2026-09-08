package validatebinarysearchtree

import (
	"math"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// IsValidBST 验证二叉搜索树
// 判断给定二叉树是否为合法的二叉搜索树：左子树所有节点 < 根 < 右子树所有节点。
// 使用上下界递归：每个节点必须落在 (lower, upper) 开区间内。
// 时间复杂度: O(n) 每个节点访问一次  空间复杂度: O(h) 递归栈深度为树高
func IsValidBST(root *datastructures.TreeNode) bool {
	var check func(node *datastructures.TreeNode, lower, upper int64) bool
	check = func(node *datastructures.TreeNode, lower, upper int64) bool {
		if node == nil {
			return true
		}
		val := int64(node.Val)
		if val <= lower || val >= upper { // 超出上下界即不合法
			return false
		}
		// 左子树的所有节点必须小于 val，右子树的所有节点必须大于 val
		return check(node.Left, lower, val) && check(node.Right, val, upper)
	}
	return check(root, math.MinInt64, math.MaxInt64)
}
