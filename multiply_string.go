package main

import (
	"fmt"
)

/*
*
Given two non-negative integers num1 and num2 represented as strings,
return the product of num1 and num2, also represented as a string.
Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.
*/

// worst case could cause integer overflow...
// O(n)
// for int 64
// 498828660196 * 498828660196 = -3269442614257959980 // integer overflow right here
func Multiply(num1 string, num2 string) string {
	strMap := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}
	var va int64
	var vap int64

	num1Len := len(num1)
	num2Len := len(num2)

	i, j := 0, 0

	for i < num1Len || j < num2Len {
		if i < num1Len {
			va = va*10 + int64(strMap[string(num1[i])])
		}
		if j < num2Len {
			vap = vap*10 + int64(strMap[string(num2[j])])
		}
		i++
		j++
	}
	return fmt.Sprintf("%d", va*vap)
}

func Multiply2(num1 string, num2 string) string {

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	n, m := len(num1), len(num2)
	result := make([]int, n+m)
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			d1 := int(num1[i] - '0')
			d2 := int(num2[j] - '0')
			mul := d1 * d2
			p1, p2 := i+j, i+j+1
			sum := mul + result[p2]
			result[p2] = sum % 10
			result[p1] += sum / 10
		}
	}
	res := ""
	for _, digit := range result {
		if !(len(res) == 0 && digit == 0) {
			res += fmt.Sprint(digit)
		}
	}
	return res
}


/*
*

	apparently this function is illegal to use
	sybau...
*/
func convertToInt(str string) int {
	if str == "" || str == " " || len(str) == 0 {
		return 0
	}

	num := 0
	sign := 1

	if str[0] == '-' {
		sign = -1
		str = str[1:]
	}

	for _, c := range str {
		switch c {
		case '0':
			num = (num * 10)
		case '1':
			num = (num * 10) + 1
		case '2':
			num = (num * 10) + 2
		case '3':
			num = (num * 10) + 3
		case '4':
			num = (num * 10) + 4
		case '5':
			num = (num * 10) + 5
		case '6':
			num = (num * 10) + 6
		case '7':
			num = (num * 10) + 7
		case '8':
			num = (num * 10) + 8
		case '9':
			num = (num * 10) + 9
		default:
			fmt.Println("bleh bleh bleeeeh bleeeeeeeeeeeh")
			return 10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 // don't ask me why this here...
		}
	}
	return num * sign
}
