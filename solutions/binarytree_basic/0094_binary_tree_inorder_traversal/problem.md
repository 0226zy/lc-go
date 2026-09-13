# 94. 二叉树的中序遍历 (Binary Tree Inorder Traversal)

## 题目描述

给定一个二叉树的根节点 `root`，返回 *它的 **中序** 遍历*。

中序遍历的访问顺序为：**左子树 -> 根节点 -> 右子树**。

### 示例 1

```
输入: root = [1,null,2,3]
输出: [1,3,2]
```

### 示例 2

```
输入: root = []
输出: []
```

### 示例 3

```
输入: root = [1]
输出: [1]
```

## 提示

- 树中节点数目在范围 `[0, 100]` 内
- `-100 <= Node.val <= 100`

**进阶**: 递归算法很简单，你可以通过迭代算法完成吗？

## 题目解析

### 核心思路

中序遍历的定义本身就是递归的：**先中序遍历左子树，访问根节点，再中序遍历右子树**。因此递归实现几乎是定义的直译：

- 递归函数 `dfs(node)` 表示「对以 node 为根的子树做中序遍历，把节点值依次追加到结果中」；
- 递归基：`node == nil` 时直接返回，什么都不做；
- 递推过程：`dfs(node.Left)` → 记录 `node.Val` → `dfs(node.Right)`。

递归不携带任何额外状态，遍历结果通过闭包外的 `result` 切片收集。

迭代版本用**显式栈**模拟递归调用栈：从当前节点出发一路向左，把沿途节点全部入栈；左链走到底后弹出栈顶并访问，再转向其右子树重复上述过程。栈里保存的正是「等待被访问的祖先节点」，与递归栈中的待返回帧一一对应。

### 算法步骤

递归法：

1. 定义结果切片 `result` 和递归闭包 `dfs(node)`。
2. `dfs(node)`：若 `node == nil` 则返回；否则依次执行 `dfs(node.Left)`、`append(result, node.Val)`、`dfs(node.Right)`。
3. 调用 `dfs(root)`，返回 `result`。

迭代法：

1. 初始化空栈和指针 `cur = root`。
2. 当 `cur != nil` 或栈非空时循环：
   - 内层循环把 `cur` 沿左链一路入栈，直到 `cur == nil`；
   - 弹出栈顶节点，将其值追加到结果；
   - `cur` 指向弹出节点的右子树。
3. 返回 `result`。

### 复杂度分析

- **时间复杂度**: O(n)，每个节点恰好入栈/访问一次。
- **空间复杂度**: O(h)，h 为树高。递归法的递归调用栈与迭代法的显式栈深度均为树高；最坏情况（链式树）为 O(n)，最好情况（平衡树）为 O(log n)。

## 代码实现

```go
package inordertraversal

import "github.com/0226zy/lc-go/pkg/datastructures"

type TreeNode = datastructures.TreeNode

// InorderTraversal 二叉树的中序遍历
// 给定二叉树的根节点 root，按「左子树 -> 根节点 -> 右子树」的顺序返回所有节点值。
// 时间复杂度: O(n)  空间复杂度: O(h)，h 为树高（递归调用栈）
func InorderTraversal(root *TreeNode) []int {
	var result []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)                    // 先递归左子树
		result = append(result, node.Val) // 再访问根节点
		dfs(node.Right)                   // 最后递归右子树
	}
	dfs(root)
	return result
}

// InorderTraversalIterative 迭代法实现二叉树中序遍历（显式栈）
// 用栈模拟递归过程：一路向左入栈，弹栈访问后再转向右子树。
// 时间复杂度: O(n)  空间复杂度: O(h)，h 为树高
func InorderTraversalIterative(root *TreeNode) []int {
	var result []int
	stack := []*TreeNode{}
	cur := root

	for cur != nil || len(stack) > 0 {
		for cur != nil { // 沿左链一路入栈，直到最左节点
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1] // 弹出栈顶并访问
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		cur = cur.Right // 转向右子树
	}

	return result
}
```

**执行过程示例**（`root = [1,null,2,3]`，即根 1 无左子、右子为 2，2 的左子为 3）：

```
递归法 dfs 调用过程：
  dfs(1)
    dfs(1.Left=nil)        → 直接返回
    访问 1                 → result=[1]
    dfs(1.Right=2)
      dfs(2.Left=3)
        dfs(3.Left=nil)    → 直接返回
        访问 3             → result=[1,3]
        dfs(3.Right=nil)   → 直接返回
      访问 2               → result=[1,3,2]
      dfs(2.Right=nil)     → 直接返回
最终结果: [1,3,2]

迭代法栈变化过程：
  cur=1: 左链入栈 → stack=[1]，cur=nil
  弹出 1，访问     → result=[1]，cur=1.Right=2
  cur=2: 左链入栈 → stack=[2,3]，cur=nil（3 无左子）
  弹出 3，访问     → result=[1,3]，cur=3.Right=nil
  弹出 2，访问     → result=[1,3,2]，cur=2.Right=nil
  栈空且 cur=nil，循环结束
最终结果: [1,3,2]
```
