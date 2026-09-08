package simplifypath

import "strings"

// SimplifyPath 简化路径
// 给定 Unix 风格的绝对路径 path，返回其规范路径（去掉多余斜杠、处理 . 和 ..）。
// 时间复杂度: O(n) 切分与拼接各遍历一次  空间复杂度: O(n) 栈最多存放所有目录名
func SimplifyPath(path string) string {
	parts := strings.Split(path, "/")
	stack := make([]string, 0, len(parts))
	for _, seg := range parts {
		switch seg {
		case "", ".":
			// 空段（连续斜杠）和当前目录：不产生任何效果
		case "..":
			// 回到上一级：弹出最近进入的目录（根目录下无效果）
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			// 进入该目录
			stack = append(stack, seg)
		}
	}

	// 按顺序拼接为规范路径，栈空时为根目录 "/"
	if len(stack) == 0 {
		return "/"
	}
	var sb strings.Builder
	sb.Grow(len(path))
	for _, seg := range stack {
		sb.WriteString("/")
		sb.WriteString(seg)
	}
	return sb.String()
}
