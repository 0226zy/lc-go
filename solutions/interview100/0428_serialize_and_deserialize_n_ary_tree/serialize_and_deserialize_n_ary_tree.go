package serializeanddeserializenarytree

import (
	"strconv"
	"strings"
)

// Node N 叉树节点定义
type Node struct {
	Val      int
	Children []*Node
}

// Codec N 叉树序列化/反序列化编解码器
// 编码格式：先序遍历，每个节点依次写入「节点值 孩子个数」，空格分隔。
// 孩子个数使反序列化时能唯一确定每个节点的子树边界。
type Codec struct {
}

// Constructor 创建编解码器实例
func Constructor() Codec {
	return Codec{}
}

// serialize 将 N 叉树序列化为字符串
// 时间复杂度: O(n) 每个节点访问一次  空间复杂度: O(h) 递归栈深度，h 为树高
func (c *Codec) serialize(root *Node) string {
	var sb strings.Builder
	var dfs func(node *Node)
	dfs = func(node *Node) {
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteByte(' ')
		sb.WriteString(strconv.Itoa(len(node.Children)))
		for _, child := range node.Children {
			sb.WriteByte(' ')
			dfs(child)
		}
	}
	if root != nil {
		dfs(root)
	}
	return sb.String()
}

// deserialize 将字符串反序列化为 N 叉树
// 时间复杂度: O(n) 每个数字项读取一次  空间复杂度: O(h) 递归栈深度
func (c *Codec) deserialize(data string) *Node {
	if data == "" {
		return nil
	}
	tokens := strings.Split(data, " ")
	idx := 0
	var dfs func() *Node
	dfs = func() *Node {
		val, _ := strconv.Atoi(tokens[idx])
		idx++
		cnt, _ := strconv.Atoi(tokens[idx])
		idx++
		node := &Node{Val: val}
		for i := 0; i < cnt; i++ {
			node.Children = append(node.Children, dfs())
		}
		return node
	}
	return dfs()
}
