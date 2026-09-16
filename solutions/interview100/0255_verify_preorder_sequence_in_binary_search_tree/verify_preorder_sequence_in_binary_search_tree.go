package verifypreordersequenceinbinarysearchtree

import "math"

// VerifyPreorder 验证二叉搜索树的前序遍历序列
// 给定一个整数数组 preorder，判断它是否是一棵二叉搜索树的前序遍历结果。
// 二叉搜索树中左子树的值都小于根，右子树的值都大于根，且所有值互不相同。
// 时间复杂度: O(n) 每个元素最多入栈、出栈各一次  空间复杂度: O(n) 最坏情况栈深为 n
func VerifyPreorder(preorder []int) bool {
	// stack 维护单调递减的祖先链，low 表示当前允许出现的值的下界
	stack := make([]int, 0, len(preorder))
	low := math.MinInt
	for _, v := range preorder {
		// v 小于下界，说明它落在了某个祖先的左子树禁区中
		if v < low {
			return false
		}
		// v 比栈顶大，说明 v 进入了某个祖先的右子树，
		// 不断弹栈，最后弹出的值即为该祖先，将下界提升到它
		for len(stack) > 0 && v > stack[len(stack)-1] {
			low = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}
	return true
}

// VerifyPreorderRecursive 验证二叉搜索树的前序遍历序列（递归分治解法）
// 以 preorder[start] 为根，向右找到第一个大于根的位置划分左右子树，
// 校验右子树中所有值都大于根之后，递归验证左右子区间。
// 时间复杂度: 平均 O(n log n)，最坏 O(n²)（序列单调时）  空间复杂度: O(n) 递归栈
func VerifyPreorderRecursive(preorder []int) bool {
	var verify func(start, end int) bool
	verify = func(start, end int) bool {
		if start >= end {
			return true
		}
		root := preorder[start]
		// 左子树：所有值都小于根
		i := start + 1
		for i <= end && preorder[i] < root {
			i++
		}
		mid := i
		// 右子树：所有值都必须大于根（值互不相同，不允许等于）
		for ; i <= end; i++ {
			if preorder[i] <= root {
				return false
			}
		}
		return verify(start+1, mid-1) && verify(mid, end)
	}
	return verify(0, len(preorder)-1)
}
