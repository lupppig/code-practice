package main

import "fmt"

/*
 * Complete the 'dynamicArray' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */

func dynamicArray(n int32, queries [][]int32) []int32 {
	var arr = make([][]int32, n)
	var lastAnswer int32 = 0
	lArr := make([]int32, 0)
	for i := 0; i < len(queries); i++ {
		x, y := queries[i][1], queries[i][2]
		typev := queries[i][0]
		if typev == 2 {
			index := (x ^ lastAnswer) % n
			lastAnswer = arr[index][y%int32(len(arr[index]))]
			lArr = append(lArr, lastAnswer)
		}
		index := (x ^ lastAnswer) % n
		arr[index] = append(arr[index], y)
	}

	return lArr
}

func main() {
	var n int32 = 2
	queries := [][]int32{
		{1, 0, 5},
		{1, 1, 7},
		{2, 1, 0},
	}
	fmt.Println(dynamicArray(n, queries))

	queries = [][]int32{
		{1, 0, 5}, {1, 1, 7}, {1, 0, 3}, {2, 1, 0}, {2, 1, 1},
	}
	fmt.Println(dynamicArray(n, queries))
}
