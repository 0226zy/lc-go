package minstack

import "testing"

func TestMinStack(t *testing.T) {
	t.Run("官方示例", func(t *testing.T) {
		stack := Constructor()
		stack.Push(-2)
		stack.Push(0)
		stack.Push(-3)
		if got := stack.GetMin(); got != -3 {
			t.Errorf("GetMin() = %v, want -3", got)
		}
		stack.Pop()
		if got := stack.Top(); got != 0 {
			t.Errorf("Top() = %v, want 0", got)
		}
		if got := stack.GetMin(); got != -2 {
			t.Errorf("GetMin() = %v, want -2", got)
		}
	})

	t.Run("单元素", func(t *testing.T) {
		stack := Constructor()
		stack.Push(5)
		if got := stack.Top(); got != 5 {
			t.Errorf("Top() = %v, want 5", got)
		}
		if got := stack.GetMin(); got != 5 {
			t.Errorf("GetMin() = %v, want 5", got)
		}
	})

	t.Run("重复最小值", func(t *testing.T) {
		stack := Constructor()
		stack.Push(1)
		stack.Push(1)
		stack.Push(2)
		if got := stack.GetMin(); got != 1 {
			t.Errorf("GetMin() = %v, want 1", got)
		}
		stack.Pop()
		if got := stack.GetMin(); got != 1 {
			t.Errorf("GetMin() after pop = %v, want 1", got)
		}
		stack.Pop()
		if got := stack.GetMin(); got != 1 {
			t.Errorf("GetMin() after second pop = %v, want 1", got)
		}
	})

	t.Run("最小值更新", func(t *testing.T) {
		stack := Constructor()
		stack.Push(3)
		stack.Push(2)
		stack.Push(1)
		if got := stack.GetMin(); got != 1 {
			t.Errorf("GetMin() = %v, want 1", got)
		}
		stack.Pop()
		if got := stack.GetMin(); got != 2 {
			t.Errorf("GetMin() after pop = %v, want 2", got)
		}
	})

	t.Run("负数操作", func(t *testing.T) {
		stack := Constructor()
		stack.Push(-1)
		stack.Push(-2)
		stack.Push(-3)
		if got := stack.GetMin(); got != -3 {
			t.Errorf("GetMin() = %v, want -3", got)
		}
		stack.Pop()
		if got := stack.Top(); got != -2 {
			t.Errorf("Top() = %v, want -2", got)
		}
		if got := stack.GetMin(); got != -2 {
			t.Errorf("GetMin() = %v, want -2", got)
		}
	})

	t.Run("升降交替", func(t *testing.T) {
		stack := Constructor()
		stack.Push(5)
		stack.Push(3)
		stack.Push(4)
		if got := stack.GetMin(); got != 3 {
			t.Errorf("GetMin() = %v, want 3", got)
		}
		stack.Push(2)
		if got := stack.GetMin(); got != 2 {
			t.Errorf("GetMin() = %v, want 2", got)
		}
		stack.Pop()
		if got := stack.GetMin(); got != 3 {
			t.Errorf("GetMin() after pop = %v, want 3", got)
		}
	})
}

func TestMinStackDiff(t *testing.T) {
	t.Run("官方示例", func(t *testing.T) {
		stack := ConstructorDiff()
		stack.Push(-2)
		stack.Push(0)
		stack.Push(-3)
		if got := stack.GetMin(); got != -3 {
			t.Errorf("GetMin() = %v, want -3", got)
		}
		stack.Pop()
		if got := stack.Top(); got != 0 {
			t.Errorf("Top() = %v, want 0", got)
		}
		if got := stack.GetMin(); got != -2 {
			t.Errorf("GetMin() = %v, want -2", got)
		}
	})

	t.Run("单元素", func(t *testing.T) {
		stack := ConstructorDiff()
		stack.Push(5)
		if got := stack.Top(); got != 5 {
			t.Errorf("Top() = %v, want 5", got)
		}
		if got := stack.GetMin(); got != 5 {
			t.Errorf("GetMin() = %v, want 5", got)
		}
	})

	t.Run("重复最小值", func(t *testing.T) {
		stack := ConstructorDiff()
		stack.Push(1)
		stack.Push(1)
		stack.Push(2)
		if got := stack.GetMin(); got != 1 {
			t.Errorf("GetMin() = %v, want 1", got)
		}
		stack.Pop()
		if got := stack.GetMin(); got != 1 {
			t.Errorf("GetMin() after pop = %v, want 1", got)
		}
		stack.Pop()
		if got := stack.GetMin(); got != 1 {
			t.Errorf("GetMin() after second pop = %v, want 1", got)
		}
	})

	t.Run("最小值更新", func(t *testing.T) {
		stack := ConstructorDiff()
		stack.Push(3)
		stack.Push(2)
		stack.Push(1)
		if got := stack.GetMin(); got != 1 {
			t.Errorf("GetMin() = %v, want 1", got)
		}
		stack.Pop()
		if got := stack.GetMin(); got != 2 {
			t.Errorf("GetMin() after pop = %v, want 2", got)
		}
	})

	t.Run("负数操作", func(t *testing.T) {
		stack := ConstructorDiff()
		stack.Push(-1)
		stack.Push(-2)
		stack.Push(-3)
		if got := stack.GetMin(); got != -3 {
			t.Errorf("GetMin() = %v, want -3", got)
		}
		stack.Pop()
		if got := stack.Top(); got != -2 {
			t.Errorf("Top() = %v, want -2", got)
		}
		if got := stack.GetMin(); got != -2 {
			t.Errorf("GetMin() = %v, want -2", got)
		}
	})

	t.Run("升降交替", func(t *testing.T) {
		stack := ConstructorDiff()
		stack.Push(5)
		stack.Push(3)
		stack.Push(4)
		if got := stack.GetMin(); got != 3 {
			t.Errorf("GetMin() = %v, want 3", got)
		}
		stack.Push(2)
		if got := stack.GetMin(); got != 2 {
			t.Errorf("GetMin() = %v, want 2", got)
		}
		stack.Pop()
		if got := stack.GetMin(); got != 3 {
			t.Errorf("GetMin() after pop = %v, want 3", got)
		}
	})
}

func BenchmarkMinStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := Constructor()
		stack.Push(3)
		stack.Push(2)
		stack.Push(1)
		stack.GetMin()
		stack.Pop()
		stack.Top()
		stack.GetMin()
	}
}

func BenchmarkMinStackDiff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := ConstructorDiff()
		stack.Push(3)
		stack.Push(2)
		stack.Push(1)
		stack.GetMin()
		stack.Pop()
		stack.Top()
		stack.GetMin()
	}
}
