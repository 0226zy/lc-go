package minstack

import "testing"

func TestMinStack(t *testing.T) {
	t.Run("官方示例调用序列", func(t *testing.T) {
		// ["MinStack","push","push","push","getMin","pop","top","getMin"]
		// [[],[-2],[0],[-3],[],[],[],[]]
		s := Constructor()
		s.Push(-2)
		s.Push(0)
		s.Push(-3)
		if got := s.GetMin(); got != -3 {
			t.Errorf("GetMin() = %d, want -3", got)
		}
		s.Pop()
		if got := s.Top(); got != 0 {
			t.Errorf("Top() = %d, want 0", got)
		}
		if got := s.GetMin(); got != -2 {
			t.Errorf("GetMin() = %d, want -2", got)
		}
	})

	t.Run("重复最小值依次弹出", func(t *testing.T) {
		s := Constructor()
		s.Push(0)
		s.Push(1)
		s.Push(0)
		if got := s.GetMin(); got != 0 {
			t.Errorf("GetMin() = %d, want 0", got)
		}
		s.Pop() // 弹出 0
		if got := s.GetMin(); got != 0 {
			t.Errorf("弹出后 GetMin() = %d, want 0（重复最小值需保留副本）", got)
		}
		s.Pop() // 弹出 1
		if got := s.GetMin(); got != 0 {
			t.Errorf("GetMin() = %d, want 0", got)
		}
	})

	t.Run("递减序列", func(t *testing.T) {
		s := Constructor()
		for _, v := range []int{5, 4, 3, 2, 1} {
			s.Push(v)
			if got := s.GetMin(); got != v {
				t.Errorf("Push(%d) 后 GetMin() = %d, want %d", v, got, v)
			}
		}
		for _, v := range []int{2, 3, 4, 5} {
			s.Pop()
			if got := s.GetMin(); got != v {
				t.Errorf("Pop() 后 GetMin() = %d, want %d", got, v)
			}
		}
	})

	t.Run("递增序列辅助栈保持最小", func(t *testing.T) {
		s := Constructor()
		for _, v := range []int{1, 2, 3, 4, 5} {
			s.Push(v)
		}
		if got := s.GetMin(); got != 1 {
			t.Errorf("GetMin() = %d, want 1", got)
		}
		if got := s.Top(); got != 5 {
			t.Errorf("Top() = %d, want 5", got)
		}
	})

	t.Run("单元素与极值", func(t *testing.T) {
		s := Constructor()
		s.Push(-2147483648) // math.MinInt32
		if got := s.GetMin(); got != -2147483648 {
			t.Errorf("GetMin() = %d, want -2147483648", got)
		}
		s.Push(2147483647) // math.MaxInt32
		if got := s.Top(); got != 2147483647 {
			t.Errorf("Top() = %d, want 2147483647", got)
		}
		if got := s.GetMin(); got != -2147483648 {
			t.Errorf("GetMin() = %d, want -2147483648", got)
		}
		s.Pop()
		if got := s.GetMin(); got != -2147483648 {
			t.Errorf("Pop() 后 GetMin() = %d, want -2147483648", got)
		}
	})

	t.Run("弹出最小值后次小值接任", func(t *testing.T) {
		s := Constructor()
		s.Push(3)
		s.Push(1)
		s.Push(2)
		s.Push(1)
		s.Pop() // 弹出 1
		if got := s.GetMin(); got != 1 {
			t.Errorf("GetMin() = %d, want 1", got)
		}
		s.Pop() // 弹出 2
		s.Pop() // 弹出 1
		if got := s.GetMin(); got != 3 {
			t.Errorf("GetMin() = %d, want 3", got)
		}
	})
}

func BenchmarkMinStack(b *testing.B) {
	b.Run("交替Push与GetMin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := Constructor()
			for j := 0; j < 1000; j++ {
				s.Push(j)
				if j%10 == 0 {
					s.GetMin()
				}
			}
		}
	})
	b.Run("递减Push与Pop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := Constructor()
			for j := 0; j < 1000; j++ {
				s.Push(-j)
			}
			for j := 0; j < 1000; j++ {
				s.GetMin()
				s.Pop()
			}
		}
	})
}
