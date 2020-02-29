package common

import "math/rand"

//Max calculates maximumm number from given list
func Max(nums ...int)int{
	max := nums[0]

	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	return max
}

//Min calculates minimum number from given list
func Min(nums ...int)int{
	min := nums[0]

	for _, num := range nums {
		if min > num {
			min = num
		}
	}
	return min
}

// Random rertuns a random number over a range
func Random(min, max int) int {
	if min == max {
		return min
	}
	return rand.Intn(max-min) + min
}