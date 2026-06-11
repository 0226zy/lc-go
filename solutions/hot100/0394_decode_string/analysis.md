# 394 题当前实现问题分析

## 当前实现

```go
func DecodeString(s string) string {
	stack := []byte{}
	for _, c := range s {
		b := byte(c)
		stack = append(stack, b)
	}
	return popStack(stack)
}

func popStack(stack []byte) string {
	stack = stack[:len(stack)-1]
	curr := ""
	str := []byte{}
	for len(stack) > 0 {
		n := len(stack)
		top := stack[n-1]
		stack = stack[:n-1]
		if top == '[' || top == ']' {
			continue
		}

		if isDigit(top) {
			curr = repeate(string(str)+curr, int(top-'0'))
		} else {
			str = append([]byte{top}, str...)
		}
	}
	return curr
}

func isDigit(b byte) bool {
	if int(b-'0') >= 0 && int(b-'0') <= 9 {
		return true
	}
	return false
}

func repeate(str string, n int) string {
	ret := ""
	for i := 0; i < n; i++ {
		ret = ret + str
	}
	return ret
}
```

## 运行测试结果

执行 `go test` 后**几乎全部失败**，仅 `纯数字括号`（`3[z]`）通过。

典型错误输出：

| 输入 | 期望 | 实际输出 |
|------|------|----------|
| `3[a]2[bc]` | `aaabcbc` | `abcbcbcabcbcbcabcbcbc` |
| `3[a2[c]]` | `accaccacc` | `acccacccaccc` |
| `abc` | `abc` | `""` (空字符串) |
| `10[a]` | `aaaaaaaaaa` | `a` |
| `2[abc]3[cd]ef` | `abcabccdcdcdef` | `abccdecdecdecde...` |

---

## 问题 1：栈的语义完全用反，导致字符顺序反转

### 错误分析

`DecodeString` 将所有字符按**从左到右**的顺序压入 `stack`（切片末尾追加），然后 `popStack` 从**末尾弹出**处理。

这导致栈顶元素是字符串的最后一个字符，处理顺序完全反转。例如输入 `3[a]2[bc]`，栈内顺序为：

```
[3, [, a, ], 2, [, b, c, ]]
```

从栈顶弹出依次为：`]`, `c`, `b`, `[`, `2`, `]`, `a`, `[`, `3`。

虽然代码中用 `str = append([]byte{top}, str...)` 做 prepend 试图恢复字母顺序，但数字和括号的顺序已经无法挽回。最终 `3[a]` 和 `2[bc]` 的**处理顺序也反转了**，导致输出混乱。

### 正确思路

字符串解码的栈应该按照**正常阅读顺序**处理，或者使用递归/栈在遇到 `]` 时回退到上一个 `[` 的上下文。当前实现把所有字符一次性压栈再从尾部弹出，是一种"伪栈"用法，本质上是把整个字符串反转后再处理。

---

## 问题 2：多位数字无法解析

### 错误代码

```go
if isDigit(top) {
    curr = repeate(string(str)+curr, int(top-'0'))
}
```

### 分析

代码将每一位数字单独处理，只支持 `1~9` 的个位数。当输入为 `10[a]` 时：

1. 弹出 `]`，跳过
2. 弹出 `a`，`str = "a"`
3. 弹出 `0`，`isDigit('0') = true`，执行 `repeat("a", 0)` 返回 `""`
4. 弹出 `1`，`isDigit('1') = true`，执行 `repeat("a" + "", 1)` 返回 `"a"`

最终输出 `"a"`，而非期望的 10 个 `a`。

### 正确做法

遇到数字时，应该从栈中**连续弹出多位数字**并组合成完整的整数：

```go
num := 0
for len(stack) > 0 && isDigit(stack[len(stack)-1]) {
    num = num*10 + int(stack[len(stack)-1]-'0')
    stack = stack[:len(stack)-1]
}
```

---

## 问题 3：嵌套括号结构完全无法处理

### 分析

`popStack` 试图用一个单层循环处理整个栈，没有嵌套上下文的切换机制。

以 `3[a2[c]]` 为例：
- 正确语义：先解析内层 `2[c]` 得到 `"cc"`，再解析外层 `3[acc]` 得到 `"accaccacc"`
- 当前代码：从内层 `]` 开始一路向外弹出，遇到 `2` 就重复当前字符串，但此时根本不知道哪里是内层的边界。最终把 `2` 和 `3` 混为一谈，输出 `acccacccaccc`

### 根本原因

缺少一个结构来保存**外层上下文**（外层的重复次数和外层已解析的字符串）。这正是栈结构在解码问题中的核心作用：

```
遇到 '[' 时：将当前的重复次数和已解析字符串压栈，开始新上下文
遇到 ']' 时：弹出栈顶，恢复外层上下文，将当前字符串重复后追加到外层
```

