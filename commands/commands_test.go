package commands

import "testing"

type commandTestCase struct {
	input    string
	expected string
}

func TestReadCommand(t *testing.T) {
	tests := []commandTestCase{
		{
			input:    "smurk",
			expected: "This is Smurk.",
		},
		{
			input:    "unknown",
			expected: "Smurk does not know command \"unknown\"",
		},
	}

	for _, tt := range tests {
		actual := ReadCommand(tt.input)
		if actual != tt.expected {
			t.Errorf("Command %q has returned wrong output.\nGot=%q, \nexpected=%q", tt.input, actual, tt.expected)
		}
	}
}
