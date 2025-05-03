package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {

	var sums = []struct {
		array    []int
		target   int
		expected [2]int
	}{
		{
			array:    []int{1, 2, 3},
			target:   4,
			expected: [2]int{0, 2},
		},
		{
			array:    []int{4, 5, 6, 7},
			target:   11,
			expected: [2]int{0, 3},
		},
		{
			array:    []int{3, 4, 5, 8},
			target:   12,
			expected: [2]int{1, 3},
		},
	}
	for _, sum := range sums {
		got := TwoSum(sum.array, sum.target)

		if got != sum.expected {
			t.Errorf("TwoSums(%v, %d) = %v; want %v", sum.array, sum.target, got, sum.expected)
		}
	}

}

func TestTwoSums(t *testing.T) {
	tests := []struct {
		numbers []int
		target  int
		want    [2]int
	}{
		{[]int{1, 4, 5, 3, 2}, 6, [2]int{0, 2}},       // 1 + 5
		{[]int{2, 7, 11, 15}, 9, [2]int{0, 1}},        // 2 + 7
		{[]int{3, 2, 4}, 6, [2]int{1, 2}},             // 2 + 4
		{[]int{-1, -2, -3, -4, -5}, -8, [2]int{2, 4}}, // -3 + -5
	}

	for _, tt := range tests {
		got := TwoSums(tt.numbers, tt.target)

		if tt.numbers[got[0]]+tt.numbers[got[1]] != tt.target {
			t.Errorf("TwoSums(%v, %d) = %v, sum mismatch", tt.numbers, tt.target, got)
		}

		if got[0] == got[1] {
			t.Errorf("TwoSums(%v, %d) = %v, indices must be different", tt.numbers, tt.target, got)
		}

		if !reflect.DeepEqual(tt.want, got) {
			t.Logf("Valid output: %v (expected one of many possible)", got)
		}
	}
}
