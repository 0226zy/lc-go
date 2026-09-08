package binarytreezigzaglevelordertraversal

import "github.com/0226zy/lc-go/pkg/datastructures"

// ZigzagLevelOrder 二叉树的锯齿形层序遍历
// 返回二叉树层序遍历结果，奇数层从左到右，偶数层从右到左，依次交替。
// 时间复杂度: O(n) 每个节点恰好入队出队一次  空间复杂度: O(n) 队列最多同时存放一层的节点
func ZigzagLevelOrder(root *datastructures.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	queue := []*datastructures.TreeNode{root}
	leftToRight := true // 当前层是否从左到右输出
	for len(queue) > 0 {
		size := len(queue) // 当前层的节点个数
		level := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			// 根据方向决定写入位置：从左到右正序写，从右到左倒序写
			if leftToRight {
				level[i] = node.Val
			} else {
				level[size-1-i] = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
		leftToRight = !leftToRight // 下一层反转方向
	}
	return result
}
