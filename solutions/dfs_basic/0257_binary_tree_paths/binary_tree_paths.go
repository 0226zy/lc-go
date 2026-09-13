package binarytreepaths

import (
	"bytes"
	"strconv"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// BinaryTreePaths 二叉树的所有路径
// 给定二叉树的根节点 root，返回所有从根节点到叶子节点的路径，
// 每条路径用 "->" 连接节点值，例如 "1->2->5"。
// 时间复杂度: O(n*h) n 为节点数，h 为树高，每条路径拼接开销 O(h)
// 空间复杂度: O(h) 递归栈与路径缓冲区，最坏 O(n)（不计输出）
func BinaryTreePaths(root *datastructures.TreeNode) []string {
	var result []string
	var path bytes.Buffer

	var dfs func(node *datastructures.TreeNode)
	dfs = func(node *datastructures.TreeNode) {
		if node == nil {
			return
		}
		// 记录进入前的缓冲区长度，退出时还原，实现路径的撤销
		mark := path.Len()
		if mark > 0 {
			path.WriteString("->")
		}
		path.WriteString(strconv.Itoa(node.Val))

		if node.Left == nil && node.Right == nil {
			// 到达叶节点：收集当前完整路径
			result = append(result, path.String())
		} else {
			dfs(node.Left)
			dfs(node.Right)
		}
		// 回溯：把缓冲区截断回进入时的长度，撤销本节点的写入
		path.Truncate(mark)
	}
	dfs(root)
	return result
}
