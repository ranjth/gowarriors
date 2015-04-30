package main

import (
	"testing"
)

func TestAddIntegers(t *testing.T) {
	var a, b, s int
	a = 15
	b = 20
	s = addIntegers(a, b)
	if s != 35 {
		t.Error("Expected 35, got", s)
	}
}
