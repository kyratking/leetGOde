package main

func buildArray(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[nums[i]] = nums[i]
	}
	return result
}
