package main

import (
	"example.com/heap"
	"fmt"
	"math/rand"
	"time"
)

type Ordered interface {
	~float64 | ~int | ~string
}

func heapSort[T Ordered](input []T) []T {
	heap1 := heap.NewHeap[T](input)
	descending := []T{}
	for {
		if len(heap1.Items) > 0 {
			descending = append(descending, heap1.Largest())
			heap1.Remove()
		} else {
			break
		}
	}
	ascending := []T{}
	for i := len(descending) - 1; i >= 0; i-- {
		ascending = append(ascending, descending[i])
	}
	return ascending
}

const size = 50_000_000

func IsSorted[T Ordered](data []T) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}

func main() {
	
	slice := []float64{0.0, 2.7, -3.3, 9.6, -13.8, 26.0, 4.9, 2.6, 5.1, 1.1}
	sorted := heapSort[float64](slice)
	fmt.Println("After heapSort on slice: ", sorted)

	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = 100.0 * rand.Float64()
	}
	start := time.Now()
	largeSorted := heapSort[float64](data)
	elapsed := time.Since(start)
	fmt.Println("Time for heapSort of 50 million floats: ", elapsed)
	if !IsSorted[float64](largeSorted) {
		fmt.Println("largeSorted is not sorted.")
	}
}
/* Output
Elapsed time for regular quicksort =  5.382400384s  (from Chapter 1)
Elapsed time for concurrent quicksort =  710.431619ms (from Chapter 1)

After heapSort on slice:  [-13.8 -3.3 0 1.1 2.6 2.7 4.9 5.1 9.6 26]
Time for heapSort of 50 million floats:  19.202580166s
*/
