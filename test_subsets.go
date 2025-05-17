package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestSubsets(t *testing.T) {
	input := []int{1, 2}
	expected := [][]int{
		{},
		{1},
		{2},
		{1, 2},
	}

	result := Subsets(input)

	normalize := func(sets [][]int) {
		for _, set := range sets {
			sort.Ints(set)
		}
		sort.Slice(sets, func(i, j int) bool {
			if len(sets[i]) != len(sets[j]) {
				return len(sets[i]) < len(sets[j])
			}
			for k := range sets[i] {
				if sets[i][k] != sets[j][k] {
					return sets[i][k] < sets[j][k]
				}
			}
			return false
		})
	}

	normalize(expected)
	normalize(result)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
