// Problem
// https://www.codewars.com/kata/51b6249c4612257ac0000005

package main

func RomanToInt(roman string) int {
	NUMERALS := map[byte]int{77: 1000, 68: 500, 67: 100, 76: 50, 88: 10, 86: 5, 73: 1}

	prev := 1001
	number := 0
	for i := 0; i < len(roman); i++ {
		value := NUMERALS[roman[i]]
		if value > prev {
			number -= prev * 2
		}
		number += value
		prev = value
	}
	return number
}
