package main

import "testing"

func TestCountBits(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 0},     // 0b00000
		{1, 1},     // 0b00001
		{2, 1},     // 0b00010
		{3, 2},     // 0b00011
		{5, 2},     // 0b00101
		{18, 2},    // 0b10010
		{255, 8},   // 0b11111111
		{1023, 10}, // 0b1111111111
	}

	for _, test := range tests {
		got := CountBits(uint(test.input))
		if got != test.expected {
			t.Errorf("CountBits(%d) = %d; want %d", test.input, got, test.expected)
		}
	}

	t.Run("Counts bit2", func(t *testing.T) {
		input := 5
		expected := 2
		result := CountBits2(uint(input))
		if result != expected {
			t.Errorf("Expected %d but got %d", expected, result)
		}
	})

	t.Run("Count bits 3", func(t *testing.T) {
		input := 5
		expected := 2
		result := CountBits3(uint(input))
		if result != expected {
			t.Errorf("Expected %d but got %d", expected, result)
		}
	})
}
