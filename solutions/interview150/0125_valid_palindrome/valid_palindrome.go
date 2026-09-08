package validpalindrome

// IsPalindrome 验证回文串
// 给定一个字符串 s，只考虑其中的字母和数字字符（忽略字母大小写），判断它是否是回文串；空字符串视为回文串。
// 时间复杂度: O(n) 双指针最多各走一遍字符串  空间复杂度: O(1) 只使用常数变量，原地判断
func IsPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		// 左指针跳过非字母数字字符
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		// 右指针跳过非字母数字字符
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}
		// 大小写不敏感比较两个有效字符
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

// isAlphanumeric 判断字节是否为字母或数字（s 由可打印 ASCII 字符组成，按字节处理即可）
func isAlphanumeric(c byte) bool {
	return c >= '0' && c <= '9' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

// toLower 将大写字母转成小写，其余字符原样返回
func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}
