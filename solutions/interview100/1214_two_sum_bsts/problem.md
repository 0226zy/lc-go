# 1214. 查找两棵二叉搜索树之和

> 难度：中等 ｜ 分类：二叉搜索树 ｜ 尊享面试 100 题 · 第 48 题
> 链接：https://leetcode.cn/problems/two-sum-bsts/

## 题目描述

给出两棵二叉搜索树的根节点 `root1` 和 `root2`，以及一个整数 `target`。

如果第一棵树中存在某个节点 `a`、第二棵树中存在某个节点 `b`，使得 `a + b == target`，返回 `true`；否则返回 `false`。

### 示例 1

```
输入: root1 = [2,1,4], root2 = [1,0,3], target = 5
输出: true
解释: 2 + 3 = 5，其中 2 在第一棵树中，3 在第二棵树中。
```

### 示例 2

```
输入: root1 = [0,-10,10], root2 = [5,1,7,0,2], target = 18
输出: false
解释: 第一棵树的任一节点与第二棵树的任一节点之和都不等于 18。
```

### 提示

- 每棵树的节点数最多为 `5000`
- `-10^9 <= Node.val, target <= 10^9`

## 思路解析

### 核心思路

这是「两数之和」在树上的版本。经典做法是**哈希集合**：

1. 遍历第一棵树，把所有节点值存入哈希集合；
2. 遍历第二棵树，对每个节点值 `b` 检查 `target - b` 是否在集合中。

两棵树各遍历一次即可。题目说是二叉搜索树，但这个做法对任意二叉树都成立；若想利用 BST 性质，也可以中序遍历得到两个有序数组后用双指针，复杂度相同。

### 算法步骤

1. 用递归（或迭代）遍历 `root1`，把所有节点值加入集合 `seen`。
2. 递归遍历 `root2`：
   - 若 `target - node.Val` 存在于 `seen`，返回 `true`；
   - 否则继续遍历左右子树。
3. 遍历完仍未找到，返回 `false`。

### 复杂度分析

- **时间复杂度**: O(n1 + n2)，n1、n2 分别为两棵树的节点数，每棵树各遍历一次
- **空间复杂度**: O(n1)，哈希集合存第一棵树的所有节点值（不计递归栈）

## 代码实现

```go
func TwoSumBSTs(root1, root2 *datastructures.TreeNode, target int) bool {
	// 第一遍：收集 root1 的所有节点值
	seen := make(map[int]bool)
	var collect func(node *datastructures.TreeNode)
	collect = func(node *datastructures.TreeNode) {
		if node == nil {
			return
		}
		seen[node.Val] = true
		collect(node.Left)
		collect(node.Right)
	}
	collect(root1)

	// 第二遍：在 root2 中查找 target - Val 是否出现过
	var search func(node *datastructures.TreeNode) bool
	search = func(node *datastructures.TreeNode) bool {
		if node == nil {
			return false
		}
		if seen[target-node.Val] {
			return true
		}
		return search(node.Left) || search(node.Right)
	}
	return search(root2)
}
```
