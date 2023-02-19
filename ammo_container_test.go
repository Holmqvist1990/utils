package utils

import "testing"

func TestAmmoContainer(t *testing.T) {
	a := AmmoContainer[string]{"a": 8, "b": 12}

	// Current ammo.
	curr := a.CurrentAmmo("a")
	if curr != 8 {
		t.Fatalf("expected %v got %v", 8, curr)
	}

	// Loading all.
	loaded := a.Take("a", 0, 8)
	if loaded != 8 {
		t.Fatalf("expected %v got %v", 8, loaded)
	}
	curr = a.CurrentAmmo("a")
	if curr != 0 {
		t.Fatalf("expected to be empty (0), got %v", curr)
	}

	// Empty.
	loaded = a.Take("a", 0, 8)
	if loaded != 0 {
		t.Fatalf("expected to be empty (0), got %v", loaded)
	}
	curr = a.CurrentAmmo("a")
	if curr != 0 {
		t.Fatalf("expected to be empty (0), got %v", curr)
	}

	// Loading some.
	loaded = a.Take("b", 2, 4)
	if loaded != 2 {
		t.Fatalf("expected %v got %v", 2, loaded)
	}
	curr = a.CurrentAmmo("b")
	if curr != 10 {
		t.Fatalf("expected %v got %v", 10, curr)
	}

	// Loading less than asked for.
	loaded = a.Take("b", 0, 12)
	if loaded != 10 {
		t.Fatalf("expected %v got %v", 10, loaded)
	}
	curr = a.CurrentAmmo("b")
	if curr != 0 {
		t.Fatalf("expected to be empty (0), got %v", curr)
	}
}
