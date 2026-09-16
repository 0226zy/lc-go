package reversewordsinastringii

// ReverseWords 反转字符串中的单词 II
// 给定字符数组 s，原地反转其中单词的顺序（单词由单个空格分隔，无首尾空格）。
// 时间复杂度: O(n) 每个字符被交换常数次  空间复杂度: O(1) 原地修改
func ReverseWords(s []byte) {
	// 反转 s 的子区间 [i, j]
	reverse := func(i, j int) {
		for i < j {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
	}

	n := len(s)
	// 第一步：整体反转，单词顺序就位但单词内部被反转
	reverse(0, n-1)

	// 第二步：逐个单词反转，恢复单词内部的字符顺序
	start := 0
	for i := 0; i <= n; i++ {
		// 遇到空格或越界，说明 [start, i-1] 是一个完整单词
		if i == n || s[i] == ' ' {
			reverse(start, i-1)
			start = i + 1
		}
	}
}
