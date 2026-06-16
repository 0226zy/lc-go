# Skill: 初始化 Hot 100 题目框架

## 描述

为本项目在 `solutions/hot100/` 下创建一个新的 LeetCode Hot 100 题目框架，包括：

- **题解源文件**（仅包含**多种解法的函数签名 + 注释**，**不包含任何具体实现**）
- **测试文件**（表驱动测试结构 + **LeetCode 官方示例用例** + 边界用例）
- **题目说明文档** `problem.md`

本 Skill 专门用于 `hot100` 分类，强调**一题多解**的框架：对于适合多种思路的题目，在 `.go` 文件中预定义多个解法函数（如递归/迭代/双指针/动态规划/单调栈等），但均只保留函数签名和 TODO 注释。

支持从 LeetCode 查询题目信息（标题、描述、示例）自动填充到 `problem.md` 和测试用例中。

> **注意**：本 Skill 通过 `WriteFile` 等工具直接创建文件，**不要**调用 `cmd/init_hot100.go`（该文件已删除，CLI 不再提供 `init-hot100` 子命令）。

## 触发条件

用户消息中包含以下任一关键字时触发本 Skill：

- `初始化 hot100`
- `init hot100`
- `hot100 初始化`
- `新建 hot100`
- `创建 hot100 题目`

## 所需参数

| 参数 | 是否必填 | 说明 |
|------|---------|------|
| `题号` | **必填** | LeetCode 题号，如 `35`、`146`。若用户未提供，必须使用 `AskUserQuestion` 或文本提示要求用户输入。 |

## 执行步骤

### Step 1: 获取题号

检查用户消息中是否包含 LeetCode 题号：

- 如果提供了题号，直接提取。
- **如果未提供，必须提示用户输入**，例如：
  > "请提供要初始化的 LeetCode Hot 100 题号。"
- 题号必须是正整数，否则提示重新输入。

### Step 2: 查询 LeetCode 题目信息

使用 LeetCode GraphQL API 查询题目信息：

1. 先通过 `problemsetQuestionList` + `searchKeywords: "<题号>"` 搜索对应的 `titleSlug`。
2. 再通过 `questionData(titleSlug: ...)` 获取完整题目描述 `content`。

需要获取的信息：

- 英文标题 `title`
- 中文标题（如果可获取；否则使用英文标题）
- 题目描述 `content`
- **示例输入输出（必须获取，用于填充测试用例）**
- 约束条件

> 若 LeetCode 查询失败：
> - `problem.md` 使用 TODO 占位。
> - 测试文件仍需创建，用例使用 TODO 注释标注，**不允许生成空的 tests 切片**。

### Step 3: 确定文件结构与命名

目标目录固定为 `solutions/hot100/`。

目录命名：`{题号}_{snake_case_英文名}/`

示例：`0035_search_insert_position/`、`0437_path_sum_iii/`

文件清单：

```
solutions/hot100/<题号>_<snake_case_英文名>/
├── <snake_case_英文名>.go              # 多种解法函数存根
├── <snake_case_英文名>_test.go         # 测试文件（含官方示例 + 边界用例）
└── problem.md                           # 题目说明
```

### Step 4: 生成 Go 源文件（多种解法定义）

**Package 名**：snake_case 英文名，全部小写、**去掉下划线**。

示例：`search_insert_position` → `searchinsertposition`

**文件内容模板**：

```go
package <package_name>

// <中文名> (<英文名>)
// <简要题目描述>
// 时间复杂度: O(?)  空间复杂度: O(?)
func <主函数名>(/* TODO: 根据题目定义参数 */) /* TODO: 定义返回值 */ {
	// TODO: 实现主解法
}

// <解法2函数名> <解法2中文说明>
// 时间复杂度: O(?)  空间复杂度: O(?)
func <解法2函数名>(/* TODO: 根据题目定义参数 */) /* TODO: 定义返回值 */ {
	// TODO: 实现解法2
}

// <解法3函数名> <解法3中文说明>
// 时间复杂度: O(?)  空间复杂度: O(?)
func <解法3函数名>(/* TODO: 根据题目定义参数 */) /* TODO: 定义返回值 */ {
	// TODO: 实现解法3
}
```

**多种解法的推荐组合**（根据题目类型选择合适子集）：

| 题目类型 | 推荐解法函数名 |
|---------|---------------|
| 数组/二分查找 | `SearchInsert` / `SearchInsertBinarySearch` / `SearchInsertLinear` |
| 链表 | `XxxRecursive` / `XxxIterative` / `XxxTwoPointers` |
| 二叉树 | `XxxRecursive` / `XxxIterative` / `XxxMorris` |
| 字符串 | `XxxBruteForce` / `XxxTwoPointers` / `XxxStack` / `XxxDP` |
| 动态规划 | `XxxDP` / `XxxMemoization` / `XxxSpaceOptimized` |
| 图/BFS/DFS | `XxxBFS` / `XxxDFS` / `XxxUnionFind` |
| 滑动窗口 | `XxxSlidingWindow` / `XxxBruteForce` |
| 单调栈/单调队列 | `XxxMonotonicStack` / `XxxBruteForce` |
| 回溯 | `XxxBacktrack` / `XxxIterative` |

