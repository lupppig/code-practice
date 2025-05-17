package main

import "fmt"

/**
Given an array nums with n objects colored red, white, or blue, sort them in-place so that
objects of the same color are adjacent, with the colors in the order red, white, and blue.
*/

func sortColors(arr []int) {

	a := sorts(arr)
	fmt.Println(a)

}

func sorts(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	left, right := sorts(arr[int(len(arr)/2):]), sorts(arr[:int(len(arr)/2)])
	return merge(left, right)
}

func merge(left, right []int) []int {
	sortedArray := make([]int, 0, len(left)+len(right))
	l, r := 0, 0

	for l < len(left) && r < len(right) {

		if left[l] < right[r] {
			sortedArray = append(sortedArray, left[l])
			l++
		} else {
			sortedArray = append(sortedArray, right[r])
			r++
		}
	}

	sortedArray = append(sortedArray, left[l:]...)
	sortedArray = append(sortedArray, right[r:]...)
	return sortedArray
}

func main() {
	arr := []int{2, 0, 2, 1, 1, 0}

	sortColors(arr)
}
