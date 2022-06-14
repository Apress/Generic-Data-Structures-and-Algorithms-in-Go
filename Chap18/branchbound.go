package main

import (
	"fmt"
	"sort"
	"time"
)

const (
	NUMCITIES = 33
)

type Node struct {
	tour       []int
	lowerBound float64
	level      int
}

type Graph [][]float64

var graph Graph
var nodesGenerated int64

type Nodes []Node

// Allow nodes to be sorted
func (nodes Nodes) Len() int {
	return len(nodes)
}

func (nodes Nodes) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
}

func (nodes Nodes) Less(i, j int) bool {
	if nodes[i].level > nodes[j].level {
		return true
	}
	if nodes[i].level == nodes[j].level && nodes[i].lowerBound == nodes[j].lowerBound {
		// Return the smaller sum of cities
		tour1 := nodes[i].tour
		sum1 := 0;
		for i := 0; i < len(tour1); i++ {
			sum1 += tour1[i]
		}
		tour2 := nodes[j].tour
		sum2 := 0;
		for i := 0; i < len(tour2); i++ {
			sum2 += tour2[i]
		}
		return sum1 < sum2
	}
	if nodes[i].level == nodes[j].level && nodes[i].lowerBound != nodes[j].lowerBound {
		return nodes[i].lowerBound < nodes[j].lowerBound
	}
	return false
}

type Point struct {
	x float64
	y float64
}

type PriorityQueue struct {
	items Nodes
}

func NewPriorityQueue() PriorityQueue {
	return PriorityQueue{}
}

func (pq *PriorityQueue) Insert(node Node) {
	tourToInsert := DeepCopy(node.tour)
	nodeToInsert := Node{tourToInsert, node.lowerBound, node.level}
	pq.items = append(pq.items, nodeToInsert)
	sort.Sort(pq.items)
}

func (pq *PriorityQueue) Remove() Node {
	result := pq.items[0]
	pq.items = pq.items[1:]
	return result
}

func DeepCopy(tour []int) []int {
	result := []int{}
	for i := range tour {
		result = append(result, tour[i])
	}
	return result
}

func Minimum(values []float64) float64 {
	// This function excludes value 0
	min := 100000000.0
	for i := 0; i < len(values); i++ {
		if values[i] != 0 && values[i] < min {
			min = values[i]
		}
	}
	if min == 100000000.0 {
		return 0.0
	}
	return min
}

func In(value int, values []int) (bool, int) {
	// Returns true if value in values
	// Returns index of location or -1 if not found
	for index := 0; index < len(values); index++ {
		if values[index] == value {
			return true, index
		}
	}
	return false, -1
}

func LowerBound(tour []int) float64 {
	edges := make([]float64, 0)
	sum := 0.0
	n := len(tour)
	for city := 0; city < NUMCITIES; city++ {
		for index := 0; index < NUMCITIES; index++ {
			// index is part of tour
			found, pos := In(city, tour)
			if n > 1 && found {
				if pos == n-1 {
					edges = append(edges, graph[city][0])
				} else {
					edges = append(edges, graph[city][tour[pos+1]])
				}
				break
			}
			found, _ = In(index, tour)
			if n == 1 || !found {
				// Don't allow an index already in tour
				edges = append(edges, graph[city][index])
			}
		}
		sum += Minimum(edges)
		edges = make([]float64, 0)
	}
	return sum
}

func TSP() {
	var elapsed time.Duration
	start := time.Now()
	bestTour := []int{}
	for i := 0; i < NUMCITIES; i++ {
		bestTour = append(bestTour, i)
	}
	pq := NewPriorityQueue()
	bestCost := LowerBound(bestTour)
	tour := []int{0}
	lowerBound := LowerBound(tour)
	node := Node{tour, lowerBound, 0}
	nodesGenerated += 1
	pq.Insert(node)
	for {
		if len(pq.items) == 0 {
			break
		}
		top := pq.Remove()
		topLevel := top.level
		topTour := top.tour

		// Generate nodes at topLevel + 1
		for i := 0; i < NUMCITIES; i++ {
			tour := DeepCopy(topTour)
			found, _ := In(i, topTour)
			if !found {
				tour = append(tour, i)
				nodesGenerated += 1
				if nodesGenerated % 10_000_000 == 0 {
					fmt.Println("\nNodes generated (in millions): ", nodesGenerated / 1_000_000)
					fmt.Printf("\n\nOptimum tour cost: %0.2f  \nBest tour: %v", bestCost, bestTour)
					elapsed = time.Since(start)
					seconds := elapsed / 1_000_000_000
					rate := float64(nodesGenerated) / float64(seconds)
					fmt.Printf("\nNodes generated per second: %0.0f  Length of PQ: %d   Time elapsed: %v", rate, len(pq.items), elapsed)
				}
				if len(tour) == NUMCITIES {
					// A complete tour is obtained
					tourCost := LowerBound(tour)
					if tourCost < bestCost {
						bestTour = tour
						bestCost = tourCost
						fmt.Println("\n\nBest cost of tour so far: ", bestCost)
					}
				} else {
					tourCost := LowerBound(tour)
					if tourCost < bestCost {
						node := Node{tour, tourCost, topLevel + 1}
						pq.Insert(node)
					}
				}
			}
		}	
	}
	fmt.Printf("\n\nOptimum tour cost: %0.2f  \nBest tour: %v  \nNodes generated: %d", bestCost, bestTour, nodesGenerated)
}

