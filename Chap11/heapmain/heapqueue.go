package main 

import (
	"example.com/heap"
	"fmt"
)

type Ordered interface {
	~float64 | ~int | ~string
}

type PriorityQueue[T Ordered] struct {
	infoHeap heap.Heap[T]
}

// Methods 
func (queue *PriorityQueue[T]) Push(item T) {
	queue.infoHeap.Insert(item)
}

func (queue *PriorityQueue[T]) Pop() T {
	returnValue := queue.infoHeap.Largest()
	queue.infoHeap.Remove()
	return returnValue
}

func main() {
	myQueue := PriorityQueue[string]{}
	myQueue.Push("Helen")
	myQueue.Push("Apollo")
	myQueue.Push("Richard")
	myQueue.Push("Barbara")
	fmt.Println(myQueue)
	myQueue.Pop()
	fmt.Println(myQueue)
	myQueue.Push("Arlene")
	fmt.Println(myQueue)
	myQueue.Pop()
	myQueue.Pop()
	fmt.Println(myQueue)
}
/* Output
{{[Richard Barbara Helen Apollo]}}
{{[Helen Barbara Apollo]}}
{{[Helen Barbara Apollo Arlene]}}
{{[Arlene Apollo]}}
*/