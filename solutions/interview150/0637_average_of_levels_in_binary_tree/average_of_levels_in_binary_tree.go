package averageoflevelsinbinarytree

import "github.com/0226zy/lc-go/pkg/datastructures"

// AverageOfLevels 二叉树的层平均值
// 返回二叉树每一层所有节点值的平均值。
// 时间复杂度: O(n) 每个节点恰好入队出队一次  空间复杂度: O(n) 队列最多同时存放一层的节点
func AverageOfLevels(root *datastructures.TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var result []float64
	queue := []*datastructures.TreeNode{root}
	for len(queue) > 0 {
		size := len(queue) // 当前层的节点个数
		sum := 0.0
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += float64(node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, sum/float64(size))
	}
	return result
}
