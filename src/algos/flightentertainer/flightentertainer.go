package main

func flightentertrainer(movieLengths []int, flightLength int) bool {
	movies := map[int]int{}

	for _, v := range movieLengths {
		// return true if we've seen the movie. else add the movie in the
		// watched list with a count 1.
		matchLength := flightLength - v
		if _, ok := movies[matchLength]; ok {
			return true
		}

		movies[v] = 1
	}
	return false
}	
