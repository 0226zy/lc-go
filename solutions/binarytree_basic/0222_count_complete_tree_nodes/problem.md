# 222. 完全二叉树的节点个数 (Count Complete Tree Nodes)

## 题目描述

给你一棵 **完全二叉树** 的根节点 `root`，求出该树的节点个数。

完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 `h` 层，则该层包含 `1 ~ 2^(h-1)` 个节点。

### 示例 1

```
输入: root = [1,2,3,4,5,6]
输出: 6
```

### 示例 2

```
输入: root = []
输出: 0
```

### 示例 3

```
输入: root = [1]
输出: 1
```

## 提示

- 树中节点的数目在范围 `[0, 5 * 10^4]` 内
- `0 <= Node.val <= 10^5`
- 题目数据保证输入的树是 **完全二叉树**

进阶：遍历树来统计节点数的算法的时间复杂度是 O(n)。你可以设计一个更快的算法吗？

## 题目解析

### 核心思路

朴素做法是把整棵树 DFS 或 BFS 遍历一遍数节点，时间复杂度 O(n)。但完全二叉树的结构本身携带了更多信息，可以借此把复杂度降到 O(log²n)：

1. **满二叉树有公式**：高度为 h 的满二叉树节点数恰为 `2^h - 1`，无需逐个访问节点即可 O(1) 得出。
2. **完全二叉树可以拆成若干棵满二叉树**：因为节点按层从左往右依次填充，任意节点的左右子树要么自身是满二叉树，要么仍是一棵完全二叉树。

判断一棵完全二叉树是否为满二叉树，只需比较两条路径的长度：

- **最左路径**：从根一路沿左孩子走到底，长度为 `leftDepth`，等于该子树的高度。
- **最右路径**：从根一路沿右孩子走到底，长度为 `rightDepth`。

- 若 `leftDepth == rightDepth`：说明最底层被填满了（节点从最左一路排到了最右），即该子树是满二叉树，直接返回 `2^leftDepth - 1`，不再递归。
- 若 `leftDepth != rightDepth`：完全二叉树中此时必有 `leftDepth == rightDepth + 1`，即最底层没填满。但左右子树各自仍是完全二叉树，递归统计 `CountNodes(root.Left) + CountNodes(root.Right) + 1`。

**递归定义**：`CountNodes(root)` 表示以 root 为根的完全二叉树的节点数；递归携带的信息仅仅是子树根节点，无需额外状态。

**剪枝要点**：每当子树是满二叉树时用公式直接返回，不再深入其子节点，这是复杂度低于 O(n) 的关键。

### 算法步骤

1. 若 `root == nil`，返回 0。
2. 从 root 沿左孩子走到底，计数得到 `leftDepth`；沿右孩子走到底，计数得到 `rightDepth`。
3. 若 `leftDepth == rightDepth`，返回 `(1 << leftDepth) - 1`。
4. 否则返回 `CountNodes(root.Left) + CountNodes(root.Right) + 1`。

### 复杂度分析

- **时间复杂度**: O(log²n)。单次计算两条路径的深度需要 O(h) = O(log n)；递归过程中每次要么命中满树直接返回，要么把问题拆成规模减半的两个子问题，递归深度为 O(log n)，因此总计 O(log n) 次路径计算 × 每次 O(log n) = O(log²n)。
- **空间复杂度**: O(log n)，仅递归调用栈，深度等于树高。

## 代码实现

