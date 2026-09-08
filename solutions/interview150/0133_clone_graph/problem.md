# 133. 克隆图 (Clone Graph)

## 题目描述

给你无向**连通**图中一个节点的引用，请你返回该图的**深拷贝**（克隆）。

图中的每个节点都包含它的值 `val`（`Node.val`）和其邻居的列表（`Node.neighbors`）。

```
class Node {
    public int val;
    public List<Node> neighbors;
}
```

### 示例 1

```
输入: adjList = [[2,4],[1,3],[2,4],[1,3]]
输出: [[2,4],[1,3],[2,4],[1,3]]
解释: 图中有 4 个节点。
节点 1 的值是 1，邻居是节点 2 和 4。
节点 2 的值是 2，邻居是节点 1 和 3。
节点 3 的值是 3，邻居是节点 2 和 4。
节点 4 的值是 4，邻居是节点 1 和 3。
```

### 示例 2

```
输入: adjList = [[]]
输出: [[]]
解释: 输入包含一个空列表，图中的节点只有 1 个，它没有邻居。
```

### 示例 3

```
输入: adjList = []
输出: []
解释: 这个图是空的，它没有任何节点。
```

## 提示

- 节点数在 `[0, 100]` 范围内
- `1 <= Node.val <= 100`，`Node.val` 对每个节点都是唯一的
- 图是简单图（不存在重复边和自环）
- 图是连通图，每个节点都可以从给定节点访问到

## 题目解析

### 核心思路

先想清楚“深拷贝”到底难在哪：难点不是复制一个节点，而是**节点之间的关系网**。

比如节点 1 和节点 2 互相是邻居。你先克隆了节点 1，正在给它填邻居时会发现需要节点 2 的克隆；而在克隆节点 2 的时候又需要节点 1 的克隆——如果没有记录，就会陷入无限递归。

解决这个死结只需要一个哈希表 `visited`：**原节点 -> 克隆节点**。

- 每次要克隆一个节点前，先查表：如果已经有克隆好的，直接拿来用（关系自然接对）；
- 如果没有，立刻创建克隆节点**并先登记到表里**，然后再递归克隆它的邻居。

“先登记、后填邻居”这个顺序是关键：即使图中存在环（示例 1 就是环），由于登记发生在递归之前，当递归绕回起点时一定能从表里拿到已经创建的克隆节点，永远不会重复创建或死循环。

这个模型就是 **图遍历 + visited 哈希表** 的经典组合，DFS 和 BFS 都可以。

### 算法步骤（DFS）

1. 特判：输入节点为 `nil` 直接返回 `nil`。
2. 创建 `visited map[*Node]*Node`。
3. 递归函数 `clone(node)`：
   - 若 `node` 已在 `visited` 中，返回对应的克隆节点；
   - 否则创建克隆节点，**先登记** `visited[node] = cloned`；
   - 遍历 `node.Neighbors`，对每个邻居递归调用 `clone`，把结果追加到 `cloned.Neighbors`；
   - 返回 `cloned`。
4. 从给定节点调用 `clone` 并返回。

### 复杂度分析

- **时间复杂度**: O(V + E)。每个节点只克隆一次，每条邻边只处理一次（无向图每条边在两个端点各出现一次）。
- **空间复杂度**: O(V)。`visited` 哈希表保存每个节点的克隆副本，递归栈最深为 V。

## 代码实现

### 主解：DFS + visited 哈希表

```go
type Node struct {
    Val       int
    Neighbors []*Node
}

func CloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    visited := make(map[*Node]*Node)
    return clone(node, visited)
}

func clone(node *Node, visited map[*Node]*Node) *Node {
    if cloned, ok := visited[node]; ok {
        return cloned
    }
    // 先创建并登记，再递归填邻居（应对环）
    cloned := &Node{Val: node.Val}
    visited[node] = cloned
    for _, neighbor := range node.Neighbors {
        cloned.Neighbors = append(cloned.Neighbors, clone(neighbor, visited))
    }
    return cloned
}
```

**执行过程示例**（`adjList = [[2,4],[1,3],[2,4],[1,3]]`，从节点 1 开始）：

```
clone(1): 创建 1'，登记 visited[1]=1'，开始处理邻居 [2,4]
  clone(2): 创建 2'，登记 visited[2]=2'，开始处理邻居 [1,3]
    clone(1): 已在表中，直接返回 1'   ← 环在这里被截断
    clone(3): 创建 3'，登记 visited[3]=3'，开始处理邻居 [2,4]
      clone(2): 已在表中，返回 2'
      clone(4): 创建 4'，登记 visited[4]=4'，处理邻居 [1,3]
        clone(1): 返回 1'
        clone(3): 返回 3'
```

### 备选：BFS

用队列代替递归：出队一个节点，把它的每个邻居的克隆节点挂上去；遇到没见过的邻居就先建副本入队。逻辑和 DFS 完全等价，只是遍历顺序不同。
