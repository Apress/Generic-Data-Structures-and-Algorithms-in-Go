package main

import (
	"fmt"
	"example.com/slicequeue"
)

func main() {
	myQueue := slicequeue.Queue[int]{}
	myQueue.Insert(15) 
	myQueue.Insert(20)
	myQueue.Insert(30)
	myQueue.Remove()
	fmt.Println(myQueue.First())

	queue := slicequeue.Queue[float64]{}
	for i := 0; i < 10; i++ {
		queue.Insert(float64(i))
	}
	iterator := queue.Range()
	for {
		if iterator.Empty() {
			break
		}
		fmt.Println(iterator.Next())
	}
	fmt.Println("queue.First() = ", queue.First())
}
/* Output
20
0
1
2
3
4
5
6
7
8
9
queue.First() =  0
*/