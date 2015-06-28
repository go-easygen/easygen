package main

import (
	"testing"
)

func TestList0(t *testing.T) {
	t.Log("First and plainest list test")
	const await = "The colors are: red, blue, white, .\n"
	if got := Generate(false, "list0"); got != await {
		t.Errorf("Mismatch:, got %s instead", got)
	}
}
