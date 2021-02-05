package sort

func InsertionSort(input []int) []int {
	if len(input) == 0 || len(input) == 1 {
		return input
	}

	var result = make([]int, len(input), len(input))

	copy(result, input)

	for i := 1; i < len(result); i++ {
		key := result[i]
		j := i - 1
		for ; j >= 0 && result[j] > key; j-- {
			result[j+1] = result[j]
		}
		result[j+1] = key
	}

	return result
}
