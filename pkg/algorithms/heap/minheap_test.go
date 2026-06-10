package heap

import "testing"

func TestNewMinHeap(t *testing.T) {
	t.Run("创建空堆", func(t *testing.T) {
		h := NewMinHeap()
		if !h.IsEmpty() {
			t.Error("新堆应该为空")
		}
		if h.Size() != 0 {
			t.Errorf("Size() = %v, want 0", h.Size())
		}
	})
}

func TestNewMinHeapFromSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		wantPeek int
		wantSize int
	}{
		{
			name:     "空切片",
			input:    []int{},
			wantSize: 0,
		},
		{
			name:     "单元素",
			input:    []int{5},
			wantPeek: 5,
			wantSize: 1,
		},
		{
			name:     "有序数组",
			input:    []int{1, 2, 3, 4, 5},
			wantPeek: 1,
			wantSize: 5,
		},
		{
			name:     "逆序数组",
			input:    []int{5, 4, 3, 2, 1},
			wantPeek: 1,
			wantSize: 5,
		},
		{
			name:     "随机数组",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6},
			wantPeek: 1,
			wantSize: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewMinHeapFromSlice(tt.input)
			if h.Size() != tt.wantSize {
				t.Errorf("Size() = %v, want %v", h.Size(), tt.wantSize)
			}
			if tt.wantSize > 0 && h.Peek() != tt.wantPeek {
				t.Errorf("Peek() = %v, want %v", h.Peek(), tt.wantPeek)
			}
		})
	}
}

func TestMinHeap_Push(t *testing.T) {
	t.Run("逐个插入", func(t *testing.T) {
		h := NewMinHeap()
		h.Push(5)
		if h.Peek() != 5 {
			t.Errorf("Peek() = %v, want 5", h.Peek())
		}
		h.Push(3)
		if h.Peek() != 3 {
			t.Errorf("Peek() = %v, want 3", h.Peek())
		}
		h.Push(1)
		if h.Peek() != 1 {
			t.Errorf("Peek() = %v, want 1", h.Peek())
		}
		h.Push(4)
		if h.Peek() != 1 {
			t.Errorf("Peek() = %v, want 1", h.Peek())
		}
		if h.Size() != 4 {
			t.Errorf("Size() = %v, want 4", h.Size())
		}
	})

	t.Run("插入负数", func(t *testing.T) {
		h := NewMinHeap()
		h.Push(5)
		h.Push(-3)
		h.Push(0)
		if h.Peek() != -3 {
			t.Errorf("Peek() = %v, want -3", h.Peek())
		}
	})
}

func TestMinHeap_Pop(t *testing.T) {
	t.Run("弹出顺序", func(t *testing.T) {
		h := NewMinHeap()
		h.Push(5)
		h.Push(3)
		h.Push(1)
		h.Push(4)
		h.Push(2)

		expected := []int{1, 2, 3, 4, 5}
		for i, want := range expected {
			got := h.Pop()
			if got != want {
				t.Errorf("Pop() #%d = %v, want %v", i, got, want)
			}
		}
		if !h.IsEmpty() {
			t.Error("弹出所有元素后堆应该为空")
		}
	})

	t.Run("空堆弹出 panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("空堆 Pop 应该 panic")
			}
		}()
		h := NewMinHeap()
		h.Pop()
	})
}

func TestMinHeap_Peek(t *testing.T) {
	t.Run("查看堆顶不移除", func(t *testing.T) {
		h := NewMinHeap()
		h.Push(3)
		h.Push(1)
		h.Push(2)
		if h.Peek() != 1 {
			t.Errorf("Peek() = %v, want 1", h.Peek())
		}
		if h.Size() != 3 {
			t.Errorf("Peek 不应改变大小，Size() = %v, want 3", h.Size())
		}
	})

	t.Run("空堆 Peek panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("空堆 Peek 应该 panic")
			}
		}()
		h := NewMinHeap()
		h.Peek()
	})
}

func TestMinHeap_MixedOperations(t *testing.T) {
	t.Run("混合 Push 和 Pop", func(t *testing.T) {
		h := NewMinHeap()
		h.Push(5)
		h.Push(3)
		h.Push(7)

		// 弹出 3
		if got := h.Pop(); got != 3 {
			t.Errorf("Pop() = %v, want 3", got)
		}

		h.Push(1)
		h.Push(2)

		// 弹出 1
		if got := h.Pop(); got != 1 {
			t.Errorf("Pop() = %v, want 1", got)
		}

		// 剩余：2, 5, 7
		if got := h.Pop(); got != 2 {
			t.Errorf("Pop() = %v, want 2", got)
		}
		if got := h.Pop(); got != 5 {
			t.Errorf("Pop() = %v, want 5", got)
		}
		if got := h.Pop(); got != 7 {
			t.Errorf("Pop() = %v, want 7", got)
		}
	})

	t.Run("从切片构建后弹出", func(t *testing.T) {
		h := NewMinHeapFromSlice([]int{9, 5, 2, 7, 1, 6})
		expected := []int{1, 2, 5, 6, 7, 9}
		for i, want := range expected {
			if got := h.Pop(); got != want {
				t.Errorf("Pop() #%d = %v, want %v", i, got, want)
			}
		}
	})

	t.Run("重复元素", func(t *testing.T) {
		h := NewMinHeap()
		h.Push(3)
		h.Push(1)
		h.Push(1)
		h.Push(2)
		expected := []int{1, 1, 2, 3}
		for i, want := range expected {
			if got := h.Pop(); got != want {
				t.Errorf("Pop() #%d = %v, want %v", i, got, want)
			}
		}
	})
}

func TestMinHeap_Data(t *testing.T) {
	t.Run("Data 返回副本", func(t *testing.T) {
		h := NewMinHeapFromSlice([]int{3, 1, 2})
		data := h.Data()
		data[0] = 999
		if h.Peek() == 999 {
			t.Error("Data() 应返回副本，修改副本不应影响原堆")
		}
	})
}

func BenchmarkMinHeap_Push(b *testing.B) {
	h := NewMinHeap()
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		h.Push(i)
	}
}

func BenchmarkMinHeap_Pop(b *testing.B) {
	h := NewMinHeap()
	for i := 0; i < b.N; i++ {
		h.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Pop()
	}
}

func BenchmarkMinHeap_FromSlice(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = 1000 - i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewMinHeapFromSlice(nums)
	}
}
