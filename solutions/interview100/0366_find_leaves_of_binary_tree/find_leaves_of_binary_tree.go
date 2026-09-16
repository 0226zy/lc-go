package findleavesofbinarytree

import "github.com/0226zy/lc-go/pkg/datastructures"

// FindLeaves 寻找二叉树的叶子节点
// 每一轮收集当前树中的所有叶子节点并移除，重复直到树为空，
// 返回按收集轮次分组的节点值（同轮内按从左到右的顺序）。
// 核心观察：节点在第几轮被收集 = 它的高度（到最远叶子的距离），
// 因此一次后序 DFS 按高度分桶即可。
// 时间复杂度: O(n) 每个节点访问一次  空间复杂度: O(h) 递归栈深度，h 为树高
func FindLeaves(root *datastructures.TreeNode) [][]int {
	result := [][]int{}
	var height func(node *datastructures.TreeNode) int
	height = func(node *datastructures.TreeNode) int {
		if node == nil {
			return -1
		}
		h := max(height(node.Left), height(node.Right)) + 1
		// 第 h 轮收集的桶还不存在时先补一个
		if len(result) == h {
			result = append(result, []int{})
		}
		result[h] = append(result[h], node.Val)
		return h
	}
	height(root)
	return result
}
