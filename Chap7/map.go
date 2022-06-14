package main

import (
	"fmt"
)

type map1 map[string]string


func main() {
	nicknames := make(map1, 5)
	nicknames["Charles"] = "Chuck"
	nicknames["Robert"] = "Bob"
	nicknames["Richard"] = "Rick"
	nicknames["Teddy"] = "Ted"
	nicknames["Mohammad"] = "Mo"

	for key, value := range (nicknames) {
		fmt.Printf("\nThe nickname of %s is %s", key, value)
	}

	// Test for the presence of James in the map
	_, present := nicknames["James"]
	fmt.Println("\nThe key James is present: ", present)

	// Test for the presence of Teddy in the map
	_, present = nicknames["Teddy"]
	fmt.Println("The key Teddy is present: ", present)

	delete(nicknames, "Robert")

	// Test for the presence of Robert in the map
	_, present = nicknames["Robert"]
	fmt.Println("The key Robert is present: ", present)

	// Modify the nickname of Charles
	nicknames["Charles"] = "Charlie"
}
/* Output
he nickname of Robert is Bob
The nickname of Richard is Rick
The nickname of Teddy is Ted
The nickname of Mohammad is Mo
The nickname of Charles is Chuck
The key James is present:  false
The key Teddy is present:  true
he key Robert is present:  false
*/