# 207. 课程表 (Course Schedule)

## 题目描述

你这个学期必须选修 `numCourses` 门课程，记为 `0` 到 `numCourses - 1`。

在选修某些课程之前需要一些先修课程。先修课程按数组 `prerequisites` 给出，其中 `prerequisites[i] = [ai, bi]`，表示如果要学习课程 `ai` 则 **必须** 先学习课程 `bi`。

- 例如，先修课程对 `[0, 1]` 表示：想要学习课程 `0`，你需要先完成课程 `1`。

请你判断是否可能完成所有课程的学习？如果可以，返回 `true`；否则，返回 `false`。

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
解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
```

## 提示

- `1 <= numCourses <= 2000`
- `0 <= prerequisites.length <= 5000`
- `prerequisites[i].length == 2`
- `0 <= ai, bi < numCourses`
- `prerequisites[i]` 中的所有课程对 **互不相同**

## 题目解析

### 核心思路

把每门课程看作图中的一个节点，`[ai, bi]` 看作一条有向边 `bi -> ai`（修完 bi 才能修 ai）。
那么“能否完成所有课程”就等价于一个经典图论问题：**这个有向图中是否存在环**？

- 无环：图是一个 DAG（有向无环图），一定存在拓扑序，可以按拓扑序修完所有课程。
- 有环：环上的课程互相依赖、互为先修，永远都无法开始，返回 `false`。

本题采用 **DFS 三色标记法** 判环。DFS 判环的模板要点：

1. **状态定义（三色标记）**：每个节点有三种状态——
   - `0` 白色（未访问）：还没碰过；
   - `1` 灰色（访问中）：在**当前递归栈**上，它的后继还没搜完；
   - `2` 黑色（已完成）：以它为起点的整棵搜索子树已确认无环。
2. **判环条件**：DFS 过程中一旦遇到一个**灰色**节点，说明从某个祖先出发又回到了栈上的节点，即存在环。
3. **标记与还原策略**：进入节点时置灰（`1`），搜完所有后继后**置黑（`2`）而不是还原成白色**。
   黑色节点表示“从这里出发一定无环”，之后再遇到它直接剪枝返回——这正是三色标记相比“访问数组 + 手动回溯还原”的优势：既有判环能力，又自带记忆化，每个节点只完整搜索一次，复杂度是 O(V+E) 而不是指数级。
4. **不连通图**：先修图可能分成多个连通分量，因此外层要对每个仍是白色的节点都启动一次 DFS。

### 算法步骤

1. 用邻接表建图：遍历 `prerequisites`，对每条 `[ai, bi]` 添加边 `bi -> ai`。
2. 初始化状态数组 `state`，全部为 `0`（白色）。
3. 定义递归函数 `hasCycle(course)`：
   - 若 `state[course] == 1`（灰色）：遇到环，返回 `true`；
   - 若 `state[course] == 2`（黑色）：已确认无环，返回 `false`；
   - 否则把 `course` 置灰，依次递归它的每个后继，任一后继返回 `true` 则向上传递 `true`；
   - 所有后继搜完没有发现环，把 `course` 置黑，返回 `false`。
4. 外层循环遍历 `0` 到 `numCourses-1`，对每个白色节点调用 `hasCycle`；任一返回 `true` 则整体返回 `false`。
5. 全部搜完都没有环，返回 `true`。

### 复杂度分析

- **时间复杂度**: O(V + E)。V 为课程数（`numCourses`），E 为先修关系数（`prerequisites` 的长度）。每个节点最多经历“白→灰→黑”一次完整搜索，每条边最多被遍历一次。
- **空间复杂度**: O(V + E)。邻接表占用 O(V + E)，状态数组 O(V)，最坏情况下递归栈深度为 O(V)（链式依赖）。

## 代码实现

```go
package courseschedule

// CanFinish 课程表
// 判断先修关系构成的有向图中是否存在环，无环则可以完成所有课程。
// 时间复杂度: O(V+E)  空间复杂度: O(V+E)
func CanFinish(numCourses int, prerequisites [][]int) bool {
	// 建图：bi -> ai，即“修完 bi 后可以修 ai”
	graph := make([][]int, numCourses)
	for _, pre := range prerequisites {
		graph[pre[1]] = append(graph[pre[1]], pre[0])
	}

	// 三色标记：0=未访问（白色），1=访问中（灰色，在当前递归栈上），2=已完成（黑色）
	state := make([]int, numCourses)

	var hasCycle func(course int) bool
	hasCycle = func(course int) bool {
		if state[course] == 1 {
			// 遇到了当前递归栈上的节点，说明存在环
			return true
		}
		if state[course] == 2 {
			// 该节点的整条搜索子树已确认无环，直接剪枝
			return false
		}
		state[course] = 1 // 进入递归前标记为访问中
		for _, next := range graph[course] {
			if hasCycle(next) {
				return true
			}
		}
		state[course] = 2 // 所有后继检查完毕，标记为已完成（不回退到 0）
		return false
	}

	// 图可能不连通，对每个未访问的节点都要启动一次 DFS
	for i := 0; i < numCourses; i++ {
		if state[i] == 0 && hasCycle(i) {
			return false
		}
	}
	return true
}
```

**执行过程示例**（`numCourses = 4, prerequisites = [[1,0],[2,1],[0,2],[3,0]]`，即边 0→1、1→2、2→0、0→3）：

```
从节点 0 启动 DFS：
  state[0]=灰
    后继 1：state[1]=灰
      后继 2：state[2]=灰
        后继 0：state[0] 是灰色 → 在递归栈上！发现环，返回 true
      向上返回 true
    向上返回 true
hasCycle(0) = true → 整体返回 false（课程 0、1、2 互相成环，无法完成）
```

再对比无环的情况（`numCourses = 2, prerequisites = [[1,0]]`，即边 0→1）：

```
从节点 0 启动 DFS：
  state[0]=灰
    后继 1：state[1]=灰，1 没有后继 → state[1]=黑，返回 false
  state[0]=黑，返回 false
外层继续检查节点 1：已是黑色，跳过
没有发现环 → 返回 true
```
