package twosumbsts

import "github.com/0226zy/lc-go/pkg/datastructures"

// TwoSumBSTs 查找两棵二叉搜索树之和
// 给定两棵二叉搜索树的根节点 root1、root2 和整数 target，
// 判断是否存在 root1 中的节点 a 与 root2 中的节点 b 满足 a.Val + b.Val == target。
// 思路：先遍历 root1 把所有节点值存入哈希集合，再遍历 root2 查找 target - Val。
// 时间复杂度: O(n1 + n2)  空间复杂度: O(n1)
func TwoSumBSTs(root1, root2 *datastructures.TreeNode, target int) bool {
	// 第一遍：收集 root1 的所有节点值
	seen := make(map[int]bool)
	var collect func(node *datastructures.TreeNode)
	collect = func(node *datastructures.TreeNode) {
		if node == nil {
			return
		}
		seen[node.Val] = true
		collect(node.Left)
		collect(node.Right)
	}
	collect(root1)

	// 第二遍：在 root2 中查找 target - Val 是否出现过
	var search func(node *datastructures.TreeNode) bool
	search = func(node *datastructures.TreeNode) bool {
		if node == nil {
			return false
		}
		if seen[target-node.Val] {
			return true
		}
		return search(node.Left) || search(node.Right)
	}
	return search(root2)
}
