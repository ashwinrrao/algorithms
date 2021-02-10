package sort_test

import (
	"github.com/ashwinrrao/algorithms/sort"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mergesort", func() {

	When("an empty input is given", func() {
		var (
			numbers = []int{}
			output  []int
		)

		BeforeEach(func() {
			output = sort.Mergesort(numbers)
		})

		It("should return an empty result", func() {
			Expect(output).To(Equal(numbers))
		})

	})

	When("an input with a single number is given", func() {
		var (
			numbers = []int{42}
			output  = []int{42}
		)

		BeforeEach(func() {
			output = sort.Mergesort(numbers)
		})

		It("should return the same slice as the output", func() {
			Expect(output).To(Equal(numbers))
		})

	})

	When("an input with an odd number of elements is given", func() {
		var (
			numbers        = []int{42, 2, 21}
			expectedOutput = []int{2, 21, 42}
			output         []int
		)

		BeforeEach(func() {
			output = sort.Mergesort(numbers)
		})

		It("should return the sorted slice as the output", func() {
			Expect(output).To(Equal(expectedOutput))
		})

	})

	When("an input with an even number of elements is given", func() {
		var (
			numbers        = []int{42, 2, 21, 32, 98, 6, -29, 0}
			expectedOutput = []int{-29, 0, 2, 6, 21, 32, 42, 98}
			output         []int
		)

		BeforeEach(func() {
			output = sort.Mergesort(numbers)
		})

		It("should return the sorted slice as the output", func() {
			Expect(output).To(Equal(expectedOutput))
		})

	})

})
