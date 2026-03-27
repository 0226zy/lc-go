# lc-go

LeetCode 刷题 Go 项目

## 项目结构

```
lc-go/
├── go.mod
├── README.md
├── Makefile
├── pkg/
│   ├── utils/            # 工具函数
│   │   ├── compare.go    # 数组/切片比较
│   │   ├── convert.go    # 类型转换
│   │   └── print.go      # 打印调试
│   └── datastructures/   # 常用数据结构
│       ├── listnode.go   # 链表节点
│       ├── treenode.go   # 树节点
│       └── unionfind.go  # 并查集
├── solutions/
│   ├── easy/             # 简单题
│   ├── medium/           # 中等题
│   └── hard/             # 困难题
├── tests/                # 测试文件
│   └── solution_test.go
└── cmd/
    └── main.go           # 入口文件
```

## 快速开始

```bash
# 运行测试
make test

# 运行 main.go
make run

# 编译
make build
```

## 添加新题目

1. 在 `solutions/easy/`、`solutions/medium/` 或 `solutions/hard/` 下新建 `.go` 文件
2. 在 `tests/solution_test.go` 中添加对应测试用例
3. 运行 `make test` 验证
