# 230. 二叉搜索树中第 K 小的元素 (Kth Smallest Element in a BST)

## 题目描述

给定一个二叉搜索树的根节点 `root`，和一个整数 `k`，请你设计一个算法查找其中第 `k` 小的元素（从 1 开始计数）。

### 示例 1

```
输入: root = [3,1,4,null,2], k = 1
输出: 1
```

### 示例 2

```
输入: root = [5,3,6,2,4,null,null,1], k = 3
输出: 3
```

## 提示

- 树中的节点数为 `n`
- `1 <= k <= n <= 10^4`
- `0 <= Node.val <= 10^4`

**进阶**：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？

## 题目解析

### 核心思路

最直接的想法是把所有节点值收集起来排序，取第 k 个——O(n log n)。但 BST 有一个关键性质：**中序遍历（左 → 根 → 右）得到的就是递增序列**。所以「第 k 小」就等于「中序遍历时数到的第 k 个节点」，一遍 O(n) 遍历即可，甚至**不需要存整个序列**。

进阶优化：给每个节点额外维护「以该节点为根的子树节点数 `size`」，就能在 O(h) 内定位第 k 小——左子树节点数 `ls` 若等于 k-1，当前节点就是答案；k ≤ ls 递归左子树；否则递归右子树并令 k -= ls + 1。代价是插入删除时要维护 `size`。

### 算法步骤

1. 初始化计数器 `count = 0`，答案 `result = -1`。
2. 中序遍历，并让递归函数返回 `bool` 表示「是否已找到」（找到即提前剪枝，不再遍历剩余节点）：
   - 递归左子树，若返回 true 直接向上传递；
   - 访问当前节点：`count++`，若 `count == k`，记录 `result = node.Val` 并返回 true；
   - 递归右子树。
3. 返回 `result`。

### 复杂度分析

- **时间复杂度**: O(h + k)，最坏 O(n)；只需走到中序第 k 个节点，找到后立即停止
- **空间复杂度**: O(h)，递归栈深度为树高 h；最坏（退化链）O(n)，平衡时 O(log n)

## 代码实现

```go
func KthSmallest(root *TreeNode, k int) int {
	count := 0
	result := -1
	var inorder func(node *TreeNode) bool
	inorder = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		// 左子树中找到则直接返回
		if inorder(node.Left) {
			return true
		}
		count++
		if count == k { // 中序第 k 个节点即第 k 小
			result = node.Val
			return true
		}
		return inorder(node.Right)
	}
	inorder(root)
	return result
}
```

**执行过程示例**（`root = [3,1,4,null,2], k = 1`）：

```
中序走向: 先到最左下角
访问1: count=1 == k=1 -> result=1, 立即返回 true
（整个右子树不再遍历，提前剪枝）
结果: 1
```
