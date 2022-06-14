package main

import (
	"fmt"
	"math/rand"
	"time"
	"example.com/floatset"
)

const (
	size = 1_000_000
)

var dataSet []float64

func main() {
	mySet := floatset.NewSet()

	dataSet = make([]float64, size)
	for i := 0; i < size; i++ {
		dataSet[i] = 100.0 * rand.Float64()
	}
	// Time construction of Set
	start := time.Now()
	for i := 0; i < size; i++ {
		mySet.Add(dataSet[i])
	}
	elapsed := time.Since(start)
	fmt.Printf("\nTime to build Set with %d numbers: %s", size, elapsed)

	// Time to test the presence of all numbers in dataSet
	start = time.Now()
	for i := 0; i < len(dataSet); i++ {
		if !mySet.IsPresent(dataSet[i]) {
			fmt.Println("%f not present", dataSet[i])
		}
	}
	elapsed = time.Since(start)
	fmt.Printf("\nTime to test the presence of all numbers in Set: %s", elapsed)

	avlSet := floatset.AVLTree{nil, 0}
	// Time construction of avlSet
	start = time.Now()
	for i := 0; i < size; i++ {
		avlSet.Insert(dataSet[i])
	}
	elapsed = time.Since(start)
	fmt.Printf("\n\nTime to build avlSet with %d numbers: %s", size, elapsed)

	// Time to test the presence of all numbers in avlSet
	start = time.Now()
	for i := 0; i < len(dataSet); i++ {
		if !mySet.IsPresent(dataSet[i]) {
			fmt.Println("%f not present", dataSet[0])
		}
	}
	elapsed = time.Since(start)
	fmt.Printf("\nTime to test the presence of all numbers in avlSet: %s", elapsed)

	// Use concurrent processing to construct concurrent avl trees
	start = time.Now()
	floatset.BuildConcurrentSet(dataSet)
	elapsed = time.Since(start)
	fmt.Printf("\n\nTime to build concurrent (%d) avlSet with %d numbers: %s", floatset.Concurrent, size, elapsed)

	// Test every number in dataSet against the concurrent set
	start = time.Now()
	for i := 0; i < len(dataSet); i++ {
		if !floatset.IsPresent(dataSet[i]) {
			fmt.Println("%f not present", dataSet[i])
		}
	}
	elapsed = time.Since(start)
	fmt.Printf("\nTime to test the presence of all numbers in concurrent (%d) avlSet: %s", floatset.Concurrent, elapsed)
}

/*
On iMac Pro witg 32G Ram and 3.2 GHz 8-Core Intel Xeon W
Time to build Set with 1000000 numbers: 184.442966ms
Time to test the presence of all numbers in Set: 105.600217ms

Time to build avlSet with 1000000 numbers: 819.517251ms
Time to test the presence of all numbers in avlSet: 103.422116ms

Time to build concurrent (32) avlSet with 1000000 numbers: 184.681628ms
Time to test the presence of all numbers in concurrent (32) avlSet: 66.183935ms

On iMac Pro Apple M1 Max with 32G Ram
Time to build Set with 1000000 numbers: 90.186209ms
Time to test the presence of all numbers in Set: 44.667542ms

Time to build avlSet with 1000000 numbers: 421.970625ms
Time to test the presence of all numbers in avlSet: 39.154042ms

Time to build concurrent (32) avlSet with 1000000 numbers: 172.478583ms
Time to test the presence of all numbers in concurrent (32) avlSet: 47.972875ms
*/
