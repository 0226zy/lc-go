package maximumaveragesubtree

import "github.com/0226zy/lc-go/pkg/datastructures"

// MaximumAverageSubtree 子树的最大平均值
// 返回二叉树中所有子树（必须包含节点的全部后代）节点平均值的最大值。
// 时间复杂度: O(n)  空间复杂度: O(h)，h 为树高，最坏 O(n)
func MaximumAverageSubtree(root *datastructures.TreeNode) float64 {
	var best float64

	// 后序遍历：返回以 node 为根的子树的 (节点值之和, 节点个数)
	var dfs func(node *datastructures.TreeNode) (int, int)
	dfs = func(node *datastructures.TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		leftSum, leftCnt := dfs(node.Left)
		rightSum, rightCnt := dfs(node.Right)
		sum := leftSum + rightSum + node.Val
		cnt := leftCnt + rightCnt + 1
		// 用当前子树的平均值更新全局最大值
		if avg := float64(sum) / float64(cnt); avg > best {
			best = avg
		}
		return sum, cnt
	}
	dfs(root)
	return best
}
