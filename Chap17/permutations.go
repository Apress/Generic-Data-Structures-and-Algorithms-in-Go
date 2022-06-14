package main  

import (
	"fmt"
)

func Permutations(data []int, operation func([]int)) {
	permute(data, operation, 0)
}

func permute(data []int, operation func([]int), step int) {
	if step > len(data) {
		operation(data)
		return
	}
	permute(data, operation, step + 1)
	for k := step + 1; k < len(data); k++ {
		data[step], data[k] = data[k], data[step]
		permute(data, operation, step + 1)
		data[step], data[k] = data[k], data[step]
	}
}

func main() {
	data := []int{0, 1, 2, 3}
	Permutations(data, func(a []int) {
		fmt.Println(a)
	})
}
/* Output
[0 1 2 3]
[0 1 3 2]
[0 2 1 3]
[0 2 3 1]
[0 3 2 1]
[0 3 1 2]
[1 0 2 3]
[1 0 3 2]
[1 2 0 3]
[1 2 3 0]
[1 3 2 0]
[1 3 0 2]
[2 1 0 3]
[2 1 3 0]
[2 0 1 3]
[2 0 3 1]
[2 3 0 1]
[2 3 1 0]
[3 1 2 0]
[3 1 0 2]
[3 2 1 0]
[3 2 0 1]
[3 0 2 1]
[3 0 1 2]
*/