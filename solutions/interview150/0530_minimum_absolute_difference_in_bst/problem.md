# 530. 二叉搜索树的最小绝对差 (Minimum Absolute Difference in BST)

## 题目描述

给你一个二叉搜索树的根节点 `root`，返回树中任意两不同节点值之间的**最小差值**。

**差值**是一个正数，其数值等于两值之差的绝对值。

### 示例 1

```
输入: root = [4,2,6,1,3]
输出: 1
```

### 示例 2

```
输入: root = [1,0,48,null,null,12,49]
输出: 1
```

## 提示

- 树中节点的数目范围是 `[2, 10^4]`
- `0 <= Node.val <= 10^5`

## 题目解析

### 核心思路

暴力做法是收集所有节点值排序后两两求差，O(n log n)。但二叉搜索树（BST）有一个**杀手级性质**：**中序遍历的结果是严格递增序列**。

递增序列中，两两之差的最小值**必定出现在相邻元素之间**（相隔越远差越大），所以根本不用比较所有组合——只需要中序遍历一遍，记录「上一个访问的节点值 `prev`」和「当前最小差」，每访问一个新节点就算一次 `cur - prev`（递增序列里大减小恒为正，无需取绝对值）。

> 这里 `prev` 用一个哨兵值（`math.MinInt32`）表示「还没有前驱」，跳过第一个节点。

这个「中序 + prev」的模板是解决「BST 中第 K 小」「BST 中两数之和」「验证 BST」等一系列问题的通用骨架，值得熟练掌握。

### 算法步骤

1. 初始化 `minDiff = +∞`，`prev = math.MinInt32`（哨兵，表示无前驱）。
2. 中序遍历（左 → 根 → 右）：
   - 递归处理左子树；
   - 访问当前节点：若 `prev` 不是哨兵，用 `node.Val - prev` 更新 `minDiff`；然后把 `prev` 更新为 `node.Val`；
   - 递归处理右子树。
3. 返回 `minDiff`。

### 复杂度分析

- **时间复杂度**: O(n)，每个节点恰好访问一次
- **空间复杂度**: O(h)，递归栈深度为树高 h；最坏（退化成链）O(n)，平衡时 O(log n)

## 代码实现

```go
func GetMinimumDifference(root *TreeNode) int {
	minDiff := math.MaxInt32
	prev := math.MinInt32 // 上一个中序访问的节点值；初始为极小哨兵
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prev != math.MinInt32 { // 跳过第一个节点（没有前驱）
			if d := node.Val - prev; d < minDiff {
				minDiff = d
			}
		}
		prev = node.Val
		inorder(node.Right)
	}
	inorder(root)
	return minDiff
}
```

**执行过程示例**（`root = [4,2,6,1,3]`）：

```
中序序列: 1 -> 2 -> 3 -> 4 -> 6
访问1: prev=哨兵, 跳过, prev=1
访问2: diff=2-1=1, minDiff=1, prev=2
访问3: diff=3-2=1, minDiff=1, prev=3
访问4: diff=4-3=1, minDiff=1, prev=4
访问6: diff=6-4=2, minDiff=1, prev=6
结果: 1
```
