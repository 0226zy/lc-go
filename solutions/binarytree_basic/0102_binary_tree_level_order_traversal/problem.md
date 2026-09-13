# 102. 二叉树的层序遍历 (Binary Tree Level Order Traversal)

## 题目描述

给你二叉树的根节点 `root`，返回其节点值的**层序遍历**（即逐层地、从左到右访问所有节点）。结果是一个二维切片，每个内层切片对应一层的节点值。

### 示例 1

```
输入: root = [3,9,20,null,null,15,7]
输出: [[3],[9,20],[15,7]]
```

树形结构：

```
    3
   / \
  9  20
    /  \
   15   7
```

### 示例 2

```
输入: root = [1]
输出: [[1]]
```

### 示例 3

```
输入: root = []
输出: []
```

## 提示

- 树中节点数目在范围 `[0, 2000]` 内
- `-1000 <= Node.val <= 1000`

## 题目解析

### 核心思路

层序遍历要求「从上到下、从左到右」逐层访问节点，这正是**广度优先搜索（BFS）**的标准模型：把树当作一张特殊的图，用**队列**逐层扩展——先进先出的特性保证了先访问的节点先扩展，自然形成按层推进的顺序。

遍历顺序：第 0 层（根）→ 第 1 层 → ……，每层内部严格从左到右（入队时先左孩子后右孩子）。

唯一需要解决的问题是**层与层的边界**：如果边遍历边收集节点值，所有值会混进同一个切片，无法分清哪一层结束。解法是在弹出节点之前先记录当前队列长度 `size`——它恰好等于当前层的节点数。接下来连续出队 `size` 次，把这批节点的值收成一个切片；期间入队的孩子全部属于下一层，留待下一轮处理。如此循环，层与层被天然分隔。

状态携带方式：队列中只携带节点指针本身（`*TreeNode`），层的信息不需要显式携带，而是由「每轮循环固定消费 `size` 个节点」这一控制流隐式表达。

### 算法步骤

1. 若 `root` 为空，直接返回 `nil`。
2. 初始化队列，放入根节点。
3. 当队列非空时，重复：
   - 记录当前队列长度 `size`，新建容量为 `size` 的切片 `level`；
   - 连续出队 `size` 次：把节点值追加到 `level`，并将其非空的左、右孩子依次入队；
   - 将 `level` 追加到结果中。
4. 返回结果。

### 复杂度分析

- **时间复杂度**: O(n)，每个节点恰好入队、出队各一次
- **空间复杂度**: O(n)，队列最多同时存放一整层的节点，最坏情况（满二叉树的最底层）约 n/2 个

## 代码实现

```go
package levelorder

import "github.com/0226zy/lc-go/pkg/datastructures"

// LevelOrder 二叉树的层序遍历
// 给定二叉树的根节点 root，从上到下、从左到右逐层返回每层节点的值。
// 时间复杂度: O(n) 每个节点恰好入队、出队各一次  空间复杂度: O(n) 队列最多同时存放一层的节点
func LevelOrder(root *datastructures.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	queue := []*datastructures.TreeNode{root}
	for len(queue) > 0 {
		size := len(queue) // 当前层节点个数：层与层边界的唯一依据
		level := make([]int, 0, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}
```

**执行过程示例**（`root = [3,9,20,null,null,15,7]`）：

```
初始:  queue=[3]，result=[]

第 1 轮: size=1
  出队 3 -> level=[3]，孩子 9、20 入队 -> queue=[9,20]
  result=[[3]]

第 2 轮: size=2
  出队 9  -> level=[9]，无孩子
  出队 20 -> level=[9,20]，孩子 15、7 入队 -> queue=[15,7]
  result=[[3],[9,20]]

第 3 轮: size=2
  出队 15 -> level=[15]，无孩子
  出队 7  -> level=[15,7]，无孩子 -> queue=[]
  result=[[3],[9,20],[15,7]]

队列为空，结束。返回 [[3],[9,20],[15,7]]
```
