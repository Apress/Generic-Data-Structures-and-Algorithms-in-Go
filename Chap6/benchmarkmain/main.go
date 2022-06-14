// We compare the performance of slicequeue and nodequeue
package main

import (
	"fmt"
	"example.com/nodequeue"
	"example.com/slicequeue"
	"time"
)

const size = 1_000_000

func main() {
	sliceQueue := slicequeue.Queue[int]{}
	nodeQueue := nodequeue.Queue[int]{}
	start := time.Now()
	for i := 0; i < size; i++ {
		sliceQueue.Insert(i)
	}
	elapsed := time.Since(start)
	fmt.Println("Time for inserting 1 million ints in sliceQueue is", elapsed)


	start = time.Now()
	for i := 0; i < size; i++ {
		nodeQueue.Insert(i)
	}
	elapsed = time.Since(start)
	fmt.Println("Time for inserting 1 million ints in nodeQueue is", elapsed)

	start = time.Now()
	for i := 0; i < size; i++ {
		sliceQueue.Remove()
	}
	elapsed = time.Since(start)
	fmt.Println("Time for removing 1 million ints from sliceQueue is", elapsed)

	start = time.Now()
	for i := 0; i < size; i++ {
		nodeQueue.Remove()
	}
	elapsed = time.Since(start)
	fmt.Println("Time for removing 1 million ints from nodeQueue is", elapsed)
}
/* Output
Time for inserting 1 million ints in sliceQueue is 18.841914ms
Time for inserting 1 million ints in nodeQueue is 30.275662ms
Time for removing 1 million ints from sliceQueue is 1.413447ms
Time for removing 1 million ints from nodeQueue is 2.818313ms
*/