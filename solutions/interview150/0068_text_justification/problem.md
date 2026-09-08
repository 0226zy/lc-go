# 68. 文本左右对齐 (Text Justification)

## 题目描述

给定一个单词数组 `words` 和一个宽度 `maxWidth`，重新排版单词，使每行恰好有 `maxWidth` 个字符，并且左右两端对齐。

你应该使用「贪心」策略排版；也就是说，在每行中尽可能多地放下单词。必要时用空格 `' '` 填充，使每行恰好有 `maxWidth` 个字符。

两端对齐的要求是：单词之间的空格要**尽可能均匀**分配。如果不能均匀分配，**左侧的空格要多于右侧的空格**。

- 最后一行应为**左对齐**，单词之间不插入额外的空格，行尾用空格填充到 `maxWidth` 长度。
- 只有一个单词的行同样左对齐，行尾用空格填充。

### 示例 1

```
输入: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
输出:
[
   "This    is    an",
   "example  of text",
   "justification.  "
]
```

### 示例 2

```
输入: words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
输出:
[
  "What   must   be",
  "acknowledgment  ",
  "shall be        "
]
解释: 注意最后一行是 "shall be    " 而不是 "shall     be"，因为最后一行必须左对齐。
     第二行也是左对齐的，因为它只包含一个单词。
```

### 示例 3

```
输入: words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"], maxWidth = 20
输出:
[
  "Science  is  what we",
  "understand      well",
  "enough to explain to",
  "a  computer.  Art is",
  "everything  else we",
  "do                  "
]
```

### 提示

- `1 <= words.length <= 300`
- `1 <= words[i].length <= 20`
- `words[i]` 由小写英文字母和符号组成
- `1 <= maxWidth <= 100`
- `words[i].length <= maxWidth`

## 题目解析

### 核心思路

这道题被标记为「困难」，不是因为算法本身难，而是**细节极多**——空格怎么分、最后一行怎么处理、单词长度恰好占满一行怎么办。但只要按规则把逻辑拆清楚，它就是一道**贪心模拟**题：从左到右依次取词，每次贪心地往当前行塞尽可能多的单词，然后把这一行按规则补成 `maxWidth` 长度。

每行只有三种形态：

1. **最后一行**：不管有几个词，一律左对齐——词与词之间只放 1 个空格，剩下的空格全部填在行尾。
2. **只有一个词且不是最后一行**：无法分配「词间空格」，同样左对齐，词后补空格。
3. **普通行（多词且非末行）**：把所有空格均匀分到词间空隙。假设行内有 `cnt` 个词、总字符长为 `len`，则需要填的空格总数为 `maxWidth - len`，词间空隙有 `cnt - 1` 个：
   - 每个空隙先分到 `space = (maxWidth - len) / (cnt - 1)` 个空格；
   - 余下 `rem = (maxWidth - len) % (cnt - 1)` 个空格，**从左到右每个空隙多分 1 个**（题目要求左边的空隙不少于右边）。

### 算法步骤

1. 用 `i` 指向当前未处理的单词，循环直到所有词都排完：
   - 从 `i` 开始贪心地取词：`cnt` 为行内单词数，`lineLen` 为单词字符总长（不含空格）。只要下一个词放进来后 `lineLen + 单词长 + 词间空格数 <= maxWidth`，就继续取。
   - 判断是否最后一行（`i == len(words)`）。
2. 组装这一行：
   - **最后一行或只有一个词**：词间各放 1 个空格，末尾补齐空格。
   - **普通行**：计算 `space = (maxWidth - lineLen) / (cnt - 1)` 与 `rem = (maxWidth - lineLen) % (cnt - 1)`，前 `rem` 个空隙放 `space + 1` 个空格，其余放 `space` 个。
3. 将该行加入答案，继续处理下一行。

### 复杂度分析

- **时间复杂度**: O(L)，L 为所有输出字符总数（每个空格只生成一次，每个词只遍历一次）。
- **空间复杂度**: O(L)，存储答案所需的输出空间。

## 代码实现

### 贪心模拟

```go
func FullJustify(words []string, maxWidth int) []string {
    res := []string{}
    for i := 0; i < len(words); {
        // 1. 贪心取词：cnt 个词，lineLen 为单词字符总长（不含空格）
        cnt, lineLen := 1, len(words[i])
        for i+cnt < len(words) && lineLen+len(words[i+cnt])+cnt <= maxWidth {
            lineLen += len(words[i+cnt])
            cnt++
        }
        // 2. 组装这一行
        var sb strings.Builder
        if i+cnt == len(words) || cnt == 1 {
            // 最后一行或单行：左对齐，词间 1 空格，行尾补齐
            for j := 0; j < cnt; j++ {
                if j > 0 {
                    sb.WriteByte(' ')
                }
                sb.WriteString(words[i+j])
            }
            for sb.Len() < maxWidth {
                sb.WriteByte(' ')
            }
        } else {
            // 普通行：空格均匀分配，余数从左往右依次多分 1 个
            space, rem := (maxWidth-lineLen)/(cnt-1), (maxWidth-lineLen)%(cnt-1)
            for j := 0; j < cnt-1; j++ {
                sb.WriteString(words[i+j])
                sb.WriteString(strings.Repeat(" ", space))
                if j < rem {
                    sb.WriteByte(' ')
                }
            }
            sb.WriteString(words[i+cnt-1]) // 行尾不补空格
        }
        res = append(res, sb.String())
        i += cnt
    }
    return res
}
```

**执行过程示例**（`words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16`）：

```
第 1 行：贪心取词 "What must be"（4+1+4+1+2=12 <= 16，下一个词 "acknowledgment" 放不进）
  空格总数 = 16-10 = 6，空隙数 = 2
  space = 6/2 = 3，rem = 6%2 = 0
  → "What   must   be"

第 2 行：只有 "acknowledgment"（14 字符），cnt == 1 左对齐
  → "acknowledgment  "

第 3 行：最后一行 "shall be"，左对齐，行尾补空格
  → "shall be        "
```

**易错点提醒**：
- 余数 `rem` 个空格要从**左往右**分给前 `rem` 个空隙，不能反过来；
- 普通行**行尾不补空格**（空格全在词间），最后一行**空格全在行尾**；
- 恰好放满的行（如 `words[i]` 长度等于 `maxWidth`）直接输出该词即可，两种情况的分支都能正确处理。
