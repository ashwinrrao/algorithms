package sort_test

import (
	"github.com/ashwinrrao/algorithms/sort"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Insertionsort", func() {

	When("an empty input is given", func() {
		var (
			result []int
		)

		BeforeEach(func() {
			result = sort.InsertionSort([]int{})
		})

		It("should return an empty output", func() {
			Expect(result).To(Equal([]int{}))
		})
	})

	When("an input is given consisting of all positive integers", func() {
		var (
			expectedResult = []int{1, 2, 3, 4, 5}
			result         []int
		)

		BeforeEach(func() {
			result = sort.InsertionSort([]int{4, 1, 3, 2, 5})
		})

		It("should return the expected result", func() {
			Expect(result).To(Equal(expectedResult))
		})

	})

	When("an input is given consisting of positive numbers and zero", func() {
		var (
			expectedResult = []int{0, 1, 2, 3, 4, 5}
			result         []int
		)

		BeforeEach(func() {
			result = sort.InsertionSort([]int{4, 1, 3, 2, 5, 0})
		})

		It("should return the expected result", func() {
			Expect(result).To(Equal(expectedResult))
		})
	})

	When("an input is given consisting of positive numbers, negative numbers and zero", func() {
		var (
			// Expected result is this way because -3 is smaller than -2 and -2 is smaller than -1
			expectedResult = []int{-3, -2, -1, 0, 1, 2, 3, 4, 5}
			result         []int
		)

		BeforeEach(func() {
			result = sort.InsertionSort([]int{4, 1, 3, 2, 5, 0, -1, -2, -3})
		})

		It("should return the expected result", func() {
			Expect(result).To(Equal(expectedResult))
		})
	})
})
