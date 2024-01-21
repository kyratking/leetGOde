// Problem
// https://www.codewars.com/kata/5926624c9b424d10390000bf

package main

func SumEvenFibonacci(limit int) int {
	sum := 2
	left := 1
	right := 2
	for right <= limit {
		temp := left
		left = right
		right = temp + left
		if right%2 == 0 {
			sum += right
		}
	}
	return sum
}
