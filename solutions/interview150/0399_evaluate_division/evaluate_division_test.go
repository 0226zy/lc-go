package evaluatedivision

import (
	"math"
	"testing"
)

// floatSlicesEqual 比较两个浮点切片，误差不超过 eps
func floatSlicesEqual(a, b []float64, eps float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Abs(a[i]-b[i]) > eps {
			return false
		}
	}
	return true
}

func TestCalcEquation(t *testing.T) {
	tests := []struct {
		name      string
		equations [][]string
		values    []float64
		queries   [][]string
		want      []float64
	}{
		// LeetCode 官方示例
		{
			"示例1",
			[][]string{{"a", "b"}, {"b", "c"}},
			[]float64{2.0, 3.0},
			[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			[]float64{6.0, 0.5, -1.0, 1.0, -1.0},
		},
		{
			"示例2",
			[][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
			[]float64{1.5, 2.5, 5.0},
			[][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
			[]float64{3.75, 0.4, 5.0, 0.2},
		},
		{
			"示例3",
			[][]string{{"a", "b"}},
			[]float64{0.5},
			[][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}},
			[]float64{0.5, 2.0, -1.0, -1.0},
		},

		// 边界：自除
		{
			"自除",
			[][]string{{"a", "b"}},
			[]float64{2.0},
			[][]string{{"a", "a"}},
			[]float64{1.0},
		},

		// 边界：单个方程链
		{
			"三个变量的链",
			[][]string{{"x1", "x2"}, {"x2", "x3"}, {"x3", "x4"}, {"x4", "x5"}},
			[]float64{3.0, 4.0, 5.0, 6.0},
			[][]string{{"x1", "x5"}, {"x5", "x1"}, {"x2", "x4"}},
			[]float64{360.0, 1.0 / 360.0, 20.0},
		},

		// 边界：变量不存在
		{
			"查询未知变量",
			[][]string{{"a", "b"}},
			[]float64{1.0},
			[][]string{{"a", "z"}, {"z", "a"}, {"z", "z"}},
			[]float64{-1.0, -1.0, -1.0},
		},

		// 边界：不连通的两个分量
		{
			"两个不连通分量",
			[][]string{{"a", "b"}, {"c", "d"}},
			[]float64{2.0, 3.0},
			[][]string{{"a", "d"}},
			[]float64{-1.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalcEquation(tt.equations, tt.values, tt.queries)
			if !floatSlicesEqual(got, tt.want, 1e-5) {
				t.Errorf("CalcEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCalcEquation(b *testing.B) {
	// 10 个变量的链：a/b=2, b/c=2, ..., i/j=2
	equations := [][]string{{"a", "b"}, {"b", "c"}, {"c", "d"}, {"d", "e"},
		{"e", "f"}, {"f", "g"}, {"g", "h"}, {"h", "i"}, {"i", "j"}}
	values := []float64{2, 2, 2, 2, 2, 2, 2, 2, 2}
	queries := [][]string{{"a", "j"}, {"j", "a"}, {"a", "a"}, {"a", "x"}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CalcEquation(equations, values, queries)
	}
}
