package main 

import (
	"fmt"
	"example.com/singlylinkedlist"
)

func main() {
	cars := singlylinkedlist.List[string]{}
	cars.Append("Honda")
	cars.InsertAt(0, "Nissan")
	cars.InsertAt(0, "Chevy")
	cars.InsertAt(1, "Ford")
	cars.InsertAt(1, "Tesla")
	cars.InsertAt(0, "Audi")
	cars.InsertAt(2, "Volkswagon")
	cars.Append("Volvo")

	fmt.Println(cars.Items())
	fmt.Println("Index of Tesla: ", cars.IndexOf("Tesla"))

	cars.RemoveAt(0)
	car, _ := cars.RemoveAt(3)
	fmt.Println("car removed is: ", car)
	fmt.Println(cars.Items())
	cars.RemoveAt(cars.Size() - 1)
	fmt.Println(cars.Items())

	cars.Append("Lexus")
	fmt.Println(cars.Items())
	fmt.Println("First car in the list is: ", cars.First().Item)
	fmt.Println("Last car in the list is: ", cars.Items()[cars.Size() - 1])
}
/* Output
[Audi Chevy Volkswagon Tesla Ford Nissan Honda Volvo]
Index of Tesla:  3
car removed is:  Ford
[Chevy Volkswagon Tesla Nissan Honda Volvo]
[Chevy Volkswagon Tesla Nissan Honda]
[Chevy Volkswagon Tesla Nissan Honda Lexus]
First car in the list is:  Chevy
Last car in the list is:  Lexus
*/