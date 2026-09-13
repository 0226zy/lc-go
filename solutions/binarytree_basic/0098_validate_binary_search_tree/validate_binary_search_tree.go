package isvalidbst

import (
	"math"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// IsValidBST 验证二叉搜索树
// 判断给定二叉树是否为合法的二叉搜索树：左子树所有节点值严格小于根，右子树所有节点值严格大于根，且左右子树也必须分别是 BST。
// 采用携带合法上下界的递归：每个节点必须落在 (lower, upper) 开区间内，递归时将当前节点值收紧为子树的新边界。
// 时间复杂度: O(n) 每个节点恰好访问一次  空间复杂度: O(h) 递归栈深度，h 为树高
func IsValidBST(root *datastructures.TreeNode) bool {
	// 使用 int64 边界，避免节点值取 int32 极值时与哨兵冲突
	return checkBounds(root, math.MinInt64, math.MaxInt64)
}

// checkBounds 递归检查 node 是否落在 (lower, upper) 开区间内，
// 并向下递归：左子树继承 (lower, 当前值)，右子树继承 (当前值, upper)
func checkBounds(node *datastructures.TreeNode, lower, upper int64) bool {
	if node == nil {
		return true
	}
	val := int64(node.Val)
	if val <= lower || val >= upper { // 超出祖先限定的区间，直接判定非法
		return false
	}
	return checkBounds(node.Left, lower, val) && checkBounds(node.Right, val, upper)
}
