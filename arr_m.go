package main

import (
	"fmt"
)

/*
 * Complete the 'arrayManipulation' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */

func arrayManipulation(n int32, queries [][]int32) int64 {
	arr := make([]int64, n)

	for _, query := range queries {
		a, b, k := query[0]-1, query[1], query[2]
		arr[a] += int64(k)
		if b < n {
			arr[b] -= int64(k)
		}
	}

	var max int64 = -999999999999999
	var curr int64 = 0

	for _, val := range arr {
		curr += val
		if curr > max {
			max = curr
		}
	}
	return max
}

func main() {
	queries := [][]int32{
		{1, 2, 100},
		{2, 5, 100},
		{3, 4, 100},
	}
	fmt.Println(arrayManipulation(5, queries))
}
