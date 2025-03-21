package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HellO WORld  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		// Check if the lengths match
		if len(actual) != len(c.expected) {
			t.Errorf("Test Failed: expected slice length %d but got %d", len(c.expected), len(actual))
			continue // Skip the rest of this iteration
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			fmt.Printf("%s, %s\n", word, expectedWord)

			if word != expectedWord {
				t.Errorf("Test Failed, actual: %s; expected: %s", word, expectedWord)
			}
		}
	}
}
