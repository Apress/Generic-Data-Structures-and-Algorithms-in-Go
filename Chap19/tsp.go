package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	NUMCITIES = 29
)

type Point struct {
	x float64
	y float64
}

type Graph [][]float64

var graph Graph

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (pt Point) distance(other Point) float64 {
	dx := pt.x - other.x
	dy := pt.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func createGraph(numCities int, cities []Point, graph [][]float64) {
	for row := 0; row < numCities; row++ {
		for col := 0; col < numCities; col++ {
			if row == col {
				graph[row][col] = 0.0
			} else {
				graph[row][col] = cities[row].distance(cities[col])
			}
		}
	}
}

func cost(graph [][]float64, tour []int) float64 {
	result := 0.0
	for index := 0; index < len(tour) - 2; index++ {
		result += graph[tour[index]][tour[index+1]]
	}
	result += graph[tour[NUMCITIES - 1]][tour[0]]
	return result
}

func randomFrom(min int, max int) int {
	return rand.Intn(max - min) + min
}

func inverseOperation(tour []int) []int {
	/*
	   Choose city i randomly from 1 to count - 1.
	   Choose city j randomly from 1 to count - 1
	   let first be the minimum of index i and j.
	   let second be the larger of index i and j.
	   reverse the order of cities in the tour from index first to index second
	   Consider tour = [0, 3, 2, 1, 5, 4] and first = 1 and second = 4
	   The segment 3, 2, 1, 5 is replaced by 5, 1, 2, 3 and the new tour is
	   [0, 5, 1, 2, 3, 4].
	*/
	// Choose first and second
	firstIndex := randomFrom(1, len(tour) - 1) // number between 1 and 28
	secondIndex := randomFrom(1, len(tour) - 1)
	for firstIndex == secondIndex {
		firstIndex = randomFrom(1, len(tour) - 1)
		secondIndex = randomFrom(1, len(tour) - 1)
	}
	if firstIndex > secondIndex {
		firstIndex, secondIndex = secondIndex, firstIndex
	}
	result := deepcopy(tour[:firstIndex])
	for index := 0; index < (secondIndex - firstIndex + 1); index += 1 {
		result = append(result, tour[secondIndex - index])
	}
	for index := secondIndex + 1; index < len(tour); index += 1 {
		result = append(result, tour[index])
	}
	return result
}

func swap(tour []int) []int {
	/*
	   Swap the city in position first with city in position second
	   Consider tour [0, 3, 2, 1, 5, 4] and first = 1 and second = 4
	   The new tour would be [0, 5, 2, 1, 3, 4]
	*/

	// Choose first and second
	firstIndex := randomFrom(1, len(tour) -  1)
    secondIndex := randomFrom(1, len(tour) - 1)
    for firstIndex == secondIndex {
        firstIndex = randomFrom(1, len(tour) - 1)
        secondIndex = randomFrom(1, len(tour) - 1)
    }
    if firstIndex > secondIndex {
        firstIndex, secondIndex = secondIndex, firstIndex
    }
    result := deepcopy(tour)
    result[firstIndex], result[secondIndex] = result[secondIndex], result[firstIndex]
	return result
}

func insert(tour []int) []int {
	/*
	   It means to move the city in position second to position first.
	   Consider tour [0, 3, 2, 1, 5, 4] and first = 1 and second = 4
	   The new tour would be [0, 5, 3, 2, 1, 4]
	*/

	// Choose first and second
	// Choose first and second
    firstIndex := randomFrom(1, len(tour) - 1)
    secondIndex := randomFrom(1, len(tour) - 1)
    for firstIndex == secondIndex {
        firstIndex = randomFrom(1, len(tour) - 1)
        secondIndex = randomFrom(1, len(tour) - 1)
    }
    if firstIndex > secondIndex {
        firstIndex, secondIndex = secondIndex, firstIndex
    }
	result := []int{}
	for index := 0; index < len(tour) + 1; index += 1 {
		if index < firstIndex {
			result = append(result, tour[index])
		} else if index == firstIndex {
			result = append(result, tour[secondIndex])
		} else if index > firstIndex && index != secondIndex + 1 {
			result = append(result, tour[index-1])
		}
	}
	return result
}

type Status struct {
	tour           []int
	bestTour       []int
	bestCostToDate float64
	previousCost   float64
	temperature    float64
	downhillMoves  int
	uphillMoves    int
	rejectedMoves  int
	inverseOps     int
	swapOps        int
	insertOps      int
}

var status Status

func deepcopy(tour []int) []int {
	result := []int{}
	for i := range tour {
		result = append(result, tour[i])
	}
	return result
}

func simulatedAnnealing(graph [][]float64) {
	for i := 0; i < NUMCITIES; i++ {
		status.tour = append(status.tour, i)
	}
    status.tour = append(status.tour, 0)
    fmt.Printf("\n\nCost of initial tour %v is %f\n\n", status.tour, cost(graph, status.tour))
	status.bestTour = deepcopy(status.tour)
	status.bestCostToDate = cost(graph, status.bestTour)
	status.previousCost = status.bestCostToDate
	numberIterationsAtTemperature := 5000
	lowestTemperature := 5.0
	for status.temperature >= lowestTemperature {
		for iteration := 0; iteration < numberIterationsAtTemperature; iteration += 1 {
			tour1 := inverseOperation(status.tour)
			cost1 := cost(graph, tour1)
			tour2 := swap(status.tour)
			cost2 := cost(graph, tour2)
			tour3 := insert(status.tour)
			cost3 := cost(graph, tour3)
			newCost1 := math.Min(cost1, cost2)
			newCost := math.Min(newCost1, cost3)

			if newCost == cost1 {
				status.inverseOps += 1
				// Determine whether to accept this tour1
				if newCost < status.previousCost {
					status.downhillMoves += 1
					status.previousCost = newCost
					status.tour = deepcopy(tour1)
					if newCost < status.bestCostToDate {
						status.bestCostToDate = newCost
						status.bestTour = deepcopy(tour1)
						fmt.Printf("\nLowest cost tour to-date = %0.2f at Temperature = %0.2f  Best tour: %v", status.bestCostToDate, status.temperature, status.bestTour)
					}
				} else {
					metropolis := math.Exp((status.previousCost - newCost) / status.temperature)
					r := rand.Float64()
					if r <= metropolis { // Accept uphill move
						status.uphillMoves += 1
						status.previousCost = newCost
						status.tour = deepcopy(tour1)
						if newCost < status.bestCostToDate {
							status.bestCostToDate = newCost
							status.bestTour = deepcopy(tour1)
							fmt.Printf("\nLowest cost tour to-date = %0.2f at Temperature = %0.2f  Best tour: %v", status.bestCostToDate, status.temperature, status.bestTour)
						} else {
							status.rejectedMoves += 1
						}
					}
				}
			} else if newCost == cost2 {
				status.swapOps += 1
				// Determine whether to accept this tour2
				if newCost < status.previousCost {
					status.downhillMoves += 1
					status.previousCost = newCost
					status.tour = deepcopy(tour2)
					if newCost < status.bestCostToDate {
						status.bestCostToDate = newCost
						status.bestTour = deepcopy(tour2)
						fmt.Printf("\nLowest cost tour to-date = %0.2f at Temperature = %0.2f  Best tour: %v", status.bestCostToDate, status.temperature, status.bestTour)
                    }
				} else {
					metropolis := math.Exp((status.previousCost - newCost) / status.temperature)
					r := rand.Float64()
					if r <= metropolis { // Accept uphill move
						status.uphillMoves += 1
						status.previousCost = newCost
						status.tour = deepcopy(tour2)
						if newCost < status.bestCostToDate {
							status.bestCostToDate = newCost
							status.bestTour = deepcopy(tour2)
							fmt.Printf("\nLowest cost tour to-date = %0.2f at Temperature = %0.2f  Best tour: %v", status.bestCostToDate, status.temperature, status.bestTour)
						} else {
							status.rejectedMoves += 1
						}
					}
                }
			} else if newCost == cost3 {
				status.insertOps += 1
                // Determine whether to accept this tour3
                if newCost < status.previousCost {
                    status.downhillMoves += 1
                    status.previousCost = newCost
                    status.tour = deepcopy(tour3)
                    if newCost < status.bestCostToDate {
                        status.bestCostToDate = newCost
                        status.bestTour = deepcopy(tour3)
                        fmt.Printf("\nLowest cost tour to-date = %0.2f at Temperature = %0.2f  Best tour: %v", status.bestCostToDate, status.temperature, status.bestTour)
                    }
                } else {
                    metropolis := math.Exp((status.previousCost - newCost) / status.temperature)
                    r := rand.Float64()
                    if r <= metropolis { // Accept uphill move
                        status.uphillMoves += 1
                        status.previousCost = newCost
                        status.tour = deepcopy(tour3)
                        if newCost < status.bestCostToDate {
                            status.bestCostToDate = newCost
                            status.bestTour = deepcopy(tour3)
                            fmt.Printf("\nLowest cost tour to-date = %0.2f at Temperature = %0.2f  Best tour: %v", status.bestCostToDate, status.temperature, status.bestTour)
                        } else {
                            status.rejectedMoves += 1
                        }
                    }
                }
			}
		} 
		// Cooling curve
        if status.temperature >= 1000.0 {
            status.temperature *= 0.90
        } else if status.temperature >= 500 {
            status.temperature *= 0.94
        } else if status.temperature >= 200 {
            status.temperature *= 0.97
        } else if status.temperature >= 50 {
            status.temperature *= 0.98
        } else {
             status.temperature *= 0.99
        } 
	}
}

func main() {
	
    cities := []Point{}
    // Known solution: 9074.15
    pt1 := Point{1150.0,1760.0}
    cities = append(cities, pt1)
    pt2 := Point{630.0, 1660.0}
    cities = append(cities, pt2)
    pt3 := Point{40.0, 2090.0}
    cities = append(cities, pt3)
    pt4 := Point{750.0, 1100.0}
    cities = append(cities, pt4)
    pt5 := Point{750.0, 2030.0}
    cities = append(cities, pt5)
    pt6 := Point{1030.0, 2070.0}
    cities = append(cities, pt6)
    pt7 := Point{1650.0, 650.0}
    cities = append(cities, pt7)
    pt8 := Point{1490.0, 1630.0}
    cities = append(cities, pt8)
    pt9 := Point{790.0, 2260.0}
    cities = append(cities, pt9)
    pt10 := Point{710.0, 1310.0}
    cities = append(cities, pt10)
    pt11 := Point{840.0, 550.0}
    cities = append(cities, pt11)
    pt12 := Point{1170.0, 2300.0}
    cities = append(cities, pt12)
    pt13 := Point{970.0, 1340.0}
    cities = append(cities, pt13)
    pt14 := Point{510.0, 700.0}
    cities = append(cities, pt14)
    pt15 := Point{750.0, 900.0}
    cities = append(cities, pt15)
    pt16 := Point{1280.0, 1200.0}
    cities = append(cities, pt16)
    pt17 := Point{230.0, 590.0}
    cities = append(cities, pt17)
    pt18 := Point{460.0, 860.0}
    cities = append(cities, pt18)
    pt19 := Point{1040.0, 950.0}
    cities = append(cities, pt19)
    pt20 := Point{590.0, 1390.0}
    cities = append(cities, pt20)
    pt21 := Point{830.0, 1770.0}
    cities = append(cities, pt21)
    pt22 := Point{490.0, 500.0}
    cities = append(cities, pt22)
    pt23 := Point{1840.0, 1240.0}
    cities = append(cities, pt23)
    pt24 := Point{1260.0, 1500.0}
    cities = append(cities, pt24)
    pt25 := Point{1280.0, 790.0}
    cities = append(cities, pt25)
    pt26 := Point{490.0, 2130.0}
    cities = append(cities, pt26)
    pt27 := Point{1460.0, 1420.0}
    cities = append(cities, pt27)
    pt28 := Point{1260.0, 1910.0}
    cities = append(cities, pt28)
    pt29 := Point{360.0, 1980.0}
    cities = append(cities, pt29)

    graph := make([][]float64, NUMCITIES)
    for i:=0; i < NUMCITIES ; i++ {
        graph[i] = make([]float64, NUMCITIES)
    }
    createGraph(NUMCITIES, cities, graph)
    
    status.temperature = 2000.0
    simulatedAnnealing(graph)
    fmt.Printf("\nInverse Operations: %d  Swap Operations: %d  Insert Operations: %d  Downhill moves: %d  Uphill moves, %d", status.inverseOps, status.swapOps, status.insertOps, status.downhillMoves, status.uphillMoves)
}

/*
Cost of initial tour [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 0] is 25814.877363


Lowest cost tour to-date = 25669.20 at Temperature = 2000.00  Best tour: [0 1 2 3 26 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 27 28 0]
Lowest cost tour to-date = 25456.00 at Temperature = 2000.00  Best tour: [0 1 2 3 26 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 25 21 22 23 24 20 27 28 0]
Lowest cost tour to-date = 24872.68 at Temperature = 2000.00  Best tour: [0 1 2 3 26 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 27 21 22 23 24 20 25 28 0]
Lowest cost tour to-date = 24249.12 at Temperature = 2000.00  Best tour: [0 1 2 3 26 4 5 6 7 8 9 15 10 11 12 13 14 16 17 18 19 27 21 22 23 24 20 25 28 0]
Lowest cost tour to-date = 22921.49 at Temperature = 2000.00  Best tour: [0 1 2 3 26 4 5 6 7 8 9 15 27 19 18 17 16 14 13 12 11 10 21 22 23 24 20 25 28 0]
Lowest cost tour to-date = 22479.54 at Temperature = 2000.00  Best tour: [0 1 2 3 26 4 5 6 7 8 9 15 27 19 18 13 17 16 14 12 11 10 21 22 23 24 20 25 28 0]
Lowest cost tour to-date = 21640.15 at Temperature = 2000.00  Best tour: [0 1 2 3 26 4 5 6 24 8 9 15 27 19 18 13 17 16 14 12 11 10 21 22 23 7 20 25 28 0]
Lowest cost tour to-date = 21208.12 at Temperature = 2000.00  Best tour: [0 1 2 3 15 26 4 5 6 24 8 9 27 19 18 13 17 16 14 12 11 10 21 22 23 7 20 25 28 0]
Lowest cost tour to-date = 18984.25 at Temperature = 2000.00  Best tour: [0 1 7 23 15 22 21 10 9 12 14 16 17 13 18 19 27 11 8 24 6 5 4 26 3 2 20 25 28 0]
Lowest cost tour to-date = 18849.81 at Temperature = 2000.00  Best tour: [0 1 7 23 15 22 21 16 10 9 12 14 17 13 18 19 27 11 8 24 6 5 4 26 3 2 20 25 28 0]
Lowest cost tour to-date = 18735.36 at Temperature = 2000.00  Best tour: [0 1 27 23 18 13 17 14 12 9 5 11 7 19 15 22 21 16 10 24 6 8 4 26 3 2 20 25 28 0]
Lowest cost tour to-date = 18218.92 at Temperature = 2000.00  Best tour: [0 6 24 10 16 21 22 15 19 7 11 5 9 12 14 17 13 18 23 27 1 8 4 26 3 2 20 25 28 0]
Lowest cost tour to-date = 18110.89 at Temperature = 2000.00  Best tour: [0 6 24 10 16 21 22 15 12 7 11 1 18 13 17 14 19 9 5 23 27 8 4 26 3 2 20 25 28 0]
Lowest cost tour to-date = 17873.48 at Temperature = 2000.00  Best tour: [0 6 24 10 16 21 22 15 12 7 11 4 1 18 13 17 14 19 9 8 23 27 5 26 3 2 20 25 28 0]
Lowest cost tour to-date = 17544.27 at Temperature = 2000.00  Best tour: [0 6 24 10 16 21 17 13 18 1 4 11 7 12 15 22 14 19 9 8 23 27 5 26 3 2 20 25 28 0]
Lowest cost tour to-date = 17327.28 at Temperature = 2000.00  Best tour: [0 6 24 10 16 21 17 13 3 1 4 11 7 12 15 22 14 19 9 8 23 27 5 26 18 2 20 25 28 0]
Lowest cost tour to-date = 17300.34 at Temperature = 2000.00  Best tour: [0 6 24 10 16 21 17 13 12 3 1 8 11 7 15 22 14 19 9 4 23 27 5 26 18 25 2 20 28 0]
Lowest cost tour to-date = 16585.81 at Temperature = 2000.00  Best tour: [0 6 24 10 16 21 17 13 12 3 1 8 11 4 9 19 14 22 15 7 23 27 5 26 18 25 2 20 28 0]
Lowest cost tour to-date = 15655.98 at Temperature = 2000.00  Best tour: [0 6 24 10 16 21 17 13 12 3 1 20 8 11 4 9 19 14 22 15 7 23 27 5 26 18 25 2 28 0]
Lowest cost tour to-date = 15655.04 at Temperature = 2000.00  Best tour: [0 6 24 10 16 21 17 13 12 3 1 20 8 11 4 9 19 14 22 7 15 23 27 5 26 18 25 2 28 0]
Lowest cost tour to-date = 15439.96 at Temperature = 2000.00  Best tour: [0 6 24 10 16 21 17 13 12 3 1 18 26 5 27 23 15 7 22 14 19 9 4 11 8 20 25 2 28 0]
Lowest cost tour to-date = 15315.97 at Temperature = 2000.00  Best tour: [0 6 24 10 16 21 17 13 12 3 1 23 27 5 26 18 15 7 22 14 19 9 4 11 8 20 25 2 28 0]
Lowest cost tour to-date = 14339.52 at Temperature = 2000.00  Best tour: [0 11 7 22 14 9 3 1 28 2 25 8 20 19 12 15 26 10 16 21 17 13 24 6 18 4 5 27 23 0]
Lowest cost tour to-date = 14118.61 at Temperature = 2000.00  Best tour: [0 6 22 15 24 10 21 14 18 3 12 16 17 13 7 26 23 27 11 4 25 5 8 1 19 9 2 28 20 0]
Lowest cost tour to-date = 14082.51 at Temperature = 1800.00  Best tour: [0 27 7 6 22 26 23 4 9 13 17 12 15 24 18 10 3 14 16 21 19 11 5 8 1 28 2 25 20 0]
Lowest cost tour to-date = 14009.54 at Temperature = 1800.00  Best tour: [0 27 7 6 22 26 23 4 9 13 17 12 15 24 18 10 14 3 16 21 19 11 5 8 1 28 2 25 20 0]
Lowest cost tour to-date = 13848.14 at Temperature = 1800.00  Best tour: [0 27 7 6 22 26 23 4 11 9 13 16 21 17 18 15 24 10 14 3 19 12 25 2 1 28 8 5 20 0]
Lowest cost tour to-date = 13659.70 at Temperature = 1800.00  Best tour: [0 27 7 23 12 15 26 19 1 28 20 17 13 21 16 9 2 25 4 5 8 11 3 14 18 10 24 6 22 0]
Lowest cost tour to-date = 13386.10 at Temperature = 1800.00  Best tour: [0 27 7 23 26 15 12 1 19 28 20 17 13 21 16 9 2 25 4 5 8 11 18 3 14 10 24 6 22 0]
Lowest cost tour to-date = 13074.00 at Temperature = 1620.00  Best tour: [0 27 20 9 3 19 14 24 26 18 10 13 17 16 21 1 28 25 2 8 4 5 11 12 15 6 22 7 23 0]
Lowest cost tour to-date = 12666.24 at Temperature = 1180.98  Best tour: [0 20 28 25 2 4 7 8 5 11 27 1 19 6 22 24 18 14 10 21 13 16 17 3 9 12 15 26 23 0]
Lowest cost tour to-date = 12643.30 at Temperature = 794.53  Best tour: [0 27 5 11 8 20 25 2 28 4 19 1 9 7 26 22 24 6 14 10 13 16 21 17 3 23 15 12 18 0]
Lowest cost tour to-date = 12487.86 at Temperature = 746.86  Best tour: [0 23 15 12 9 3 18 6 24 22 26 5 11 20 1 25 8 4 28 2 19 17 21 13 16 10 14 7 27 0]
Lowest cost tour to-date = 12162.28 at Temperature = 702.05  Best tour: [0 20 5 11 4 25 2 28 8 9 19 3 17 16 21 13 10 14 15 12 1 18 6 24 26 22 7 23 27 0]
Lowest cost tour to-date = 11764.31 at Temperature = 702.05  Best tour: [0 27 5 11 8 4 28 2 25 20 3 12 15 18 7 26 23 1 19 9 16 21 10 13 17 14 24 6 22 0]
Lowest cost tour to-date = 11640.94 at Temperature = 484.32  Best tour: [0 7 15 19 14 3 12 20 8 4 5 23 26 22 6 24 18 10 21 13 16 17 9 1 28 2 25 11 27 0]
Lowest cost tour to-date = 11428.50 at Temperature = 484.32  Best tour: [0 5 4 25 2 20 19 12 3 14 9 17 16 13 21 10 18 24 6 22 7 23 26 15 1 28 8 11 27 0]
Lowest cost tour to-date = 11411.98 at Temperature = 469.79  Best tour: [0 27 23 26 7 22 15 19 3 9 14 13 17 16 21 10 6 24 18 12 20 1 11 5 4 25 28 2 8 0]
Lowest cost tour to-date = 11140.76 at Temperature = 469.79  Best tour: [0 27 23 7 26 15 22 6 24 18 10 14 21 13 16 3 17 19 9 12 1 20 11 5 4 25 2 28 8 0]
Lowest cost tour to-date = 11091.14 at Temperature = 455.70  Best tour: [0 27 11 5 8 4 28 2 25 20 1 12 24 6 10 21 16 17 13 18 9 19 3 14 15 23 7 22 26 0]
Lowest cost tour to-date = 10549.20 at Temperature = 346.44  Best tour: [0 27 11 5 4 8 25 28 2 19 9 14 3 17 16 10 21 13 18 24 6 22 26 15 23 7 12 1 20 0]
Lowest cost tour to-date = 10477.04 at Temperature = 346.44  Best tour: [0 5 27 11 4 8 25 28 2 19 9 14 3 17 21 16 13 10 18 24 6 22 26 15 7 23 12 1 20 0]
Lowest cost tour to-date = 10368.42 at Temperature = 325.96  Best tour: [0 27 5 11 8 2 28 25 4 20 1 19 9 3 18 12 10 21 17 16 13 14 24 6 22 26 7 23 15 0]
Lowest cost tour to-date = 10162.14 at Temperature = 325.96  Best tour: [0 27 11 5 8 2 28 25 4 20 12 1 19 9 3 14 10 21 17 16 13 18 24 6 22 26 7 23 15 0]
Lowest cost tour to-date = 9899.85 at Temperature = 247.81  Best tour: [0 5 27 11 8 25 2 28 4 20 1 19 9 12 15 18 14 13 3 17 16 21 10 24 6 22 26 7 23 0]
Lowest cost tour to-date = 9846.42 at Temperature = 226.17  Best tour: [0 7 23 26 15 22 6 24 18 3 13 21 16 17 10 14 12 9 19 1 20 28 2 25 4 8 5 11 27 0]
Lowest cost tour to-date = 9829.90 at Temperature = 179.14  Best tour: [0 23 15 7 26 22 6 24 18 14 10 21 16 13 17 3 9 12 19 1 20 11 8 28 2 25 4 5 27 0]
Lowest cost tour to-date = 9740.09 at Temperature = 179.14  Best tour: [0 23 7 26 15 22 6 24 18 14 10 21 16 13 17 3 9 12 19 1 20 11 8 28 2 25 4 5 27 0]
Lowest cost tour to-date = 9677.66 at Temperature = 175.56  Best tour: [0 23 7 22 26 15 12 3 14 17 16 13 21 10 6 24 18 9 19 1 20 4 28 2 25 8 11 5 27 0]
Lowest cost tour to-date = 9642.94 at Temperature = 175.56  Best tour: [0 23 7 26 22 15 12 3 14 17 16 13 21 10 6 24 18 9 19 1 20 4 28 2 25 8 11 5 27 0]
Lowest cost tour to-date = 9606.44 at Temperature = 175.56  Best tour: [0 23 7 26 22 6 24 15 12 14 21 16 17 13 10 18 3 9 19 1 20 4 28 2 25 8 11 5 27 0]
Lowest cost tour to-date = 9596.98 at Temperature = 175.56  Best tour: [0 23 7 26 22 6 24 10 13 17 16 21 14 12 15 18 3 9 19 1 20 4 28 2 25 8 11 5 27 0]
Lowest cost tour to-date = 9569.98 at Temperature = 175.56  Best tour: [0 27 11 5 8 4 28 2 25 20 1 19 9 3 10 21 16 13 17 14 18 12 15 24 6 22 7 26 23 0]
Lowest cost tour to-date = 9490.31 at Temperature = 172.05  Best tour: [0 27 20 5 11 8 4 25 2 28 1 19 9 3 12 15 18 14 17 16 21 13 10 24 6 22 26 7 23 0]
Lowest cost tour to-date = 9406.00 at Temperature = 161.93  Best tour: [0 23 7 26 22 6 24 15 18 10 16 21 13 17 14 3 9 12 19 1 20 4 28 2 25 8 11 5 27 0]
Lowest cost tour to-date = 9248.08 at Temperature = 161.93  Best tour: [0 23 7 26 22 6 24 15 18 10 21 16 13 17 14 3 9 12 19 1 20 4 28 2 25 8 11 5 27 0]
Lowest cost tour to-date = 9248.08 at Temperature = 161.93  Best tour: [0 27 5 11 8 25 2 28 4 20 1 19 12 9 3 14 17 13 16 21 10 18 15 24 6 22 26 7 23 0]
Lowest cost tour to-date = 9120.82 at Temperature = 161.93  Best tour: [0 27 5 11 8 25 2 28 4 20 1 19 9 12 3 14 17 13 16 21 10 18 15 24 6 22 26 7 23 0]
Lowest cost tour to-date = 9107.19 at Temperature = 93.85  Best tour: [0 27 5 11 8 25 2 28 4 20 1 19 9 12 3 14 17 13 16 21 10 18 24 6 22 26 15 23 7 0]
Lowest cost tour to-date = 9077.92 at Temperature = 93.85  Best tour: [0 27 5 11 8 25 2 28 4 20 1 19 9 12 3 14 17 13 16 21 10 18 24 6 22 15 26 7 23 0]
Lowest cost tour to-date = 9076.98 at Temperature = 72.17  Best tour: [0 23 15 26 7 22 6 24 18 10 21 16 13 17 14 3 12 9 19 1 20 4 28 2 25 8 11 5 27 0]
Lowest cost tour to-date = 9074.15 at Temperature = 72.17  Best tour: [0 27 5 11 8 25 2 28 4 20 1 19 9 3 14 17 13 16 21 10 18 24 6 22 7 26 15 12 23 0]
Inverse Operations: 748742  Swap Operations: 345285  Insert Operations: 625973  Downhill moves: 95259  Uphill moves, 96014
> Elapsed: 2.288s
*/
