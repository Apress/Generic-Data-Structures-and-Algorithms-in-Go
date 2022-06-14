package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"image/color"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

const (
	NUMCITIES = 29
)

type Point struct {
	x float64
	y float64
}

var cities []Point

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
	DrawTour(cities, status.bestTour)
}

func definePoints(cities []Point, tour []int) plotter.XYs {
	pts := make(plotter.XYs, len(cities) + 1)
	pts[0].X = cities[0].x
	pts[0].Y = cities[0].y
	for i := 1; i < len(cities); i++ {
		pts[i].X = cities[tour[i]].x
		pts[i].Y = cities[tour[i]].y
	}
	pts[len(cities)].X = cities[0].x
	pts[len(cities)].Y = cities[0].y
	return pts
}

func DrawTour(cities []Point, tour []int) {
	data := definePoints(cities, tour) // plotter.XYs
	p := plot.New()
	p.Title.Text = "TSP Tour"
	lines, points, err := plotter.NewLinePoints(data)
	if err != nil {
		panic(err)
	}
	lines.Color = color.RGBA{R: 255, A: 255}
	points.Shape = draw.PyramidGlyph{}
	points.Color = color.RGBA{B: 255, A: 255}
	p.Add(lines, points)
	// Save the plot to a PNG file.
	if err := p.Save(6*vg.Inch, 6*vg.Inch, "tour.png"); err != nil {
		panic(err)
	}
}

func main() {
    cities = []Point{}
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
    fmt.Printf("\nInverse Operations: %d  Swap Operations: %d  Insert Operations: %d  Downhill moves: %d  Uphill moves, %d", 
    		status.inverseOps, status.swapOps, status.insertOps, status.downhillMoves, status.uphillMoves)
}