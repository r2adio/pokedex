package main

import (
	"testing"
)

// standard way to implement unit tests in go
func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello WORLD",
			expected: []string{"hello", "world"},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("The lengths are not equal: %v vs %v", len(actual), len(cs.expected))
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("%v doesn't equal %v", actualWord, expectedWord)
			}
		}
	}
}
