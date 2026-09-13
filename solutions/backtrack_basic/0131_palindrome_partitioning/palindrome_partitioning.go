package palindromepartitioning

// Partition 分割回文串
// 给定字符串 s，将 s 分割成若干子串，使每个子串都是回文串，返回所有可能的分割方案。
// 时间复杂度: O(n × 2^n)  n 为字符串长度，每个位置都可切/不切  空间复杂度: O(n) 递归栈深度（不计输出）
func Partition(s string) [][]string {
	var result [][]string
	path := []string{}

	// backtrack 表示当前准备从 s[start:] 中切出下一个回文子串
	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(s) {
			// 已经切到字符串末尾，当前 path 是一种完整分割方案，拷贝后收集
			partition := make([]string, len(path))
			copy(partition, path)
			result = append(result, partition)
			return
		}
		// 枚举下一个子串的结束位置 end，子串为 s[start:end+1]
		for end := start; end < len(s); end++ {
			if !isPalindrome(s, start, end) {
				continue // 不是回文串不能作为一段，直接跳过
			}
			path = append(path, s[start:end+1])
			backtrack(end + 1)
			path = path[:len(path)-1] // 回溯：撤销本次切割
		}
	}
	backtrack(0)
	return result
}

// isPalindrome 双指针判断 s[left:right+1] 是否为回文串
func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
