package main 

import (
	"example.com/set"
	"fmt"
)

func main() {
	set1 := set.Set[int]{}
	set1.Insert(3)
	set1.Insert(5)
	set1.Insert(7)
	set1.Insert(9)
	set2 := set.Set[int]{}
	set2.Insert(3)
	set2.Insert(6)
	set2.Insert(8)
	set2.Insert(9)
	set2.Insert(11)
	set2.Delete(11)
	fmt.Println("Items in set2: ", set2.Items())

	fmt.Println("5 in set1: ", set1.In(5))
	fmt.Println("5 in set2: ", set2.In(5))

	fmt.Println("Union of set1 and set2: ", set1.Union(set2).Items())
	fmt.Println("Intersection of set1 and set2: ", set1.Intersection(set2).Items())
	fmt.Println("Difference of set2 with respect to set1: ", set2.Difference(set1).Items())
	fmt.Println("Size of this difference: ", set1.Intersection(set2).Size())
}
/* Output 
Items in set2:  [6 8 9 3]
5 in set1:  true
5 in set2:  false
Union of set1 and set2:  [9 3 5 7 6 8]
Intersection of set1 and set2:  [3 9]
Difference of set2 with respect to set1:  [6 8]
Size of this difference:  2
*/
