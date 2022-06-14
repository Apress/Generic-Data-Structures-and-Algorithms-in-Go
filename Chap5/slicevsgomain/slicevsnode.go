package main

import (
	"example.com/nodestack"
	"example.com/slicestack"
	"time"
	"fmt"
)

const size = 10_000_000

func main() {
	nodeStack := nodestack.Stack[int]{}
	sliceStack := slicestack.Stack[int]{}

	// Benchmark nodeStack
	start := time.Now()
	for i := 0; i < size; i++ {
		nodeStack.Push(i)
	}
	elapsed := time.Since(start)
	fmt.Println("\nTime for 10 million Push() operations on nodeStack: ", elapsed)

	start = time.Now()
	for i := 0; i < size; i++ {
		nodeStack.Pop()
	}
	elapsed = time.Since(start)
	fmt.Println("\nTime for 10 million Pop() operations on nodeStack: ", elapsed)

	// Benchmark sliceStack
	start = time.Now()
	for i := 0; i < size; i++ {
		sliceStack.Push(i)
	}
	elapsed = time.Since(start)
	fmt.Println("\nTime for 10 million Push() operations on sliceStack: ", elapsed)

	start = time.Now()
	for i := 0; i < size; i++ {
		sliceStack.Pop()
	}
	elapsed = time.Since(start)
	fmt.Println("\nTime for 10 million Pop() operations on sliceStack: ", elapsed)
}
/* Output
Time for 10 million Push() operations on nodeStack:  616.365084ms

Time for 10 million Pop() operations on nodeStack:  29.104829ms

Time for 10 million Push() operations on sliceStack:  148.623915ms

Time for 10 million Pop() operations on sliceStack:  11.485335ms
*/