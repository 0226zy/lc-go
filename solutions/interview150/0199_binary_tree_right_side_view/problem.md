# 199. 二叉树的右视图 (Binary Tree Right Side View)

## 题目描述

给定一个二叉树的**根节点** `root`，想象自己站在它的右侧，按照从顶部到底部的顺序，返回你能从右侧看到的节点的值。

### 示例 1

```
输入: root = [1,2,3,null,5,null,4]
输出: [1,3,4]
```

### 示例 2

```
输入: root = [1,null,3]
输出: [1,3]
```

### 示例 3

```
输入: root = []
输出: []
```

## 提示

- 二叉树的节点个数的范围是 `[0, 100]`
- `-100 <= Node.val <= 100`

## 题目解析

### 核心思路

站在树的右边看过去，每一层**最多只能看到一个节点**——就是最靠右的那个。所以右视图本质上是：**每一层的最后一个节点**组成的序列。

「按层处理」立刻让人想到 BFS 层序遍历模板（见 [102. 二叉树的层序遍历](../0102_binary_tree_level_order_traversal/problem.md)）。在层序遍历中，我们用「出队前记录队列长度 `size`」的技巧确定层的边界；那么「每层最后一个节点」就是本层 `size` 个节点中**最后一个出队的那个**（下标 `size-1`）。

一个容易误解的点：最右节点**不一定是右孩子**。比如某节点的右孩子为空、左孩子很深，站在右侧看到的仍是这个左子树里的节点（本实现按层处理天然正确；若用「先递归右子树」的 DFS 写法，必须右、左都递归）。

### 算法步骤

1. 若 `root` 为空，返回空结果。
2. 初始化队列放入根节点。
3. 当队列非空时循环：
   - 记录当前层节点个数 `size`。
   - 连续出队 `size` 次；当出队的是本层第 `size-1` 个（最后一个）节点时，把它的值加入结果。
   - 左右孩子按顺序入队。
4. 返回结果。

### 复杂度分析

- **时间复杂度**: O(n)，每个节点恰好入队、出队一次
- **空间复杂度**: O(n)，队列最多同时存放最底层一层的节点

## 代码实现

```go
func RightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue) // 当前层的节点个数
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == size-1 { // 本层最后一个出队的节点，即最右侧节点
				result = append(result, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}
```

**执行过程示例**（`root = [1,2,3,null,5,null,4]`）：

```
第1层: size=1, 出队1(最后一个) -> result=[1], 入队2,3
第2层: size=2, 出队2, 出队3(最后一个) -> result=[1,3], 入队5,4
第3层: size=2, 出队5, 出队4(最后一个) -> result=[1,3,4]
结果: [1,3,4]
```
