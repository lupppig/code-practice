package main

import (
	"math"
)

// given an array of example [1, 2, 3]
// find every possible subsets of the given array
// answer -> [[], [1], [1, 2], [1, 3], [2], [2, 3][3], [1, 2, 3]]
// 8 possible subsets....
//
// O(2^n)

func Subsets(nums []int) [][]int {
	capped := int(math.Pow(2, float64(len(nums))))
	resulArr := make([][]int, 0, capped)
	subArrays := make([]int, 0)
	generateSubsets(0, nums, subArrays, &resulArr)

	return resulArr
}

func generateSubsets(index int, nums []int, subArray []int, resultArr *[][]int) {
	res := *resultArr
	copied := make([]int, len(subArray))
	copy(copied, subArray)
	res = append(res, copied)
	*resultArr = res
	for i := index; i < len(nums); i++ {
		subArray = append(subArray, nums[i])
		generateSubsets(i+1, nums, subArray, resultArr)
		subArray = subArray[:len(subArray)-1]
	}

}
