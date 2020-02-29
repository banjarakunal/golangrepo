package algos


func getProductOfOthers(list []int)[]int{
	if len(list) < 2 {
		return []int{}
	}

	out := make([]int, len(list))

	// get product of all other numbers before their indices.
	start1 := 1
	for i := 0; i < len(list); i++ {
		out[i] = start1
		start1 *= list[i]
	}

	// get product of all other numbers after their indices then multiply them
	// with their corresponding values.
	start2 := 1
	for i := len(list) - 1; i > -1; i-- {
		out[i] *= start2
		start2 *= list[i]
	}

	return out
}

