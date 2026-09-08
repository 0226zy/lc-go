package insertdeletegetrandomo1

import "testing"

// TestRandomizedSet 按 LeetCode 官方示例的调用序列测试
func TestRandomizedSet(t *testing.T) {
	s := Constructor()

	t.Run("插入1返回true", func(t *testing.T) {
		if !s.Insert(1) {
			t.Error("Insert(1) = false, want true")
		}
	})

	t.Run("删除不存在的2返回false", func(t *testing.T) {
		if s.Remove(2) {
			t.Error("Remove(2) = true, want false")
		}
	})

	t.Run("插入2返回true", func(t *testing.T) {
		if !s.Insert(2) {
			t.Error("Insert(2) = false, want true")
		}
	})

	t.Run("getRandom返回1或2", func(t *testing.T) {
		r := s.GetRandom()
		if r != 1 && r != 2 {
			t.Errorf("GetRandom() = %d, want 1 or 2", r)
		}
	})

	t.Run("删除1返回true", func(t *testing.T) {
		if !s.Remove(1) {
			t.Error("Remove(1) = false, want true")
		}
	})

	t.Run("重复插入2返回false", func(t *testing.T) {
		if s.Insert(2) {
			t.Error("Insert(2) = true, want false")
		}
	})

	t.Run("getRandom只能返回2", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			if r := s.GetRandom(); r != 2 {
				t.Fatalf("GetRandom() = %d, want 2", r)
			}
		}
	})
}

// TestRandomizedSetRemoveSwap 验证交换删除后集合内容正确
func TestRandomizedSetRemoveSwap(t *testing.T) {
	s := Constructor()
	for _, v := range []int{10, 20, 30, 40, 50} {
		s.Insert(v)
	}

	// 删除中间元素 20，应触发交换删除
	if !s.Remove(20) {
		t.Fatal("Remove(20) = false, want true")
	}

	// 剩余元素都应存在：Remove 返回 true 即存在，随后重新插回
	for _, v := range []int{10, 30, 40, 50} {
		if !s.Remove(v) {
			t.Errorf("Remove(%d) = false, 说明集合中缺少该元素", v)
		}
		s.Insert(v)
	}
	if s.Remove(20) {
		t.Error("Remove(20) 第二次 = true, 说明 20 未被删除")
	}

	// 删除末尾元素后再删除被搬动的元素
	if !s.Remove(50) {
		t.Fatal("Remove(50) = false, want true")
	}
	if !s.Remove(40) {
		t.Fatal("Remove(40) = false, want true")
	}
	if s.Remove(50) {
		t.Error("Remove(50) 第二次 = true, want false")
	}

	// 清空后再插入，下标应从 0 重新开始
	for len(s.vals) > 0 {
		s.Remove(s.vals[0])
	}
	if len(s.vals) != 0 {
		t.Fatalf("清空后 len(vals) = %d, want 0", len(s.vals))
	}
	if !s.Insert(99) {
		t.Error("清空后 Insert(99) = false, want true")
	}
	if r := s.GetRandom(); r != 99 {
		t.Errorf("GetRandom() = %d, want 99", r)
	}
}

// TestGetRandomUniform 粗略验证随机性：重复调用应覆盖集合中的所有元素
func TestGetRandomUniform(t *testing.T) {
	s := Constructor()
	for v := 0; v < 5; v++ {
		s.Insert(v)
	}
	seen := make(map[int]bool)
	for i := 0; i < 2000; i++ {
		seen[s.GetRandom()] = true
	}
	for v := 0; v < 5; v++ {
		if !seen[v] {
			t.Errorf("调用 2000 次 getRandom 从未返回 %d", v)
		}
	}
}

func BenchmarkRandomizedSet(b *testing.B) {
	b.Run("Insert", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := Constructor()
			for v := 0; v < 1000; v++ {
				s.Insert(v)
			}
		}
	})

	b.Run("Remove", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := Constructor()
			for v := 0; v < 1000; v++ {
				s.Insert(v)
			}
			b.StartTimer()
			for v := 0; v < 1000; v++ {
				s.Remove(v)
			}
			b.StopTimer()
		}
	})

	b.Run("GetRandom", func(b *testing.B) {
		s := Constructor()
		for v := 0; v < 1000; v++ {
			s.Insert(v)
		}
		b.ResetTimer()
		sum := 0
		for i := 0; i < b.N; i++ {
			sum += s.GetRandom()
		}
		_ = sum
	})
}
