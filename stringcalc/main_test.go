package main

import "testing"

func TestAddEmptyString(t *testing.T) {
	if result := Add(""); result != 0 {
		t.Fatalf("Empty string should return 0")
	}
}

func TestAddSingleNumber(t *testing.T) {
	if result := Add("1"); result != 1 {
		t.Fatalf("Should return 1")
	}
}

func TestAddMulitpleNumbers(t *testing.T) {
	if result := Add("1,2"); result != 3 {
		t.Fatalf("Should return 3, but returned <%d>", result)
	}
}
