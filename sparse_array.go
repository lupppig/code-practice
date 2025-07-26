package main

/*
 * Complete the 'matchingStrings' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY stringList
 *  2. STRING_ARRAY queries
 */

func matchingStrings(stringList []string, queries []string) []int32 {
	var matchA []int32 = make([]int32, len(queries))
	for i := 0; i < len(stringList); i++ {
		for j := 0; j < len(queries); j++ {
			if stringList[i] == queries[j] {
				matchA[j]++
			}
		}
	}
	return matchA
}

// func main() {
// 	stringList := []string{"ab", "ab", "abc"}
// 	queries := []string{"ab", "abc", "bc"}
// 	matchA := matchingStrings(stringList, queries)

// 	fmt.Println(matchA)
// 	stringList = []string{"aba", "baba","aba" , "xzxb"}
// 	queries = []string{"aba", "xzxb", "ab"}

// 	fmt.Println(matchingStrings(stringList,queries))
// }
