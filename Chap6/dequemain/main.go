package main

import (
	"fmt"
	"example.com/deque"
)

func main() {
	myDeque := deque.Deque[int]{}
	myDeque.InsertFront(5)
	myDeque.InsertBack(10)
	myDeque.InsertFront(2)
	myDeque.InsertBack(12) // 2 5 10 12
	fmt.Println("myDeque.First() = ", myDeque.First())
	fmt.Println("myDeque.Last() = ", myDeque.Last())

	myDeque.RemoveLast()
	myDeque.RemoveFirst()
	fmt.Println("myDeque.First() = ", myDeque.First())
	fmt.Println("myDeque.Last() = ", myDeque.Last())
}
/* Output
myDeque.First() =  2
myDeque.Last() =  12
myDeque.First() =  5
myDeque.Last() =  10
*/
