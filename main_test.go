package main

import "testing"

func TestHeight(t *testing.T) {
	if got := Height(nil); got != 0 {
		t.Errorf("Height(nil) = %d, want 0", got)
	}
	if got := Height(tn(1, nil, nil)); got != 1 {
		t.Errorf("single node = %d, want 1", got)
	}
	root := tn(3, tn(9, nil, nil), tn(20, tn(15, nil, nil), tn(7, nil, nil)))
	if got := Height(root); got != 3 {
		t.Errorf("Height = %d, want 3", got)
	}
	// Skewed chain of 4.
	chain := tn(1, tn(2, tn(3, tn(4, nil, nil), nil), nil), nil)
	if got := Height(chain); got != 4 {
		t.Errorf("chain = %d, want 4", got)
	}
}