---

## 问题 4：纯字母输入（无数字无括号）返回空字符串

### 错误代码

```go
func popStack(stack []byte) string {
    stack = stack[:len(stack)-1]  // 无条件弹出最后一个字符！
    curr := ""
    // ...
    return curr  // curr 仅在遇到数字时被赋值
}
```

### 分析

1. `popStack` 第一句 `stack = stack[:len(stack)-1]` 就**无条件弹出了最后一个字符**，无论它是什么。输入 `abc` 时，`c` 直接被丢弃。
2. 循环中只有遇到数字才会给 `curr` 赋值。如果输入中没有数字，`curr` 永远是 `""`。
3. 循环中收集到 `str` 中的普通字母从未被返回。

这导致 `"abc"` -> `""`，`"ef"` -> `""` 等全部失败。

### 正确做法

普通字母应该按顺序直接拼接到当前正在构建的字符串中，并在最后返回完整的字符串。

---

## 问题 5：`popStack` 假设输入以 `]` 结尾

### 分析

`stack = stack[:len(stack)-1]` 的意图是弹出末尾的 `]`，但如果输入不以 `]` 结尾（如 `2[abc]3[cd]ef`、`abc`），就会**丢失最后一个有效字符**。

例如 `2[abc]3[cd]ef`：
- 最后一个字符是 `f`，直接被丢弃
- 倒数第二个 `e` 进入循环处理
- 最终输出中完全不包含 `f`

---

## 问题 6：括号配对逻辑错误

### 错误代码

```go
if top == '[' || top == ']' {
    continue
}
```

### 分析

代码把 `[` 和 `]` 一视同仁，全部跳过。但 `[` 和 `]` 在解码中的语义截然不同：

- `[` 标志着一个新编码段的**开始**，前面紧跟的是重复次数
- `]` 标志着一个编码段的**结束**，此时应该取重复次数并重复当前字符串

简单跳过 `]` 还能勉强蒙混（如 `3[z]` 可以通过），但跳过 `[` 会导致无法区分不同编码段的边界。

---

## 问题 7：`repeate` 函数使用字符串拼接，效率低下

### 错误代码

```go
func repeate(str string, n int) string {
    ret := ""
    for i := 0; i < n; i++ {
        ret = ret + str
    }
    return ret
}
```

### 分析

Go 中字符串是不可变的，`ret = ret + str` 每次都会创建新字符串并拷贝旧内容，时间复杂度为 O(n^2)。虽然题目数据量小（输出不超过 10^5），但使用 `strings.Builder` 是更好的做法：

```go
var sb strings.Builder
for i := 0; i < n; i++ {
    sb.WriteString(str)
}
return sb.String()
```

---

## 问题 8：函数名拼写错误

`repeate` 应为 `repeat`。这不会导致功能错误，但影响代码可读性。

---

## 修正后的实现思路（递归法）

```go
func DecodeString(s string) string {
    var dfs func(i int) (string, int)
    dfs = func(i int) (string, int) {
        var result strings.Builder
        for i < len(s) {
            if s[i] == ']' {
                return result.String(), i
            }
            if isDigit(s[i]) {
                num := 0
                for i < len(s) && isDigit(s[i]) {
                    num = num*10 + int(s[i]-'0')
                    i++
                }
                i++ // skip '['
                subStr, nextIdx := dfs(i)
                i = nextIdx + 1 // skip ']'
                for j := 0; j < num; j++ {
                    result.WriteString(subStr)
                }
            } else {
                result.WriteByte(s[i])
                i++
            }
        }
        return result.String(), i
    }
    res, _ := dfs(0)
    return res
}
```

### 复杂度分析

- **时间复杂度**: O(n * k)，n 为输入长度，k 为最大重复次数。每个字符在最终结果中出现一次，构建结果的总时间为输出长度。
- **空间复杂度**: O(d)，d 为最大嵌套深度，递归栈的深度。

---

## 总结

当前实现的根本性错误在于：**试图用一个无状态的倒序遍历替代栈/递归的嵌套状态管理**。具体表现为：

| 问题 | 影响 |
|------|------|
| 栈语义用反 | 处理顺序完全颠倒，输出乱序 |
| 多位数字无法解析 | `10[a]` -> `"a"`，`100[a]` -> `""` |
| 无嵌套处理能力 | `3[a2[c]]` 输出错误 |
| 纯字母返回空 | `"abc"` -> `""` |
| 无条件弹栈尾 | 不以 `]` 结尾的输入丢失最后一个字符 |
| 括号语义混乱 | `[` 和 `]` 被同等跳过，无法区分编码段边界 |
| 字符串拼接低效 | 重复次数大时性能差 |

核心修正方向：**引入栈或递归来保存嵌套上下文，正序处理输入，在遇到 `]` 时回退到上一层上下文继续构建字符串。**
