package main

import (
	"testing"
)

func TestToFilter(t *testing.T) {
	// Given
	args := []string{"hello=world", "foo=bar"}

	// When
	filter := toFilter(args)

	// Then
	if len(filter) != 3 {
		t.Fatalf("expected 3 params; actual was %d\n", len(filter))
	}
	if filter["hello"] != "world" {
		t.Fatalf("expected hello => world ; actual => %s\n", filter["hello"])
	}
	if filter["foo"] != "bar" {
		t.Fatalf("expected foo => bar ; actual => %s\n", filter["foo"])
	}
}
