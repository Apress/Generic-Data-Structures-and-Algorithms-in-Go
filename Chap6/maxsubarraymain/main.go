package main 

import (
	"fmt"
	"example.com/deque"
	"time"
	"math/rand"
)

const size = 1_000_000

func MaxSubarrayBruteForce(input []int, k int) (output []int) {
	for first := 0; first <= len(input) - k; first++ {
		max := input[first]
		for second := 0; second < k; second++ {
			if input[first + second] > max {
				max = input[first + second]
			}
		}
		output = append(output, max)
	}
	return output
}

func MaxSubarrayUsingDeque(input []int, k int) (output []int) {
	
	deque := deque.Deque[int]{} 


	var index int
	// First window
	for index = 0; index < k;  index++ {
		for {
			if deque.Empty() || input[index] < input[deque.Last()] { 
				break
			}
			deque.RemoveLast()
		}
		deque.InsertBack(index)
	}

	for ; index < len(input); index++ {
		output = append(output, input[deque.First()])

		// Remove elements out of the window
		for {
			if deque.Empty() || deque.First() > index - k {
				break
			}
			deque.RemoveFirst()
		}
		// Remove values smaller than the element currently being added
		for {
			if deque.Empty() || input[index] < input[deque.Last()] {
				break
			}
			deque.RemoveLast()
		}
		deque.InsertBack(index)
	}
	output = append(output, input[deque.First()])
	return output
}

func main() {
	input := []int{9, 1, 1, 0, 0, 0, 1, 0, 6, 8}
	output1 := MaxSubarrayBruteForce(input, 3)
	fmt.Println("Output = ", output1)

	output2 := MaxSubarrayUsingDeque(input, 3) 
	fmt.Println("Output = ", output2) 

	// Benchmark performance of two algorithms
	input = []int{}
	for i := 0; i < size; i++ {
		input = append(input, rand.Intn(1000))
	}
	start := time.Now()
	MaxSubarrayUsingDeque(input, 10000)
	elapsed := time.Since(start)
	fmt.Println("Using Deque: ", elapsed)

	start = time.Now()
	MaxSubarrayBruteForce(input, 10000)
	elapsed = time.Since(start)
	fmt.Println("Using Brute Force: ", elapsed)
}
/* Output
Output =  [9 1 1 0 1 1 6 8]
Output =  [9 1 1 0 1 1 6 8]
Using Deque:  21.873658ms
Using Brute Force:  6.042102028s
*/