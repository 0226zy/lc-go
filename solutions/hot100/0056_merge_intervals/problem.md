# 56. 合并区间 (Merge Intervals)

## 题目描述

给定一个数组 `intervals`，其中 `intervals[i] = [starti, endi]` 表示第 `i` 个区间是 `[starti, endi)`。

返回 **合并所有重叠区间后** 的区间列表。这些区间必须按顺序存放。

### 示例 1

```
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 合并为 [1,6]。
```

### 示例 2

```
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
```

## 提示

- `1 <= intervals.length <= 10^4`
- `intervals[i].length == 2`
- `0 <= starti <= endi <= 10^4`

## 题目解析

### 核心思路

这道题的核心是 **排序 + 贪心合并**：

1. **排序**：按区间起点排序，这样所有可能重叠的区间都会相邻。
2. **合并**：遍历排序后的区间，若当前区间与结果列表中最后一个区间重叠，则合并；否则加入结果列表。

### 算法步骤

1. 若区间数为 0，返回空列表
2. 按区间起点排序
3. 初始化结果列表 `merged`，加入第一个区间
4. 从第二个区间开始遍历：
   - 若当前区间起点 `<= merged 最后一个区间的终点`，则合并（更新终点为 `max(终点, 当前终点)`）
   - 否则，将当前区间加入 `merged`
5. 返回 `merged`

### 复杂度分析

- **时间复杂度**: O(n log n)，排序占主导
- **空间复杂度**: O(log n) 或 O(n)，取决于排序算法

## 代码实现

### 最优解：排序 + 贪心合并

```go
func merge(intervals [][]int) [][]int {
    // 按起点排序
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    merged := [][]int{intervals[0]}
    for _, interval := range intervals[1:] {
        last := merged[len(merged)-1]
        if interval[0] <= last[1] {
            // 重叠，合并
            if interval[1] > last[1] {
                last[1] = interval[1]
            }
        } else {
            // 不重叠，加入
            merged = append(merged, interval)
        }
    }
    return merged
}
```

**执行过程示例**（`intervals = [[1,3],[2,6],[8,10],[15,18]]`）：

```
排序后：[[1,3],[2,6],[8,10],[15,18]]
merged = [[1,3]]
i=[2,6]: 2<=3 → 合并 → merged = [[1,6]]
i=[8,10]: 8>6 → 加入 → merged = [[1,6],[8,10]]
i=[15,18]: 15>10 → 加入 → merged = [[1,6],[8,10],[15,18]]
结果：[[1,6],[8,10],[15,18]]
```

**性能优势**：
- O(n log n) 时间复杂度，排序是瓶颈
- 原地修改，减少内存分配
- 单次遍历，高效合并
