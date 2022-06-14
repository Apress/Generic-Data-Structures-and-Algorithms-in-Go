package main 

import (
	"fmt" 
	"time"
)

// Brute Force solution
func KnapSackBF(weightLimit int, weights []int, profits []int, n int) int {
	if n == 0 || weightLimit == 0 {
		return 0
	}
	if weights[n - 1] > weightLimit {
		return KnapSackBF(weightLimit, weights, profits, n - 1)
	} else {
		// Assume that we include object n - 1
		value1 := profits[n - 1] + KnapSackBF(weightLimit - weights[n - 1], weights, profits, n - 1)
		// Assume that we do not include object n - 1
		value2 := KnapSackBF(weightLimit, weights, profits, n - 1)
		if value1 >= value2 {
			return value1
		} else {
			return value2
		}
	}
}

// Dynamic Programming solution
func KnapSackDP(weightLimit int, weights []int, profits []int) int {
	n := len(weights)
	if weightLimit <= 0 || n == 0 || len(profits) != n {
		return 0
	}

	// Create a (n + 1 x weighlimit + 1) table 
	table := make([][]int, n + 1)
	for row := 0; row < n + 1; row++ {
		table[row] = make([]int, weightLimit + 1)
	}

	for i := 0; i < n + 1; i++ {
		for w := 0; w < weightLimit + 1; w++ {
			if i == 0 || w == 0 {
				table[i][w] = 0 
			} else if weights[i - 1] <= w {
				// Include item i
				wt := w - weights[i - 1]
				profit1 := profits[i - 1] + table[i - 1][wt]
				// Exclude item i
				profit2 := table[i  - 1][w]
				if profit1 >= profit2 {
					table[i][w] = profit1
				} else {
					table[i][w] = profit2
				}
			} else {
				// Exclude item
				table[i][w] = table[i - 1][w]
			}
		}
	}
	return table[n][weightLimit]
}

func main() {
	weights := []int{4, 6, 2, 8}
	profits := []int{12, 15, 9, 21}
	fmt.Println("Solution 1 = ", KnapSackBF(10, weights, profits, 4))

	weights1 := []int{4, 6, 2, 8, 1, 17, 23, 10, 4, 8}
	profits1 := []int{12, 15, 9, 21, 5, 8, 20, 6, 1, 15}
	result := KnapSackBF(20, weights1, profits1, 10)
	fmt.Println("Solution 2 = ", result)

	weights2 := []int{}
	for i := 0; i < 800; i++ {
		weights2 = append(weights2, 2 * i)
	}
	profits2 := []int{}
	for i := 0; i < 800; i++ {
		profits2 = append(profits2, 3 * i)
	}
	
	start := time.Now()
	result2 := KnapSackBF(400, weights2, profits2, 800)
	elapsed := time.Since(start)
	fmt.Println("Solution 3 = ", result2)
	fmt.Println("Time for solution3 (brute force): ", elapsed)

	start = time.Now()
	result3 := KnapSackDP(400, weights2, profits2)
	elapsed = time.Since(start)
	fmt.Println("Solution 3 = ", result3)
	fmt.Println("Time for solution3 (dynamic programming): ", elapsed)
}
/* Output
Solution 1 =  30
Solution 2 =  57
Solution 3 =  600
Time for solution3 (brute force):  1m10.248200934s
Solution 3 =  600
Time for solution3 (dynamic prograamming):  1.621038ms
*/
