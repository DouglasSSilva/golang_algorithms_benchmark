package sort_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "golang_studies/golang_algorithms_benchmark/sort"
	"golang_studies/golang_algorithms_benchmark/utils"
)

var _ = Describe("Bubblesort", func() {
	var arr []int
	var size int
	Context("Benchmarking bubblesort efficacy", func() {
		size = utils.RandomInt(50000, 100000)
		Measure("It should be fast", func(b Benchmarker) {
			arr = utils.RandomArrayCreator(size, 1, 1000)
			arrLen := len(arr)

			runtime := b.Time("runtime", func() {
				arr = BubbleSort(arr)
			})

			Expect(runtime.Seconds()).Should(BeNumerically("<", 10))
			fmt.Println("array length = ", arrLen, " rutime = ", runtime.Seconds())
			for i := 1; i < arrLen; i++ {
				Expect(arr[i-1]).Should(BeNumerically("<=", arr[i]))
			}
		}, utils.RandomInt(10, 20))
	})
})
