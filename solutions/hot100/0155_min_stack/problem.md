# 155. 最小栈

## 题目描述

设计一个支持 `push`，`pop`，`top`，`getMin` 操作的栈，且每个操作的时间复杂度均为 O(1)。

- `MinStack()` 初始化堆栈对象
- `void push(int val)` 将元素 val 推入堆栈
- `void pop()` 删除堆栈顶部的元素
- `int top()` 获取堆栈顶部的元素
- `int getMin()` 获取堆栈中的最小元素

**示例 1：**

```
输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
```

**提示：**

- `-2^31 <= val <= 2^31 - 1`
- `pop`、`top` 和 `getMin` 操作总是在非空栈上调用

## 解题思路

### 方法一：辅助栈

使用两个栈，一个数据栈存储实际值，一个辅助栈存储当前最小值。

- `push`：数据栈正常入栈，辅助栈入栈 min(当前值, 辅助栈顶)
- `pop`：两个栈同时出栈
- `top`：返回数据栈顶
- `getMin`：返回辅助栈顶

### 方法二：单栈存储差值

使用一个栈存储当前值与最小值的差值，同时维护一个全局最小值变量。

- `push`：如果栈为空，直接入栈 val，更新 minVal = val；否则入栈 val - minVal，更新 minVal
- `pop`：如果栈顶 < 0，说明该元素入栈时更新了最小值，需要回退 minVal
- `top`：如果栈顶 < 0，返回 minVal；否则返回 minVal + 栈顶
- `getMin`：返回 minVal

## 复杂度分析

### 方法一：辅助栈

- **时间复杂度**：每个操作均为 O(1)
- **空间复杂度**：O(n)，辅助栈与数据栈大小相同

### 方法二：单栈存储差值

- **时间复杂度**：每个操作均为 O(1)
- **空间复杂度**：O(n)，只需一个栈

## 代码实现

见 `min_stack.go`
