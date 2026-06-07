package lrucache

import "testing"

func TestLRUCache_OfficialExample(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)

	if got := lru.Get(1); got != 1 {
		t.Errorf("Get(1) = %d, want 1", got)
	}

	lru.Put(3, 3)
	if got := lru.Get(2); got != -1 {
		t.Errorf("Get(2) = %d, want -1", got)
	}

	lru.Put(4, 4)
	if got := lru.Get(1); got != -1 {
		t.Errorf("Get(1) = %d, want -1", got)
	}
	if got := lru.Get(3); got != 3 {
		t.Errorf("Get(3) = %d, want 3", got)
	}
	if got := lru.Get(4); got != 4 {
		t.Errorf("Get(4) = %d, want 4", got)
	}
}

func TestLRUCache_CapacityOne(t *testing.T) {
	lru := Constructor(1)
	lru.Put(1, 1)
	if got := lru.Get(1); got != 1 {
		t.Errorf("Get(1) = %d, want 1", got)
	}

	lru.Put(2, 2)
	if got := lru.Get(1); got != -1 {
		t.Errorf("Get(1) after eviction = %d, want -1", got)
	}
	if got := lru.Get(2); got != 2 {
		t.Errorf("Get(2) = %d, want 2", got)
	}
}

func TestLRUCache_UpdateExistingKey(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(1, 10)

	if got := lru.Get(1); got != 10 {
		t.Errorf("Get(1) after update = %d, want 10", got)
	}
	if got := lru.Get(2); got != 2 {
		t.Errorf("Get(2) = %d, want 2", got)
	}
}

func TestLRUCache_GetUpdatesOrder(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)

	// 访问 1，使其变为最近使用
	lru.Get(1)

	// 插入 3，应淘汰最久未使用的 2
	lru.Put(3, 3)
	if got := lru.Get(2); got != -1 {
		t.Errorf("Get(2) = %d, want -1 (should be evicted)", got)
	}
	if got := lru.Get(1); got != 1 {
		t.Errorf("Get(1) = %d, want 1", got)
	}
	if got := lru.Get(3); got != 3 {
		t.Errorf("Get(3) = %d, want 3", got)
	}
}
