package designcompressedstringiterator

import "testing"

func TestStringIterator(t *testing.T) {
	tests := []struct {
		name       string
		compressed string
		ops        []string // 操作序列："next" 或 "hasNext"
		want       []string // 期望输出序列，next 用单字符字符串表示
	}{
		{
			"LeetCode官方示例",
			"L1e2t1C1o1d1e1",
			[]string{"next", "next", "next", "next", "next", "next", "hasNext", "next", "hasNext"},
			[]string{"L", "e", "e", "t", "C", "o", "true", "d", "true"},
		},
		{
			"单个字符重复10次",
			"x10",
			[]string{"next", "hasNext", "next"},
			[]string{"x", "true", "x"},
		},
		{
			"耗尽后next返回空格",
			"a1",
			[]string{"next", "next", "hasNext"},
			[]string{"a", " ", "false"},
		},
		{
			"空字符串",
			"",
			[]string{"hasNext", "next"},
			[]string{"false", " "},
		},
		{
			"大重复次数惰性解压",
			"b1000000000",
			[]string{"next", "next", "hasNext"},
			[]string{"b", "b", "true"},
		},
		{
			"多段混合大小写",
			"a2B3c1",
			[]string{"next", "next", "next", "next", "next", "next", "next"},
			[]string{"a", "a", "B", "B", "B", "c", " "},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := Constructor(tt.compressed)
			if len(tt.ops) != len(tt.want) {
				t.Fatalf("操作数与期望数不一致: %d vs %d", len(tt.ops), len(tt.want))
			}
			for i, op := range tt.ops {
				switch op {
				case "next":
					if got := string(it.Next()); got != tt.want[i] {
						t.Errorf("第 %d 步 next() = %q, want %q", i, got, tt.want[i])
					}
				case "hasNext":
					got := "false"
					if it.HasNext() {
						got = "true"
					}
					if got != tt.want[i] {
						t.Errorf("第 %d 步 hasNext() = %q, want %q", i, got, tt.want[i])
					}
				default:
					t.Fatalf("未知操作 %q", op)
				}
			}
		})
	}
}

func BenchmarkStringIterator(b *testing.B) {
	// 构造包含 100 段的压缩字符串，每段重复次数为 10^9
	compressed := ""
	for c := byte('a'); c <= byte('z'); c++ {
		compressed += string(c) + "1000000000"
	}
	for i := 0; i < b.N; i++ {
		it := Constructor(compressed)
		// 模拟真实使用：交错调用 hasNext 与 next
		for j := 0; j < 1000; j++ {
			if it.HasNext() {
				it.Next()
			}
		}
	}
}
