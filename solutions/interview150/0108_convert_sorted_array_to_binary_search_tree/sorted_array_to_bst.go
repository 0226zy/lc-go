package sortedarraytobst

import "github.com/0226zy/lc-go/pkg/datastructures"

// SortedArrayToBST 将有序数组转换为二叉搜索树
// 给你一个整数数组 nums，其中元素已经按升序排列，请你将其转换为一棵
// 高度平衡的二叉搜索树（左右子树高度差不超过 1）。
// 核心思路：每次取数组中点作为根节点，中点左侧递归构造左子树，右侧递归构造右子树，
// 这样左右子树的节点数最多相差 1，天然保证高度平衡。
// 时间复杂度: O(n) 每个元素访问一次  空间复杂度: O(log n) 递归栈深度（平衡树）
func SortedArrayToBST(nums []int) *datastructures.TreeNode {
	return build(nums, 0, len(nums)-1)
}

// build 递归构造 nums[low..high] 对应的平衡二叉搜索树
// 取上中点作为根：左子树节点数 ≤ 右子树节点数，与 LeetCode 官方示例的树形一致
func build(nums []int, low, high int) *datastructures.TreeNode {
	if low > high {
		return nil
	}
	mid := low + (high-low+1)/2
	return &datastructures.TreeNode{
		Val:   nums[mid],
		Left:  build(nums, low, mid-1),
		Right: build(nums, mid+1, high),
	}
}
