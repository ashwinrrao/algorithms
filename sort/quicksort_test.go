package sort_test

import (
	"github.com/ashwinrrao/algorithms/sort"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Quicksort", func() {

	When("asked to sort an empty slice", func() {
		var (
			output []int
		)

		BeforeEach(func() {
			output = sort.QuickSort([]int{})
		})

		It("should return an empty slice as the output", func() {
			Expect(output).To(Equal([]int{}))
		})

	})

	When("asked to sort a slice containing one number", func() {
		var (
			input          = []int{8}
			output         []int
			expectedOutput = []int{8}
		)

		BeforeEach(func() {
			output = sort.QuickSort(input)
		})

		It("should return the same slice as the output", func() {
			Expect(output).To(Equal(expectedOutput))
		})

	})

	When("asked to sort a slice of numbers", func() {

		var (
			input          = []int{99, 4, 23, 5}
			output         []int
			expectedOutput = []int{4, 5, 23, 99}
		)

		BeforeEach(func() {
			output = sort.QuickSort(input)
		})

		It("sort the given input slice", func() {
			Expect(output).To(Equal(expectedOutput))
		})

	})

	When("asked to sort a slice of numbers", func() {

		var (
			input          = []int{99, -1, 0, 4, 23, 5, 87}
			output         []int
			expectedOutput = []int{-1, 0, 4, 5, 23, 87, 99}
		)

		BeforeEach(func() {
			output = sort.QuickSort(input)
		})

		It("sort the given input slice", func() {
			Expect(output).To(Equal(expectedOutput))
		})

	})
})
