# 155. 最小栈 (Min Stack)

## 题目描述

设计一个支持 `push`、`pop`、`top` 操作，并能在常数时间内检索到最小元素的栈。

实现 `MinStack` 类：

- `MinStack()` 初始化堆栈对象。
- `void push(int val)` 将元素 val 推入堆栈。
- `void pop()` 删除堆栈顶部的元素。
- `int top()` 获取堆栈顶部的元素。
- `int getMin()` 获取堆栈中的最小元素。

### 示例 1

```
输入:
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出:
[null,null,null,null,-3,null,0,-2]

解释:
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   // 返回 -3
minStack.pop();
minStack.top();      // 返回 0
minStack.getMin();   // 返回 -2
```

### 提示

- `-2^31 <= val <= 2^31 - 1`
- `pop`、`top` 和 `getMin` 操作总是在非空栈上调用
- `push`、`pop`、`top` 和 `getMin` 操作最多调用 `3 * 10^4` 次

## 题目解析

### 核心思路

这道题的关键矛盾在于：`getMin` 要求 O(1) 时间，而栈是后进先出的，最小值会随着 `push`/`pop` 不断变化。如果只存一个「当前最小值」，一旦这个最小值被弹出，我们就丢失了「次小值」的信息。

解决办法是**辅助栈（同步最小值栈）**——典型的「空间换时间」设计：

- 主栈 `data`：正常存放所有元素，支持 `push`/`pop`/`top`。
- 辅助栈 `minStack`：**与主栈同步操作**，栈顶永远是主栈中当前所有元素的最小值。

同步规则：

- `push(val)`：`val` 进主栈；同时如果 `val <= 辅助栈顶`（或辅助栈为空），`val` 也进辅助栈。
- `pop()`：主栈弹出的元素如果等于辅助栈顶，辅助栈也弹出一个——弹出的恰好就是当前最小值，辅助栈的新栈顶自然就是剩余元素的最小值。

正确性论证：辅助栈从栈底到栈顶是一个**非严格递减**序列，它记录了主栈中「曾经成为过最小值」的那些值的副本，且每个值出现的次数恰好等于它在主栈中作为最小值的覆盖范围。因此任意时刻辅助栈顶就是主栈当前的最小值。

> 为什么用 `<=` 而不是 `<`？因为栈中可以有重复元素。比如连续 `push` 两个 `-3`，弹出第一个 `-3` 后最小值仍然是 `-3`，只有 `<=` 才能保证两个副本都进了辅助栈、弹出时才一一对应。

### 算法步骤

1. `Constructor`：初始化主栈 `data` 与辅助栈 `minStack`。
2. `Push(val)`：`val` 压入主栈；若 `minStack` 为空或 `val <= minStack` 栈顶，`val` 也压入 `minStack`。
3. `Pop`：取主栈栈顶元素 `top`；主栈弹出；若 `top == minStack` 栈顶，辅助栈也弹出。
4. `Top`：返回主栈栈顶元素。
5. `GetMin`：返回辅助栈栈顶元素。

### 复杂度分析

- **时间复杂度**: `push`、`pop`、`top`、`getMin` 全部为 O(1)
- **空间复杂度**: O(n)，辅助栈最坏与主栈等大；当元素非递增时辅助栈元素最多

## 代码实现

```go
type MinStack struct {
    data    []int // 主栈，存放所有元素
    minData []int // 辅助栈，栈顶为当前主栈的最小值
}

func Constructor() MinStack {
    return MinStack{}
}

func (s *MinStack) Push(val int) {
    s.data = append(s.data, val)
    // 注意用 <=，保证重复的最小值会重复入辅助栈，弹出时一一对应
    if len(s.minData) == 0 || val <= s.minData[len(s.minData)-1] {
        s.minData = append(s.minData, val)
    }
}

func (s *MinStack) Pop() {
    top := s.data[len(s.data)-1]
    s.data = s.data[:len(s.data)-1]
    if top == s.minData[len(s.minData)-1] {
        s.minData = s.minData[:len(s.minData)-1]
    }
}

func (s *MinStack) Top() int {
    return s.data[len(s.data)-1]
}

func (s *MinStack) GetMin() int {
    return s.minData[len(s.minData)-1]
}
```

**执行过程示例**（官方示例调用序列）：

```
push(-2): data = [-2]            minData = [-2]
push(0):  data = [-2, 0]         minData = [-2]          (0 > -2，不入辅助栈)
push(-3): data = [-2, 0, -3]     minData = [-2, -3]      (-3 <= -2)
getMin(): 返回 minData 栈顶 -3 ✓
pop():    弹出 -3，== 辅助栈顶，辅助栈也弹出
          data = [-2, 0]         minData = [-2]
top():    返回 0 ✓
getMin(): 返回 -2 ✓
```
