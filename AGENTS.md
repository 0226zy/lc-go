# AGENTS.md

> 本文件面向 AI 编程助手。如果你正在阅读本文件，说明你可能需要在这个 LeetCode 刷题 Go 项目上添加新题目、修改现有代码或运行测试。本文档假设你对项目一无所知，以下是所有你需要了解的上下文。

## 项目概述

`lc-go` 是一个使用 **Go 1.22** 编写的 LeetCode 刷题项目，模块路径为 `github.com/0226zy/lc-go`。
项目结构遵循标准 Go 项目布局，所有题目实现、单元测试、性能基准测试和题解分析均集中在仓库中。

项目完全依赖 Go 标准库，**没有引入任何第三方依赖**。

## 目录结构

```
lc-go/
├── go.mod                  # 模块定义，Go 1.22
├── Makefile                # 常用构建/测试命令
├── README.md               # 面向人类的项目简介
├── AGENTS.md               # 本文件
├── cmd/
│   └── main.go             # CLI 入口，支持运行已注册题目的示例
├── pkg/
│   ├── utils/              # 测试与调试工具函数
│   │   ├── compare.go      # 切片比较（EqualIntSlice, Equal2DIntSlice 等）
│   │   ├── convert.go      # 类型转换（IntSliceToString, StringToIntSlice 等）
│   │   └── print.go        # 打印辅助（PrintIntSlice, PrintMatrix 等）
│   └── datastructures/     # 常用数据结构
│       ├── listnode.go     # 链表节点 + 构造/转换辅助方法
│       ├── treenode.go     # 二叉树节点 + 层序构造/打印方法
│       └── unionfind.go    # 并查集（路径压缩 + 按秩合并）
└── solutions/              # 题解代码
    ├── easy/               # 简单题
    │   └── 0001_two_sum/
    │       ├── problem.md
    │       ├── two_sum.go
    │       └── two_sum_test.go
    ├── medium/             # 中等题
    ├── hard/               # 困难题
    ├── interview100/       # LeetCode 会员题单「尊享面试 100 题」（当前仅有 problem.md 占位，实现待补充）
    └── hot100/             # LeetCode Hot 100
        └── 0002_add_two_numbers/
            ├── problem.md          # 题目描述、思路解析、复杂度分析
            ├── add_two_numbers.go   # 题解实现
            └── add_two_numbers_test.go  # 单元测试 + Benchmark
```

### 题目组织方式

`solutions/` 下所有分类目录（`easy/`、`medium/`、`hard/`、`hot100/` 等）采用**统一的子目录结构**：

- 每道题一个**独立子目录**，目录名格式为 `<题号>_<下划线分隔的描述>`，例如 `0001_two_sum`、`0002_add_two_numbers`。
- 每个子目录下包含：
  - `problem.md`：题目描述、核心思路、算法步骤、复杂度分析（**中文撰写**）。
  - `<name>.go`：题解代码，包名使用**小写、无下划线**的缩写或简写，例如 `twosum`、`addtwonumbers`。
  - `<name>_test.go`：测试文件，包名与实现文件相同。

## 构建与测试命令

项目根目录提供 `Makefile`，常用命令如下：

| 命令 | 作用 |
|------|------|
| `make test` | 运行所有测试（`go test ./... -v`） |
| `make run` | 运行 CLI 入口（`go run ./cmd/main.go`） |
| `make build` | 编译可执行文件到 `bin/lc-go` |
| `make clean` | 删除编译产物 `bin/` |

直接运行 `go test ./solutions/... -v` 也可以。

## 代码风格规范

- **语言**：所有注释、文档、`problem.md`、测试用例名称均使用**中文**。
- **导出函数**：题解函数必须是导出形式（PascalCase），因为它们会被 `cmd/main.go` 或其他包导入调用。
- **复杂度注释**：每个题解函数上方必须标注时间复杂度和空间复杂度，格式如下：
  ```go
  // TwoSum 两数之和
  // 给定一个整数数组 nums 和一个整数目标值 target，找出和为目标值的两个整数的下标。
  // 时间复杂度: O(n)  空间复杂度: O(n)
  func TwoSum(nums []int, target int) []int { ... }
  ```
- **包命名**：
  - 所有子目录均使用描述性小写包名，不含下划线（如 `package twosum`、`package longestpalindromicsubstring`）。
- **数据结构**：链表、树等题目不要自己重新定义节点结构，统一使用 `pkg/datastructures` 中的类型：
  - `datastructures.ListNode`
  - `datastructures.TreeNode`
  - `datastructures.UnionFind`
- **工具函数**：测试中需要比较切片或二维切片时，优先使用 `pkg/utils` 中的辅助函数，避免手写重复比较逻辑。

## 测试策略

- 使用标准库 `testing`，**表驱动测试**（table-driven tests）。
- 每个测试用例使用 `t.Run(tt.name, func(t *testing.T) { ... })` 命名子测试，方便定位失败。
- 测试用例名称使用中文，覆盖：
  - LeetCode 官方示例
  - 边界情况（空输入、单元素、极大/极小值等）
  - 性能压力场景（如超长字符串、大数组）
- **Benchmark**：对核心算法鼓励编写 `BenchmarkXxx` 函数，比较不同解法或不同数据规模下的性能。
- 浮点数比较使用 epsilon（如 `eps = 1e-5`），避免直接相等判断。

## 如何添加新题目

1. 根据题目难度或分类，选择放入 `solutions/easy/`、`solutions/medium/`、`solutions/hard/` 或 `solutions/hot100/`。
2. 新建目录：`solutions/<分类>/000X_problem_name/`
3. 添加 `problem.md`（题目描述 + 思路解析 + 复杂度分析）
4. 添加 `<name>.go`（题解实现）
5. 添加 `<name>_test.go`（充分覆盖的单元测试 + Benchmark）
6. 如果需要在 CLI 中运行该题，需在 `cmd/main.go` 中注册对应的 `command` 条目（目前为手动维护）。
7. 运行 `make test` 确保全部通过。

## 注意事项

- 项目**没有外部依赖**，不要引入第三方包。
- `cmd/main.go` 中已注册的命令列表需要手动同步，新增题目后如果希望 CLI 支持，请一并更新。
- `.gitignore` 仅忽略了 `.cursor/`，没有其他特殊排除规则。
- 本项目为本地学习/刷题用途，**没有网络服务、数据库或部署流程**，不涉及传统意义上的安全漏洞（如 SQL 注入、XSS 等），但仍应保持代码健壮性（边界检查、避免 panic 等）。
