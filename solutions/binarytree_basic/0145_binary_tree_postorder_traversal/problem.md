# 145. 二叉树的后序遍历 (Binary Tree Postorder Traversal)

## 题目描述

给你一棵二叉树的根节点 `root`，返回其节点值的 **后序遍历** 序列。

后序遍历的顺序为：**左子树 -> 右子树 -> 根节点**，即根节点最后访问。

### 示例 1

```
输入: root = [1,null,2,3]
输出: [3,2,1]
```

树结构：

```
  1
   \
    2
   /
  3
```

### 示例 2

```
输入: root = [1,2,3,4,5,null,8,null,null,6,7,9]
输出: [4,6,7,5,2,9,8,3,1]
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

- 树中节点的数目在范围 `[0, 100]` 内
- `-100 <= Node.val <= 100`

**进阶**：递归算法很简单，你可以通过迭代算法完成吗？

## 题目解析

### 核心思路

后序遍历的递归定义非常直接：对任意节点，先递归遍历其**左子树**，再递归遍历其**右子树**，最后才访问**根节点**自身。递归边界是空节点（`node == nil`），直接返回。

由于根节点要最后输出，每个节点都必须等左右子树全部处理完才能加入结果——这正是递归调用栈天然携带的「状态」：调用 `dfs(node.Left)` 和 `dfs(node.Right)` 返回后，当前节点的值才被 `append`，无需额外记录。

**迭代法**的巧妙之处：后序「左 -> 右 -> 根」的逆序是「根 -> 右 -> 左」，而这与前序遍历（根 -> 左 -> 右）形式一致，只需把入栈顺序从「先左后右」换成「先右后左」。因此：

1. 用栈按「根 -> 右 -> 左」的顺序遍历整棵树；
2. 将得到的序列**整体反转**，即为后序结果。

这种方式不需要给节点打「是否已访问」的标记，实现简洁。

### 算法步骤

**递归法**：

1. 定义内部递归函数 `dfs(node)`：
   - 若 `node == nil`，直接返回；
   - 递归 `dfs(node.Left)`；
   - 递归 `dfs(node.Right)`；
   - 将 `node.Val` 追加到结果切片。
2. 从根节点调用 `dfs(root)`，返回结果。

**迭代法**：

1. 若 `root == nil`，返回空切片。
2. 初始化栈，压入 `root`。
3. 循环弹出栈顶节点，将其值追加到结果；然后**先压入左子节点、再压入右子节点**（栈是后进先出，右子节点会先被弹出，形成「根 -> 右 -> 左」的访问序）。
4. 栈空后，将结果切片反转，得到「左 -> 右 -> 根」。

### 复杂度分析

| 方法 | 时间复杂度 | 空间复杂度 |
|------|----------|----------|
| 递归法 | O(n) | O(h)，递归栈深度，最坏 O(n)（链式树） |
| 迭代法 | O(n) | O(h)，显式栈大小，最坏 O(n) |

- **n** 为二叉树节点个数，**h** 为树高。
- 迭代法最后的反转是 O(n)，不改变总时间复杂度量级。

## 代码实现

```go
package postordertraversal

import "github.com/0226zy/lc-go/pkg/datastructures"

type TreeNode = datastructures.TreeNode

// PostorderTraversal 二叉树的后序遍历
// 给定二叉树的根节点 root，返回其节点值的后序遍历序列（左子树 -> 右子树 -> 根节点）。
// 时间复杂度: O(n)  空间复杂度: O(h)，h 为树高，最坏 O(n)
func PostorderTraversal(root *TreeNode) []int {
	result := []int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		result = append(result, node.Val)
	}
	dfs(root)
	return result
}

// PostorderTraversalIterative 迭代法实现二叉树后序遍历（显式栈 + 结果逆置）
// 先按「根 -> 右 -> 左」的顺序入栈遍历，最后将结果反转，即得到「左 -> 右 -> 根」。
// 时间复杂度: O(n)  空间复杂度: O(h)，h 为树高，最坏 O(n)
func PostorderTraversalIterative(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	// 反转结果：「根 -> 右 -> 左」逆置后即为「左 -> 右 -> 根」
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
```

**执行过程示例**（`root = [1,null,2,3]`，即根 1，右子 2，2 的左子 3）：

递归法推演：

```
dfs(1)
├─ dfs(nil)          // 1 的左子为空，直接返回
├─ dfs(2)
│  ├─ dfs(3)
│  │  ├─ dfs(nil)    // 3 的左子为空
│  │  ├─ dfs(nil)    // 3 的右子为空
│  │  └─ 收集 3      // result = [3]
│  ├─ dfs(nil)       // 2 的右子为空
│  └─ 收集 2         // result = [3, 2]
└─ 收集 1            // result = [3, 2, 1]
返回 [3, 2, 1]
```

迭代法推演（同一棵树）：

```
栈=[1]            → 弹出 1，收集；压入右子 2        result=[1]，栈=[2]
栈=[2]            → 弹出 2，收集；压入左子 3        result=[1,2]，栈=[3]
栈=[3]            → 弹出 3，收集；无子节点          result=[1,2,3]，栈=[]
反转 result: [1,2,3] → [3,2,1]
返回 [3, 2, 1]
```
