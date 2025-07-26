package main

func isValid(s string) bool {
	mapBrack := map[rune]rune{')': '(', '}': '{', ']': '['}
	stack := []rune{}
	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != mapBrack[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// func main() {
// 	fmt.Println(isValid("()") == true) // true
// 	fmt.Println(isValid("([])"))       // true
// 	fmt.Println(isValid("{[()]}"))     // true
// 	fmt.Println(isValid("{[()]}}"))    // false
// 	fmt.Println(isValid("({[]})"))     // true
// 	fmt.Println(isValid("[{()}]"))     // true
// 	fmt.Println(isValid("{({})}"))     // true
// 	fmt.Println(isValid("(}"))         // false
// 	fmt.Println(isValid("[(])"))       // false
// 	fmt.Println(isValid("("))          // false
// 	fmt.Println(isValid(""))           // false
// }
