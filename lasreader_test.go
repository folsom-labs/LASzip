package main

import "testing"

// TestEatBits tests eatBits function
func TestEatBits(t *testing.T) {
	var b uint8 = 0xff
	var testData = []struct {
		nBits uint
		exp   int
	}{
		{3, 7},
		{1, 1},
		{2, 3},
		{2, 3},
	}
	var got int
	for _, d := range testData {
		got, b = eatBits(b, d.nBits)
		if got != d.exp {
			t.Fatalf("expected: %d, got: %d", d.exp, got)
		}
	}
}
