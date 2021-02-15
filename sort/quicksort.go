package sort

func QuickSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}
	var result = make([]int, len(numbers), len(numbers))

	copy(result, numbers)

	result = quicksort(result)

	return result
}

func quicksort(numbers []int) []int {
	if len(numbers) > 1 {
		mid := numbers[len(numbers)/2]

		left := findAllNumbersLessThan(numbers, mid)
		right := findAllNumbersGreaterThan(numbers, mid)

		return append(quicksort(left), append([]int{mid}, quicksort(right)...)...)
	}
	return numbers
}

func findAllNumbersLessThan(numbers []int, pivot int) []int {
	var result []int

	for _, num := range numbers {
		if num < pivot {
			result = append(result, num)
		}
	}

	return result
}

func findAllNumbersGreaterThan(numbers []int, pivot int) []int {
	var result []int

	for _, num := range numbers {
		if num > pivot {
			result = append(result, num)
		}
	}

	return result
}
