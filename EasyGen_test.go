package main

import (
	"testing"
)

//==========================================================================
// list0

func TestList0(t *testing.T) {
	t.Log("First and plainest list test")
	const await = "The colors are: red, blue, white, .\n"
	if got := Generate(false, "list0"); got != await {
		t.Errorf("Mismatch:, got '%s' instead", got)
	}
}

//==========================================================================
// list1

func TestList1Text(t *testing.T) {
	t.Log("Second test, with text template")
	const await = "The quoted colors are: \"red\", \"blue\", \"white\", .\n"
	if got := Generate(false, "list1"); got != await {
		t.Errorf("Mismatch:, got '%s' instead", got)
	}
}

func TestList1HTML(t *testing.T) {
	t.Log("Second test, with html template")
	const await = "The quoted colors are: &#34;red&#34;, &#34;blue&#34;, &#34;white&#34;, .\n"
	if got := Generate(true, "list1"); got != await {
		t.Errorf("Mismatch:, got '%s' instead", got)
	}
}

//==========================================================================
// listfunc1

func TestListFunc1(t *testing.T) {
	t.Log("Test custom template function - minus1")
	const await = "red, blue, white"
	if got := Generate(false, "listfunc1"); got != await {
		t.Errorf("Mismatch:, got '%s' instead", got)
	}
}
