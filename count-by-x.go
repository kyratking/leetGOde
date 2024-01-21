// Problem
// https://www.codewars.com/kata/5513795bd3fafb56c200049e

package main

func CountByX(x, n int) []int {
	result := []int{}
	for i := 1; i <= n; i++ {
		result = append(result, i*x)
	}
	return result
}
