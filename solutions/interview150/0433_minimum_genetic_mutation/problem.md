# 433. 最小基因变化 (Minimum Genetic Mutation)

## 题目描述

基因序列由 8 个字符组成，每个字符都是 `A`、`C`、`G`、`T` 之一。

基因变化是指从原始基因序列中**一次改变一个字符**产生一个新基因序列。

给你两个基因序列 `startGene` 和 `endGene`，以及一个基因库 `bank`（存储了所有合法的有效基因变化）。只有当基因变化后的结果在 `bank` 中，这次变化才合法。

返回从 `startGene` 到 `endGene` 所需的**最少变化次数**。如果无法到达，返回 `-1`。

注意：起始基因序列 `startGene` 不一定是 `bank` 中的基因。

### 示例 1

```
输入: startGene = "AACCGGTT", endGene = "AACCGGTA", bank = ["AACCGGTA"]
输出: 1
解释: 只需改变最后一个字符 T -> A。
```

### 示例 2

```
输入: startGene = "AACCGGTT", endGene = "AAACGGTA", bank = ["AACCGGTA","AACCGCTA","AAACGGTA"]
输出: 2
解释: AACCGGTT -> AACCGGTA -> AAACGGTA，共两次变化（每次只改一个字符，且中间结果都在基因库中）。
```

## 提示

- `0 <= bank.length <= 10`
- `startGene.length == endGene.length == 8`
- `startGene`、`endGene` 和 `bank[i]` 仅由字符 `A`、`C`、`G`、`T` 组成

## 题目解析

### 核心思路

把每个基因序列看成图中的一个点，"一次合法变化"就是一条边（只能差一个字符且结果在基因库中），题目就变成了**无权图最短路**——直接套 BFS 模板。

与「单词接龙」是同一类模型，区别是：基因只有固定的 4 种字符，所以"邻居"不需要预建图，而是**逐位尝试替换为其余 3 个字符**，看是否在基因库中即可。基因库很小（≤10），用哈希集合判存在性即可。

为什么 BFS 是对的：BFS 按变化次数一层层扩展，第一次碰到 `endGene` 时，对应的层数就是最少变化次数。

### 算法步骤

1. 特判：`startGene == endGene` 时返回 0（不需要任何变化）。
2. 把 `bank` 放入哈希集合 `bankSet`；若 `endGene` 不在其中，直接返回 `-1`（永远不可能合法地变成它）。
3. 队列初始放入 `startGene`，标记已访问，`steps = 0`。
4. BFS 逐层处理，取出队首基因 `cur`：
   - 若 `cur == endGene`，返回 `steps`；
   - 否则对 8 个位置中的每一位，尝试替换成另外 3 种字符，若新基因在 `bankSet` 且未访问，则标记并入队。
5. 一层处理完 `steps++`；队列空了还没碰到终点，返回 `-1`。

### 复杂度分析

- **时间复杂度**: O(n × 8 × 3)，n 为基因库大小（≤10），每个基因最多出队一次，每次出队尝试 8 个位置 × 3 种替换
- **空间复杂度**: O(n)，visited 集合与队列

## 代码实现

```go
func MinMutation(startGene string, endGene string, bank []string) int {
	bankSet := make(map[string]bool, len(bank))
	for _, g := range bank {
		bankSet[g] = true
	}
	if !bankSet[endGene] {
		return -1
	}

	const genes = "ACGT"
	visited := map[string]bool{startGene: true}
	queue := []string{startGene}
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == endGene {
				return steps
			}
			// 逐位尝试替换为其他 3 种字符
			for j := 0; j < len(cur); j++ {
				for _, ch := range genes {
					if byte(ch) == cur[j] {
						continue
					}
					next := cur[:j] + string(ch) + cur[j+1:]
					if bankSet[next] && !visited[next] {
						visited[next] = true
						queue = append(queue, next)
					}
				}
			}
		}
		steps++
	}
	return -1
}
```

**执行过程示例**（示例 2）：

```
初始: queue=["AACCGGTT"], steps=0
第 0 层: cur="AACCGGTT"，逐位替换后在 bank 中的邻居有 "AACCGGTA"、"AACCGCTA"，入队
第 1 层: cur="AACCGGTA"，其合法邻居中 "AAACGGTA" 是终点，入队
第 2 层: cur="AAACGGTA" == endGene，返回 2
```
