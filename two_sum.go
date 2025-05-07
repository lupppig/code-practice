package main

import "fmt"

func TwoSum(numbers []int, target int) [2]int {
	var indexes [2]int

	left := 0
	right := len(numbers) - 1

	for right > left {
		current_sum := numbers[left] + numbers[right]
		if current_sum == target {
			indexes[0] = left
			indexes[1] = right
			break
		}
		if current_sum < target {
			left++
		} else {
			right--
		}
	}
	return indexes
}

// if array is not sorted o(n)
func TwoSums(numbers []int, target int) [2]int {
	var seen = make(map[int]int)
	for ind, num := range numbers {
		comp := target - num
		fmt.Println(seen)
		if idx, ok := seen[comp]; ok {
			fmt.Println("---->", idx)
			return [2]int{seen[comp], ind}
		}
		seen[num] = ind
	}

	return [2]int{-1, -1}
}
