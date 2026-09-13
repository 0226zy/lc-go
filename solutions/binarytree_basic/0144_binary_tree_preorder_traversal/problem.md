# 144. 二叉树的前序遍历 (Binary Tree Preorder Traversal)

## 题目描述

给你二叉树的根节点 `root`，返回它节点值的**前序**遍历。

前序遍历的顺序为：**根节点 -> 左子树 -> 右子树**。

### 示例 1

```
输入: root = [1,null,2,3]

     1
      \
       2
      /
     3

输出: [1,2,3]
```

### 示例 2

```
输入: root = [1,2,3,4,5,null,8,null,null,6,7,9]

         1
        / \
       2   3
      / \   \
     4   5   8
        / \  /
       6   7 9

输出: [1,2,4,5,6,7,3,8,9]
```

### 示例 3

```
输入: root = []
输出: []
```

### 示例 4

```
输入: root = [1]
输出: [1]
```

## 提示

- 树中节点数目在范围 `[0, 100]` 内
- `-100 <= Node.val <= 100`

**进阶**：递归算法很简单，你可以通过迭代算法完成吗？

## 题目解析

### 核心思路

前序遍历的递归定义非常直白：对任意一棵子树，**先访问根节点，再前序遍历左子树，最后前序遍历右子树**。整棵树的遍历结果就是这三个部分按顺序拼接。

两种实现方式：

1. **递归法**：直接翻译递归定义。每进入一个节点立刻把它的值加入结果（这就是「前序」的含义——根在最前面），然后依次递归左、右子树。状态通过递归调用栈隐式携带，无需额外维护。
2. **迭代法（显式栈）**：用栈模拟递归。关键是栈是「后进先出」，所以**先把右子节点入栈、再把左子节点入栈**，弹出时才能先处理左子树。每弹出一个节点就访问它（根优先），然后压入它的右、左子节点。本题不涉及回溯或剪枝，每个节点恰好入栈、出栈一次。

### 算法步骤

**递归法：**

1. 若当前节点为 `nil`，直接返回（递归边界）。
2. 把当前节点的值追加到结果切片。
3. 递归遍历左子树。
4. 递归遍历右子树。

**迭代法：**

1. 若 `root` 为 `nil`，返回空切片。
2. 初始化栈，将 `root` 入栈。
3. 循环直到栈为空：
   - 弹出栈顶节点，把它的值追加到结果切片；
   - 若右子节点非空，右子节点入栈；
   - 若左子节点非空，左子节点入栈（后入栈的先弹出，保证左先于右被访问）。
4. 返回结果切片。

### 复杂度分析

| 方法 | 时间复杂度 | 空间复杂度 |
|------|----------|----------|
| 递归法 | O(n) | O(h)，递归调用栈深度为树高，最坏（链式树）O(n) |
| 迭代法 | O(n) | O(h)，显式栈最大深度为树高，最坏 O(n) |

- **n** 为二叉树节点个数
- **h** 为二叉树高度

## 代码实现

```go
package preordertraversal

import "github.com/0226zy/lc-go/pkg/datastructures"

type TreeNode = datastructures.TreeNode

// PreorderTraversal 二叉树的前序遍历
// 给定二叉树的根节点 root，返回其节点值的前序遍历结果（根 -> 左 -> 右）。
// 时间复杂度: O(n)  空间复杂度: O(h)，h 为树高，最坏 O(n)
func PreorderTraversal(root *TreeNode) []int {
	result := []int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return result
}

// PreorderTraversalIterative 迭代法实现二叉树前序遍历（显式栈）
// 栈中先压入右子节点再压入左子节点，保证弹出顺序为「根 -> 左 -> 右」。
// 时间复杂度: O(n)  空间复杂度: O(h)，h 为树高，最坏 O(n)
func PreorderTraversalIterative(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result
}
```

## 执行过程示例

以示例 1 的树 `[1,null,2,3]` 为例（1 为根，2 是 1 的右子节点，3 是 2 的左子节点）：

**递归法调用过程：**

1. `dfs(1)`：1 非空，`result = [1]`；递归 `dfs(1.Left)` 即 `dfs(nil)`，直接返回；递归 `dfs(1.Right)` 即 `dfs(2)`。
2. `dfs(2)`：2 非空，`result = [1, 2]`；递归 `dfs(2.Left)` 即 `dfs(3)`。
3. `dfs(3)`：3 非空，`result = [1, 2, 3]`；`dfs(3.Left)` 与 `dfs(3.Right)` 均为 `nil`，直接返回。
4. 回到 `dfs(2)`，递归 `dfs(2.Right)` 即 `dfs(nil)`，返回；`dfs(2)` 结束，`dfs(1)` 结束。
5. 最终结果：`[1, 2, 3]`。

**迭代法栈变化过程：**

| 步骤 | 弹出节点 | result | 操作后栈（栈顶在右） |
|------|---------|--------|---------------------|
| 初始 | - | `[]` | `[1]` |
| 1 | 1 | `[1]` | 右子 2 入栈 → `[2]` |
| 2 | 2 | `[1,2]` | 左子 3 入栈 → `[3]` |
| 3 | 3 | `[1,2,3]` | 无子节点 → `[]` |

栈空，循环结束，返回 `[1, 2, 3]`。可以看到每个节点一出栈就被访问，且由于右子节点先入栈，左子节点总是先被弹出处理。
