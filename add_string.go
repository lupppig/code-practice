package main

import "fmt"

func addStrings(num1 string, num2 string) string {

	if num1 == "0" && num2 == "0" {
		return "0"
	}
	i := len(num1) - 1
	j := len(num2) - 1

	carry := 0
	result := make([]byte, 0)

	for i >= 0 || j >= 0 || carry > 0 {
		var n1, n2 int
		if i >= 0 {
			n1 = int(num1[i] - '0')
			i--
		}

		if j >= 0 {
			n2 = int(num2[j] - '0')
			j--
		}

		sum := n1 + n2 + carry

		result = append(result, byte(sum%10)+'0')
		carry = sum / 10
	}

	i, j = 0, len(num2)
	for i, j := 0, len(result)-1; i < j; {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}
	return string(result)
}

func main() {

	fmt.Printf("%s", addStrings("456", "77"))

}
