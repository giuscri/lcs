package main

import "testing"

func TestLcs(t *testing.T) {
	if lcs("abc", "abc") != 3 {
		t.Fail()
	}

	if lcs("abzzzzc", "abczzzz") != 6 {
		t.Fail()
	}

	if lcs("", "a") != 0 {
		t.Fail()
	}
}
