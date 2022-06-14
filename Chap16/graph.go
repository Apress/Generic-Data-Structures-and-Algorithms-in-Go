
package main

import (
	"fmt"
)

type OrderedStringer interface {
	comparable // instead of constraints
	String() string
}

type Graph[T OrderedStringer] struct {
	Vertices map[T]*Vertex[T]
}

type Vertex[T OrderedStringer] struct {
	Key T
	Neighbors map[T]*Vertex[T]
}

var visitation []string

func NewVertex[T OrderedStringer](key T) *Vertex[T] {
	return &Vertex[T]{
		Key: key,
		Neighbors: map[T]*Vertex[T]{},
	}
}

func NewGraph[T OrderedStringer]() *Graph[T] {
	return &Graph[T]{Vertices: map[T]*Vertex[T]{}}
}

func (g *Graph[T]) AddVertex(key T) {
	vertex := NewVertex(key)
	g.Vertices[key] = vertex
}

func (g *Graph[T]) AddEdge(key1, key2 T, edgeValue int) {
	vertex1 := g.Vertices[key1]
	vertex2 := g.Vertices[key2]
	if vertex1 == nil || vertex2 == nil {
		return
	}
	vertex1.Neighbors[vertex2.Key] = vertex2
	g.Vertices[vertex1.Key] = vertex1 
	g.Vertices[vertex2.Key] = vertex2
}

func (g *Graph[T]) DepthFirstSearch(start *Vertex[T], visited map[T]bool) {
	if start == nil {
		return
	}
	visited[start.Key] = true
	visitation = append(visitation, start.Key.String())

	// for each of the adjacent vertices, call the function recursively
	// if it hasn't yet been visited
	for _, v := range start.Neighbors {
		// The sequence of v may change from run to run
		if visited[v.Key] {
			continue
		}
		g.DepthFirstSearch(v, visited)
	}
}

type Queue[T any] struct {
	items []T
}

// Methods
func (queue *Queue[T]) Insert(item T) {
	// item is added to the right-most position in the slice
	queue.items = append(queue.items, item)
}

func (queue *Queue[T]) Remove() T {
	returnValue := queue.items[0]
	queue.items = queue.items[1:]
	return returnValue
}


func (g *Graph[T]) BreadthFirstSearch(start *Vertex[T], visited map[T]bool) {
	if start == nil {
		return
	}
	queue := Queue[*Vertex[T]]{} // Queue hold pointers to Vertex
	current := start 
	for { 
		if !visited[current.Key] {
			visitation = append(visitation, current.Key.String())
		}
		visited[current.Key] = true
		// Insert each neighboring vertex not visited onto the queue 
		for _, v := range current.Neighbors {
			if !visited[v.Key] {
				queue.Insert(v)
			} 
		}
		// Grab first vertex in the queue and remove it
		if len(queue.items) > 0 {
			current = queue.Remove()
		} else {
			break
		}
	}
}

func (g *Graph[T]) String() string {
	result := ""
	for i := 0; i < len(visitation); i++ {
		result += " " + visitation[i]
	}
	return result
}

// Make String implement Stringer
type String string 

func (str String) String() string {
	return string(str)
}

func main() {
	g := NewGraph[String]()
	g.AddVertex("A")
	start := g.Vertices["A"]
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddVertex("E")
	g.AddVertex("F")
	g.AddVertex("G")
	g.AddEdge("A", "B", 2)
	g.AddEdge("A", "C", 5)
	g.AddEdge("A", "D", 9)
	g.AddEdge("B", "D", 3)
	g.AddEdge("C", "F", 9)
	g.AddEdge("D", "E", 4)
	g.AddEdge("E", "D", 4)
	g.AddEdge("F", "E", 6)
	g.AddEdge("E", "G", 7)
	g.AddEdge("F", "G", 3) 
	visited := make(map[String]bool)
	visitation = []string{}
	g.DepthFirstSearch(start, visited)
	fmt.Println("Depth First Search:", g.String())
	visited = make(map[String]bool)
	visitation = []string{}
	g.BreadthFirstSearch(start, visited)
	fmt.Println("Breadth First Search:", g.String())
}
/* Output
Depth First Search:  A B D E G C F
Breadth First Search:  A C D B F E G
*/