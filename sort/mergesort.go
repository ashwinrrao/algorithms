package sort

func Mergesort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	var result = make([]int, len(numbers), len(numbers))
	copy(result, numbers)

	mergesort(result)

	return result
}

func mergesort(numbers []int) {
	if len(numbers) > 1 {
		var mid int = len(numbers) / 2

		left := make([]int, len(numbers[:mid]), len(numbers[:mid]))
		right := make([]int, len(numbers[mid:]), len(numbers[mid:]))

		copy(left, numbers[:mid])
		copy(right, numbers[mid:])

		mergesort(left)
		mergesort(right)

		i, j, k := 0, 0, 0
		for ; i < len(left) && j < len(right); {
			if left[i] < right[j] {
				numbers[k] = left[i]
				i++
				k++
			} else {
				numbers[k] = right[j]
				j++
				k++
			}
		}

		for ; i < len(left); i++ {
			numbers[k] = left[i]
			k++
		}

		for ; j < len(right); j++ {
			numbers[k] = right[j]
			k++
		}
	}
}
