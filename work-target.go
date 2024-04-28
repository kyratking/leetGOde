package main

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	result := 0

	for i := range hours {
		if hours[i] >= target {
			result++
		}
	}

	return result
}
