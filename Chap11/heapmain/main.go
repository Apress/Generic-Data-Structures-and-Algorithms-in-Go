package main 

import (
	"fmt"
	"example.com/heap"
)

func main() {
	slice1 := []int{100, 60, 80, 50, 30, 75, 40, 10, 35}
	heap1 := heap.NewHeap[int](slice1)
	heap1.Insert(90)
	fmt.Println("heap1 after inserting 90")
	fmt.Println(heap1.Items)
	fmt.Println("Largest item in heap: ", heap1.Largest())
	
	heap1.Remove()
	fmt.Println("Removing largest item from heap yielding the heap: ")
	fmt.Println(heap1.Items)
	fmt.Println("Largest item in heap: ", heap1.Largest())
	
	slice2 := []int{10, 35, 100, 80, 30, 75, 40, 50, 60}
	heap2 := heap.NewHeap[int](slice2)
	heap2.Insert(90)
	fmt.Println("heap2 with rearranged slice2 after inserting 90")
	fmt.Println(heap2.Items)
}
/* Output
heap1 after inserting 90
[100 90 80 50 60 75 40 10 35 30]
Largest item in heap:  100
Removing largest item from heap yielding the heap: 
[90 60 80 50 30 75 40 10 35]
Largest item in heap:  90
heap2 with rearranged slice2 after inserting 90
[100 90 75 60 80 35 40 10 50 30]
*/
