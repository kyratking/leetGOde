// Problem
//

package main

func FibStrings(n int) string {
	left := "0"
	right := "01"
	for n != 0 {
		temp := left
		left = right
		right = left + temp
		n--
	}
	return right
}
