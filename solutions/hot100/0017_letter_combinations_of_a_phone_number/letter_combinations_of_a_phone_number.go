package lettercombinations

// LetterCombinations 电话号码的字母组合（回溯法优化版）
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
// 使用回溯法 + 字节切片路径，原地修改避免中间字符串频繁分配。
// 时间复杂度: O(3^m * 4^n * L)  空间复杂度: O(L)
// 其中 m 为对应 3 个字母的数字个数，n 为对应 4 个字母的数字个数，L = len(digits)
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	// 数字 '2'-'9' 对应的字母映射
	mapping := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	// 预计算结果容量，避免切片扩容带来的额外分配
	total := 1
	for i := 0; i < len(digits); i++ {
		total *= len(mapping[digits[i]-'2'])
	}

	res := make([]string, 0, total)
	path := make([]byte, len(digits))

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(digits) {
			res = append(res, string(path))
			return
		}
		letters := mapping[digits[idx]-'2']
		for i := 0; i < len(letters); i++ {
			path[idx] = letters[i]
			dfs(idx + 1)
		}
	}

	dfs(0)
	return res
}

// LetterCombinationsIterative 电话号码的字母组合（迭代笛卡尔积版）
// 逐层构建结果，通过不断的切片追加与字符串拼接生成所有组合。
// 时间复杂度: O(3^m * 4^n * L)  空间复杂度: O(3^m * 4^n * L)
func LetterCombinationsIterative(digits string) []string {
	if len(digits) < 1 {
		return nil
	}
	s := str(digits[0])

	for i := 1; i < len(digits); i++ {
		s = combinations(s, digits[i])
	}
	return s
}

func str(c byte) []string {
	n := int(c - '0')
	idx := (n - 2) * 3
	if n >= 8 {
		idx = idx + 1
	}
	aCode := int('a')
	s := []string{
		string(byte(aCode + idx)),
		string(byte(aCode + idx + 1)),
		string(byte(aCode + idx + 2)),
	}

	if n == 7 || n == 9 {
		s = append(s, string(byte(aCode+idx+3)))
	}
	return s
}

func combinations(strs []string, c byte) []string {
	strs1 := str(c)
	ret := []string{}
	for _, str := range strs {
		for _, str1 := range strs1 {
			ret = append(ret, str+str1)
		}
	}
	return ret
}
