package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Graph [][]int

type TourCost struct {
	cost int
	tour []int
}

var minimumTourCost TourCost
var graph Graph

func Permutations(data []int, operation func([]int)) {
	permute(data, operation, 0)
}

func permute(data []int, operation func([]int), step int) {
	if step > len(data) {
		operation(data)
		return
	}
	permute(data, operation, step+1)
	for k := step + 1; k < len(data); k++ {
		data[step], data[k] = data[k], data[step]
		permute(data, operation, step+1)
		data[step], data[k] = data[k], data[step]
	}
}

func TSP(graph Graph, numCities int) {
	tour := []int{}
	for i := 1; i < numCities; i++ {
		tour = append(tour, i)
	}
	minimumTourCost = TourCost{32767, []int{}}
	Permutations(tour, func(tour []int) {
		// Compute cost of tour
		cost := graph[0][tour[0]]
		for i := 0; i < len(tour)-1; i++ {
			cost += graph[tour[i]][tour[i+1]]
		}
		cost += graph[tour[len(tour)-1]][0]
		if cost < minimumTourCost.cost {
			minimumTourCost.cost = cost

			var tourCopy []int
			tourCopy = append(tourCopy, 0)
			tourCopy = append(tourCopy, tour...)
			tourCopy = append(tourCopy, 0)
			minimumTourCost.tour = tourCopy
		}
	})
}

func main() {
	graph = Graph{{0, 5, 3, 9}, {5, 0, 2, 1}, {3, 2, 0, 6}, {9, 1, 6, 0}}
	TSP(graph, 4)
	fmt.Printf("\nOptimum tour cost: %d  An Optimum Tour %v", minimumTourCost.cost, minimumTourCost.tour)

	numCities := 14
	graph2 := make([][]int, numCities)
	for i := 0; i < numCities; i++ {
		graph2[i] = make([]int, numCities)
	}
	for row := 0; row < numCities; row++ {
		for col := 0; col < numCities; col++ {
			graph2[row][col] = rand.Intn(9) + 2
		}
	}
	// Create a short path for test purposes
	for i := 0; i < numCities-1; i++ {
		graph2[i][i+1] = 1
	}
	graph2[numCities-1][0] = 1

	start := time.Now()
	TSP(graph2, numCities)
	elapsed := time.Since(start)
	fmt.Printf("\nOptimum tour cost: %d  An Optimum Tour %v", minimumTourCost.cost, minimumTourCost.tour)
	fmt.Println("\nComputation time: ", elapsed)
}
/* Output
Optimum tour cost: 15  An Optimum Tour [0 1 3 2 0]
Optimum tour cost: 14  An Optimum Tour [0 1 2 3 4 5 6 7 8 9 10 11 12 13 0]
Computation time:  2m15.918717943s
*/
