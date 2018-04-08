package main

import "testing"

func TestNewDeck(t *testing.T) { // t: test handller
	d := newDeck()

	if len(d) != 16 {

		// %v : space
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}
