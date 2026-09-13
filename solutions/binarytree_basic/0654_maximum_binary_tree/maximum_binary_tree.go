package maximumbinarytree

import "github.com/0226zy/lc-go/pkg/datastructures"

// ConstructMaximumBinaryTree 最大二叉树
// 给定一个不含重复元素的整数数组 nums，构建最大二叉树：
// 根节点为数组中的最大值，左子树由最大值左侧的区间递归构建，右子树由最大值右侧的区间递归构建。
// 时间复杂度: O(n²)  最坏情况（单调数组退化为链式树）每次线性扫描找最大值
// 空间复杂度: O(n)   递归栈深度最坏为 n
func ConstructMaximumBinaryTree(nums []int) *datastructures.TreeNode {
	return build(nums, 0, len(nums)-1)
}

// build 递归构建 nums[lo..hi] 区间对应的最大二叉树
func build(nums []int, lo, hi int) *datastructures.TreeNode {
	if lo > hi {
		return nil // 空区间不产生节点
	}
	// 在区间 [lo, hi] 内线性扫描找最大值的下标
	maxIdx := lo
	for i := lo + 1; i <= hi; i++ {
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
	}
	root := &datastructures.TreeNode{Val: nums[maxIdx]}
	root.Left = build(nums, lo, maxIdx-1)  // 左侧区间递归构建左子树
	root.Right = build(nums, maxIdx+1, hi) // 右侧区间递归构建右子树
	return root
}
