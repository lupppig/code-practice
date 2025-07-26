package main

/*
 * Complete the 'rotateLeft' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER d
 *  2. INTEGER_ARRAY arr
 */

func rotateLeft(d int32, arr []int32) []int32 {
	rotate := d % int32(len(arr))
	lenArr := len(arr) - 1
	var count int32 = 0
	for {
		var i int32 = 0
		if count == rotate {
			break
		}
		if i > int32(lenArr) {
			i = 0
		}
		ap := arr[i]
		arr[i] = arr[i+1]
		for j := i; j < int32(lenArr); j++ {
			a := j + 1
			if a > int32(lenArr)-1 {
				break
			}
			arr[a] = arr[a+1]
		}
		arr[lenArr] = ap
		i++
		count++
	}
	return arr
}

// func main() {
// 	arr := []int32{1, 2, 3, 4, 5}
// 	fmt.Println(rotateLeft(1, arr))
// }
