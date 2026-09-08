# 274. H 指数 (H-Index)

## 题目描述

给你一个整数数组 `citations`，其中 `citations[i]` 表示研究者的第 `i` 篇论文被引用的次数。计算并返回该研究者的 **h 指数**。

h 指数的定义：h 代表"高引用论文的数量"。如果一名研究者有 `h` 篇论文每篇**至少**被引用 `h` 次，并且其余 `n - h` 篇论文**每篇至多**被引用 `h` 次，那么这名研究者的 h 指数就是 `h`。

### 示例 1

```
输入: citations = [3,0,6,1,5]
输出: 3
解释: 有 3 篇论文引用次数 >= 3（3, 6, 5 分别被引用 3、6、5 次），
     另外 2 篇引用次数 <= 3（0, 1 次）。满足条件的最大 h 为 3。
```

### 示例 2

```
输入: citations = [1,3,1]
输出: 1
```

## 提示

- `1 <= citations.length <= 5000`
- `0 <= citations[i] <= 1000`

## 题目解析

### 核心思路

把定义翻译一下：h 指数是满足"**引用次数 ≥ h 的论文至少有 h 篇**"的最大的 `h`。

关键观察是：**把论文按引用次数从大到小排序后，如果第 `i` 篇（从 1 开始数）论文的引用次数 `citations[i-1]` 还 ≥ `i`，就说明"引用 ≥ i 的论文至少有 i 篇"**。从前往后扫描，找到第一个"引用次数 < 篇数"的位置，答案就是它前面的篇数。

注意 h 指数有个隐含上界：`h` 不可能超过论文总数 `n`（总共才 n 篇，不可能有 n+1 篇论文）。利用这个上界还可以得到 O(n) 的桶计数解法：引用次数 ≥ n 的论文全部归入第 `n` 号桶，然后从桶 n 往桶 0 累加论文数，首次出现"累计论文数 ≥ 桶编号"时，桶编号就是 h。

本题属于 **"排序 + 一次扫描判定"** 的经典模型（也可看作计数排序/桶计数的应用）。

### 算法步骤（排序法）

1. 将 `citations` 按从大到小排序；
2. 从左到右扫描（`i` 从 0 到 n-1）：
   - 如果 `citations[i] < i+1`，说明引用次数 ≥ `i+1` 的论文不足 `i+1` 篇，返回 `i`；
3. 全部满足则返回 `n`。

### 复杂度分析

- **排序法**：时间复杂度 O(n log n)，空间复杂度 O(log n)（排序递归栈）
- **计数桶法**：时间复杂度 O(n)，空间复杂度 O(n)

## 代码实现

```go
// 排序法：O(n log n)
func HIndex(citations []int) int {
	n := len(citations)
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	for i := 0; i < n; i++ {
		if citations[i] < i+1 {
			return i
		}
	}
	return n
}

// 计数桶法：O(n)
func HIndexBucket(citations []int) int {
	n := len(citations)
	buckets := make([]int, n+1)
	for _, c := range citations {
		if c >= n {
			buckets[n]++
		} else {
			buckets[c]++
		}
	}
	papers := 0
	for h := n; h >= 0; h-- {
		papers += buckets[h]
		if papers >= h {
			return h
		}
	}
	return 0
}
```

**执行过程示例**（`citations = [3,0,6,1,5]`，排序法）：

```
排序后: [6,5,3,1,0]
i=0: citations[0]=6 >= 1 ✓
i=1: citations[1]=5 >= 2 ✓
i=2: citations[2]=3 >= 3 ✓
i=3: citations[3]=1 < 4 ✗ → 返回 3
```
