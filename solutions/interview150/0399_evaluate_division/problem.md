# 399. 除法求值 (Evaluate Division)

## 题目描述

给你一个变量对数组 `equations` 和一个实数值数组 `values` 作为已知条件，其中 `equations[i] = [Ai, Bi]` 表示 `Ai / Bi = values[i]`。每个变量对最多出现一次。

请根据已知条件求出 `queries[j] = [Cj, Dj]` 的结果，即 `Cj / Dj` 的值。如果结果不存在，返回 `-1.0`。题目保证不会有重复的方程或矛盾的方程。

### 示例 1

```
输入: equations = [["a","b"],["b","c"]], values = [2.0,3.0],
      queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
输出: [6.0,0.5,-1.0,1.0,-1.0]
解释: a/c = (a/b)*(b/c) = 2*3 = 6; b/a = 1/(a/b) = 0.5
```

### 示例 2

```
输入: equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0],
      queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
输出: [3.75,0.4,5.0,0.2]
```

### 示例 3

```
输入: equations = [["a","b"]], values = [0.5],
      queries = [["a","b"],["b","a"],["a","c"],["x","y"]]
输出: [0.5,2.0,-1.0,-1.0]
```

## 提示

- `1 <= equations.length <= 20`
- `equations[i].length == 2`
- `1 <= Ai.length, Bi.length <= 5`
- `values.length == equations.length`
- `0.0 < values[i] <= 20.0`
- `1 <= queries.length <= 20`
- `Ai、Bi、Cj、Dj` 由小写英文字母与数字组成

## 题目解析

### 核心思路

把每个变量看成一个**节点**，每个方程 `a / b = v` 看成一条**有向加权边**：

- 边 `a -> b`，权值 `v`（a 是 b 的 v 倍）；
- 边 `b -> a`，权值 `1/v`。

这样一来，查询 `c / d` 等价于：**在图中找一条从 c 到 d 的路径，把路径上所有边权乘起来**。

为什么乘起来就对了？因为路径 `c -> x -> y -> d` 表示

```
c/d = (c/x) * (x/y) * (y/d)
```

每一步的比值都恰好是某条边的权值。这正是除法的传递性——也可以理解为在图上做“带权最短路”的弱化版：不需要最短，只需要任意一条可达路径。

于是每个查询就是一次 **BFS / DFS 找路径**：

- 起点或终点不在图中 → `-1.0`；
- 起点 == 终点（且存在）→ `1.0`；
- 否则 BFS，沿途累乘边权，到达终点时的累乘积就是答案；队列耗尽还没到达 → `-1.0`。

### 算法步骤

1. **建图**：遍历 `equations`，为每对方程添加正反两条带权边。
2. **逐个处理查询**：
   - 检查 `from`、`to` 是否在图中，不在则记录 `-1.0`；
   - `from == to` 时记录 `1.0`；
   - 否则从 `from` 开始 BFS：队列里存 `(当前节点, 起点到当前节点的比值)`。每走到一条新边，比值乘上边权；到达 `to` 时停止并记录答案。
3. 返回答案数组。

### 复杂度分析

- **时间复杂度**: O((E + Q) × V)。建图 O(E)；每次查询一次 BFS 最坏访问所有节点和边，共 Q 次查询。由于本题约束很小（E、Q ≤ 20），完全够用。
- **空间复杂度**: O(V + E)。邻接表占用的空间。

## 代码实现

```go
func CalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    // 建图：边 a->b 权值为 v，边 b->a 权值为 1/v
    type edge struct{ to string; w float64 }
    graph := make(map[string][]edge)
    for i, eq := range equations {
        a, b := eq[0], eq[1]
        graph[a] = append(graph[a], edge{b, values[i]})
        graph[b] = append(graph[b], edge{a, 1.0 / values[i]})
    }

    results := make([]float64, len(queries))
    for i, q := range queries {
        from, to := q[0], q[1]
        if len(graph[from]) == 0 || len(graph[to]) == 0 {
            results[i] = -1.0 // 变量不存在
            continue
        }
        if from == to {
            results[i] = 1.0
            continue
        }
        // BFS：搜索 from 到 to 的路径，累乘边权
        visited := map[string]bool{from: true}
        type state struct{ node string; prod float64 }
        queue := []state{{from, 1.0}}
        found := -1.0
        for len(queue) > 0 && found < 0 {
            curr := queue[0]
            queue = queue[1:]
            for _, e := range graph[curr.node] {
                if visited[e.to] {
                    continue
                }
                prod := curr.prod * e.w
                if e.to == to {
                    found = prod
                    break
                }
                visited[e.to] = true
                queue = append(queue, state{e.to, prod})
            }
        }
        results[i] = found
    }
    return results
}
```

**执行过程示例**（`equations = [["a","b"],["b","c"]], values = [2,3]`，查询 `a/c`）：

```
建图: a --2--> b   b --1/2--> a
      b --3--> c   c --1/3--> b

查询 a/c:
  队列初始: [(a, 1.0)]
  出队 (a, 1.0): 邻居 b，比值 1.0*2 = 2，入队 (b, 2.0)
  出队 (b, 2.0): 邻居 a(已访问)，邻居 c，比值 2.0*3 = 6，c 即目标！
  结果: a/c = 6.0
```
