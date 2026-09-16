package designinmemoryfilesystem

import (
	"strings"
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestFileSystem(t *testing.T) {
	tests := []struct {
		name string
		ops  []string   // 操作序列："ls" / "mkdir" / "addContentToFile" / "readContentFromFile"
		args [][]string // 每个操作的参数
		want []string   // 期望输出：ls 用逗号连接，readContentFromFile 为内容，mkdir/add 用 "null"
	}{
		{
			"LeetCode官方示例",
			[]string{"ls", "mkdir", "addContentToFile", "ls", "readContentFromFile"},
			[][]string{{"/"}, {"/a/b/c"}, {"/a/b/c/d", "hello"}, {"/"}, {"/a/b/c/d"}},
			[]string{"", "null", "null", "a", "hello"},
		},
		{
			"空目录与尾部斜杠",
			[]string{"ls", "mkdir", "ls", "ls", "ls"},
			[][]string{{"/"}, {"/zijzllb"}, {"/"}, {"/zijzllb"}, {"/zijzllb/"}},
			[]string{"", "null", "zijzllb", "", ""},
		},
		{
			"多次追加文件内容",
			[]string{"addContentToFile", "addContentToFile", "readContentFromFile"},
			[][]string{{"/file", "hello"}, {"/file", "world"}, {"/file"}},
			[]string{"null", "null", "helloworld"},
		},
		{
			"ls文件返回自身名称",
			[]string{"addContentToFile", "ls", "ls"},
			[][]string{{"/a/b/c", "x"}, {"/a/b/c"}, {"/a/b"}},
			[]string{"null", "c", "c"},
		},
		{
			"ls结果按字典序排列",
			[]string{"mkdir", "mkdir", "mkdir", "addContentToFile", "ls"},
			[][]string{{"/m"}, {"/a"}, {"/z"}, {"/b", "1"}, {"/"}},
			[]string{"null", "null", "null", "null", "a,b,m,z"},
		},
		{
			"重复mkdir已有目录无副作用",
			[]string{"mkdir", "addContentToFile", "mkdir", "ls", "readContentFromFile"},
			[][]string{{"/a"}, {"/a/f", "data"}, {"/a"}, {"/a"}, {"/a/f"}},
			[]string{"null", "null", "null", "f", "data"},
		},
		{
			"深层路径",
			[]string{"mkdir", "addContentToFile", "readContentFromFile", "ls"},
			[][]string{{"/a/b/c/d/e"}, {"/a/b/c/d/e/f", "deep"}, {"/a/b/c/d/e/f"}, {"/a/b/c/d/e"}},
			[]string{"null", "null", "deep", "f"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := Constructor()
			if len(tt.ops) != len(tt.args) || len(tt.ops) != len(tt.want) {
				t.Fatalf("操作/参数/期望数量不一致: %d/%d/%d", len(tt.ops), len(tt.args), len(tt.want))
			}
			for i, op := range tt.ops {
				switch op {
				case "ls":
					got := fs.Ls(tt.args[i][0])
					want := []string{}
					if tt.want[i] != "" {
						want = strings.Split(tt.want[i], ",")
					}
					if !utils.EqualStringSlice(got, want) {
						t.Errorf("第 %d 步 ls(%q) = %v, want %v", i, tt.args[i][0], got, want)
					}
				case "mkdir":
					fs.Mkdir(tt.args[i][0])
				case "addContentToFile":
					fs.AddContentToFile(tt.args[i][0], tt.args[i][1])
				case "readContentFromFile":
					if got := fs.ReadContentFromFile(tt.args[i][0]); got != tt.want[i] {
						t.Errorf("第 %d 步 readContentFromFile(%q) = %q, want %q", i, tt.args[i][0], got, tt.want[i])
					}
				default:
					t.Fatalf("未知操作 %q", op)
				}
			}
		})
	}
}

func BenchmarkFileSystem(b *testing.B) {
	b.Run("mkdir与ls混合", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fs := Constructor()
			for j := 0; j < 50; j++ {
				fs.Mkdir("/a/b/c/d/e/f/g/h")
				fs.AddContentToFile("/a/b/c/d/e/f/g/h/file", "hello")
			}
			fs.Ls("/")
			fs.ReadContentFromFile("/a/b/c/d/e/f/g/h/file")
		}
	})
}
