package zigzaglevelorder

import "github.com/0226zy/lc-go/pkg/datastructures"

// ZigzagLevelOrder 二叉树的锯齿形层序遍历
// 逐层返回节点值，第 1 层从左到右，第 2 层从右到左，依次交替。
// 时间复杂度: O(n) 每个节点恰好入队、出队各一次  空间复杂度: O(n) 队列最多同时存放一层的节点
func ZigzagLevelOrder(root *datastructures.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	queue := []*datastructures.TreeNode{root}
	leftToRight := true // 当前层的输出方向：true 从左到右，false 从右到左
	for len(queue) > 0 {
		size := len(queue) // 当前层节点个数
		level := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			// 方向只影响写入下标：正向正序填，反向从尾部倒序填
			if leftToRight {
				level[i] = node.Val
			} else {
				level[size-1-i] = node.Val
			}
			// 孩子入队顺序固定为「先左后右」，与输出方向无关
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
		leftToRight = !leftToRight // 下一层方向取反
	}
	return result
}
