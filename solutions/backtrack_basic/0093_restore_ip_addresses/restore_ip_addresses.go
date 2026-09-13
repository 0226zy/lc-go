package restoreipaddresses

import (
	"strconv"
	"strings"
)

// RestoreIpAddresses 复原 IP 地址
// 给定一个只包含数字的字符串 s，在 s 中插入 '.' 将其切分为 4 段，
// 返回所有可能的有效 IP 地址（每段 0~255，长度 1~3，无前导零）。
// 时间复杂度: O(3^4 · n) 切分方案数不超过 81 种  空间复杂度: O(n) 递归栈与路径（不计输出）
func RestoreIpAddresses(s string) []string {
	var result []string
	var path []string
	n := len(s)

	var backtrack func(start, cnt int)
	backtrack = func(start, cnt int) {
		// 已切出 4 段：恰好用完整个字符串才是合法方案
		if cnt == 4 {
			if start == n {
				result = append(result, strings.Join(path, "."))
			}
			return
		}
		// 剩余字符数剪枝：不足以切完剩余段，或超出剩余段数的最大容量
		remain := n - start
		if remain < 4-cnt || remain > (4-cnt)*3 {
			return
		}
		// 枚举当前段长度 1~3
		for l := 1; l <= 3 && start+l <= n; l++ {
			seg := s[start : start+l]
			if l > 1 && s[start] == '0' {
				break // 前导零不合法，更长的段同样以 '0' 开头，直接剪掉
			}
			if val, _ := strconv.Atoi(seg); val > 255 {
				break // 数值超界，更长的段只会更大
			}
			path = append(path, seg)
			backtrack(start+l, cnt+1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, 0)
	return result
}
