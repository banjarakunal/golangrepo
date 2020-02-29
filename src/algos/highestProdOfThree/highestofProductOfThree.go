package algos

import (
	"algos/common"
)


func highestProduct(list []int) int{
	if len(list) <= 3{
		return 0
	}
	
	highest := common.Max(list[0], list[1])
	//lowest := common.Min(list[0], list[1])
	highestOfTwo := list[0]*list[1]
	//lowestofTwo := list[0]*list[1]
	highestOfThree :=  highestOfTwo * list[2]

	for i := 2; i < len(list) ; i++{
		current := list[i]

		highestOfThree = common.Max(highestOfThree, current*highestOfTwo)
		highestOfTwo = common.Max(highestOfTwo, current * highest)
		//lowestofTwo = common.Min(lowestofTwo, current * lowest, current * highest)
		highest = common.Max(highest, current)
		//lowest = common.Min(lowest, current)
	}
	return highestOfThree

}