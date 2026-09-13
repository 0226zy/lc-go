package lettercombinations

// 数字到字母的映射表，索引即数字，0 和 1 不对应任何字母
var phoneMap = [10]string{
	"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
}

// LetterCombinations 电话号码的字母组合
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合，
// 数字到字母的映射与电话按键相同，答案可以按任意顺序返回。
// 时间复杂度: O(4^n * n)  空间复杂度: O(n) 递归栈深度
func LetterCombinations(digits string) []string {
	result := []string{}
	if len(digits) == 0 {
		return result
	}
	path := make([]byte, len(digits))
	var backtrack func(index int)
	backtrack = func(index int) {
		// 结束条件：所有数字都选完了
		if index == len(digits) {
			result = append(result, string(path))
			return
		}
		letters := phoneMap[digits[index]-'0']
		for i := 0; i < len(letters); i++ {
			path[index] = letters[i] // 做选择
			backtrack(index + 1)
			// path[index] 会被下一轮循环覆盖，无需显式撤销
		}
	}
	backtrack(0)
	return result
}
