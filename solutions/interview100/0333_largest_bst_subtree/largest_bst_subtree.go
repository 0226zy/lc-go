package largestbstsubtree

import (
	"math"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// LargestBSTSubtree 最大二叉搜索子树
// 给定二叉树根节点，返回其中节点数最多的二叉搜索子树的节点个数。
// 后序遍历：每个节点向上返回（是否BST, 最小值, 最大值, 节点数），
// 当左右子树均为 BST 且 左子树最大值 < 当前值 < 右子树最小值 时当前子树为 BST。
// 时间复杂度: O(n) 每个节点访问一次  空间复杂度: O(h) 递归栈深度
func LargestBSTSubtree(root *datastructures.TreeNode) int {
	ans := 0

	// dfs 返回以 node 为根的子树信息：isBST 是否为二叉搜索树，
	// minVal/maxVal 为子树中的最小/最大值，size 为节点数
	var dfs func(node *datastructures.TreeNode) (isBST bool, minVal, maxVal, size int)
	dfs = func(node *datastructures.TreeNode) (bool, int, int, int) {
		if node == nil {
			// 空树视为 BST：最小值取 +∞、最大值取 -∞，使父节点比较天然成立
			return true, math.MaxInt32, math.MinInt32, 0
		}

		lBST, lMin, lMax, lSize := dfs(node.Left)
		rBST, rMin, rMax, rSize := dfs(node.Right)

		// 左右子树都是 BST，且满足 左子树最大值 < 当前值 < 右子树最小值
		if lBST && rBST && lMax < node.Val && node.Val < rMin {
			size := lSize + rSize + 1
			if size > ans {
				ans = size
			}
			return true, min(lMin, node.Val), max(rMax, node.Val), size
		}
		// 当前子树不是 BST，向上传递失败标记，答案已在子树递归中更新
		return false, 0, 0, 0
	}

	dfs(root)
	return ans
}
