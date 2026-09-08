package lrucache

import "testing"

// op 描述一次对 LRUCache 的调用
type op struct {
	method string // "get" 或 "put"
	key    int
	value  int
	want   int // 仅 get 有期望值
}

func runLRUCase(t *testing.T, capacity int, ops []op) {
	t.Helper()
	cache := Constructor(capacity)
	for i, o := range ops {
		switch o.method {
		case "get":
			if got := cache.Get(o.key); got != o.want {
				t.Fatalf("第 %d 步 Get(%d) = %d, want %d", i+1, o.key, got, o.want)
			}
		case "put":
			cache.Put(o.key, o.value)
		}
	}
}

func TestLRUCache(t *testing.T) {
	t.Run("示例1: 容量2的完整调用序列", func(t *testing.T) {
		runLRUCase(t, 2, []op{
			{"put", 1, 1, 0},
			{"put", 2, 2, 0},
			{"get", 1, 0, 1},
			{"put", 3, 3, 0}, // 淘汰 key=2
			{"get", 2, 0, -1},
			{"put", 4, 4, 0}, // 淘汰 key=1
			{"get", 1, 0, -1},
			{"get", 3, 0, 3},
			{"get", 4, 0, 4},
		})
	})

	t.Run("示例2: get触发淘汰", func(t *testing.T) {
		runLRUCase(t, 2, []op{
			{"put", 1, 1, 0},
			{"put", 2, 2, 0},
			{"get", 1, 0, 1},
			{"put", 3, 3, 0},
			{"get", 2, 0, -1},
			{"put", 4, 4, 0},
			{"get", 1, 0, -1},
			{"get", 3, 0, 3},
			{"get", 4, 0, 4},
		})
	})

	t.Run("示例3: 更新已有key不增加占用", func(t *testing.T) {
		runLRUCase(t, 2, []op{
			{"put", 2, 1, 0},
			{"put", 1, 1, 0},
			{"put", 2, 3, 0}, // 更新已有 key=2，不应淘汰任何 key
			{"put", 4, 1, 0}, // 此时最久未使用的是 key=1，淘汰它
			{"get", 1, 0, -1},
			{"get", 2, 0, 3},
		})
	})

	t.Run("容量为1", func(t *testing.T) {
		runLRUCase(t, 1, []op{
			{"put", 1, 10, 0},
			{"get", 1, 0, 10},
			{"put", 2, 20, 0}, // 淘汰 key=1
			{"get", 1, 0, -1},
			{"get", 2, 0, 20},
		})
	})

	t.Run("get不存在返回-1", func(t *testing.T) {
		cache := Constructor(3)
		if got := cache.Get(99); got != -1 {
			t.Errorf("Get(99) = %d, want -1", got)
		}
	})

	t.Run("get命中后成为最近使用", func(t *testing.T) {
		runLRUCase(t, 2, []op{
			{"put", 1, 1, 0},
			{"put", 2, 2, 0},
			{"get", 1, 0, 1},  // key=1 变为最近使用
			{"put", 3, 3, 0},  // 淘汰 key=2
			{"get", 1, 0, 1},  // key=1 仍在
			{"get", 2, 0, -1}, // key=2 已被淘汰
		})
	})
}

func BenchmarkLRUCache(b *testing.B) {
	// 基准：容量 100，循环写入 1000 个 key，触发持续的淘汰
	b.Run("put淘汰", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cache := Constructor(100)
			for k := 0; k < 1000; k++ {
				cache.Put(k, k)
			}
		}
	})
	b.Run("get命中", func(b *testing.B) {
		cache := Constructor(100)
		for k := 0; k < 100; k++ {
			cache.Put(k, k)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cache.Get(i % 100)
		}
	})
}
