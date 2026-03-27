package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input		string
		expected	[]string
	} {
		{
			input:		"  hello  world  ",
			expected:	[]string{"hello", "world"},
		},
		{
			input:		"word",
			expected:	[]string{"word"},
		},
		{
			input:		"one\nper\nline",
			expected:	[]string{"one", "per", "line"},
		},
		{
			input:		"sPoNgE bOb ! !!  !",
			expected:	[]string{"sponge", "bob", "!", "!!", "!"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("The amount of words in the output does not match the expected amount of words!\nExpected %v, got %v", len(c.expected), len(actual))
		}

		mismatches := []struct {
			is		string
			should	string
		}{}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				mismatches = append(mismatches, struct {
					is		string
					should	string
				} {
					is: word,
					should: expectedWord,
				})
			}
		}

		if len(mismatches) != 0 {
			t.Errorf("These words in the output don't mach the expected words:\n%v", mismatches)
		}
	}
}
