package main

import "testing"



func TestIsValid(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"([])", true},
		{"({[]})", true},
		{"[{()}]", true},
		{"{({})}", true},
		{"(}", false},
		{"[(])", false},
		{"(", false},
		{"", true},
		{"{[()]}", true},
		{"[", false},
	}

	for _, tt := range tests {
		got := isValid(tt.input)
		if got != tt.expected {
			t.Errorf("isValid(%q) = %v; want %v", tt.input, got, tt.expected)
		}
	}
}
