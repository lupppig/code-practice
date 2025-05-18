package main

import "slices"

func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0, len(nums))
	slices.Sort(nums)

	backtrack(nums, &result, []int{}, 0)

	return result
}

func backtrack(arr []int, result *[][]int, current []int, start int) {
	cur := make([]int, len(current))
	copy(cur, current)

	*result = append(*result, cur)

	for i := start; i < len(arr); i++ {
		if i > start && arr[i] == arr[i-1] {
			continue
		}

		current = append(current, arr[i])
		backtrack(arr, result, current, i+1)

		current = current[:len(current)-1]
	}
}