至少预定义 **2 种** 不同思路的解法函数；如果题目明显只有一种常见解法，则至少保留主函数 + 一种变体（如递归/迭代）。

### Step 5: 生成测试文件

测试文件必须包含：

1. **LeetCode 官方示例用例**（至少一个，必须能从 content 中解析出输入输出）
2. **边界情况用例**（如空输入、单元素、极大/极小值等）
3. 针对**每种解法**分别编写子测试，命名格式：
   - `Test<函数名>_OfficialExamples`
   - `Test<函数名>_EdgeCases`
   - `Test<函数名>_MultipleSolutions`（在同一用例中分别调用多个解法并比较结果）

**测试文件模板**：

```go
package <package_name>

import (
	"testing"
	// TODO: 按需导入 pkg/utils 或 pkg/datastructures
)

func Test<函数名>_OfficialExamples(t *testing.T) {
	tests := []struct {
		name string
		// TODO: 根据题目定义输入字段
		want interface{}
	}{
		{"官方示例1", <输入>, <期望输出>},
		{"官方示例2", <输入>, <期望输出>},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: 调用函数并断言
		})
	}
}

func Test<函数名>_EdgeCases(t *testing.T) {
	// TODO: 补充边界测试
}

func Test<函数名>_AllSolutions(t *testing.T) {
	// TODO: 对每个官方示例，分别调用所有解法并验证结果一致
}

func Benchmark<函数名>_Main(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// TODO: 调用主解法
	}
}

func Benchmark<函数名>_Solution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// TODO: 调用解法2
	}
}
```

**测试用例要求**：

- `tests` 切片不能为空。
- 至少包含从 LeetCode 官方示例解析出的一个用例。
- 至少补充 1 个边界情况用例。
- 使用 `t.Run(tt.name, func(t *testing.T){ ... })` 命名子测试。
- 断言失败时使用 `t.Errorf` 打印输入、got、want。
- 若无法从 LeetCode 解析具体用例，使用 TODO 注释占位，但 `tests` 切片中仍需保留带 TODO 的示例结构。

### Step 6: 生成 problem.md

```markdown
# <题号>. <中文名> (<英文名>)

## 题目描述

<从 LeetCode 查询到的题目描述，或 TODO 占位>

### 示例 1

```
输入：
输出：
解释：
```

## 约束条件

- <从 LeetCode 查询到的约束条件，或 TODO 占位>

## 核心思路

TODO: 实现后补充

### 解法 1：<主解法>

TODO

### 解法 2：<另一种解法>

TODO

### 解法 3：<再一种解法，可选>

TODO

## 复杂度分析

| 解法 | 时间复杂度 | 空间复杂度 |
|------|----------|----------|
| <解法1> | O(?) | O(?) |
| <解法2> | O(?) | O(?) |
| <解法3> | O(?) | O(?) |
```

### Step 7: 创建文件并汇报

1. 使用 `WriteFile` 工具创建所有文件。
2. **不需要**修改 `cmd/main.go`（本 Skill 仅初始化框架，不注册运行示例）。
3. **不要**调用或引用 `cmd/init_hot100.go`：该文件已从仓库中删除，不存在 `lc-go init-hot100` 命令。
3. 向用户汇报创建结果：
   - 列出创建的所有文件路径
   - 说明预定义的多种解法函数名
   - 说明 TODO 占位符的位置
   - 提示用户运行 `go test ./solutions/hot100/<目录>/ -v` 验证编译（预期测试失败，但应能正常编译）

## 命名规范

- **目录/文件名**：全部小写，单词间用下划线 `_` 连接（snake_case）
- **主函数名**：使用 CamelCase，与题目英文名对应（如 `SearchInsert`、`PathSum`、`LowestCommonAncestor`）
- **多种解法函数名**：在主函数名后追加解法后缀（如 `SearchInsertLinear`、`PathSumPrefixSum`、`LowestCommonAncestorParentPointer`）
- **Package 名**：snake_case 英文名，全部小写、无下划线（如 `searchinsertposition`、`pathsumiii`）

## 注意事项

1. **绝不自动生成题解代码**：所有 `.go` 文件中的函数只能包含签名、`TODO` 注释和必要的 `panic("not implemented")`。
2. **必须生成含示例的测试用例**：`_test.go` 中必须包含从 LeetCode 获取的官方示例用例，并补充边界情况；`tests` 切片不能为空。
3. 如果题目信息不足以确定函数签名（参数类型、返回值类型），在 `.go` 文件中使用 `/* TODO: ... */` 注释提示用户补充。
4. 查询 LeetCode 时优先使用国际站（`leetcode.com/graphql`），若失败可尝试中文站。
5. 创建文件后，建议用户运行 `go test ./solutions/hot100/<目录>/ -v` 验证编译是否通过。
