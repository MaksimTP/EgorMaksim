package main

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func findClosestNumber(nums []int) int {
	closest := abs(nums[0])

	for _, num := range nums {
		if abs(num) < abs(closest) {
			closest = abs(num)
		}
	}

	for _, num := range nums {
		if closest == num {
			return closest
		}
	}
	return closest * -1
}
