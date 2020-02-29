package algo

func insertionSort(in []int) []int {

	for i := 1; i < len(in); i++ {
		current := in[i]
		j := i - 1

		for j >= 0 && in[j] > current {
			in[j+1] = in[j]
			j--
		}
		in[j+1] = current
	}
	return in
}
