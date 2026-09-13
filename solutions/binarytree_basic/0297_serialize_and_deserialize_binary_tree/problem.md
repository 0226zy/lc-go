# 297. 二叉树的序列化与反序列化 (Serialize and Deserialize Binary Tree)

## 题目描述

序列化是将一个数据结构或对象转换为一系列比特的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，以便稍后在同一个或另一个计算机环境中重建。

设计一个算法来序列化和反序列化**二叉树**。保证二叉树可以序列化为字符串，并且可以将字符串反序列化为原始树结构。

你不必将序列化 / 反序列化算法限制为某种特定格式，只需保证能正确地重建即可。

### 示例 1

```
输入: root = [1,2,3,null,null,4,5]
输出: [1,2,3,null,null,4,5]
```

### 示例 2

```
输入: root = []
输出: []
```

## 提示

- 树中结点数在范围 `[0, 10^4]` 内
- `-1000 <= Node.val <= 1000`

## 题目解析

### 核心思路

任意一种能唯一确定树结构的遍历都可以：层序、前序、后序等。本题采用**前序遍历 + 空节点标记**：

- 序列化：按「根 → 左 → 右」输出节点值，空孩子记为特殊标记 `"N"`，用逗号拼接。
- 反序列化：按同样的前序顺序消费字符串 token；读到 `"N"` 返回 `nil`，否则创建节点后先递归左再递归右。

为什么前序特别好写？因为根总在子树之前出现，反序列化时「读一个 token → 建根 → 再建左右」与序列化顺序完全对称，用一个全局下标即可线性还原。

空节点必须显式编码：否则无法区分「左孩子为空、右孩子非空」与「只有左孩子」等情况。

### 算法步骤

**Serialize**

1. DFS 前序遍历：非空节点写数值，空节点写 `"N"`，token 之间用 `,` 分隔。
2. 返回拼接后的字符串（空树为单个 `"N"`）。

**Deserialize**

1. 按 `,` 切分得到 token 数组，维护消费下标 `idx`。
2. 递归：取当前 token 并 `idx++`；若为 `"N"` 返回 `nil`；否则创建节点，先递归赋值 `Left`，再递归赋值 `Right`。
3. 返回重建的根。

### 复杂度分析

- **时间复杂度**: O(n)，每个节点序列化 / 反序列化各一次
- **空间复杂度**: O(n)，字符串长度与递归栈（最坏链状 O(n)）

## 代码实现

```go
const nilMark = "N"

func (c *Codec) Serialize(root *TreeNode) string {
	var sb strings.Builder
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
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

func (c *Codec) Deserialize(data string) *TreeNode {
	tokens := strings.Split(data, ",")
	idx := 0
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if idx >= len(tokens) {
			return nil
		}
		tok := tokens[idx]
		idx++
		if tok == nilMark {
			return nil
		}
		val, _ := strconv.Atoi(tok)
		node := &TreeNode{Val: val}
		node.Left = dfs()
		node.Right = dfs()
		return node
	}
	return dfs()
}
```
