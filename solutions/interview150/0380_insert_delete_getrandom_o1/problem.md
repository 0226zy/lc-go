# 380. O(1) 时间插入、删除和获取随机元素 (Insert Delete GetRandom O(1))

## 题目描述

实现 `RandomizedSet` 类：

- `RandomizedSet()` 初始化一个空的集合。
- `bool insert(int val)` 当元素 `val` 不存在时，向集合中插入该项，并返回 `true`；否则返回 `false`。
- `bool remove(int val)` 当元素 `val` 存在时，从集合中移除该项，并返回 `true`；否则返回 `false`。
- `int getRandom()` 随机返回现有集合中的一项（保证调用此方法时集合中至少存在一个元素）。每个元素应该有**相同的概率**被返回。

你必须实现类的所有函数，并满足每个函数的**平均**时间复杂度为 O(1)。

### 示例 1

```
输入: ["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
     [[], [1], [2], [2], [], [1], [2], []]
输出: [null, true, false, true, 2, true, false, 2]
解释:
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1);    // 向集合中插入 1，返回 true
randomizedSet.remove(2);    // 集合中不存在 2，返回 false
randomizedSet.insert(2);    // 向集合中插入 2，返回 true
randomizedSet.getRandom();  // getRandom 应随机返回 1 或 2（本题输出 2）
randomizedSet.remove(1);    // 从集合中移除 1，返回 true
randomizedSet.insert(2);    // 2 已在集合中，返回 false
randomizedSet.getRandom();  // getRandom 应随机返回 2
```

## 提示

- `-2^31 <= val <= 2^31 - 1`
- 最多调用 `insert`、`remove` 和 `getRandom` 共 `2 * 10^5` 次
- 在调用 `getRandom` 方法时，集合中至少存在一个元素

## 题目解析

### 核心思路

三个操作各有天然的 O(1) 结构，但各有短板：

- **哈希表**：插入、删除 O(1)，但无法"等概率随机取一个"；
- **动态数组**：按下标随机取 O(1)、尾部插入 O(1)，但删除指定值要 O(n) 查找，删除中间元素还要 O(n) 搬移。

把两者结合起来：**切片负责存值，哈希表负责定位**。哈希表存"值 → 切片下标"，这样每个值都能 O(1) 找到它在切片里的位置。

关键在 **remove 怎么做到 O(1)**。数组删除慢，是因为要挪走中间的空洞；那我们就不留空洞——用**切片最后一个元素覆盖待删元素的位置，然后截断末尾**：

```
vals = [10, 20, 30, 40]   要删除 20（下标 1）
  1. 把末尾元素 40 放到下标 1:  [10, 40, 30, 40]
  2. 截断末尾:                 [10, 40, 30]
  3. 更新哈希表中 40 的下标为 1
```

这样切片永远保持紧凑，没有空洞。由于切片里存的就是当前所有元素、且每个位置等概率，`getRandom` 只需 `vals[rand.Intn(len(vals))]`，天然满足等概率。

本题属于 **"哈希表 + 动态数组（交换删除）"** 的经典设计模型。

### 算法步骤

- **insert(val)**：若 `val` 在哈希表中已存在，返回 `false`；否则追加到切片尾部，哈希表记录 `val -> 下标`，返回 `true`。
- **remove(val)**：若哈希表中没有 `val`，返回 `false`；否则：
  1. 查哈希表得到下标 `i`；
  2. 取末尾元素 `lastVal`，把它写到 `vals[i]`，哈希表更新 `lastVal -> i`；
  3. 切片截掉末尾；哈希表删除 `val`；返回 `true`。
- **getRandom**：返回 `vals[rand.Intn(len(vals))]`。

### 复杂度分析

- **时间复杂度**：三个操作均为 O(1)（均摊）。哈希表操作 O(1)，切片尾部插入/截断 O(1)，随机下标访问 O(1)。
- **空间复杂度**: O(n)，n 为集合中元素个数（切片 + 哈希表各存一份）。

## 代码实现

```go
type RandomizedSet struct {
	vals  []int       // 存储所有元素，保持紧凑
	index map[int]int // 值 -> 在 vals 中的下标
}

func Constructor() RandomizedSet {
	return RandomizedSet{vals: make([]int, 0), index: make(map[int]int)}
}

func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.index[val]; ok {
		return false
	}
	s.index[val] = len(s.vals)
	s.vals = append(s.vals, val)
	return true
}

func (s *RandomizedSet) Remove(val int) bool {
	i, ok := s.index[val]
	if !ok {
		return false
	}
	last := len(s.vals) - 1
	s.vals[i] = s.vals[last]       // 末尾元素覆盖待删位置
	s.index[s.vals[last]] = i      // 更新被搬动元素的下标
	s.vals = s.vals[:last]         // 截断末尾
	delete(s.index, val)
	return true
}

func (s *RandomizedSet) GetRandom() int {
	return s.vals[rand.Intn(len(s.vals))]
}
```

**remove 执行过程示例**（`vals = [10,20,30,40]`，删除 `20`）：

```
删除前: vals=[10,20,30,40], index={10:0, 20:1, 30:2, 40:3}
i = index[20] = 1
vals[1] = vals[3] → vals=[10,40,30,40]
index[40] = 1 → index={10:0, 20:1, 30:2, 40:1}
截断末尾 → vals=[10,40,30]
删除 index[20] → index={10:0, 40:1, 30:2}
```
