package main

import (
	"testing"
)

func TestLetterSumEmptyString(t *testing.T) {
	value := ""
	want := 0
	result := letterSum(value)

	if result != want {
		t.Fatalf(`letterSum(%q) did not result in %q`, value, want)
	}
}

func TestLetterSumSingleCharacterString(t *testing.T) {
	value := "a"
	want := 1
	result := letterSum(value)

	if result != want {
		t.Fatalf(`letterSum(%q) did not result in %q`, value, want)
	}
}

func TestLetterSumMultiCharacterString(t *testing.T) {
	value := "abc"
	want := 6
	result := letterSum(value)

	if result != want {
		t.Fatalf(`letterSum(%q) did not result in %q`, value, want)
	}
}
