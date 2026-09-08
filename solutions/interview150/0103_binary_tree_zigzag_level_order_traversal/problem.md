# 103. 二叉树的锯齿形层序遍历 (Binary Tree Zigzag Level Order Traversal)

## 题目描述

给你二叉树的根节点 `root`，返回其节点值的**锯齿形层序遍历**（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

### 示例 1

```
输入: root = [3,9,20,null,null,15,7]
输出: [[3],[20,9],[15,7]]
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

这道题就是 [102. 二叉树的层序遍历](../0102_binary_tree_level_order_traversal/problem.md) 的「加料版」：BFS 框架完全不变，只是在**输出某些层时要倒过来**。

最容易想到的两种做法：

1. **层内反转**：正常 BFS，每层收集完后，如果当前层是偶数层（从 0 开始计），就反转切片。
2. **下标镜像写入**（本实现采用）：在收集本层节点时，根据方向直接决定写入位置——从左到右的层按 `level[0..size-1]` 正序写，从右到左的层按 `level[size-1..0]` 倒序写。这样省掉一次反转（虽然复杂度不变，但更直观），且**入队顺序始终不变**（孩子永远按「先左后右」入队，保证下一层的相对次序正确）。

方向用一个布尔标记 `leftToRight` 维护，每处理完一层就取反。

### 算法步骤

1. 若 `root` 为空，直接返回空结果。
2. 初始化队列放入根节点，`leftToRight = true`。
3. 当队列非空时循环：
   - 记录当前层节点个数 `size`，建一个长度为 `size` 的切片 `level`。
   - 连续出队 `size` 次：若 `leftToRight` 为真，把节点值写到 `level[i]`，否则写到 `level[size-1-i]`；左右孩子按「先左后右」入队。
   - 把 `level` 追加到结果，`leftToRight` 取反。
4. 返回结果。

### 复杂度分析

- **时间复杂度**: O(n)，每个节点恰好入队、出队一次
- **空间复杂度**: O(n)，队列最多同时存放最底层一层的节点

## 代码实现

```go
func ZigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	queue := []*TreeNode{root}
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
```

**执行过程示例**（`root = [3,9,20,null,null,15,7]`）：

```
第1层(leftToRight=true):  出队3  -> level=[3]                       -> queue=[9,20]
第2层(leftToRight=false): 出队9  写到 level[1], 出队20 写到 level[0]
                          -> level=[20,9], 入队15,7               -> queue=[15,7]
第3层(leftToRight=true):  出队15 写到 level[0], 出队7 写到 level[1]
                          -> level=[15,7]                         -> queue=[]
结果: [[3],[20,9],[15,7]]
```

> 注意：无论输出方向如何，孩子入队永远是「先左后右」，这样下一层在队列中的相对次序才是从左到右，反转后才正确。
