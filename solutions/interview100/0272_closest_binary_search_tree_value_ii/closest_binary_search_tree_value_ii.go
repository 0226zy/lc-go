package closestbinarysearchtreevalueii

import (
	"math"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// ClosestKValues 最接近的二叉搜索树值 II
// 给定二叉搜索树的根节点 root、目标值 target 和整数 k，
// 返回树中最接近 target 的 k 个节点值（顺序不限）。
// 思路：中序遍历得到升序序列，答案必为其中长度 k 的连续窗口，
// 遍历过程中维护滑动窗口并在越过最优位置时提前终止。
// 时间复杂度: O(n)  空间复杂度: O(h + k)，h 为树高
func ClosestKValues(root *datastructures.TreeNode, target float64, k int) []int {
	// queue 维护中序序列上长度不超过 k 的滑动窗口
	queue := make([]int, 0, k)
	// stack 用于迭代中序遍历
	var stack []*datastructures.TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		// 一路向左，把沿途节点压栈
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		val := node.Val
		if len(queue) < k {
			// 窗口未满，直接入队
			queue = append(queue, val)
		} else {
			// 中序值单调递增：若当前值不比队首更接近 target，
			// 说明窗口已越过最优位置，继续遍历只会更远，提前终止
			if math.Abs(float64(val)-target) >= math.Abs(float64(queue[0])-target) {
				break
			}
			// 窗口右移一格
			queue = append(queue[1:], val)
		}
		node = node.Right
	}
	return queue
}