func main() {
	
	graph = Graph{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{184,  0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{292,  195,  0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{449,  310,  215,  0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{670,  540,  380,  288,  0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{516,  357,  232,  200,  211,  0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{598,  514,  434,  566,  436,  381,  0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{618,  434,  493,  787,  814,  642,  295,  0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{881,  697,  719,  790,  632,  697,  224,  320,  0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{909,  964,  955,  1020, 974,  952,  541,  341,  318,  0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{978,  892,  1031, 1246, 1352, 1180, 843,  538,  747,  441,  0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{654,  597,  803,  1018, 1154, 1104, 766,  461,  749,  634,  380,  0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{504,  503,  722,  937,  1043, 806,  986,  722,  1042, 954,  784,  404,  0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{276,  460,  568,  725,  946,  817,  874,  894,  1214, 1185, 1218, 660,  452, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{780,  964,  1072, 1229, 1450, 1321, 1378, 1326, 1646, 1672, 1410, 1030, 626, 476, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{529,  644,  789,  1004, 1184, 1001, 1214, 950,  1270, 1213, 1043, 632,  219, 436, 419, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{805,  698,  917,  1132, 1238, 1055, 1113, 842,  1162, 1027, 779,  473,  195, 637, 634, 256, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1181, 1007, 1226, 1441, 1547, 1364, 1375, 1080, 1134, 1138, 783,  611,  563, 1046,759, 624,  368, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1548, 1444, 1630, 1845, 1984, 1801, 1726, 1431, 1685, 1477, 1134, 1033, 906, 1389, 1094, 967,  711, 404, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1547, 1454, 1668, 1883, 1994 ,1811, 1879, 1584, 1776, 1632, 1267, 1053, 944, 1427, 1196, 1005, 749, 442, 251,  0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1239, 1167, 1353, 1568, 1707, 1524, 1584, 1313, 1633, 1498, 1151, 979,  614, 988,  600,  525,  471, 368,  512,  507,  0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1538, 1733, 1830, 2045, 2208, 2090, 2136, 2078, 2398, 2332, 2110, 1782, 1378, 1300, 760,  1229, 1382, 1319, 1163, 930,  910,  0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1999, 2158, 2291, 2448, 2669, 2515, 2597, 2500, 2820, 2675, 2336, 2164, 1707, 1860, 1375, 1582, 1658, 1545, 1389, 1156, 1237, 436,  0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1716, 1875, 2008, 2165, 2386, 2232, 2488, 2217, 2537, 2392, 2053, 1881, 1422, 1577, 1106, 1244, 1375, 1262, 1106, 873,  954,  483,  283,  0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1580, 1738, 1872, 2029, 2250, 2095, 2352, 2081, 2401, 2256, 1917, 1745, 1286, 1473, 988,  1147, 1239, 1126, 970,  737,  818,  379,  419,  136,  0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1425, 1569, 1717, 1874, 2109, 1926, 2115, 1844, 2164, 2019, 1680, 1508, 1118, 1335, 862,  913,  1002, 889,  733,  500,  581,  430,  656,  373,  237,  0, 0, 0, 0, 0, 0, 0, 0},
		{1560, 1549, 1852, 2009, 2089, 1906, 2063, 1792, 2112, 1967, 1456, 1274, 1032, 1540, 1068, 1007, 944,  665,  521,  282,  491,  768,  994,  711,  575,  358,  0, 0, 0, 0, 0, 0, 0},
		{1918, 1744, 1963, 2178, 2284, 2101, 2174, 1879, 2071, 1892, 1562, 1348, 1239, 1722, 1258, 1300, 1044, 737,  526,  295,  802,  816,  1022, 739,  603,  386,  545,  0, 0, 0, 0, 0, 0},
		{2065, 2102, 2357, 2514, 2642, 2459, 2626, 2355, 2675, 2530, 2191, 2019, 1673, 1905, 1432, 1570, 1507, 1320, 1109, 878,  1124, 842,  715,  432,  465,  533,  849,  739,  0, 0, 0, 0, 0},
		{2284, 2131, 2326, 2441, 2671, 2488, 2418, 2133, 2437, 2259, 1929, 1715, 1606, 1924, 1451, 1589, 1411, 1104, 893,  662,  1143, 1004, 981,  698,  599,  589,  768,  523,  266,  0, 0, 0, 0},
		{2340, 2348, 2543, 2658, 2888, 2705, 2869, 2598, 2918, 2773, 2434, 2262, 1916, 2148, 1675, 1813, 1750, 1468, 1102, 871,  1367, 1085, 958,  675,  1033, 778,  1092, 982,  243,  349,  0, 0, 0},
		{2247, 2327, 2539, 2696, 2867, 2684, 2851, 2580, 2900, 2755, 2416, 2244, 1898, 2130, 1657, 1795, 1732, 1545, 1334, 1103, 1349, 1014, 814,  531,  1015, 760,  1047, 964,  225,  497,  266,  0, 0},
		{2163, 2322, 2455, 2612, 2833, 2679, 2761, 2664, 2984, 2839, 2500, 2328, 1871, 2024, 1539, 1746, 1822, 1709, 1553, 1320, 1391, 693,  346,  447,  583,  820,  1158, 1355, 581,  847,  710,  444, 0}, 
	}
	for row := 0; row < NUMCITIES; row++ {
		for col := row + 1; col < NUMCITIES; col++ {
			graph[row][col] = graph[col][row]
		}
	}
	TSP()	
}
	 
