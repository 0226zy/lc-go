# 102. 二叉树的层序遍历 (Binary Tree Level Order Traversal)

## 题目描述

给你二叉树的根节点 `root`，返回其节点值的**层序遍历**（即逐层地、从左到右访问所有节点）。

### 示例 1

```
输入: root = [3,9,20,null,null,15,7]
输出: [[3],[9,20],[15,7]]
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

层序遍历就是「从上到下、从左到右」一层一层地访问节点，这和**广度优先搜索（BFS）**的模型完全一致：把树看作一张图，用队列逐层扩展即可。

关键难点只有一个：**怎么知道当前这一层到哪里结束**。如果边遍历边往结果里加，所有节点会混在同一个切片里，分不清层与层的边界。

解决办法很朴素：在弹出节点之前，先**记录当前队列的长度 `size`**——这个长度恰好就是当前层的节点个数。然后连续出队 `size` 次，把这一层的节点值收集为一个切片；这些节点入队的孩子自然就是下一层。循环往复，层与层天然被分开。

### 算法步骤

1. 若 `root` 为空，直接返回空结果。
2. 初始化队列，放入根节点。
3. 当队列非空时，循环以下过程：
   - 记录当前队列长度 `size`，新建一个空切片 `level`。
   - 连续出队 `size` 次：把节点值追加到 `level`，并把它的左右孩子（非空时）入队。
   - 把 `level` 追加到结果中。
4. 返回结果。

### 复杂度分析

- **时间复杂度**: O(n)，每个节点恰好入队、出队一次
- **空间复杂度**: O(n)，队列最多同时存放最底层一层的节点，最坏（满二叉树）约 n/2 个

## 代码实现

```go
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue) // 当前层的节点个数
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
初始: queue=[3]
第1层: size=1, 出队3, level=[3], 入队9,20 -> queue=[9,20]
第2层: size=2, 出队9(无孩子), 出队20(入队15,7), level=[9,20] -> queue=[15,7]
第3层: size=2, 出队15, 出队7, level=[15,7] -> queue=[]
结果: [[3],[9,20],[15,7]]
```
