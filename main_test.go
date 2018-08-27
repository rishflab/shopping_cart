package main

import "testing"

func testInitialiseCart(t *testing.T) {
	total := 9
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