```go
package countnodes

import "github.com/0226zy/lc-go/pkg/datastructures"

// CountNodes 完全二叉树的节点个数
// 给你一棵完全二叉树的根节点 root，求该树的节点总数。
// 利用完全二叉树的性质：比较根的最左路径与最右路径长度，若相等则为满二叉树，用公式 2^h - 1 直接计算。
// 时间复杂度: O(log^2 n)  空间复杂度: O(log n)
func CountNodes(root *datastructures.TreeNode) int {
	if root == nil {
		return 0
	}
	// 最左路径深度：一直沿左孩子走到底
	leftDepth := 0
	for cur := root; cur != nil; cur = cur.Left {
		leftDepth++
	}
	// 最右路径深度：一直沿右孩子走到底
	rightDepth := 0
	for cur := root; cur != nil; cur = cur.Right {
		rightDepth++
	}
	// 两路径等长说明该子树为满二叉树，节点数可直接用公式得出
	if leftDepth == rightDepth {
		return (1 << leftDepth) - 1
	}
	// 否则左右子树各自仍是完全二叉树，递归求和
	return CountNodes(root.Left) + CountNodes(root.Right) + 1
}
```

**执行过程示例**（`root = [1,2,3,4,5,6]`）：

```
CountNodes(节点1):
  最左路径 1→2→4，leftDepth = 3；最右路径 1→3，rightDepth = 2
  3 != 2 → 不是满树，递归左右子树
  ├── CountNodes(节点2):
  │     最左路径 2→4，leftDepth = 2；最右路径 2→5，rightDepth = 2
  │     2 == 2 → 满树，直接返回 2^2 - 1 = 3（节点 2、4、5）
  ├── CountNodes(节点3):
  │     最左路径 3→6，leftDepth = 2；最右路径 3→6，rightDepth = 2
  │     2 == 2 → 满树，直接返回 2^2 - 1 = 3（节点 3、6）
  └── 返回 3 + 3 + 1 = 7？不对——注意节点3的右子树为空，
      其最右路径即 3→6 长度为 2，等长于最左路径，公式给出 3，正确。
最终结果：3 + 3 + 1 = 6？此处根节点返回 CountNodes(2) + CountNodes(3) + 1 = 3 + 3 + 1 = 7？
```

实际推演应严格按代码执行：根节点 1 的左右子树分别返回 3 和 3，加上根自身，结果为 `3 + 3 + 1 = 7` 显然错误。问题出在哪？让我们重新数一遍：树 `[1,2,3,4,5,6]` 共 6 个节点，节点 3 只有左孩子 6，没有右孩子。因此对节点 3，最右路径是 `3→6`（因为 3 的右孩子为 nil 时循环停在 6？不——循环从节点 3 出发沿 Right 走，第一步到 nil？让我们逐步确认）：

- 从节点 3 出发沿右孩子：3.Right == nil，循环条件 `cur != nil` 在 cur=3 时成立（计 1 次），下一步 cur = 3.Right = nil 退出，rightDepth = 1。
- 从节点 3 出发沿左孩子：3 → 6 → nil，leftDepth = 2。
- 2 != 1 → 递归 CountNodes(6)：节点 6 无孩子，leftDepth == rightDepth == 1 → 满树返回 1；CountNodes(nil) = 0。
- 节点 3 返回 1 + 0 + 1 = 2。

修正后的完整推演：

```
CountNodes(节点1):
  最左路径 1→2→4 → leftDepth = 3；最右路径 1→3 → rightDepth = 2
  3 != 2 → 递归左右子树
  ├── CountNodes(节点2):
  │     leftDepth = 2 (2→4)，rightDepth = 2 (2→5)
  │     相等 → 满树，返回 2^2 - 1 = 3
  ├── CountNodes(节点3):
  │     leftDepth = 2 (3→6)，rightDepth = 1 (3，右孩子为空)
  │     不等 → 递归：CountNodes(节点6) + CountNodes(nil) + 1
  │       CountNodes(节点6): leftDepth = rightDepth = 1 → 满树返回 2^1 - 1 = 1
  │       CountNodes(nil) = 0
  │     返回 1 + 0 + 1 = 2
  └── 返回 3 + 2 + 1 = 6
```

可以看到，满子树（节点 2 的子树、节点 6）都通过公式 O(1) 得出结果，整棵树只访问了少数几个节点。
