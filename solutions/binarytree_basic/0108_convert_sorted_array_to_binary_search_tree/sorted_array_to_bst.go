package sortedarraytobst

import "github.com/0226zy/lc-go/pkg/datastructures"

// SortedArrayToBST 将有序数组转换为二叉搜索树
// 给定一个按升序排列的整数数组 nums，将其转换为一棵高度平衡的二叉搜索树。
// 核心思路：分治。每次取当前区间的中间元素作为根节点，左半区间递归构造左子树、
// 右半区间递归构造右子树，左右子树节点数最多相差 1，天然保证高度平衡。
// 时间复杂度: O(n) 每个元素恰好访问一次  空间复杂度: O(log n) 递归栈深度（平衡树）
func SortedArrayToBST(nums []int) *datastructures.TreeNode {
	return buildBST(nums, 0, len(nums)-1)
}

// buildBST 递归地用 nums[low..high] 构造高度平衡二叉搜索树
// 取下中点作为根：当区间长度为偶数时，左子树节点数比右子树多 1
func buildBST(nums []int, low, high int) *datastructures.TreeNode {
	if low > high {
		return nil
	}
	mid := low + (high-low)/2
	return &datastructures.TreeNode{
		Val:   nums[mid],
		Left:  buildBST(nums, low, mid-1),
		Right: buildBST(nums, mid+1, high),
	}
}
