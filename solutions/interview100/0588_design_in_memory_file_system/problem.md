# 588. 设计内存文件系统

> 难度：困难 ｜ 分类：前缀树 ｜ 尊享面试 100 题 · 第 68 题
> 链接：https://leetcode.cn/problems/design-in-memory-file-system/

## 题目描述

设计一个内存文件系统，模拟以下功能：

- `FileSystem()`：初始化文件系统，初始时只有一个根目录 `/`。
- `List<String> ls(String path)`：
  - 如果 `path` 是一个**文件**路径，返回仅包含该文件名称的列表；
  - 如果 `path` 是一个**目录**路径，返回该目录下所有文件和子目录的名称列表。
  - 返回结果按**字典序**排列。
- `void mkdir(String path)`：按给定路径创建一个新目录。如果路径中的中间目录不存在，需要一并创建。如果目录已存在则不做任何操作。
- `void addContentToFile(String filePath, String content)`：将 `content` **追加**到文件 `filePath` 末尾。如果文件不存在，则创建该文件（题目保证其父目录已存在）。
- `String readContentFromFile(String filePath)`：返回文件 `filePath` 的全部内容。

### 示例 1

```
输入:
  ["FileSystem", "ls", "mkdir", "addContentToFile", "ls", "readContentFromFile"]
  [[], ["/"], ["/a/b/c"], ["/a/b/c/d", "hello"], ["/"], ["/a/b/c/d"]]
输出:
  [null, [], null, null, ["a"], "hello"]
解释:
  初始根目录为空，ls("/") 返回空列表；
  mkdir("/a/b/c") 依次创建 a、b、c 三层目录；
  addContentToFile 创建文件 d 并写入 "hello"；
  ls("/") 返回 ["a"]；readContentFromFile("/a/b/c/d") 返回 "hello"。
```

### 示例 2

```
输入:
  ["FileSystem", "ls", "mkdir", "ls", "ls", "ls"]
  [[], ["/"], ["/zijzllb"], ["/"], ["/zijzllb"], ["/zijzllb/"]]
输出:
  [null, [], null, ["zijzllb"], [], []]
解释:
  路径末尾可以带 "/"，"ls" 对存在但为空的目录返回空列表。
```

### 提示

- `1 <= path.length, filePath.length <= 100`
- 路径格式为 `/dir1/dir2/...`，可能以 `/` 结尾；目录名和文件名只包含小写字母
- 对 `addContentToFile` 和 `readContentFromFile`，给定的 `filePath` 一定有效（文件存在或父目录存在）
- `content` 只包含小写字母，`1 <= content.length <= 50`
- 所有操作的调用次数总和不超过 `300`

## 思路解析

### 核心思路

文件系统的目录树天然是一棵**前缀树（Trie）**：每个节点代表一个目录或文件，节点的 `children` 是从「名称」到子节点的映射；节点上额外记录 `isFile` 和 `content`。路径 `/a/b/c` 按 `/` 切分后就是从根出发的一条路径。

- `mkdir`：沿路径逐层向下走，缺哪层建哪层（天然支持「创建中间目录」，重复创建也无副作用）。
- `addContentToFile`：走到文件节点，若不存在则在父目录下新建并标记为文件，然后追加内容。
- `ls`：走到目标节点；若它是文件，返回它自己的名字；否则把 `children` 的所有键取出来排序返回。
- `readContentFromFile`：走到文件节点，直接返回内容。

### 算法步骤

1. 定义节点结构：`children map[string]*node`、`isFile bool`、`content string`（或 `strings.Builder`）。
2. 实现辅助函数 `split(path)`：按 `/` 切分并过滤空段，得到路径分量列表；空路径（即 `/`）返回空列表，表示根节点。
3. 实现辅助函数 `walk(segs)`：从根节点出发，沿分量逐层查找，返回最终节点（`ls` / `readContentFromFile` 使用）。
4. `mkdir` 与 `addContentToFile` 使用「查找或创建」版本的 walk：每层若子节点不存在则新建。
5. `ls` 中遇到目录时，将 `children` 的键收集到切片并 `sort.Strings` 排序后返回。

## 复杂度分析

设路径分量的平均长度为 `d`（即目录深度），名称平均长度为 `L`，目录下条目数为 `k`。

- **时间复杂度**:
  - `mkdir`：O(d·L)，逐层哈希查找；
  - `ls`：O(d·L + k·log k)，查找 O(d·L)，结果排序 O(k·log k)；
  - `addContentToFile`：O(d·L + |content|)；
  - `readContentFromFile`：O(d·L)。
- **空间复杂度**: O(所有路径分量总数 + 所有文件内容总长)，每个目录/文件名在 Trie 中只存一份。
