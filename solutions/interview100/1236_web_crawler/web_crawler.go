package webcrawler

import "strings"

// HtmlParser 是题目给定的接口，用于获取某个页面中包含的所有链接。
// 实际评测时由判题系统实现，本地测试用 mock 实现驱动。
type HtmlParser interface {
	// GetUrls 返回给定页面中包含的所有链接
	GetUrls(url string) []string
}

// Crawl 网络爬虫
// 从 startUrl 出发，爬取所有与 startUrl 主机名相同的链接，以任意顺序返回。
// 时间复杂度: O(n + m)，n 为可达页面数，m 为这些页面中的链接总数  空间复杂度: O(n)
func Crawl(startUrl string, htmlParser HtmlParser) []string {
	host := hostname(startUrl)
	visited := make(map[string]bool)
	var ans []string

	var dfs func(url string)
	dfs = func(url string) {
		if visited[url] {
			return
		}
		visited[url] = true
		ans = append(ans, url)
		for _, next := range htmlParser.GetUrls(url) {
			// 只爬取与 startUrl 同一主机名的链接
			if hostname(next) == host {
				dfs(next)
			}
		}
	}
	dfs(startUrl)
	return ans
}

// hostname 提取链接的主机名：去掉 "http://" 前缀后，第一个 '/' 之前的部分
func hostname(url string) string {
	return strings.SplitN(url[7:], "/", 2)[0]
}
