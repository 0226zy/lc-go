# 637. 二叉树的层平均值 (Average of Levels in Binary Tree)

## 题目描述

给定一个非空二叉树的根节点 `root`，以数组的形式返回每一层节点值的**平均值**。

### 示例 1

```
输入: root = [3,9,20,15,7]
输出: [3.00000,14.50000,11.00000]
解释: 第 0 层的平均值为 3，第 1 层的平均值为 (9 + 20) / 2 = 14.5，第 2 层的平均值为 (15 + 7) / 2 = 11。
```

### 示例 2

```
输入: root = [3,9,20,null,null,15,7]
输出: [3.00000,14.50000,11.00000]
```

## 提示

- 树中节点数量在 `[1, 10^4]` 范围内
- `-2^31 <= Node.val <= 2^31 - 1`

## 题目解析

### 核心思路

「每一层」这个字眼直接对应 **BFS 层序遍历模板**（见 [102. 二叉树的层序遍历](../0102_binary_tree_level_order_traversal/problem.md)）：按层出队时天然能把同一层的节点聚在一起，我们要做的只是**加一层求和再除以个数**。

有两个实现细节值得注意：

1. **先记 `size` 再出队**：出队前记录当前队列长度，这就是本层节点个数，连续出队 `size` 次即完成一层。
2. **浮点累加**：平均值不是整数，求和与除法都要用 `float64`。本题节点值上限约 `10^4`、节点数上限 `10^4`，单层和约 `10^8`，`float64` 的 53 位尾数（约 15~16 位十进制精度）完全够用；按提示的极端值也可用 `int64` 求和再转浮点，这里从简直接用 `float64`。

### 算法步骤

1. 初始化队列放入根节点。
2. 当队列非空时循环：
   - 记录当前层节点个数 `size`，令 `sum = 0`。
   - 连续出队 `size` 次：把节点值累加到 `sum`，左右孩子入队。
   - 把 `sum / size` 追加到结果。
3. 返回结果。

### 复杂度分析

- **时间复杂度**: O(n)，每个节点恰好入队、出队一次
- **空间复杂度**: O(n)，队列最多同时存放最底层一层的节点

## 代码实现

```go
func AverageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var result []float64
	queue := []*TreeNode{root}
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
```

**执行过程示例**（`root = [3,9,20,15,7]`）：

```
第0层: size=1, sum=3     -> avg=3.0      -> queue=[9,20]
第1层: size=2, sum=29    -> avg=14.5     -> queue=[15,7]
第2层: size=2, sum=22    -> avg=11.0     -> queue=[]
结果: [3.0, 14.5, 11.0]
```
