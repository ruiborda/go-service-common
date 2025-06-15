package types

import (
	"testing"
)

func TestPtr(t *testing.T) {
	type testCase[T any] struct {
		input    T
		expected T
	}

	intTest := testCase[int]{input: 42, expected: 42}
	strTest := testCase[string]{input: "hello", expected: "hello"}
	boolTest := testCase[bool]{input: true, expected: true}

	if got := Pointer(intTest.input); *got != intTest.expected {
		t.Errorf("Pointer(int): got %v, want %v", *got, intTest.expected)
	}

	if got := Pointer(strTest.input); *got != strTest.expected {
		t.Errorf("Pointer(string): got %v, want %v", *got, strTest.expected)
	}

	if got := Pointer(boolTest.input); *got != boolTest.expected {
		t.Errorf("Pointer(bool): got %v, want %v", *got, boolTest.expected)
	}
}
