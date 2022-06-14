package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

const (
	NUMCITIES          = 29
	EliteNumber        =  2 
	ToursPerGeneration = 100
	NumberGenerations  = 50000
	TournamentNumber   = 4
	ProbMutation       = 0.25
)

type Point struct {
	x float64
	y float64
}

type Graph [][]float64

var graph Graph
var population [][]int
var newpopulation [][]int

func (pt Point) distance(other Point) float64 {
	dx := pt.x - other.x
	dy := pt.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func CreateGraph(numCities int, cities []Point, graph [][]float64) {
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

func DeepCopy(tour []int) []int {
	result := []int{}
	for i := range tour {
		result = append(result, tour[i])
	}
	return result
}

func In(value int, values []int) (bool, int) {
	// Returns true if value in values
	// returns index of location or -1 if not found
	for index := 0; index < len(values); index++ {
		if values[index] == value {
			return true, index
		}
	}
	return false, -1
}

func Cost(tour []int) float64 {
	result := 0.0
	for index := 0; index < len(tour)-2; index++ {
		result += graph[tour[index]][tour[index+1]]
	}
	result += graph[tour[NUMCITIES-1]][tour[0]]
	return result
}

func CreateInitialPopulation() {
	firstCities := make([]int, NUMCITIES-1)
	for i := 1; i < NUMCITIES; i++ {
		firstCities[i-1] = i
	}
	for row := 0; row < ToursPerGeneration; row++ {
		rand.Shuffle(len(firstCities), func(i, j int) {
			firstCities[i], firstCities[j] = firstCities[j], firstCities[i]
		})
		population[row] = []int{0}
		for col := 1; col < NUMCITIES; col++ {
			population[row] = append(population[row], firstCities[col-1])
		}
	}
}

func ChooseEliteGroup() (elite [][]int) {
	// The population is sorted prior calling
    // this function 

	// Initialize elite
	elite = make([][]int, EliteNumber)
	for row := 0; row < EliteNumber; row++ {
		elite[row] = make([]int, EliteNumber)
	}

	for row := 0; row < EliteNumber; row++ {
		elite[row] = DeepCopy(population[row])
	}
	return elite
}

func FormNextGeneration() {
	elite := ChooseEliteGroup()
	// Move elite into newpopulation
	row := 0 // index into newpopulaton
	for ; row < EliteNumber; row++ {
		newpopulation[row] = DeepCopy(elite[row])
	}
	// Remove the first EliteNumber rows from population
	population = population[EliteNumber:]

	// Initialize group1 and group2
	group1 := make([][]int, TournamentNumber)
	for i := 0; i < TournamentNumber; i++ {
		group1[i] = make([]int, NUMCITIES)
	}
	group2 := make([][]int, TournamentNumber)
	for i := 0; i < TournamentNumber; i++ {
		group2[i] = make([]int, NUMCITIES)
	}

	MatingPoolSize := (ToursPerGeneration - EliteNumber) / 2
	for index := 0; index < MatingPoolSize; index++ {
		// Grap first group 
		indicesChosen := []int{}
		rowsChosen := 0;
		for {
			randomRow := rand.Intn(TournamentNumber)
			found, _ := In(randomRow, indicesChosen)
			if !found {
				indicesChosen = append(indicesChosen, randomRow)
				group1[rowsChosen] = DeepCopy(population[randomRow])
				rowsChosen += 1
			}
			if rowsChosen == TournamentNumber {
				break
			}
		}
		// Grap second group 
		indicesChosen = []int{}
		rowsChosen = 0;
		for {
			randomRow := rand.Intn(TournamentNumber)
			found, _ := In(randomRow, indicesChosen)
			if !found {
				indicesChosen = append(indicesChosen, randomRow)
				group2[rowsChosen] = DeepCopy(population[randomRow])
				rowsChosen += 1
			}
			if rowsChosen == TournamentNumber {
				break
			}
		}
		// Sort group1 and group2
		sort.Slice(group1, func(i, j int) bool {
			return Cost(group1[i]) < Cost(group1[j])
		})
		sort.Slice(group2, func(i, j int) bool {
			return Cost(group2[i]) < Cost(group2[j])
		})
		parent1 := group1[0] // The best from group1
		parent2 := group2[0] // The best from group2
		child1, child2 := OrderedCrossOver(parent1, parent2)
		newpopulation[row] = child1 
		row += 1 
		newpopulation[row] = child2 
		row += 1
	}
	// Perform mutations 
	for row := 0; row < ToursPerGeneration; row++ {
		r := rand.Float64()
		if r <= ProbMutation {
			SwapMutation(newpopulation[row])
		}
	}
	population = make([][]int, ToursPerGeneration)
	for i := 0; i < ToursPerGeneration; i++ {
		population[i] = make([]int, NUMCITIES)
	}
	// Copy newpopulation to population
	for row := 0; row < ToursPerGeneration; row++ {
		for col := 0; col < NUMCITIES; col++ {
			population[row][col] = newpopulation[row][col]
		}
	}
}

func SwapMutation(tour []int) {
	var index1, index2 int
	n := len(tour)
	for {
		index1 = 1 + rand.Intn(n-1)
		index2 = 1 + rand.Intn(n-1)
		if index2 != index1 + 4 {
			break // the two indices are different
		}
	}
	if index1 > index2 {
		index1, index2 = index2, index1
	}
	tour[index1], tour[index2] = tour[index2], tour[index1]
}

func OrderedCrossOver(parent1, parent2 []int) (child1, child2 []int) {
	var index1, index2 int
	n := len(parent1)
	for {
		index1 = 1 + rand.Intn(len(parent1)-1)
		index2 = 1 + rand.Intn(len(parent1)-1)
		if index1 != index2 {
			break // the two indices are different
		}
	}
	if index1 > index2 {
		index1, index2 = index2, index1
	}
	child1 = make([]int, len(parent1))
	child2 = make([]int, len(parent1))
	for i := 0; i < len(parent1); i++ {
		// Since 0 is a legal value
		child1[i] = -1
		child2[i] = -1
	}

	// Logic for child1
	for i := index1; i <= index2; i++ {
		child1[i] = parent1[i]
	}
	k := index2 + 1 // index for child1
	for i := index2 + 1; i < len(parent1); i++ {
		found, _ := In(parent2[i], child1)
		if !found {
			child1[k%n] = parent2[i]
			k += 1
		}
	}
	for i := 0; i <= index2; i++ {
		found, _ := In(parent2[i], child1)
		if !found {
			// j := (i + index2 + 1) % n
			child1[k%n] = parent2[i]
			k += 1
		}
	}

	// Logic for child2
	for i := index1; i <= index2; i++ {
		child2[i] = parent2[i]
	}
	k = index2 + 1 // index for child2
	for i := index2 + 1; i < len(parent2); i++ {
		found, _ := In(parent1[i], child2)
		if !found {
			child2[k%n] = parent1[i]
			k += 1
		}
	}
	for i := 0; i <= index2; i++ {
		found, _ := In(parent1[i], child2)
		if !found {
			// j := (i + index2 + 1) % n
			child2[k%n] = parent1[i]
			k += 1
		}
	}
	// Form child11 and child22
	// so they both start at 0
	child11 := []int{}
	child22 := []int{}
	_, index0 := In(0, child1)
	for i := index0; i < len(child1); i++ {
		child11 = append(child11, child1[i])
	}
	for i := 0; i < index0; i++ {
		child11 = append(child11, child1[i])
	}

	_, index0 = In(0, child2)
	for i := index0; i < len(child2); i++ {
		child22 = append(child22, child2[i])
	}
	for i := 0; i < index0; i++ {
		child22 = append(child22, child2[i])
	}
	return child11, child22
}

func GeneticAlgorithm() {
	generation := 0
	population = make([][]int, ToursPerGeneration)
	for i := 0; i < ToursPerGeneration; i++ {
		population[i] = make([]int, NUMCITIES)
	}
	newpopulation = make([][]int, ToursPerGeneration)
	for i := 0; i < ToursPerGeneration; i++ {
		newpopulation[i] = make([]int, NUMCITIES)
	}
	lowestCostTour := 1000000000.0
	CreateInitialPopulation()
	for {
		if generation == NumberGenerations {
			break
		}
		// Sort the population based on tour cost
		sort.Slice(population, func(i, j int) bool {
			return Cost(population[i]) < Cost(population[j])
		})
		bestCost := Cost(population[0])
		if bestCost < lowestCostTour {
			lowestCostTour = bestCost
			fmt.Printf("\nLowest cost tour at generation %d = %0.2f", generation, lowestCostTour)
		}
		FormNextGeneration()
		generation += 1
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
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

	graph = make([][]float64, NUMCITIES)
	for i:=0; i < NUMCITIES ; i++ {
	    graph[i] = make([]float64, NUMCITIES)
	}
	CreateGraph(NUMCITIES, cities, graph)

	GeneticAlgorithm()	
}
	   
/* Output
Lowest cost tour at generation 0 = 22019.11
Lowest cost tour at generation 1 = 20169.35
Lowest cost tour at generation 5 = 20017.31
Lowest cost tour at generation 6 = 19545.05
Lowest cost tour at generation 7 = 18447.20
Lowest cost tour at generation 12 = 18340.78
Lowest cost tour at generation 13 = 17953.87
Lowest cost tour at generation 15 = 17350.68
Lowest cost tour at generation 16 = 17095.07
Lowest cost tour at generation 18 = 16612.21
Lowest cost tour at generation 19 = 16425.63
Lowest cost tour at generation 20 = 16299.86
Lowest cost tour at generation 24 = 16002.17
Lowest cost tour at generation 28 = 15749.40
Lowest cost tour at generation 30 = 14754.66
Lowest cost tour at generation 53 = 13900.84
Lowest cost tour at generation 68 = 13831.31
Lowest cost tour at generation 72 = 13668.22
Lowest cost tour at generation 73 = 13636.80
Lowest cost tour at generation 77 = 13392.64
Lowest cost tour at generation 92 = 12979.84
Lowest cost tour at generation 103 = 12200.31
Lowest cost tour at generation 123 = 12030.21
Lowest cost tour at generation 186 = 11960.10
Lowest cost tour at generation 191 = 11860.86
Lowest cost tour at generation 204 = 11647.36
Lowest cost tour at generation 209 = 11639.41
Lowest cost tour at generation 215 = 11582.62
Lowest cost tour at generation 218 = 11580.22
Lowest cost tour at generation 224 = 11255.27
Lowest cost tour at generation 280 = 11150.08
Lowest cost tour at generation 344 = 11099.42
Lowest cost tour at generation 423 = 10775.75
Lowest cost tour at generation 482 = 10717.58
Lowest cost tour at generation 492 = 10592.38
Lowest cost tour at generation 496 = 10587.10
Lowest cost tour at generation 503 = 10556.30
Lowest cost tour at generation 508 = 10489.54
Lowest cost tour at generation 513 = 10415.89
Lowest cost tour at generation 519 = 10409.44
Lowest cost tour at generation 527 = 10292.43
Lowest cost tour at generation 536 = 10256.38
Lowest cost tour at generation 561 = 9990.04
Lowest cost tour at generation 795 = 9936.06
Lowest cost tour at generation 810 = 9869.37
Lowest cost tour at generation 883 = 9817.69
Lowest cost tour at generation 891 = 9694.19
Lowest cost tour at generation 909 = 9616.14
Lowest cost tour at generation 956 = 9541.14
Lowest cost tour at generation 965 = 9456.43
Lowest cost tour at generation 970 = 9362.68
Lowest cost tour at generation 1179 = 9285.43
*/
