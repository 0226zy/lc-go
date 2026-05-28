# Skill: 初始化 LeetCode 题目

## 描述

为本项目（`lc-go`）在 `solutions/` 下的指定分类目录中创建一个新的 LeetCode 题目框架，包括：
- 题解源文件（仅函数存根 + 注释，**不包含具体实现**）
- 测试文件（表驱动测试结构 + **LeetCode 官方示例用例**，用例不能为空）
- 题目说明文档 `problem.md`（框架，待后续补充解析）

支持从 LeetCode 查询题目信息（标题、描述、示例）自动填充到 `problem.md` 中。

## 触发条件

用户消息中包含 **"初始化题目"** 关键字时触发本 Skill。

## 所需参数

触发后，从用户输入中提取以下信息：

| 参数 | 是否必填 | 说明 |
|------|---------|------|
| `题号` | 可选 | LeetCode 题号，如 `1`、`146` |
| `英文名称` | **必填** | 题目的英文名称，如 `Two Sum`、`LRU Cache` |
| `中文名称` | 可选 | 题目的中文名称，如 "两数之和" |
| `题目描述/链接` | **至少提供一个** | 用于去 LeetCode 查询题目信息 |
| `目标目录` | **必填** | `solutions/` 下的分类目录，如 `easy`、`medium`、`hard`、`hot100` |

## 执行步骤

### Step 1: 解析用户输入

从用户消息中提取上述参数。如果用户提供了 LeetCode 链接（如 `https://leetcode.cn/problems/...` 或 `https://leetcode.com/problems/...`），直接从中解析题号和英文名称。

### Step 2: 检查目标目录

检查用户是否明确指定了 `solutions/` 下的目标目录。

- **如果已指定**：验证该目录是否存在于 `solutions/` 下。
  - 若不存在，询问用户是**创建新目录**还是**选择现有目录**。
- **如果未指定**：
  1. 扫描 `solutions/` 下所有现有目录（排除 `.gitkeep`）。
  2. 使用 `AskUserQuestion` 向用户展示可选目录列表，并提供 "创建新目录" 选项。
  3. 如果用户选择创建新目录，要求输入新目录名（只允许小写字母、数字、下划线、连字符）。

### Step 3: 查询 LeetCode 题目信息

如果用户提供了题目描述或链接，使用 `WebSearch` 或 `FetchURL` 访问 LeetCode 获取题目信息：
- 中文/英文标题
- 题目描述
- **示例输入输出（必须获取，用于填充测试用例）**
- 提示/约束条件

> 若查询失败或用户未提供任何题目信息：
> - `problem.md` 使用占位符填充，提示用户后续手动补充。
> - **测试文件仍需创建，但用例使用 TODO 注释标注预期值，并提示用户根据题目补充；绝不允许生成空的 tests 切片。**

### Step 4: 确定文件结构与命名

根据目标目录的类型，决定文件组织方式：

#### 方式 A：`hot100` 目录（子目录结构）

目录命名：`{题号}_{snake_case_英文名}/`

示例：`0001_two_sum/`、`0146_lru_cache/`

文件清单：
```
solutions/hot100/0001_two_sum/
├── two_sum.go              # 题解存根
├── two_sum_test.go         # 测试文件（含 LeetCode 官方示例用例）
└── problem.md              # 题目说明
```

#### 方式 B：`easy` / `medium` / `hard` 等（扁平结构）

文件直接放在目标目录下：
```
solutions/easy/
├── two_sum.go
└── two_sum_test.go
```

> **注意**：扁平结构下**不创建** `problem.md`，题目信息通过注释写在 `.go` 文件头部。

### Step 5: 生成文件内容

#### 5.1 Go 源文件模板

**`hot100` 子目录结构的 package 名**：使用目录名的 snake_case（不含题号前缀）。

**`easy`/`medium`/`hard` 扁平结构的 package 名**：使用目录名本身（如 `easy`、`medium`）。

```go
package <package_name>

// <中文名> (<英文名>)
// <简要题目描述，从 LeetCode 获取或占位>
// 时间复杂度:   空间复杂度:
func <CamelCase函数名>(/* TODO: 根据题目定义参数 */) /* TODO: 定义返回值 */ {
    // TODO: 实现题解
    panic("not implemented")
}
```

#### 5.2 测试文件模板

测试用例必须从 LeetCode 获取官方示例填充，**不得为空**。

```go
package <package_name>

import (
    "testing"
    // TODO: 按需导入 pkg/utils 或 pkg/datastructures
)

func Test<CamelCase函数名>(t *testing.T) {
    tests := []struct {
        name string
        // 根据题目定义输入字段和 want 字段
    }{
        {"官方示例1", <输入>, <期望输出>},
        {"官方示例2", <输入>, <期望输出>},
        {"空输入", <输入>, <期望输出>},
        {"单元素", <输入>, <期望输出>},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // 调用待实现函数并断言结果
            got := <函数名>(tt.<输入字段>)
            if !<比较结果>(got, tt.want) {
                t.Errorf("<函数名>(%v) = %v, want %v", tt.<输入字段>, got, tt.want)
            }
        })
    }
}
```

**测试用例要求**：
- 至少包含 LeetCode 官方给出的所有示例
- 至少补充 1 个边界情况用例（如空输入、单元素等）
- 使用 `t.Run(tt.name, func(t *testing.T){ ... })` 命名子测试
- 断言失败时使用 `t.Errorf` 打印输入、got、want

#### 5.3 `problem.md` 模板（仅 `hot100` 子目录结构创建）

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

## 提示

- <从 LeetCode 查询到的约束条件，或 TODO 占位>

## 题目解析

### 核心思路

// TODO: 分析核心思路

### 算法步骤

1. // TODO: 描述步骤

### 复杂度分析

- **时间复杂度**: 
- **空间复杂度**: 
```

### Step 6: 创建文件并汇报

1. 使用 `Write` 工具创建所有文件。
2. 如果 `hot100` 子目录结构，还需在 `cmd/main.go` 中**追加**该题目的命令行示例（参考现有格式，添加一个新的 `command` 结构体到 `commands` 切片中，并添加对应 import）。
3. 向用户汇报创建结果：
   - 列出创建的所有文件路径
   - 说明 TODO 占位符的位置，提示用户后续补充
   - 如果是 `hot100` 结构，提示已更新 `cmd/main.go`

## 命名规范

- **目录/文件名**：全部小写，单词间用下划线 `_` 连接（snake_case）
- **函数名**：使用 CamelCase，与题目英文名对应（如 `LRUCache`、`AddTwoNumbers`）
- **Package 名**：
  - `hot100` 子目录：snake_case 英文名（如 `twosum`、`lrucache`）
  - 扁平目录：目录名本身（如 `easy`、`medium`、`hard`）

## 注意事项

1. **绝不自动生成题解代码**：`.go` 文件中只能包含函数签名、`panic("not implemented")` 和 TODO 注释。
2. **必须生成含示例的测试用例**：`_test.go` 中必须包含从 LeetCode 获取的官方示例用例（至少一个），并补充边界情况用例；tests 切片不能为空。
3. 如果用户提供的题目信息不足以确定函数签名（参数类型、返回值类型），在 `.go` 文件中使用 `/* TODO: ... */` 注释提示用户补充。
4. 查询 LeetCode 时优先尝试中文站（`leetcode.cn`），若失败则尝试国际站（`leetcode.com`）。
5. 创建文件后，建议用户运行 `go test ./...` 验证编译是否通过（预期会因 `panic` 而测试失败，但应能正常编译）。
