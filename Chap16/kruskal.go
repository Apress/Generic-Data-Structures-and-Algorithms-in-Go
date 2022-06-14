package main 

import (
	"fmt"
	"sort"
)

type Edge struct {
	weight int 
	node1 Node
	node2 Node
}

type Node = string

type EdgeSlice []Edge

// Infrastructure to allow []Edges to be sorted
func (edges EdgeSlice) Len() int {
	return len(edges)
}

func (edges EdgeSlice) Swap(i, j int) {
	edges[i], edges[j] = edges[j], edges[i]
}

func (edges EdgeSlice) Less(i, j int) bool {
	return edges[i].weight < edges[j].weight
}

var connection map[Node]Node

/*
	The initial level of each Node is 0.
	If the node is node2 of an Edge, 
	increase its level by 1.
*/
var end map[Node]int

func Initialize(node Node) {
	connection[node] = node 
	end[node]  = 0
}

func Find(node Node) Node {
	// Stops a cycle 
	if connection[node] != node {
		connection[node] = Find(connection[node])
	}
	return connection[node]
}

func Connect(node1, node2 Node) {
	n1 := Find(node1) 
	n2 := Find(node2)
	fmt.Printf("\nFind(%s) = %s", node1, n1)
	fmt.Printf("\nFind(%s) = %s", node2, n2)
	if n1 != n2 {
		fmt.Printf("\nend[%s] = %d", n1, end[n1])
		fmt.Printf("\nend[%s] = %d", n2, end[n2])
		if end[n1] > end[n2] {
			connection[n2] = n1
			fmt.Printf("\nconnection[%s] = %s", n2, n1)
		} else {
			connection[n1] = n2 
			fmt.Printf("\nconnection[%s] = %s", n1, n2)
			if end[n1] == end[n2] {
				end[n2] += 1
				fmt.Printf("\nend[%s] = 1", n2)
			}
		}
	}
}

func Kruskal(nodes []Node, edges EdgeSlice) []Edge {
	for _, node := range nodes {
		Initialize(node)
	}
	spanningTree := []Edge{}
	sort.Sort(edges)
	for _, edge := range edges {
		node1 := edge.node1 
		node2 := edge.node2 
		n1 := Find(node1) 
		n2 := Find(node2) 
		fmt.Printf("\nFind(%s) = %s", node1, n1)
		fmt.Printf("\nFind(%s) = %s", node2, n2)
		if n1 != n2 {
			Connect(node1, node2) 
			fmt.Printf("\nConnect(%s, %s)", node1, node2)
			spanningTree = append(spanningTree, edge)
		} else {
			fmt.Printf("\nReject edge %s and %s", node1, node2)
		}
	}
	return spanningTree
}

func main() {
	connection = make(map[Node]Node)
	end = make(map[Node]int)
	// Define the graph by its nodes and edges
	nodes := []Node{"A", "B", "C", "D", "E", "F", "G"}
	edges := []Edge{ {1, "A", "B"}, {1, "A", "C"}, 
                     {4, "B", "C"}, {20, "C", "D"},
                     {2, "D", "E"}, {3, "E", "F"},
                     {6, "G", "D"}, {1, "C", "G"}, 
                     {5, "D", "F"} }
    spanningTree := Kruskal(nodes, edges)
    fmt.Println("\n", spanningTree)
}
/* Output 
Find(A) = A
Find(B) = B
Find(A) = A
Find(B) = B
end[A] = 0
end[B] = 0
connection[A] = B
end[B] = 1
Find(A) = B
Find(C) = C
Find(A) = B
Find(C) = C
end[B] = 1
end[C] = 0
connection[C] = B
Find(C) = B
Find(G) = G
Find(C) = B
Find(G) = G
end[B] = 1
end[G] = 0
connection[G] = B
Find(D) = D
Find(E) = E
Find(D) = D
Find(E) = E
end[D] = 0
end[E] = 0
connection[D] = E
end[E] = 1
Find(E) = E
Find(F) = F
Find(E) = E
Find(F) = F
end[E] = 1
end[F] = 0
connection[F] = E
Find(B) = B
Find(C) = B
Reject edge B and C
Find(D) = E
Find(F) = E
Reject edge D and F
Find(G) = B
Find(D) = E
Find(G) = B
Find(D) = E
end[B] = 1
end[E] = 1
connection[B] = E
end[E] = 1
Find(C) = E
Find(D) = E
Reject edge C and D
 [{1 A B} {1 A C} {1 C G} {2 D E} {3 E F} {6 G D}]
*/
