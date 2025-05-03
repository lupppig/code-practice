package main

import "math/bits"

func CountBits(n uint) int {
	var count int

	for n > 0 {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}

	return count
}

// or
func CountBits2(n uint) int {
	count := 0
	for n > 0 {
		count += int(n & 1)
		n >>= 1
	}
	return count
}

// or
func CountBits3(n uint) int {
	return bits.OnesCount(n)
}
