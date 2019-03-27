package main

import "testing"

func TestF(t *testing.T) {
	result := f(0, 0)
	if result != 0 {
		t.Fatalf("%v", result)
	}
}

