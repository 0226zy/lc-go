# 210. 课程表 II (Course Schedule II)

## 题目描述

你这个学期必须选修 `numCourses` 门课程，记为 `0` 到 `numCourses - 1`。

在选修某些课程之前需要一些先修课程。先修课程以数组 `prerequisites` 给出，其中 `prerequisites[i] = [ai, bi]` 表示如果要学习课程 `ai` 则**必须先**学习课程 `bi`。

- 例如，先修课程对 `[0, 1]` 表示：想要学习课程 `0`，你需要先完成课程 `1`。

返回一个合理的修课顺序，使得在选修每门课程前都完成了它的所有先修课程。如果不可能完成所有课程，返回一个**空数组**。

### 示例 1

```
输入: numCourses = 2, prerequisites = [[1,0]]
输出: [0,1]
解释: 总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，一个正确的课程顺序是 [0,1] 。
```

### 示例 2

```
输入: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
输出: [0,2,1,3]
解释: 总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。一个正确的课程顺序是 [0,1,2,3] ，另一个正确的顺序是 [0,2,1,3] 。
```

### 示例 3

```
输入: numCourses = 1, prerequisites = []
输出: [0]
```

## 提示

- `1 <= numCourses <= 2000`
- `0 <= prerequisites.length <= numCourses * (numCourses - 1)`
- `prerequisites[i].length == 2`
- `0 <= ai, bi < numCourses`
- `ai != bi`
- `prerequisites[i]` 中的所有课程对互不相同

## 题目解析

### 核心思路

和 207（课程表）完全同一套建模：课程是节点，先修关系是有向边（`bi -> ai`），整张图是**有向图**。

207 只问“能不能修完”（有没有环），本题进一步要求**给出具体的修课顺序**——这正是**拓扑排序**的定义：对一个有向无环图（DAG）的节点排序，使得每条边的起点都排在终点之前。

同样用 **Kahn 算法（BFS 拓扑排序）**：

- 入度为 0 的课（没有先修要求）随时可以修，先入队；
- 每次出队一门课，把它记入答案，并把它的所有后继课程入度减一；减到 0 表示该课的所有先修都已完成，入队；
- 出队的先后顺序就是一个合法的拓扑序。

如果队列处理完后，记录的课数少于总课数，说明图中存在环（剩下的课互相卡死），返回空数组。

### 算法步骤

1. 建图：对每条 `[ai, bi]`，添加边 `bi -> ai`，同时 `indegree[ai]++`。
2. 所有入度为 0 的课程入队。
3. BFS：出队课程 `c`，把 `c` 追加到 `order`；遍历 `c` 的后继，入度减一，减到 0 入队。
4. 若 `len(order) == numCourses`，返回 `order`；否则返回空数组。

### 复杂度分析

- **时间复杂度**: O(V + E)。每个节点出入队一次，每条边只被检查一次。
- **空间复杂度**: O(V + E)。邻接表、入度数组与结果数组。

## 代码实现

```go
func FindOrder(numCourses int, prerequisites [][]int) []int {
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

    // 出队的先后顺序就是拓扑序
    order := make([]int, 0, numCourses)
    for len(queue) > 0 {
        course := queue[0]
        queue = queue[1:]
        order = append(order, course)
        for _, next := range graph[course] {
            indegree[next]--
            if indegree[next] == 0 {
                queue = append(queue, next)
            }
        }
    }

    // 修不完所有课程说明有环，返回空数组
    if len(order) != numCourses {
        return []int{}
    }
    return order
}
```

**执行过程示例**（示例 2：`numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]`）：

```
建图: 0 → 1, 0 → 2, 1 → 3, 2 → 3，入度: [0,1,1,2]
初始队列: [0]（只有课程 0 入度为 0）
出队 0: order=[0]，课程 1、2 入度减为 0，队列 [1,2]
出队 1: order=[0,1]，课程 3 入度 2→1，队列 [2]
出队 2: order=[0,1,2]，课程 3 入度 1→0，队列 [3]
出队 3: order=[0,1,2,3]
len(order)==4 → 返回 [0,1,2,3]
```

注意：合法的拓扑序往往不止一种（示例 2 的 `[0,2,1,3]` 也正确），因此测试时不能逐个比对答案，而要校验“顺序是否合法”：每门课都恰好出现一次，且每门课都排在它所有先修课之后。
