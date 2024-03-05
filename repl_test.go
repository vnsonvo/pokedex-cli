package main

import "testing"

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
			input:    "Testing clean input func",
			expected: []string{"testing", "clean", "input", "func"},
		},
	}

	for _, testCase := range cases {
		actual := cleanInput(testCase.input)
		if len(actual) != len(testCase.expected) {
			t.Error("The lengths are not equal")
			continue
		}

		for i, word := range actual {
			expectedWord := testCase.expected[i]
			if word != expectedWord {
				t.Errorf("%v does not equal %v", word, expectedWord)
			}

		}
	}
}
