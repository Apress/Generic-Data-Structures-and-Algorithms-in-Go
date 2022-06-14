package main  

import (
	"fmt"
	"time"
)

func FibonacciTopDown(n int) int64 {
    firstTwoCases := map[int]int64{
        0: 0,
        1: 1,
    }
    return computeFromCache(n, firstTwoCases)
}

func computeFromCache(n int, cache map[int]int64) int64 {
	// If answer already found for n, return it
    if val, found := cache[n]; found {
        return val
    }
    cache[n] = computeFromCache(n - 1, cache) + computeFromCache(n - 2, cache)
    return cache[n]
}

func FibonacciBottomUp(n int) int64 {
	table := []int64{0, 1}
	for i := 2; i <= n; i++ {
		table = append(table, table[i - 1] + table[i - 2])
	}
	return table[n]
}

func Fib(n int64) int64 {
	if n < 2 {
		return n
	}
	return Fib(n - 1) + Fib(n - 2)
}


func main() {
	fmt.Println("fib(7) = ", FibonacciTopDown(7))
	start := time.Now()
	fib40 := FibonacciTopDown(40)
	elapsed := time.Since(start)
	fmt.Println("Value of FibonacciTopDown(40): ", fib40)
	fmt.Println("Computation time: ", elapsed)

	fmt.Println("fib(7) = ", FibonacciBottomUp(7))
	start = time.Now()
	fib40 = FibonacciBottomUp(40)
	elapsed = time.Since(start)
	fmt.Println("\nValue of FibonacciBottomUp(40): ", fib40)
	fmt.Println("Computation time: ", elapsed)

	fmt.Println("fib(7) = ", Fib(7))
	start = time.Now()
	fib40 = Fib(40)
	elapsed = time.Since(start)
	fmt.Println("\nValue of Fib(40): ", fib40)
	fmt.Println("Computation time: ", elapsed)
}
/* Output
fib(7) =  13
Value of FibonacciTopDown(40):  102334155
Computation time:  36.136µs
fib(7) =  13

Value of FibonacciBottomUp(40):  102334155
Computation time:  7.377µs
fib(7) =  13

Value of Fib(40):  102334155
Computation time:  424.44211ms
*/
