# 103. 二叉树的锯齿形层序遍历 (Binary Tree Zigzag Level Order Traversal)

## 题目描述

给你二叉树的根节点 `root`，返回其节点值的**锯齿形层序遍历**。即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行。

### 示例 1

```
输入: root = [3,9,20,null,null,15,7]
输出: [[3],[20,9],[15,7]]
```

```
        3
       / \
      9   20
         /  \
        15   7
```

- 第 1 层从左往右：`[3]`
- 第 2 层从右往左：`[20, 9]`
- 第 3 层从左往右：`[15, 7]`

### 示例 2

```
输入: root = [1]
输出: [[1]]
```

### 示例 3

```
输入: root = []
输出: []
```

## 提示

- 树中节点数目在范围 `[0, 2000]` 内
- `-1000 <= Node.val <= 1000`

## 题目解析

### 核心思路

本题是普通层序遍历（102 题）的变体：BFS 按层访问的框架完全保留，唯一的差别是**每一层收集结果的方向交替变化**。

关键点有两处：

1. **如何表达「交替方向」**：用一个布尔标记 `leftToRight` 记录当前层的输出方向，每处理完一层取反一次。它只影响**结果的写入顺序**，不影响节点的访问顺序。
2. **孩子入队顺序保持不变**：无论当前层是正向还是反向输出，左右孩子始终按「先左后右」入队。这样队列里下一层节点的相对次序永远是「从左到右」，方向标记只需要在写结果时倒序填下标即可，无需真的去反转队列。

具体到写入方式，预先知道本层有 `size` 个节点，可以直接开一个长度恰为 `size` 的切片，按下标镜像填充：

- 正向层：第 `i` 个出队的节点写到 `level[i]`；
- 反向层：第 `i` 个出队的节点写到 `level[size-1-i]`。

这样省掉了「先正序收集、再整体反转」的一次额外遍历。

另一种等价做法是维护一个双端队列（deque），反向层时从队尾取节点、孩子从队头插入，但实现更繁琐，本题数据规模下没有必要。

### 算法步骤

1. 若 `root == nil`，返回空结果 `nil`。
2. 队列初始化为 `[root]`，方向标记 `leftToRight = true`。
3. 循环直到队列为空：
   - 取当前层节点数 `size = len(queue)`，创建长度 `size` 的切片 `level`；
   - 连续出队 `size` 个节点，第 `i` 个出队的节点按方向写入 `level[i]` 或 `level[size-1-i]`；
   - 每个出队节点的左右孩子（若存在）按「先左后右」入队；
   - `level` 追加进结果集，`leftToRight` 取反。
4. 返回结果集。

### 复杂度分析

- **时间复杂度**: O(n)，n 为节点总数，每个节点恰好入队、出队各一次，下标填充也是 O(1)。
- **空间复杂度**: O(n)，队列在最坏情况下（满二叉树的最后一层）同时存放约 n/2 个节点。

## 代码实现

```go
package zigzaglevelorder

import "github.com/0226zy/lc-go/pkg/datastructures"

// ZigzagLevelOrder 二叉树的锯齿形层序遍历
// 逐层返回节点值，第 1 层从左到右，第 2 层从右到左，依次交替。
// 时间复杂度: O(n) 每个节点恰好入队、出队各一次  空间复杂度: O(n) 队列最多同时存放一层的节点
func ZigzagLevelOrder(root *datastructures.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	queue := []*datastructures.TreeNode{root}
	leftToRight := true // 当前层的输出方向：true 从左到右，false 从右到左
	for len(queue) > 0 {
		size := len(queue) // 当前层节点个数
		level := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			// 方向只影响写入下标：正向正序填，反向从尾部倒序填
			if leftToRight {
				level[i] = node.Val
			} else {
				level[size-1-i] = node.Val
			}
			// 孩子入队顺序固定为「先左后右」，与输出方向无关
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
		leftToRight = !leftToRight // 下一层方向取反
	}
	return result
}
```

## 执行过程示例

以 `root = [1,2,3,4,5,null,6]` 为例：

```
        1
       / \
      2   3
     / \   \
    4   5   6
```

```
初始: queue=[1], leftToRight=true

第1层 (正向): size=1
  出队1 -> 写入 level[0]            level=[1]
  孩子2,3入队                        queue=[2,3]
  leftToRight -> false

第2层 (反向): size=2
  出队2 -> 写入 level[2-1-0]=[1]    level=[_,2]
  出队3 -> 写入 level[2-1-1]=[0]    level=[3,2]
  孩子4,5,6依次入队                  queue=[4,5,6]
  leftToRight -> true

第3层 (正向): size=3
  出队4 -> level[0]，出队5 -> level[1]，出队6 -> level[2]
                                     level=[4,5,6]
  无孩子入队                         queue=[]
  leftToRight -> false

队列空，循环结束。
结果: [[1],[3,2],[4,5,6]]
```

可以看到：第 2 层节点在队列中仍是 `2` 在 `3` 之前（从左到右的物理次序），只是写结果时倒序填充，于是得到 `[3,2]`；方向标记翻转后第 3 层又恢复正向。这正是「入队顺序不变、只改写入下标」的精髓所在。
