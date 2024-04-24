package main

func getConcatenation(nums []int) []int {
	n := len(nums)
	result := make([]int, 2*n)
	copy(result, nums)
	copy(result[n:], nums)
	return result
}
