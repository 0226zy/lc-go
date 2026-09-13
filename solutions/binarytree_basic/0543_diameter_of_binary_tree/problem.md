# 543. 二叉树的直径 (Diameter of Binary Tree)

## 题目描述

给你一棵二叉树的根节点 `root`，返回该树的 **直径**。

二叉树的 **直径** 是指树中任意两个节点之间最长路径的 **长度**（以经过的边数计）。这条路径可能穿过也可能不穿过根节点。

两节点之间路径的长度由它们之间边的数目表示。

### 示例 1

```
    1
   / \
  2   3
 / \
4   5

输入: root = [1,2,3,4,5]
输出: 3
解释: 3 是路径 [4,2,1,3] 或 [5,2,1,3] 的长度（边数）。
```

### 示例 2

```
输入: root = [1,2]
输出: 1
```

## 提示

- 树中节点数目在范围 `[0, 10^4]` 内
- `-100 <= Node.val <= 100`

## 题目解析

### 核心思路

关键观察：**任意一条最长路径必然存在一个「最高点」节点**，路径从该节点的左子树一路向下到最深叶子，再经该节点折向右子树最深叶子。因此：

> 以某节点为「拐点」的直径候选 = 左子树高度 + 右子树高度（高度以边数计）。

由于题目中的直径是边数，我们定义 `height(node)` 为以 `node` 为根的子树的**边数高度**：

- 空节点高度为 `-1`（这样叶子节点高度恰为 `0`，即叶子向下没有边）；
- 非空节点高度为 `max(左高度, 右高度) + 1`。

采用**后序 DFS**：先递归拿到左右子树高度，再回到当前节点做两件事——

1. **更新答案**：`leftH + rightH + 2` 是经过当前节点的路径边数（`+2` 是当前节点到左右孩子的两条边），与全局最大值比较；
2. **向上传递信息**：返回 `max(leftH, rightH) + 1` 作为当前子树高度，供父节点使用。

注意「更新答案」与「返回值」是两份不同的信息：答案要的是 `左+右`（路径同时利用两侧），返回给父节点的只能是 `max(左,右)`（父路径只能走一侧）。这正是本题「一次递归携带两种语义」的要点。

由于每个节点的候选值都被计算过，全局最大值即为整棵树的直径——**无需假定路径穿过根节点**。

### 算法步骤

1. 初始化全局变量 `diameter = 0`。
2. 后序递归 `height(node)`：
   - 若 `node == nil`，返回 `-1`；
   - 递归求得 `leftH = height(node.Left)`、`rightH = height(node.Right)`；
   - 用 `leftH + rightH + 2` 更新 `diameter`；
   - 返回 `max(leftH, rightH) + 1`。
3. 递归结束后返回 `diameter`。

### 复杂度分析

- **时间复杂度**: O(n)，n 为节点数，每个节点恰好访问一次。
- **空间复杂度**: O(h)，h 为树高，即递归栈深度；最坏情况（链式树）为 O(n)。

## 代码实现

```go
package diameterofbinarytree

import "github.com/0226zy/lc-go/pkg/datastructures"

// DiameterOfBinaryTree 二叉树的直径
// 给定一棵二叉树，计算它的直径长度：任意两个节点之间最长路径的边数，
// 这条路径可能穿过也可能不穿过根节点。
// 时间复杂度: O(n)  n 为节点数，每个节点恰好访问一次
// 空间复杂度: O(h)  h 为树高，递归栈深度，最坏退化为链式树时 O(n)
func DiameterOfBinaryTree(root *datastructures.TreeNode) int {
	diameter := 0

	// height 返回以 node 为根的子树高度（边数），同时更新全局最大直径
	var height func(node *datastructures.TreeNode) int
	height = func(node *datastructures.TreeNode) int {
		if node == nil {
			return -1 // 空节点高度定义为 -1，使叶子节点高度为 0（叶子没有向下的边）
		}
		leftH := height(node.Left)
		rightH := height(node.Right)
		// 经过当前节点的最长路径 = 左子树高度 + 右子树高度 + 2 条边
		if leftH+rightH+2 > diameter {
			diameter = leftH + rightH + 2
		}
		// 向上返回当前子树的高度，供父节点继续累加
		return max(leftH, rightH) + 1
	}
	height(root)
	return diameter
}
```

**执行过程示例**（`root = [1,2,3,4,5]`，即示例 1）：

后序遍历顺序为 4 → 5 → 2 → 3 → 1，逐步推演（高度以边数计）：

```
height(4): 左右均为空(leftH=rightH=-1)
           候选直径 = -1 + -1 + 2 = 0，diameter = 0
           返回高度 max(-1,-1)+1 = 0
height(5): 同节点4，候选 0，返回 0
height(2): leftH=0, rightH=0
           候选直径 = 0 + 0 + 2 = 2（路径 4-2-5），diameter = 2
           返回高度 max(0,0)+1 = 1
height(3): 左右均为空，候选 0，返回 0
height(1): leftH=1, rightH=0
           候选直径 = 1 + 0 + 2 = 3（路径 4-2-1-3 或 5-2-1-3），diameter = 3
           返回高度 max(1,0)+1 = 2
最终结果: diameter = 3
```
