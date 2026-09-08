# 173. 二叉搜索树迭代器 (Binary Search Tree Iterator)

## 题目描述

实现一个二叉搜索树迭代器类 `BSTIterator`，表示二叉搜索树（BST）中值的**中序遍历**：

- `BSTIterator(TreeNode root)` 初始化一个 `BSTIterator` 类对象。BST 的根节点 `root` 会作为构造函数的一部分给出。指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。
- `boolean hasNext()` 如果向指针右侧遍历存在数字，则返回 `true`；否则返回 `false`。
- `int next()` 将指针向右移动，然后返回指针处的数字。

可以认为 `next()` 的每次调用都是有效的，也就是说，当调用 `next()` 时，BST 的中序遍历中至少存在一个下一个数字。

**注意**：指针初始化为一个小于 BST 中任意元素的值，第一次调用 `next()` 会返回中序遍历的第一个元素。

### 示例 1

```
输入：
["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
[[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]

输出：
[null, 3, 7, true, 9, true, 15, true, 20, false]

解释：
BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
bSTIterator.next();    // 返回 3
bSTIterator.next();    // 返回 7
bSTIterator.hasNext(); // 返回 True
bSTIterator.next();    // 返回 9
bSTIterator.hasNext(); // 返回 True
bSTIterator.next();    // 返回 15
bSTIterator.hasNext(); // 返回 True
bSTIterator.next();    // 返回 20
bSTIterator.hasNext(); // 返回 False
```

补充说明：本题只有这一个官方示例；当 `hasNext` 返回 `false` 后不会再调用 `next`。

## 提示

- 树中节点的数目在范围 `[0, 10^5]` 内
- `0 <= Node.val <= 10^6`
- 最多可以调用 `10^5` 次 `hasNext` 和 `next` 操作

## 题目解析

### 核心思路

BST 的中序遍历（左 → 根 → 右）会得到一个**升序序列**。最直觉的做法是构造函数里一次性遍历整棵树，把所有值存进一个数组，然后用下标模拟指针。

但题目进阶要求：**能否设计一个平均 O(1) 时间、O(h) 空间的算法**（h 是树高）？一次存整棵树需要 O(n) 空间，不满足。

关键在于“**受控递归**”：标准的中序遍历是用栈实现的迭代版——先一路把左子树压栈，弹出一个节点时输出它，然后转向它的右子树。我们把这个迭代过程**拆散到每次 `Next()` 调用里**：

- 栈里始终保存“从根到当前待输出节点”的左链路上的节点。
- `Next()` 弹出栈顶（即当前最小值），返回它之前，先把它的右子树的左链路全部压栈（相当于递归进入右子树）。
- `HasNext()` 只需判断栈是否为空。

这样栈中最多同时保存树高 h 个节点，空间 O(h)；每个节点恰好入栈一次、出栈一次，`Next()` 的**均摊**时间复杂度是 O(1)（虽然单次弹栈后可能压入一整条左链是 O(h)，但均摊到每个节点是 O(1)）。

这属于“**迭代器设计 = 把递归/遍历过程拆成可暂停的步骤 + 栈保存现场**”这一模板，和第 284 题“顶端迭代器”思路一致。

### 算法步骤

1. `Constructor(root)`：从 `root` 出发，沿左孩子一路压栈，直到遇到 `nil`。
2. `Next()`：
   - 弹出栈顶节点 `node`（中序遍历的下一个元素）；
   - 令 `cur = node.Right`，沿左孩子一路压栈；
   - 返回 `node.Val`。
3. `HasNext()`：返回 `len(stack) > 0`。

### 复杂度分析

- **时间复杂度**: `Next()` 均摊 O(1)，`HasNext()` O(1)。因为每个节点恰好入栈、出栈各一次，n 次 `Next()` 总共 O(n) 工作量。
- **空间复杂度**: O(h)，栈中最多存放树高个节点（平衡 BST 为 O(log n)，退化为链表时 O(n)）。

## 代码实现

```go
type BSTIterator struct {
	stack []*datastructures.TreeNode // 保存尚未输出的节点
}

// Constructor 初始化迭代器，把根节点所在左链路压栈
func Constructor(root *datastructures.TreeNode) BSTIterator {
	it := BSTIterator{}
	it.pushLeft(root)
	return it
}

// pushLeft 把 node 及其所有左孩子一路压栈
func (it *BSTIterator) pushLeft(node *datastructures.TreeNode) {
	for node != nil {
		it.stack = append(it.stack, node)
		node = node.Left
	}
}

// Next 返回中序遍历的下一个元素
func (it *BSTIterator) Next() int {
	n := len(it.stack)
	top := it.stack[n-1]   // 栈顶即当前最小值
	it.stack = it.stack[:n-1]
	it.pushLeft(top.Right) // 转向右子树，继续压左链
	return top.Val
}

// HasNext 判断是否还有下一个元素
func (it *BSTIterator) HasNext() bool {
	return len(it.stack) > 0
}
```

**执行过程示例**（树 `[7,3,15,null,null,9,20]`，中序应为 `[3,7,9,15,20]`）：

```
Constructor: 从 7 压左链 → 栈 [7, 3]（栈顶在右，3 是栈顶）
Next():      弹出 3，3.Right 为空，栈 [7]，返回 3
Next():      弹出 7，压 7.Right(15) 的左链 → 栈 [15, 9]，返回 7
HasNext():   栈非空 → true
Next():      弹出 9，9.Right 为空，栈 [15]，返回 9
Next():      弹出 15，压 15.Right(20) 的左链 → 栈 [20]，返回 15
Next():      弹出 20，栈空，返回 20
HasNext():   栈空 → false
```
