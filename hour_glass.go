package main

func hourglassSum(arr [][]int32) int32 {
	var max int32 = -10000 // set max value...
	for i := 0; i < len(arr); i++ {
		var sum int32
		for j := i + 1; j < len(arr[i]); j++ {
			for k := j + 1; k < len(arr[i]); k++ {
				for m := 0; m < len(arr[i]); m++ {
					a := m + 1
					b := a + 1
					if a > len(arr)-1 || b > len(arr)-1 {
						break
					}
					sum = arr[i][m] + arr[i][a] + arr[i][b] + arr[j][a] + arr[k][m] + arr[k][a] + arr[k][b]
					if max < sum {
						max = sum
					}
				}
				break
			}
			break
		}

	}
	return max
}

// func main() {
// 	arr := [][]int32{
// 		{-9, -9, -9, 1, 1, 1},
// 		{0, -9, 0, 4, 3, 2},
// 		{-9, -9, -9, 1, 2, 3},
// 		{0, 0, 8, 6, 6, 0},
// 		{0, 0, 0, -2, 0, 0},
// 		{0, 0, 1, 2, 4, 0},
// 	}

// 	arr2 := [][]int32{
// 		{1, 1, 1, 0, 0, 0},
// 		{0, 1, 0, 0, 0, 0},
// 		{1, 1, 1, 0, 0, 0},
// 		{0, 0, 2, 4, 4, 0},
// 		{0, 0, 0, 2, 0, 0},
// 		{0, 0, 1, 2, 4, 0},
// 	}
// 	fmt.Println(hourglassSum(arr), "--")
// 	fmt.Println(hourglassSum(arr2))
// }
