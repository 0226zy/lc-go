package mapsumpairs

import "testing"

func TestMapSum(t *testing.T) {
	t.Skip("未实现")

	m := Constructor()
	m.Insert("apple", 3)
	if got := m.Sum("ap"); got != 3 {
		t.Errorf("Sum(ap) = %d, want 3", got)
	}
	m.Insert("app", 2)
	if got := m.Sum("ap"); got != 5 {
		t.Errorf("Sum(ap) = %d, want 5", got)
	}
}
