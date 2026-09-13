package serializebinarytree

import (
	"strconv"
	"strings"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// 序列化时的空节点标记
const nilMark = "N"

// Codec 二叉树序列化与反序列化器
// 采用「前序遍历 + nil 标记」格式：根、左子树、右子树依次输出，空节点记为 "N"
type Codec struct{}

// Constructor 构造一个 Codec 实例
func Constructor() Codec {
	return Codec{}
}

// Serialize 二叉树的序列化
// 将一棵二叉树序列化为字符串，空树序列化为单个空节点标记。
// 时间复杂度: O(n)  空间复杂度: O(n)
func (this *Codec) Serialize(root *datastructures.TreeNode) string {
	var sb strings.Builder
	var dfs func(node *datastructures.TreeNode)
	dfs = func(node *datastructures.TreeNode) {
		if sb.Len() > 0 {
			sb.WriteByte(',')
		}
		if node == nil {
			sb.WriteString(nilMark)
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return sb.String()
}

// Deserialize 二叉树的反序列化
// 将 Serialize 生成的字符串还原为二叉树，空节点标记还原为 nil。
// 时间复杂度: O(n)  空间复杂度: O(n)
func (this *Codec) Deserialize(data string) *datastructures.TreeNode {
	tokens := strings.Split(data, ",")
	idx := 0 // 全局消费指针，始终指向下一个待读取的标记
	var dfs func() *datastructures.TreeNode
	dfs = func() *datastructures.TreeNode {
		if idx >= len(tokens) {
			return nil
		}
		tok := tokens[idx]
		idx++
		if tok == nilMark {
			return nil
		}
		val, _ := strconv.Atoi(tok)
		node := &datastructures.TreeNode{Val: val}
		// 前序序列中左子树排在右子树之前，必须先左后右
		node.Left = dfs()
		node.Right = dfs()
		return node
	}
	return dfs()
}
