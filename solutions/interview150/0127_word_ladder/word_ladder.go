package wordladder

// LadderLength 单词接龙
// 从 beginWord 出发，每次只能改一个字母且中间单词必须在 wordList 中，
// 返回从 beginWord 到 endWord 的最短转换序列中的单词数目（含首尾），不存在返回 0。
// 时间复杂度: O(n * L^2) n 为单词数，L 为单词长度（构建通配符映射 + BFS）  空间复杂度: O(n * L) 通配符映射表
func LadderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool, len(wordList))
	for _, w := range wordList {
		wordSet[w] = true
	}
	if !wordSet[endWord] {
		return 0
	}

	// 预处理：把所有单词按“去掉第 i 个字母”的通配符模式分组，
	// 例如 hot -> *ot, h*t, ho*，同一组内任意两个单词只差一个字母
	patternMap := make(map[string][]string)
	for _, w := range wordList {
		for i := 0; i < len(w); i++ {
			pattern := w[:i] + "*" + w[i+1:]
			patternMap[pattern] = append(patternMap[pattern], w)
		}
	}

	visited := map[string]bool{beginWord: true}
	queue := []string{beginWord}
	depth := 1 // beginWord 本身算第 1 个单词

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			word := queue[0]
			queue = queue[1:]
			if word == endWord {
				return depth
			}
			for j := 0; j < len(word); j++ {
				pattern := word[:j] + "*" + word[j+1:]
				for _, next := range patternMap[pattern] {
					if !visited[next] {
						visited[next] = true
						queue = append(queue, next)
					}
				}
			}
		}
		depth++
	}
	return 0
}
