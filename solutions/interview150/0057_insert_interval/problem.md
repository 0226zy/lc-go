# 57. 插入区间 (Insert Interval)

## 题目描述

给你一个 **无重叠的**、按照区间起始端点排序的区间列表 `intervals`，其中 `intervals[i] = [start_i, end_i]` 表示第 `i` 个区间的开始和结束，并且这些区间之间是互不重叠的。同样给你一个区间 `newInterval = [start, end]`，表示另一个区间的开始和结束。

在 `intervals` 中插入区间 `newInterval`，使得 `intervals` 依然按照 `start_i` 升序排列，且区间之间不重叠（如果有必要的话，可以合并区间）。

返回插入之后的 `intervals`。

**注意**：你不需要原地修改 `intervals`。你可以创建一个新数组然后返回它。

### 示例 1

```
输入: intervals = [[1,3],[6,9]], newInterval = [2,5]
输出: [[1,5],[6,9]]
```

### 示例 2

```
输入: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出: [[1,2],[3,10],[12,16]]
解释: 新区间 [4,8] 与 [3,5]、[6,7]、[8,10] 重叠。
```

### 提示

- `0 <= intervals.length <= 10^4`
- `intervals[i].length == 2`
- `0 <= start_i <= end_i <= 10^5`
- `intervals` 根据 `start_i` 按 **升序** 排列
- `newInterval.length == 2`
- `0 <= start <= end <= 10^5`

## 题目解析

### 核心思路

输入区间列表**已经按左端点排好序且不重叠**，这是最大的利好。把插入位置想清楚后，整个过程可以分成三段：

1. **左边的区间**：右端点 `< newInterval.start` 的区间，与新区间**不可能重叠**，原样放进答案。
2. **中间的区间**：与新区间有重叠的（左端点 `<= newInterval.end` 的那些），逐个和新区间**合并**，把合并结果作为一个区间。
3. **右边的区间**：左端点 `> newInterval.end` 的区间，与新区间不可能重叠，原样放进答案。

这就是“**左中右三段划分 + 中间局部合并**”的模板。因为输入有序，每段内的区间各处理一次即可，整体 O(n)。也可以用“直接插入再整体排序合并”的 O(n log n) 暴力写法，但没有必要。

### 算法步骤

1. 初始化空答案 `res`，下标 `i = 0`，`n = len(intervals)`。
2. **左段**：当 `i < n` 且 `intervals[i][1] < newInterval[0]` 时，把 `intervals[i]` 追加进 `res`，`i++`。
3. **中段（合并）**：当 `i < n` 且 `intervals[i][0] <= newInterval[1]` 时，循环合并：
   - `newInterval[0] = min(newInterval[0], intervals[i][0])`
   - `newInterval[1] = max(newInterval[1], intervals[i][1])`
   - `i++`
   循环结束后，把合并后的 `newInterval` 追加进 `res`。
4. **右段**：把剩余的 `intervals[i:]` 全部追加进 `res`。
5. 返回 `res`。

### 复杂度分析

- **时间复杂度**: O(n)，最多遍历一遍区间列表
- **空间复杂度**: O(n)，需要返回新数组

## 代码实现

```go
// Insert 插入区间
// 在无重叠且按左端点排序的区间列表中插入新区间，保持有序且不重叠（必要时合并）。
// 时间复杂度: O(n)  空间复杂度: O(n) 输出数组
func Insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	i, n := 0, len(intervals)

	// 左段：右端点 < newInterval 左端点，不可能重叠，直接放入
	for i < n && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}
	// 中段：与 newInterval 有重叠（含端点相接），逐个合并
	for i < n && intervals[i][0] <= newInterval[1] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}
	res = append(res, newInterval) // 合并后的新区间
	// 右段：剩余区间不可能重叠，直接放入
	for i < n {
		res = append(res, intervals[i])
		i++
	}
	return res
}
```

**执行过程示例**（`intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]]`, `newInterval = [4,8]`）：

```
左段: [1,2] 右端点 2 < 4, 放入 -> res = [[1,2]]
中段: [3,5] 左端点 3 <= 8, 合并 -> newInterval = [3,8]
      [6,7] 左端点 6 <= 8, 合并 -> newInterval = [3,8]
      [8,10] 左端点 8 <= 8, 合并 -> newInterval = [3,10]
      [12,16] 左端点 12 > 8, 停止
      放入合并结果 -> res = [[1,2],[3,10]]
右段: 放入 [12,16] -> res = [[1,2],[3,10],[12,16]]
结果: [[1,2],[3,10],[12,16]]
```
