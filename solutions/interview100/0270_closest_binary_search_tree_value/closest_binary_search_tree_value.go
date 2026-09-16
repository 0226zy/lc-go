package closestbinarysearchtreevalue

import (
	"math"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// ClosestValue 最接近的二叉搜索树值
// 给定一个不为空的二叉搜索树的根节点 root 和一个目标值 target，
// 请在该二叉搜索树中找到最接近目标值 target 的节点值。
// 时间复杂度: O(h) h 为树高  空间复杂度: O(1)
func ClosestValue(root *datastructures.TreeNode, target float64) int {
	// ans 记录当前找到的与 target 最接近的节点值
	ans := root.Val
	node := root
	for node != nil {
		// 沿途更新最优解：差值更小者胜出
		if math.Abs(float64(node.Val)-target) < math.Abs(float64(ans)-target) {
			ans = node.Val
		}
		// 利用 BST 性质决定向左还是向右：另一侧只会离 target 更远
		if target < float64(node.Val) {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return ans
}
