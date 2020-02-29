package algo

func bubbleSort(list []int) []int {
	length := len(list)

	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}

		}
	}
	return list
}
