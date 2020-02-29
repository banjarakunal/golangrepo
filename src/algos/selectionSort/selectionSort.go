package algo

func selectionSort(in []int) []int {
	minIndex := 0

	for i := 0; i < len(in)-1; i++ {
		minIndex = i
		for j := i + 1; j < len(in); j++ {
			if in[j] < in[minIndex] {
				minIndex = j
			}
		}
		in[i], in[minIndex] = in[minIndex], in[i]
	}
	return in
}
