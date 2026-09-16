# 1236. 网络爬虫

> 难度：中等 ｜ 分类：图 ｜ 尊享面试 100 题 · 第 57 题
> 链接：https://leetcode.cn/problems/web-crawler/

## 题目描述

给定一个链接 `startUrl` 和一个接口 `HtmlParser`，请你实现一个网络爬虫，从 `startUrl` 开始爬取所有与 `startUrl` 处于**同一主机名（hostname）**下的链接。

以**任意顺序**返回爬虫获取的所有链接。

你的爬虫应当：

- 从页面 `startUrl` 开始爬取；
- 调用 `HtmlParser.GetUrls(url)` 来获取某个页面中的所有链接；
- 同一个链接不要爬取两次；
- 只爬取与 `startUrl` **主机名相同**的链接。

如下图所示，链接 `http://example.org/test` 的主机名是 `example.org`。简单起见，你可以假设所有链接都使用 **http 协议**且**没有指定端口**。例如 `http://leetcode.com/problems` 和 `http://leetcode.com/contest` 处于同一主机名下，而 `http://example.org/test` 和 `http://example.com/abc` 不是。

`HtmlParser` 接口定义如下：

```go
type HtmlParser interface {
	// GetUrls 返回给定页面中包含的所有链接
	GetUrls(url string) []string
}
```

注意：结尾带斜杠 `/` 的链接与不带斜杠的链接视为**不同**的链接，例如 `http://news.yahoo.com` 和 `http://news.yahoo.com/` 是不同的链接。

### 示例 1

```
输入:
urls = [
  "http://news.yahoo.com",
  "http://news.yahoo.com/news",
  "http://news.yahoo.com/news/topics/",
  "http://news.google.com",
  "http://news.yahoo.com/us"
]
edges = [[2,0],[2,1],[3,2],[3,1],[0,4]]
startUrl = "http://news.yahoo.com/news/topics/"
输出: [
  "http://news.yahoo.com",
  "http://news.yahoo.com/news",
  "http://news.yahoo.com/news/topics/",
  "http://news.yahoo.com/us"
]
解释: edges[i] = [x, y] 表示链接 urls[x] 对应的页面中包含指向 urls[y] 的链接。
     从 startUrl（下标 2）出发可以到达下标 0 和 1，从下标 0 又能到达下标 4；
     下标 3 的主机名是 news.google.com，与 startUrl 不同，被跳过。
```

### 示例 2

```
输入:
urls = [
  "http://news.yahoo.com",
  "http://news.yahoo.com/news",
  "http://news.yahoo.com/news/topics/",
  "http://news.google.com"
]
edges = [[0,2],[2,1],[3,2],[3,1],[3,0]]
startUrl = "http://news.google.com"
输出: ["http://news.google.com"]
解释: startUrl 链到了所有其它页面，但它们的主机名都与 startUrl 不同，因此只返回 startUrl 本身。
```

### 提示

- `1 <= urls.length <= 1000`
- `1 <= urls[i].length <= 300`
- `startUrl` 是 `urls` 中的一个
- 主机名长度在 1 到 63 个字符之间（含点号），只包含小写字母 `a` 到 `z`、数字 `0` 到 `9` 以及连字符 `-`
- 主机名不会以连字符开头或结尾
- 可以假设链接库中没有重复的链接

## 思路解析

### 核心思路

把每个页面看成一个节点、页面中的链接看成有向边，本题就是**从起点出发的图遍历**（DFS 或 BFS 均可），外加两个约束：不重复访问、不跨主机名。

用哈希集合 `visited` 记录已经爬过的链接，保证每个链接只调用一次 `GetUrls`；遍历到一个新链接时，先比较它与 `startUrl` 的主机名，不同就直接跳过。主机名的提取很简单：所有链接都以 `http://` 开头，去掉前 7 个字符后，第一个 `/` 之前的部分就是主机名。

### 算法步骤

1. 提取 `startUrl` 的主机名 `host`；初始化空集合 `visited` 与结果切片。
2. 从 `startUrl` 开始 DFS（或 BFS）：
   - 若当前链接已在 `visited` 中，直接返回；
   - 否则加入 `visited` 并计入结果；
   - 调用 `htmlParser.GetUrls(url)` 获取页面中的所有链接，对每个链接：若主机名与 `host` 相同则递归爬取。
3. 遍历结束后返回结果切片（题目允许任意顺序）。

## 复杂度分析

- **时间复杂度**: O(n + m)。`n` 为同一主机名下可达的页面数，`m` 为这些页面中链接的总数；每个页面最多调用一次 `GetUrls`，每条链接最多被检查一次。
- **空间复杂度**: O(n)。`visited` 集合与递归调用栈（DFS）或队列（BFS）最多存储 n 个链接。
