# 637. 二叉树的层平均值 (Average of Levels in Binary Tree)

## 题目描述

给定一个非空二叉树的根节点 `root`，以数组的形式返回每一层节点值的 **平均值**。（与实际答案误差不超过 `10^-5` 的答案均可接受。）

### 示例 1

```
输入: root = [3,9,20,15,7]
输出: [3.00000,14.50000,11.00000]
解释: 第 0 层的平均值为 3，第 1 层的平均值为 (9 + 20) / 2 = 14.5，第 2 层的平均值为 (15 + 7) / 2 = 11。
```

### 示例 2

```
输入: root = [3,9,20,null,null,15,7]
输出: [3.00000,14.50000,11.00000]
```

## 提示

- 树中节点数量在 `[1, 10^4]` 范围内
- `-2^31 <= Node.val <= 2^31 - 1`

## 题目解析

### 核心思路

要求「每一层」的统计量，最自然的做法就是 **BFS 层序遍历**：用队列保存待访问节点，每次循环开始时队列里恰好是完整的一层，把这层的节点一次性出队、累加求和，再除以节点数即得该层平均值。

遍历顺序与状态携带方式：

- **按层批量出队**：每轮循环先记录 `n = len(queue)`，然后连续出队 `n` 次。出队的 `n` 个节点就是当前层，入队的左右孩子则构成下一层——层与层之间严格隔离，不会混在一起。
- **层内只携带一个累加器 `levelSum`**：每出队一个节点就把它的值加入 `levelSum`，整层处理完计算 `levelSum / n` 追加到结果。

两个实现要点：

1. **层数统计要在出队前固定**：`n` 必须在循环一开始取出，否则边出队边入队会让 `len(queue)` 不断变化，无法确定本层边界。
2. **浮点精度**：题目允许 `Node.val` 达到 `2^31 - 1`，但单层最多 `10^4` 个节点，`float64`（53 位尾数）累加单层和的误差远小于题目允许的 `10^-5`，直接用 `float64` 求和即可，无需担心溢出或精度问题。

也可以用 DFS（递归时携带深度 `depth`，把节点值累加进 `sums[depth]`、计数加进 `counts[depth]`，最后逐层相除），同样 O(n)，但 BFS 一趟出结果，更贴合题意。

### 算法步骤

1. 若 `root` 为空，直接返回 `nil`。
2. 初始化队列 `queue = [root]`，结果切片 `answer`。
3. 当队列非空时循环：
   - 记 `n = len(queue)`（当前层节点数），`levelSum = 0`；
   - 连续出队 `n` 次：累加节点值到 `levelSum`，把非空左右孩子入队；
   - 把 `levelSum / n` 追加到 `answer`。
4. 返回 `answer`。

### 复杂度分析

- **时间复杂度**: O(n)，n 为节点数，每个节点恰好入队、出队各一次
- **空间复杂度**: O(n)，队列中最多同时存放一层的节点（满二叉树底层约为 n/2）

## 代码实现

```go
func AverageOfLevels(root *datastructures.TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var answer []float64
	queue := []*datastructures.TreeNode{root}
	for len(queue) > 0 {
		n := len(queue) // 当前层节点数，出队前先固定下来
		levelSum := 0.0
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			levelSum += float64(node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		answer = append(answer, levelSum/float64(n))
	}
	return answer
}
```

**执行过程示例**（`root = [3,9,20,null,null,15,7]`，即 3 的左孩子为 9、右孩子为 20，20 的左右孩子为 15 和 7）：

```
初始: queue=[3], answer=[]
第1轮: n=1, 出队 3, levelSum=3, 入队 9、20
       -> avg=3/1=3.0, answer=[3.0], queue=[9,20]
第2轮: n=2, 出队 9（无孩子）, 出队 20（入队 15、7）, levelSum=29
       -> avg=29/2=14.5, answer=[3.0,14.5], queue=[15,7]
第3轮: n=2, 出队 15、7（均无孩子）, levelSum=22
       -> avg=22/2=11.0, answer=[3.0,14.5,11.0], queue=[]
队列空，返回 [3.0, 14.5, 11.0]
```
