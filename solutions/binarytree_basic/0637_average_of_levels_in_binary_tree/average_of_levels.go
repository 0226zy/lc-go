package averageoflevels

import "github.com/0226zy/lc-go/pkg/datastructures"

// AverageOfLevels 二叉树的层平均值
// 给定一个二叉树的根节点，返回每一层节点值的平均值组成的数组。
// 时间复杂度: O(n)  空间复杂度: O(n)
func AverageOfLevels(root *datastructures.TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var answer []float64
	queue := []*datastructures.TreeNode{root}
	for len(queue) > 0 {
		n := len(queue) // 当前层节点数，出队前先固定下来
		levelSum := 0.0
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			levelSum += float64(node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		answer = append(answer, levelSum/float64(n))
	}
	return answer
}
