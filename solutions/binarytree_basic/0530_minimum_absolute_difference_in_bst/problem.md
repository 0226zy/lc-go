# 530. 二叉搜索树的最小绝对差 (Minimum Absolute Difference in BST)

## 题目描述

给你一个二叉搜索树的根节点 `root`，返回**树中任意两不同节点的值之间的最小差值**。

差值是一个正数，其数值等于两值之差的绝对值。

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

如果把所有节点值收集起来排序再两两求差，需要 O(n log n)。但二叉搜索树（BST）自带一个关键性质：**中序遍历（左 → 根 → 右）得到的序列是严格递增的**。

在递增序列中，任意两数之差的最小值**一定出现在相邻元素之间**——相隔越远的两个数差值越大。因此只需一次中序遍历，边遍历边比较「当前节点值」与「上一个访问节点值」的差即可，无需保存整个序列。

遍历过程中用闭包携带两份状态：

- `prev`：上一个中序访问节点的值，用 `*int` 指针表示，`nil` 表示「还没有前驱」（第一个被访问的节点跳过比较）；
- `minDiff`：迄今为止的最小差，初始为正无穷（`math.MaxInt32`）。

由于是递增序列，`node.Val - *prev` 恒为非负数，直接相减即可，不用再取绝对值。本题没有额外剪枝空间——每个节点都必须访问一次才能确认最小差。

> 「中序遍历 + prev 指针」是 BST 题的通用模板，验证 BST、找第 K 小、众数等题目共用这一骨架。

### 算法步骤

1. 初始化 `minDiff = math.MaxInt32`，`prev = nil`。
2. 递归中序遍历：
   - 先递归左子树；
   - 访问当前节点：若 `prev != nil`，计算 `node.Val - *prev` 并更新 `minDiff`；随后令 `prev` 指向当前节点值；
   - 再递归右子树。
3. 遍历结束返回 `minDiff`。

### 复杂度分析

- **时间复杂度**: O(n)，n 为节点数，每个节点恰好访问一次。
- **空间复杂度**: O(h)，h 为树高，即递归调用栈深度；平衡树为 O(log n)，退化成链时为 O(n)。

## 代码实现

```go
package minimumabsolutedifferenceinbst

import (
	"math"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// GetMinimumDifference 二叉搜索树的最小绝对差
// 给定二叉搜索树的根节点，返回树中任意两个不同节点值之差的最小绝对值。
// 时间复杂度: O(n)  空间复杂度: O(h)，h 为树高（递归栈深度）
func GetMinimumDifference(root *datastructures.TreeNode) int {
	minDiff := math.MaxInt32
	// prev 指向上一个中序访问节点的值；用指针的 nil 区分"尚未访问任何节点"，
	// 避免哨兵值与节点真实取值冲突
	var prev *int

	var inorder func(node *datastructures.TreeNode)
	inorder = func(node *datastructures.TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prev != nil {
			if diff := node.Val - *prev; diff < minDiff {
				minDiff = diff
			}
		}
		val := node.Val
		prev = &val
		inorder(node.Right)
	}
	inorder(root)

	return minDiff
}
```

## 执行过程示例

以 `root = [4,2,6,1,3]` 为例，树的形状：

```
      4
     / \
    2   6
   / \
  1   3
```

中序遍历（左 → 根 → 右）的访问过程：

```
中序序列: 1 -> 2 -> 3 -> 4 -> 6

访问 1: prev == nil，跳过比较，prev = 1
访问 2: diff = 2 - 1 = 1，minDiff = 1，prev = 2
访问 3: diff = 3 - 2 = 1，minDiff = 1，prev = 3
访问 4: diff = 4 - 3 = 1，minDiff = 1，prev = 4
访问 6: diff = 6 - 4 = 2，minDiff 保持 1，prev = 6

返回: 1
```
