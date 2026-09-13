# 404. 左叶子之和 (Sum of Left Leaves)

## 题目描述

给定二叉树的 **根节点** `root`，返回所有 **左叶子** 之和。

**左叶子** 的定义：一个节点是另一个节点的 **左孩子**，并且它本身 **没有任何子节点**（左右孩子都为空）。注意区分“左孩子”与“左叶子”——左孩子如果不是叶子，不能计入答案。

### 示例 1

```
输入: root = [3,9,20,null,null,15,7]
输出: 24
解释: 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24。
```

### 示例 2

```
输入: root = [1]
输出: 0
```

## 提示

- 树中节点的数目在范围 `[1, 1000]` 内
- `-1000 <= Node.val <= 1000`

## 题目解析

### 核心思路

本题的关键在于：**一个节点本身无法判断自己是不是左叶子**，因为“是左孩子”这个信息在父节点那里，而“是叶子”这个信息在自己这里。

因此使用 **DFS 携带“是否为左孩子”标记** 的写法：

- 递归函数除了接收当前节点 `node`，再携带一个布尔参数 `isLeft`，表示当前节点是其父节点的左孩子。
- 当遍历到某个节点时，若 `node` 是叶子（左右孩子都为空）**且** `isLeft == true`，说明它是左叶子，把 `node.Val` 累加进答案。
- 否则继续向下递归：左子树传 `true`，右子树传 `false`。
- 根节点没有父节点，约定以 `isLeft = false` 进入递归，因此单节点树的答案是 0（根节点不算左叶子）。

本质上是 **先序遍历** 的一个变体：访问当前节点时做“左叶子判定”，不满足则递归左右子树。

### 算法步骤

1. 定义辅助递归函数 `dfs(node *TreeNode, isLeft bool) int`：
   - 若 `node == nil`，返回 `0`。
   - 若 `node` 是叶子（`node.Left == nil && node.Right == nil`）：
     - 若 `isLeft == true`，返回 `node.Val`；
     - 否则返回 `0`。
   - 否则返回 `dfs(node.Left, true) + dfs(node.Right, false)`。
2. 入口调用 `dfs(root, false)`（根节点没有父节点，不可能是左叶子）。
3. 返回值即为所有左叶子之和。

### 复杂度分析

- **时间复杂度**: O(n)，n 为树中节点数，每个节点恰好被访问一次。
- **空间复杂度**: O(h)，h 为树的高度，即递归栈的深度；最坏情况（退化为链式树）为 O(n)，平衡二叉树为 O(log n)。

## 代码实现

```go
// SumOfLeftLeaves 左叶子之和
// 给定二叉树的根节点 root，返回所有左叶子之和。
// 左叶子指“是父节点的左孩子、且自身没有子节点”的节点。
// 时间复杂度: O(n)  空间复杂度: O(h) 递归栈深度，h 为树高
func SumOfLeftLeaves(root *TreeNode) int {
	var dfs func(node *TreeNode, isLeft bool) int
	dfs = func(node *TreeNode, isLeft bool) int {
		if node == nil {
			return 0
		}
		// 叶子节点：只有同时是左孩子才计入答案
		if node.Left == nil && node.Right == nil {
			if isLeft {
				return node.Val
			}
			return 0
		}
		// 非叶子节点：累加左右子树中的左叶子之和
		return dfs(node.Left, true) + dfs(node.Right, false)
	}
	return dfs(root, false) // 根节点没有父节点，不是左叶子
}
```

**执行过程示例**（`root = [3,9,20,null,null,15,7]`）：

```
树形结构:
        3
       / \
      9   20
         /  \
        15    7

递归过程（dfs(node, isLeft)）:
dfs(3, false)
├─ 3 不是叶子 → 递归左右子树
├─ dfs(9, true)
│  └─ 9 是叶子且 isLeft=true → 返回 9        （左叶子 9）
└─ dfs(20, false)
   ├─ 20 不是叶子 → 递归左右子树
   ├─ dfs(15, true)
   │  └─ 15 是叶子且 isLeft=true → 返回 15   （左叶子 15）
   └─ dfs(7, false)
      └─ 7 是叶子但 isLeft=false → 返回 0    （右叶子，不计入）

结果: 9 + 15 + 0 = 24
```
