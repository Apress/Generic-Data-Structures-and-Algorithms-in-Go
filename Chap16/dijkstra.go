package main

import (
	"fmt"
	"github.com/jiaxwu/container/heap"
)

// PriorityQueue Priority queue
type PriorityQueue[T any] struct {
	h *heap.Heap[T]
}

func New[T any](less func(e1 T, e2 T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		h: heap.New(less),
	}
}

func (p *PriorityQueue[T]) Add(elem T) {
	p.h.Push(elem)
}

func (p *PriorityQueue[T]) Remove() T {
	return p.h.Pop()
}

func (p *PriorityQueue[T]) Len() int {
	return p.h.Len()
}

func (p *PriorityQueue[T]) Empty() bool {
	return p.Len() == 0
}

func Less(a, b tuple) bool {
	return a.weight < b.weight
}
// end priority queue

type edges = map[rune]int

type Graph map[rune]edges

type tuple struct {
	node   rune
	weight int
}

func convert(r rune) int {
	return int(r) - 65
}

func Dijkastra(graph Graph) []tuple {
	distances := make([]tuple, len(graph))

	for i := 0; i < len(graph); i++ {
		distances[i] = tuple{'A', 32767}
	}

	distances[0] = tuple{'A', 0}
	heapQueue := New[tuple](Less)
	t := tuple{'A', 0}
	heapQueue.Add(t)
	for {
		if heapQueue.Len() == 0 {
			break
		}
		t = heapQueue.Remove()
		currentNode := t.node
		currentDistance := t.weight
		if currentDistance > distances[convert(currentNode)].weight {
			continue
		}
		for t, w := range graph[currentNode] {
			neighbor := t
			weight := w
			distance := currentDistance + weight
			/*
			   Only consider this new path if it's
			   better than any path we've already found
			*/
			if distance < distances[convert(neighbor)].weight {
				distances[convert(neighbor)] = tuple{neighbor, distance}
				heapQueue.Add(tuple{neighbor, distance})
			}
		}
	}
	return distances
}

func main() {
	graph := make(map[rune]edges)
	graph['A'] = edges{'B': 4, 'H': 1}
	graph['B'] = edges{'A': 4, 'C': 1, 'H': 11}
	graph['C'] = edges{'B': 1, 'I': 2, 'F': 1, 'D': 7}
	graph['D'] = edges{'C': 7, 'F': 8, 'E': 1}
	graph['E'] = edges{'D': 1, 'F': 10}
	graph['F'] = edges{'G': 2, 'C': 1, 'D': 8, 'E': 10}
	graph['G'] = edges{'F': 2, 'I': 1, 'H': 1}
	graph['H'] = edges{'G': 1, 'I': 7, 'B': 11, 'A': 1}
	graph['I'] = edges{'C': 2, 'H': 7, 'G': 1}

	solution := Dijkastra(graph)
	for node, weight := range solution {
		fmt.Printf("%s %d ", string(node + 65), weight)
	}
}
/* Output
A {65 0} B {66 4} C {67 5} D {68 12} E {69 13} F {70 4} G {71 2} H {72 1} I {73 3} 
*/

