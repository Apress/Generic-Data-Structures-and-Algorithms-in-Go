package main

import (
	"fmt"
	"example.com/nodequeue"
)

func main() {
	myQueue := nodequeue.Queue[int]{}
	myQueue.Insert(15) 
	myQueue.Insert(20)
	myQueue.Insert(30)
	myQueue.Remove()
	fmt.Println(myQueue.First())

	queue := nodequeue.Queue[float64]{}
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