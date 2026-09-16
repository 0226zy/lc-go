package sentencesimilarity

// AreSentencesSimilar 句子相似性
// 判断两个句子（字符串数组）是否相似：长度相同，且每个对应位置的单词相同或相似。
// 相似关系双向但不可传递，similarPairs 中的 [x, y] 表示 x 与 y 直接相似。
// 时间复杂度: O(n+p*L)，n 为句子长度，p 为词对数量，L 为单词平均长度  空间复杂度: O(p*L)
func AreSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	// 长度不同直接不相似
	if len(sentence1) != len(sentence2) {
		return false
	}

	// 建哈希表：单词 -> 所有与它直接相似的单词集合（双向插入）
	sim := make(map[string]map[string]bool)
	for _, pair := range similarPairs {
		x, y := pair[0], pair[1]
		if sim[x] == nil {
			sim[x] = make(map[string]bool)
		}
		sim[x][y] = true
		if sim[y] == nil {
			sim[y] = make(map[string]bool)
		}
		sim[y][x] = true
	}

	// 逐位置比较
	for i := range sentence1 {
		a, b := sentence1[i], sentence2[i]
		if a == b {
			continue // 单词与自身总是相似
		}
		if !sim[a][b] {
			return false
		}
	}
	return true
}
