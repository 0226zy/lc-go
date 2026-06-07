package sort

import (
	"math"
	"testing"
)

// TestMinHeapInt 测试整数类型的最小堆
func TestMinHeapInt(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "普通情况",
			nums:     []int{5, 3, 8, 1, 2},
			expected: []int{1, 2, 3, 5, 8},
		},
		{
			name:     "已排序",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "逆序",
			nums:     []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "有重复元素",
			nums:     []int{3, 1, 3, 2, 1},
			expected: []int{1, 1, 2, 3, 3},
		},
		{
			name:     "单个元素",
			nums:     []int{42},
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewMinHeap(func(a, b int) bool {
				return a < b
			})

			// 插入所有元素
			for _, num := range tt.nums {
				h.Push(num)
			}

			// 依次弹出，验证顺序
			for _, expected := range tt.expected {
				actual := h.Pop()
				if actual != expected {
					t.Errorf("期望 %d，实际 %d", expected, actual)
				}
			}

			// 堆应该为空
			if h.Len() != 0 {
				t.Errorf("堆应该为空，实际长度 %d", h.Len())
			}
		})
	}
}

// TestMinHeapString 测试字符串类型的最小堆
func TestMinHeapString(t *testing.T) {
	h := NewMinHeap(func(a, b string) bool {
		return a < b
	})

	words := []string{"banana", "apple", "cherry", "date"}
	for _, word := range words {
		h.Push(word)
	}

	expected := []string{"apple", "banana", "cherry", "date"}
	for _, exp := range expected {
		actual := h.Pop()
		if actual != exp {
			t.Errorf("期望 %s，实际 %s", exp, actual)
		}
	}
}

// TestMinHeapPeek 测试 Peek 方法
func TestMinHeapPeek(t *testing.T) {
	h := NewMinHeap(func(a, b int) bool {
		return a < b
	})

	// 空堆 Peek 应该返回零值
	zero := h.Peek()
	if zero != 0 {
		t.Errorf("空堆 Peek 应该返回 0，实际 %d", zero)
	}

	// 插入元素后 Peek 应该返回最小值
	h.Push(5)
	h.Push(3)
	h.Push(8)

	minVal := h.Peek()
	if minVal != 3 {
		t.Errorf("Peek 应该返回 3，实际 %d", minVal)
	}

	// Peek 不应该移除元素
	if h.Len() != 3 {
		t.Errorf("Peek 不应该移除元素，堆大小应该为 3，实际 %d", h.Len())
	}
}

// TestMinHeapPopEmpty 测试从空堆弹出
func TestMinHeapPopEmpty(t *testing.T) {
	h := NewMinHeap(func(a, b int) bool {
		return a < b
	})

	// 从空堆弹出应该返回零值
	zero := h.Pop()
	if zero != 0 {
		t.Errorf("从空堆弹出应该返回 0，实际 %d", zero)
	}
}

// TestMinHeapLen 测试 Len 方法
func TestMinHeapLen(t *testing.T) {
	h := NewMinHeap(func(a, b int) bool {
		return a < b
	})

	if h.Len() != 0 {
		t.Errorf("新创建的堆长度应该为 0，实际 %d", h.Len())
	}

	h.Push(1)
	h.Push(2)
	if h.Len() != 2 {
		t.Errorf("插入两个元素后长度应该为 2，实际 %d", h.Len())
	}

	h.Pop()
	if h.Len() != 1 {
		t.Errorf("弹出一个元素后长度应该为 1，实际 %d", h.Len())
	}

	h.Pop()
	if h.Len() != 0 {
		t.Errorf("弹出所有元素后长度应该为 0，实际 %d", h.Len())
	}
}

// TestMinHeapComplexStruct 测试复杂结构体
func TestMinHeapComplexStruct(t *testing.T) {
	type Point struct {
		X, Y int
	}

	// 按 X 坐标排序
	h := NewMinHeap(func(a, b Point) bool {
		return a.X < b.X
	})

	points := []Point{{5, 10}, {1, 20}, {3, 30}}
	for _, p := range points {
		h.Push(p)
	}

	expected := []Point{{1, 20}, {3, 30}, {5, 10}}
	for _, exp := range expected {
		actual := h.Pop()
		if actual != exp {
			t.Errorf("期望 %v，实际 %v", exp, actual)
		}
	}
}

// TestMinHeapWithFloat 测试浮点数类型的最小堆
func TestMinHeapWithFloat(t *testing.T) {
	eps := 1e-9
	h := NewMinHeap(func(a, b float64) bool {
		return a < b
	})

	values := []float64{3.14, 2.71, 1.41, 0.99}
	for _, v := range values {
		h.Push(v)
	}

	expected := []float64{0.99, 1.41, 2.71, 3.14}
	for _, exp := range expected {
		actual := h.Pop()
		if math.Abs(actual-exp) > eps {
			t.Errorf("期望 %f，实际 %f", exp, actual)
		}
	}
}

// BenchmarkMinHeapPushPop 性能测试：Push 和 Pop 操作
func BenchmarkMinHeapPushPop(b *testing.B) {
	h := NewMinHeap(func(a, b int) bool {
		return a < b
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Push(i)
		h.Pop()
	}
}

// BenchmarkMinHeapLarge 性能测试：大量数据
func BenchmarkMinHeapLarge(b *testing.B) {
	for n := 1000; n <= 100000; n *= 10 {
		b.Run("", func(b *testing.B) {
			h := NewMinHeap(func(a, b int) bool {
				return a < b
			})

			// 预填充数据
			for i := 0; i < n; i++ {
				h.Push(i)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				h.Pop()
				h.Push(i)
			}
		})
	}
}
