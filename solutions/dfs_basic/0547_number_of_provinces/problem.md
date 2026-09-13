# 547. 省份数量 (Number of Provinces)

## 题目描述

有 `n` 个城市，其中一些彼此相连，另一些没有相连。如果城市 `a` 与城市 `b` 直接相连，且城市 `b` 与城市 `c` 直接相连，那么城市 `a` 与城市 `c` **间接相连**。

**省份** 是一组直接或间接相连的城市，组内不含其他没有相连的城市。

给你一个 `n x n` 的矩阵 `isConnected`，其中 `isConnected[i][j] = 1` 表示第 `i` 个城市和第 `j` 个城市直接相连，而 `isConnected[i][j] = 0` 表示二者不直接相连。

返回矩阵中 **省份** 的数量。

### 示例 1

```
输入: isConnected = [[1,1,0],[1,1,0],[0,0,1]]
输出: 2
```

### 示例 2

```
输入: isConnected = [[1,0,0],[0,1,0],[0,0,1]]
输出: 3
```

## 提示

- `1 <= n <= 200`
- `n == isConnected.length`
- `n == isConnected[i].length`
- `isConnected[i][j]` 为 `1` 或 `0`
- `isConnected[i][i] == 1`
- `isConnected[i][j] == isConnected[j][i]`

## 题目解析

### 核心思路

把城市看成图的节点，`isConnected` 就是图的**邻接矩阵**。题目要求的“省份数量”其实就是这张无向图的**连通分量个数**——和 200 题「岛屿数量」是同一种模型，只是邻接关系从“上下左右四个方向”换成了“矩阵中值为 1 的位置”。

关键观察：**从一个未访问的城市出发做一次 DFS，恰好能走遍它所在的整个省份**。因为 DFS 会沿着所有 `isConnected[city][next] == 1` 的边不断扩展，直到这个连通分量里的所有城市都被标记。

DFS 模板要点：

- **标记策略**：使用一个 `visited` 数组，进入递归时立刻把当前城市标记为已访问（“先标记，后扩展”），避免通过环或互连边重复进入；
- **扩展方式**：扫描邻接矩阵的整行 `isConnected[city][:]`，对所有值为 1 且未访问的城市递归；
- **计数时机**：外层循环每遇到一个未访问的城市，说明发现了一个新省份，计数加一，然后从这个城市发起 DFS 把全省标记掉。

另一种思路是 **并查集**：把每条 `isConnected[i][j] = 1` 的两个城市合并到同一集合，最终集合个数就是省份数量。两种方法复杂度同级，DFS 代码更短、更直观。

### 算法步骤（DFS 主解）

1. 初始化 `visited` 数组，全部为 false；省份计数 `count = 0`。
2. 依次遍历每个城市 `city`：
   - 若 `city` 未被访问过：`count++`，并从 `city` 开始 DFS。
   - DFS 过程：标记 `visited[city] = true`，然后扫描第 `city` 行，对所有 `isConnected[city][next] == 1 && !visited[next]` 的城市 `next` 递归调用 `dfs(next)`。
3. 遍历结束后，`count` 即为省份数量。

### 复杂度分析

- **时间复杂度**: O(n²)。外层循环 n 次，每个城市被访问一次，每次 DFS 需要扫描整行邻接矩阵（长度为 n），总计 O(n²)。
- **空间复杂度**: O(n)。`visited` 数组占 O(n)，递归栈深度最坏为 n（所有城市串成一条链）。

## 代码实现

### 主解：DFS 数连通分量

```go
func FindCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    visited := make([]bool, n)
    count := 0

    var dfs func(city int)
    dfs = func(city int) {
        visited[city] = true // 先标记，防止重复访问
        for next := 0; next < n; next++ {
            if isConnected[city][next] == 1 && !visited[next] {
                dfs(next)
            }
        }
    }

    for city := 0; city < n; city++ {
        if !visited[city] {
            count++   // 发现新省份
            dfs(city) // 把全省标记掉
        }
    }
    return count
}
```

**执行过程示例**（示例 1：`isConnected = [[1,1,0],[1,1,0],[0,0,1]]`）：

```
扫描到城市 0：未访问 → count=1，发起 dfs(0)
    dfs(0)：标记 0；isConnected[0][1]=1 且 1 未访问 → dfs(1)
        dfs(1)：标记 1；城市 0 已访问，城市 2 不相连 → 返回
    城市 2 与 0 不相连 → 返回
扫描到城市 1：已访问，跳过
扫描到城市 2：未访问 → count=2，发起 dfs(2)
    dfs(2)：标记 2；城市 0、1 都不相连 → 返回
最终 count = 2
```

### 备选：并查集

每个城市是一个元素，遍历邻接矩阵上三角，把 `isConnected[i][j] = 1` 的城市对合并；合并结束后并查集内部的连通分量数 `Count` 就是省份数量。路径压缩 + 按秩合并让每次操作接近 O(1)，总复杂度 O(n²·α(n))。
