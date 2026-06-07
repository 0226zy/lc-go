package sort

import (
	"math"
	"testing"
)

// ============================================================
// 最小堆测试（MinHeap）
// ============================================================

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

// ============================================================
// 最大堆测试（MaxHeap）
// ============================================================

// TestMaxHeapInt 测试整数类型的最大堆
func TestMaxHeapInt(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "普通情况",
			nums:     []int{5, 3, 8, 1, 2},
			expected: []int{8, 5, 3, 2, 1},
		},
		{
			name:     "已排序（逆序插入）",
			nums:     []int{5, 4, 3, 2, 1},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "有重复元素",
			nums:     []int{3, 1, 3, 2, 1},
			expected: []int{3, 3, 2, 1, 1},
		},
		{
			name:     "单个元素",
			nums:     []int{42},
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewMaxHeap(func(a, b int) bool {
				return a > b
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

// TestMaxHeapString 测试字符串类型的最大堆
func TestMaxHeapString(t *testing.T) {
	h := NewMaxHeap(func(a, b string) bool {
		return a > b
	})

	words := []string{"banana", "apple", "cherry", "date"}
	for _, word := range words {
		h.Push(word)
	}

	expected := []string{"date", "cherry", "banana", "apple"}
	for _, exp := range expected {
		actual := h.Pop()
		if actual != exp {
			t.Errorf("期望 %s，实际 %s", exp, actual)
		}
	}
}

// ============================================================
// 通用方法测试
// ============================================================

// TestHeapPeek 测试 Peek 方法（最小堆和最大堆）
func TestHeapPeek(t *testing.T) {
	// 最小堆的 Peek
	minH := NewMinHeap(func(a, b int) bool { return a < b })
	if zero := minH.Peek(); zero != 0 {
		t.Errorf("空最小堆 Peek 应该返回 0，实际 %d", zero)
	}
	minH.Push(5)
	minH.Push(3)
	minH.Push(8)
	if val := minH.Peek(); val != 3 {
		t.Errorf("最小堆 Peek 应该返回 3，实际 %d", val)
	}
	if minH.Len() != 3 {
		t.Errorf("Peek 不应该移除元素，堆大小应该为 3，实际 %d", minH.Len())
	}

	// 最大堆的 Peek
	maxH := NewMaxHeap(func(a, b int) bool { return a > b })
	maxH.Push(5)
	maxH.Push(3)
	maxH.Push(8)
	if val := maxH.Peek(); val != 8 {
		t.Errorf("最大堆 Peek 应该返回 8，实际 %d", val)
	}
}

// TestHeapPopEmpty 测试从空堆弹出
func TestHeapPopEmpty(t *testing.T) {
	h := NewMinHeap(func(a, b int) bool { return a < b })
	zero := h.Pop()
	if zero != 0 {
		t.Errorf("从空堆弹出应该返回 0，实际 %d", zero)
	}
}

// TestHeapLen 测试 Len 方法
func TestHeapLen(t *testing.T) {
	h := NewMinHeap(func(a, b int) bool { return a < b })
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

// TestHeapComplexStruct 测试复杂结构体
func TestHeapComplexStruct(t *testing.T) {
	type Point struct {
		X, Y int
	}

	// 按 X 坐标排序的最小堆
	minH := NewMinHeap(func(a, b Point) bool {
		return a.X < b.X
	})
	points := []Point{{5, 10}, {1, 20}, {3, 30}}
	for _, p := range points {
		minH.Push(p)
	}
	expected := []Point{{1, 20}, {3, 30}, {5, 10}}
	for _, exp := range expected {
		actual := minH.Pop()
		if actual != exp {
			t.Errorf("期望 %v，实际 %v", exp, actual)
		}
	}

	// 按 X 坐标排序的最大堆
	maxH := NewMaxHeap(func(a, b Point) bool {
		return a.X > b.X
	})
	for _, p := range points {
		maxH.Push(p)
	}
	expectedMax := []Point{{5, 10}, {3, 30}, {1, 20}}
	for _, exp := range expectedMax {
		actual := maxH.Pop()
		if actual != exp {
			t.Errorf("期望 %v，实际 %v", exp, actual)
		}
	}
}

// TestHeapWithFloat 测试浮点数类型
func TestHeapWithFloat(t *testing.T) {
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

// TestHeapNewHeap 测试 NewHeap 通用构造函数
func TestHeapNewHeap(t *testing.T) {
	// 使用 NewHeap 创建最小堆
	minH := NewHeap(func(a, b int) bool { return a < b })
	minH.Push(3)
	minH.Push(1)
	minH.Push(2)
	if val := minH.Pop(); val != 1 {
		t.Errorf("最小堆应该弹出 1，实际 %d", val)
	}

	// 使用 NewHeap 创建最大堆
	maxH := NewHeap(func(a, b int) bool { return a > b })
	maxH.Push(3)
	maxH.Push(1)
	maxH.Push(2)
	if val := maxH.Pop(); val != 3 {
		t.Errorf("最大堆应该弹出 3，实际 %d", val)
	}
}

// ============================================================
// 性能测试
// ============================================================

// BenchmarkHeapPushPop 性能测试：Push 和 Pop 操作
func BenchmarkHeapPushPop(b *testing.B) {
	h := NewMinHeap(func(a, b int) bool {
		return a < b
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Push(i)
		h.Pop()
	}
}

// BenchmarkHeapLarge 性能测试：大量数据
func BenchmarkHeapLarge(b *testing.B) {
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
