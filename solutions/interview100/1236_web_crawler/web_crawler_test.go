package webcrawler

import (
	"sort"
	"strconv"
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

// mockHtmlParser 用 urls + edges（有向边，edges[i] = [x, y] 表示 urls[x] 页面中包含指向 urls[y] 的链接）
// 模拟题目中的 HtmlParser 接口
type mockHtmlParser struct {
	graph map[string][]string
}

func newMockHtmlParser(urls []string, edges [][2]int) *mockHtmlParser {
	graph := make(map[string][]string)
	for _, e := range edges {
		graph[urls[e[0]]] = append(graph[urls[e[0]]], urls[e[1]])
	}
	return &mockHtmlParser{graph: graph}
}

func (m *mockHtmlParser) GetUrls(url string) []string {
	return m.graph[url]
}

func TestCrawl(t *testing.T) {
	tests := []struct {
		name     string
		urls     []string
		edges    [][2]int
		startUrl string
		want     []string
	}{
		{
			// LeetCode 官方示例1
			name: "示例1: 跳过不同主机名的链接",
			urls: []string{
				"http://news.yahoo.com",
				"http://news.yahoo.com/news",
				"http://news.yahoo.com/news/topics/",
				"http://news.google.com",
				"http://news.yahoo.com/us",
			},
			edges:    [][2]int{{2, 0}, {2, 1}, {3, 2}, {3, 1}, {0, 4}},
			startUrl: "http://news.yahoo.com/news/topics/",
			want: []string{
				"http://news.yahoo.com",
				"http://news.yahoo.com/news",
				"http://news.yahoo.com/news/topics/",
				"http://news.yahoo.com/us",
			},
		},
		{
			// LeetCode 官方示例2
			name: "示例2: 起点链接到的页面主机名都不同",
			urls: []string{
				"http://news.yahoo.com",
				"http://news.yahoo.com/news",
				"http://news.yahoo.com/news/topics/",
				"http://news.google.com",
			},
			edges:    [][2]int{{0, 2}, {2, 1}, {3, 2}, {3, 1}, {3, 0}},
			startUrl: "http://news.google.com",
			want:     []string{"http://news.google.com"},
		},

		// 边界：孤立页面，没有任何出边
		{
			name:     "孤立页面",
			urls:     []string{"http://a.com"},
			edges:    nil,
			startUrl: "http://a.com",
			want:     []string{"http://a.com"},
		},

		// 边界：图中存在环，不能重复爬取
		{
			name: "存在环不重复访问",
			urls: []string{
				"http://a.com/1",
				"http://a.com/2",
				"http://a.com/3",
			},
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 0}},
			startUrl: "http://a.com/1",
			want: []string{
				"http://a.com/1",
				"http://a.com/2",
				"http://a.com/3",
			},
		},

		// 边界：结尾带斜杠与不带斜杠是不同链接
		{
			name: "带斜杠与不带斜杠视为不同链接",
			urls: []string{
				"http://a.com",
				"http://a.com/",
			},
			edges:    [][2]int{{0, 1}},
			startUrl: "http://a.com",
			want: []string{
				"http://a.com",
				"http://a.com/",
			},
		},

		// 典型场景：同主机名的不可达页面不应出现
		{
			name: "同主机名但不可达的页面不返回",
			urls: []string{
				"http://a.com/1",
				"http://a.com/2",
			},
			edges:    nil,
			startUrl: "http://a.com/1",
			want:     []string{"http://a.com/1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := newMockHtmlParser(tt.urls, tt.edges)
			got := Crawl(tt.startUrl, parser)
			// 题目允许任意顺序返回，排序后比较
			sort.Strings(got)
			sort.Strings(tt.want)
			if !utils.EqualStringSlice(got, tt.want) {
				t.Errorf("Crawl(%q) = %v, want %v", tt.startUrl, got, tt.want)
			}
		})
	}
}

func BenchmarkCrawl(b *testing.B) {
	// 构造一条含 1000 个页面的链：urls[i] -> urls[i+1]
	const n = 1000
	urls := make([]string, n)
	for i := range urls {
		urls[i] = "http://a.com/page/" + strconv.Itoa(i)
	}
	edges := make([][2]int, 0, n-1)
	for i := 0; i+1 < n; i++ {
		edges = append(edges, [2]int{i, i + 1})
	}
	parser := newMockHtmlParser(urls, edges)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Crawl(urls[0], parser)
	}
}
