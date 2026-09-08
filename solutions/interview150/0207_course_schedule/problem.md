# 207. 课程表 (Course Schedule)

## 题目描述

你这个学期必须选修 `numCourses` 门课程，记为 `0` 到 `numCourses - 1`。

在选修某些课程之前需要一些先修课程。先修课程以数组 `prerequisites` 给出，其中 `prerequisites[i] = [ai, bi]` 表示如果要学习课程 `ai` 则**必须先**学习课程 `bi`。

- 例如，先修课程对 `[0, 1]` 表示：想要学习课程 `0`，你需要先完成课程 `1`。

请你判断是否可能完成所有课程的学习？

### 示例 1

```
输入: numCourses = 2, prerequisites = [[1,0]]
输出: true
解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。这是可能的。
```

### 示例 2

```
输入: numCourses = 2, prerequisites = [[1,0],[0,1]]
输出: false
解释: 存在循环依赖：课程 0 依赖课程 1，课程 1 也依赖课程 0，无法完成。
```

## 提示

- `1 <= numCourses <= 2000`
- `0 <= prerequisites.length <= 5000`
- `prerequisites[i].length == 2`
- `0 <= ai, bi < numCourses`
- `prerequisites[i]` 中的所有课程对互不相同

## 题目解析

### 核心思路

把课程看成**节点**，把“先修关系”看成**有向边**（`bi -> ai` 表示学完 bi 才能学 ai），整个先修关系就是一个**有向图**。

题目“能否修完所有课程”就等价于一个图论经典问题：**这个有向图有没有环？**

- 有环 = 课程之间互相依赖成死结，永远修不完（如示例 2）；
- 无环（DAG，有向无环图）= 总能排出一个合法的修课顺序，可以修完。

判断有向图是否有环，标准做法是**拓扑排序**，有两种经典实现：

**方法一：BFS（Kahn 算法）**——利用“入度”概念：

- 没有先修要求的课程入度为 0，可以直接修；
- 每修完一门课，就把以它为先修的课程入度减一；减到 0 说明这门课的所有先修都修完了，可以入队；
- 如果最终修完的课数等于总课数，说明无环。

**方法二：DFS 三色标记**：

- 白色（未访问）、灰色（正在递归栈上访问）、黑色（已彻底完成）；
- 遍历时一旦遇到**灰色**节点，说明沿当前路径绕回了正在访问的节点——有环。

两种方法都是 O(V + E)。主解采用 Kahn 算法，更直观。

### 算法步骤（Kahn 算法）

1. 建图：对每条 `[ai, bi]`，添加边 `bi -> ai`，同时 `indegree[ai]++`。
2. 把所有入度为 0 的课程入队。
3. 循环出队一门课：`finished++`，遍历它的所有后继课程，入度减一；减到 0 则入队。
4. 循环结束后，若 `finished == numCourses` 则无环（返回 true），否则有环（返回 false）。

### 复杂度分析

- **时间复杂度**: O(V + E)。建图 O(E)，每个节点出入队一次、每条边只被检查一次。
- **空间复杂度**: O(V + E)。邻接表和入度数组。

## 代码实现

### 主解：BFS 拓扑排序（Kahn 算法）

```go
func CanFinish(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)
    indegree := make([]int, numCourses)
    for _, pre := range prerequisites {
        ai, bi := pre[0], pre[1]
        graph[bi] = append(graph[bi], ai)
        indegree[ai]++
    }

    // 所有没有先修要求的课程入队
    queue := make([]int, 0, numCourses)
    for i := 0; i < numCourses; i++ {
        if indegree[i] == 0 {
            queue = append(queue, i)
        }
    }

    finished := 0
    for len(queue) > 0 {
        course := queue[0]
        queue = queue[1:]
        finished++
        for _, next := range graph[course] {
            indegree[next]--
            if indegree[next] == 0 {
                queue = append(queue, next)
            }
        }
    }
    return finished == numCourses
}
```

**执行过程示例**（`numCourses = 4, prerequisites = [[1,0],[2,1],[3,2]]`）：

```
建图: 0 → 1 → 2 → 3，入度: [0,1,1,1]
初始队列: [0]（只有课程 0 入度为 0）
出队 0: finished=1，课程 1 入度 0，队列 [1]
出队 1: finished=2，课程 2 入度 0，队列 [2]
出队 2: finished=3，课程 3 入度 0，队列 [3]
出队 3: finished=4
finished == numCourses → 无环，返回 true
```

若是 `[[1,0],[0,1]]`：入度都是 1，初始队列为空，finished=0 ≠ 2，返回 false。

### 备选：DFS 三色标记

递归时把节点标灰，回溯完标黑；一旦发现灰色节点即存在环。同样 O(V + E)。
