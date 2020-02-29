package algos

import "algos/common"

func inplaceShuffle(list []int)[]int{
	if len(list) <= 1 {
		return list
	}

	lastIndex := len(list) - 1
	for i := 0; i < len(list); i++ {
		// get a andomized index that is between the current and last index.
		randomIndex := common.Random(i, lastIndex)

		// swap current value.
		if i != randomIndex {
			list[i], list[randomIndex] = list[randomIndex], list[i]
		}
	}

	return list


}