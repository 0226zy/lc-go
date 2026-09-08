# 98. 验证二叉搜索树 (Validate Binary Search Tree)

## 题目描述

给你一个二叉树的根节点 `root`，判断其是否是一个有效的二叉搜索树（BST）。

**有效** 二叉搜索树定义如下：

- 节点的左子树只包含**小于**当前节点的数。
- 节点的右子树只包含**大于**当前节点的数。
- 所有左子树和右子树自身必须也是二叉搜索树。

### 示例 1

```
输入: root = [2,1,3]
输出: true
```

### 示例 2

```
输入: root = [5,1,4,null,null,3,6]
输出: false
解释: 根节点的值是 5，但是右子节点的值是 4。
```

## 提示

- 树中节点数目范围在 `[1, 10^4]` 内
- `-2^31 <= Node.val <= 2^31 - 1`

## 题目解析

### 核心思路

新手最容易踩的坑：只比较每个节点和它的**直接父节点**（`node.Left.Val < node.Val < node.Right.Val`）。这是**错误的**——BST 要求的是「左子树**所有**节点 < 根 < 右子树**所有**节点」。反例就是示例 2：树 `5 / 1 4 / 3 6`，每个父子关系都满足大小关系，但 3 在 5 的右子树里却比 5 小，不是合法 BST。

正确思路有两种：

1. **中序遍历 + 比较前驱**：BST 的中序序列必须严格递增。遍历时记录 `prev`，出现 `node.Val <= prev` 即为非法。（同 [530. 二叉搜索树的最小绝对差](../0530_minimum_absolute_difference_in_bst/problem.md) 的模板）
2. **上下界递归**（本实现采用）：给每个节点划定一个合法值域开区间 `(lower, upper)`：
   - 根节点的值域是 `(-∞, +∞)`；
   - 递归左子树时，上界收缩为当前节点的值；递归右子树时，下界收缩为当前节点的值；
   - 任何节点值落在值域外即返回 false。

上下界递归一遍就能想清楚「整棵子树都要满足」这件事，且可以**提前剪枝**，是最标准的写法。

> 实现细节：题目值域覆盖 int32 全范围（包括 `math.MinInt32` 本身），所以边界要用 `int64` 表示，避免「节点的值恰好等于哨兵边界」造成误判。

### 算法步骤

1. 定义递归函数 `check(node, lower, upper)`：
   - `node` 为空返回 true；
   - 若 `node.Val <= lower` 或 `node.Val >= upper`，返回 false；
   - 返回 `check(node.Left, lower, node.Val) && check(node.Right, node.Val, upper)`。
2. 从 `check(root, math.MinInt64, math.MaxInt64)` 开始。

### 复杂度分析

- **时间复杂度**: O(n)，每个节点恰好访问一次（遇到非法立即剪枝）
- **空间复杂度**: O(h)，递归栈深度为树高 h；最坏（退化链）O(n)，平衡时 O(log n)

## 代码实现

```go
func IsValidBST(root *TreeNode) bool {
	var check func(node *TreeNode, lower, upper int64) bool
	check = func(node *TreeNode, lower, upper int64) bool {
		if node == nil {
			return true
		}
		val := int64(node.Val)
		if val <= lower || val >= upper { // 超出上下界即不合法
			return false
		}
		// 左子树的所有节点必须小于 val，右子树的所有节点必须大于 val
		return check(node.Left, lower, val) && check(node.Right, val, upper)
	}
	return check(root, math.MinInt64, math.MaxInt64)
}
```

**执行过程示例**（`root = [5,1,4,null,null,3,6]`，这是一个陷阱树）：

```
check(5, -∞, +∞): 5 在区间内
  ├─ check(1, -∞, 5):   1 < 5 合法，左右为空 -> true
  └─ check(4, 5, +∞):   4 <= lower(5) -> false!
（4 在 5 的右子树中却比 5 小，整棵树非法）
结果: false
```

对比「只比较父子」的错误写法：5>1、5<4 都满足，会误判为 true——可见**传递上下界**才是完整的约束。
