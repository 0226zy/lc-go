package simplifypath

import (
	"strings"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		// LeetCode 官方示例
		{"示例1: 末尾多余斜杠", "/home/", "/home"},
		{"示例2: 根目录上一级", "/../", "/"},
		{"示例3: 连续斜杠", "/home//foo/", "/home/foo"},

		// 边界：基本路径
		{"单个目录", "/a", "/a"},
		{"根目录", "/", "/"},
		{"仅多级斜杠", "///", "/"},

		// 边界：点与双点
		{"当前目录无效果", "/a/./b", "/a/b"},
		{"双点回退", "/a/b/..", "/a"},
		{"双点连续回退到根", "/a/b/../../", "/"},
		{"根目录双点无效果", "/..", "/"},
		{"根目录连续双点无效果", "/../../", "/"},
		{"当前目录与双点混合", "/a/./b/../../c/", "/c"},

		// 边界：多个点当作普通文件名
		{"三个点是文件名", "/...", "/..."},
		{"点与双点拼接是文件名", "/..hidden", "/..hidden"},
		{"点后跟字符是文件名", "/.a", "/.a"},

		// 边界：长路径与下划线
		{"多段长目录名", "/home/user/documents/./pictures/.././music", "/home/user/documents/music"},
		{"含下划线的目录名", "/_a_/_b_", "/_a_/_b_"},
		{"数字目录名", "/1/2/3/..", "/1/2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimplifyPath(tt.path); got != tt.want {
				t.Errorf("SimplifyPath(%q) = %q, want %q", tt.path, got, tt.want)
			}
		})
	}
}

func BenchmarkSimplifyPath(b *testing.B) {
	benchmarks := []struct {
		name string
		path string
	}{
		{"短路径_含双点", "/a/./b/../../c/"},
		{"长路径_大量当前目录", generatePath(1000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SimplifyPath(bm.path)
			}
		})
	}
}

// generatePath 生成包含 n 层目录的长路径，穿插 "." 段
func generatePath(n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString("/dir")
		sb.WriteString(itoa(i))
		if i%2 == 0 {
			sb.WriteString("/.")
		}
	}
	sb.WriteString("/")
	return sb.String()
}

// itoa 简单的非负整数转字符串，避免在生成器中重复引入 strconv
func itoa(x int) string {
	if x == 0 {
		return "0"
	}
	var buf [20]byte
	i := len(buf)
	for x > 0 {
		i--
		buf[i] = byte('0' + x%10)
		x /= 10
	}
	return string(buf[i:])
}
