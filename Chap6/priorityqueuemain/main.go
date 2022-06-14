package main

import (
	"example.com/nodequeue"
	"fmt"
)

type Passenger struct {
	name string 
	priority int
}


type PriorityQueue[T any] struct {
	q []nodequeue.Queue[T]
	size int 
}

func NewPriorityQueue[T any](numberPriorities int) (pq PriorityQueue[T]) {
	pq.q = make([]nodequeue.Queue[T], numberPriorities)
	return pq
}

// Methods for priority queue
func (pq *PriorityQueue[T]) Insert(item T, priority int) {
	pq.q[priority - 1].Insert(item)
	pq.size++
}

func (pq *PriorityQueue[T]) Remove() T {
	pq.size--
	for i := 0; i < len(pq.q); i++ {
		if pq.q[i].Size() > 0 {
			return pq.q[i].Remove()
		}
	}
	var zero T
	return zero
}

func (pq *PriorityQueue[T]) First() T {
	for _, queue := range(pq.q) {
		if queue.Size() > 0 {
			return queue.First()
		}
	}
	var zero T 
	return zero
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	result := true
	for _, queue := range(pq.q) {
		if queue.Size() > 0 {
			result = false 
			break
		}
	}
	return result
}

func main() {
	airlineQueue := NewPriorityQueue[Passenger](3)
	passengers := []Passenger{ {"Erika", 3},{"Robert", 3}, {"Danielle", 3}, {"Madison", 1},
							   {"Frederik", 1}, {"James", 2}, {"Dante", 2}, {"Shelley", 3} }
	fmt.Println("Passsengers: ",passengers)
	for i := 0; i < len(passengers); i++ {
		airlineQueue.Insert(passengers[i], passengers[i].priority)
	}
	fmt.Println("First passenger in line: ", airlineQueue.First())
	airlineQueue.Remove()
	airlineQueue.Remove()
	airlineQueue.Remove()
	fmt.Println("First passenger in line after three Removes: ", airlineQueue.First())
}
/* Output
Passsengers:  [{Erika 3} {Robert 3} {Danielle 3} {Madison 1} {Frederik 1} {James 2} {Dante 2} {Shelley 3}]
First passenger in line:  {Madison 1}
First passenger in line after three Removes:  {Dante 2}
*/