package wordbreak

import (
	"container/list"
	"strings"
)

// WordBreak 单词拆分（动态规划）
// 判断字符串 s 是否能由字典 wordDict 中的单词拼接而成。
// 时间复杂度: O(n²·L)  空间复杂度: O(n)
func WordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool, len(wordDict))
	maxLen := 0
	for _, w := range wordDict {
		wordSet[w] = true
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	n := len(s)
	// dp[i] 表示 s[0:i] 能否被字典拆分
	dp := make([]bool, n+1)
	dp[0] = true // 空串可拆分

	for i := 1; i <= n; i++ {
		// j 为最后一个分割点，只需考虑 [i-maxLen, i) 范围
		start := 0
		if i-maxLen > 0 {
			start = i - maxLen
		}
		for j := start; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break // 已找到一种拆分方式，无需继续
			}
		}
	}
	return dp[n]
}

// WordBreakBFS 单词拆分（BFS）
// 将位置建模为图节点，BFS 判断从 0 到 n 是否可达。
// 时间复杂度: O(n²·L)  空间复杂度: O(n)
func WordBreakBFS(s string, wordDict []string) bool {
	wordSet := make(map[string]bool, len(wordDict))
	maxLen := 0
	for _, w := range wordDict {
		wordSet[w] = true
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	n := len(s)
	visited := make([]bool, n+1)
	queue := list.New()
	queue.PushBack(0)
	visited[0] = true

	for queue.Len() > 0 {
		start := queue.Remove(queue.Front()).(int)
		// 尝试从 start 出发匹配所有可能长度的单词
		for end := start + 1; end <= n && end-start <= maxLen; end++ {
			if visited[end] {
				continue
			}
			if wordSet[s[start:end]] {
				if end == n {
					return true
				}
				visited[end] = true
				queue.PushBack(end)
			}
		}
	}
	return visited[n]
}

// WordBreakMemo 单词拆分（记忆化递归）
// 递归判断从位置 start 开始的后缀能否拆分，memo 缓存结果。
// 时间复杂度: O(n²·L)  空间复杂度: O(n)
func WordBreakMemo(s string, wordDict []string) bool {
	wordSet := make(map[string]bool, len(wordDict))
	maxLen := 0
	for _, w := range wordDict {
		wordSet[w] = true
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	n := len(s)
	// memo[i]: -1 未计算, 0 false, 1 true
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}

	var canBreak func(start int) bool
	canBreak = func(start int) bool {
		if start == n {
			return true
		}
		if memo[start] != -1 {
			return memo[start] == 1
		}
		// 尝试所有可能的结束位置
		for end := start + 1; end <= n && end-start <= maxLen; end++ {
			if wordSet[s[start:end]] && canBreak(end) {
				memo[start] = 1
				return true
			}
		}
		memo[start] = 0
		return false
	}
	return canBreak(0)
}

// WordBreakTrie 单词拆分（Trie 树优化字典查找）
// 用 Trie 树存储字典，DP 时沿 Trie 匹配，避免子串切片开销。
// 时间复杂度: O(n·maxLen)  空间复杂度: O(字典总长度 + n)
func WordBreakTrie(s string, wordDict []string) bool {
	// 构建 Trie
	root := &trieNode{}
	for _, w := range wordDict {
		node := root
		for i := 0; i < len(w); i++ {
			c := w[i] - 'a'
			if node.children[c] == nil {
				node.children[c] = &trieNode{}
			}
			node = node.children[c]
		}
		node.isEnd = true
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 0; i < n; i++ {
		if !dp[i] {
			continue
		}
		// 从 i 出发沿 Trie 匹配
		node := root
		for j := i; j < n; j++ {
			c := s[j] - 'a'
			if node.children[c] == nil {
				break // 无法继续匹配
			}
			node = node.children[c]
			if node.isEnd {
				dp[j+1] = true
			}
		}
	}
	return dp[n]
}

type trieNode struct {
	children [26]*trieNode
	isEnd    bool
}

// 确保 strings 包被使用（WordBreakTrie 中如需 strings 辅助可扩展）
var _ = strings.Builder{}
