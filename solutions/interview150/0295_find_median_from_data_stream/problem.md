# 0295. 数据流的中位数 (Find Median from Data Stream)

## 题目描述

中位数是有序整数列表中的中间值。如果列表大小为偶数，则没有中间值，此时中位数是两个中间数的平均值。

实现 `MedianFinder` 类：

- `MedianFinder()` 初始化 `MedianFinder` 对象。
- `void addNum(int num)` 将数据流中的整数 `num` 添加到数据结构中。
- `double findMedian()` 返回到目前为止所有元素的中位数。与实际答案误差在 `10^-5` 以内的答案将被接受。

### 示例

```
输入：
["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
[[], [1], [2], [], [3], []]
输出：
[null, null, null, 1.5, null, 2.0]

解释：
MedianFinder medianFinder = new MedianFinder();
medianFinder.addNum(1);    // arr = [1]
medianFinder.addNum(2);    // arr = [1, 2]
medianFinder.findMedian(); // 返回 1.5 ((1 + 2) / 2)
medianFinder.addNum(3);    // arr = [1, 2, 3]
medianFinder.findMedian(); // 返回 2.0
```

### 提示

- `-10^5 <= num <= 10^5`
- 在调用 `findMedian` 之前，数据结构中至少有一个元素
- 最多调用 `5 * 10^4` 次 `addNum` 和 `findMedian`

## 核心思路

动态求中位数，关键是把数据分成**较小一半**和**较大一半**，并快速拿到两半的边界：

| 堆 | 角色 | 堆顶含义 |
| --- | --- | --- |
| 大顶堆 `low` | 较小一半 | 较小半里的**最大值** |
| 小顶堆 `high` | 较大一半 | 较大半里的**最小值** |

**不变量**：

1. `low` 中所有元素 ≤ `high` 中所有元素（由堆顶关系保证）；
2. `len(low) == len(high)` 或 `len(low) == len(high) + 1`（奇数时多出的那个在 `low`）。

于是：

- 奇数个元素：中位数 = `low` 堆顶；
- 偶数个元素：中位数 = `(low 堆顶 + high 堆顶) / 2`。

## 算法步骤

**AddNum(num)**：

1. 若 `low` 为空或 `num <= low.top`，压入 `low`；否则压入 `high`。
2. 若 `len(low) > len(high) + 1`：把 `low` 堆顶弹到 `high`。
3. 若 `len(high) > len(low)`：把 `high` 堆顶弹到 `low`。

**FindMedian()**：

1. 若 `len(low) > len(high)`，返回 `float64(low.top)`；
2. 否则返回 `(low.top + high.top) / 2.0`。

Go 中用 `container/heap` 实现：`Less` 对调即可得到大顶堆。

## 复杂度分析

- **AddNum**：O(log n)，每次最多两次堆操作。
- **FindMedian**：O(1)，只读堆顶。
- **空间**：O(n)，两个堆合计存全部元素。
