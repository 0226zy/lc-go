package ternaryexpressionparser

import (
	"strings"
	"testing"
)

func TestParseTernary(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       string
	}{
		// LeetCode 官方示例
		{"示例1: 简单条件取真分支", "T?2:3", "2"},
		{"示例2: 外层取假分支再嵌套", "F?1:T?4:5", "4"},
		{"示例3: 嵌套在真分支中", "T?T?F:5:3", "F"},

		// 边界：单个值
		{"单个T", "T", "T"},
		{"单个F", "F", "F"},

		// 典型场景
		{"条件为F取假分支", "F?7:9", "9"},
		{"真分支为数字的假分支嵌套", "T?1:F?2:3", "1"},
		{"多层嵌套取最深真分支", "T?T?T?1:2:3:4", "1"},
		{"嵌套结果回传外层", "F?T?1:2:F?3:4", "4"},
		{"条件字符来自嵌套求值", "T?F?1:2:3", "2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseTernary(tt.expression); got != tt.want {
				t.Errorf("ParseTernary(%q) = %q, want %q", tt.expression, got, tt.want)
			}
		})
	}
}

func BenchmarkParseTernary(b *testing.B) {
	// 构造深层右嵌套表达式: T?T?T?...?1:2:3:4...
	var sb strings.Builder
	depth := 2000
	for i := 0; i < depth; i++ {
		sb.WriteString("T?")
	}
	sb.WriteByte('1')
	for i := 0; i < depth; i++ {
		sb.WriteString(":2")
	}
	long := sb.String()

	benchmarks := []struct {
		name       string
		expression string
	}{
		{"len=5", "T?2:3"},
		{"len=13", "F?1:T?4:5"},
		{"深层嵌套", long},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ParseTernary(bm.expression)
			}
		})
	}
}
