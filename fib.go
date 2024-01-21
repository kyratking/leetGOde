// Problem
// https://www.codewars.com/kata/57a1d5ef7cb1f3db590002af

package main

func Fib(n int) int {
	memo := make(map[int]int)
	return FibMemo(n, memo)
}

func FibMemo(n int, memo map[int]int) int {
	if n == 0 || n == 1 {
		return n
	}

	if memo[n] != 0 {
		return memo[n]
	}

	if memo[n-1] == 0 {
		memo[n-1] = FibMemo(n-1, memo)
	}

	if memo[n-2] == 0 {
		memo[n-2] = FibMemo(n-2, memo)
	}

	return memo[n-1] + memo[n-2]
}
