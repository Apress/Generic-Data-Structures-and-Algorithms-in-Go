package main

import (
	"fmt"
)

type Stack[T any] struct {
	items []T
}

// Methods
func (stack *Stack[T]) Push(item T) {
	// item is added to the right-most position in the slice
	stack.items = append(stack.items, item)
}

func (stack *Stack[T]) Pop() T {
	length := len(stack.items)
	returnValue := stack.items[length - 1]
	stack.items = stack.items[:(length - 1)]
	return returnValue
}

func (stack Stack[T]) Top() T {
	length := len(stack.items)
	return stack.items[length - 1]
}

func (stack Stack[T]) IsEmpty() bool {
	return len(stack.items) == 0
}

func main() {
	// Create a stack of names
	nameStack := Stack[string]{}
	nameStack.Push("Zachary")
	nameStack.Push("Adolf")

	if !nameStack.IsEmpty() {
		topOfStack := nameStack.Top()
		fmt.Printf("\nTop of stack is %s", topOfStack)
	}

	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nValue popped from stack is %s", poppedFromStack)
	}

	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nValue popped from stack is %s", poppedFromStack)
	}

	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nValue popped from stack is %s", poppedFromStack)
	}

	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nValue popped from stack is %s", poppedFromStack)
	}

	// Create a stack of integers
	intStack := Stack[int]{}
	intStack.Push(5)
	intStack.Push(10)
	intStack.Push(0) 

	if !intStack.IsEmpty() {
		top := intStack.Top()
		fmt.Printf("\nValue on top of intStack is %d", top)
	}

	if !intStack.IsEmpty() {
		popFromStack := intStack.Pop() 
		fmt.Printf("\nValue popped from intStack is %d", popFromStack)
	}
			
	if !intStack.IsEmpty() {
		popFromStack := intStack.Pop() 
		fmt.Printf("\nValue popped from intStack is %d", popFromStack)
	}
	
	if !intStack.IsEmpty() {
		popFromStack := intStack.Pop() 
		fmt.Printf("\nValue popped from intStack is %d", popFromStack)
	}
}
/* Output
Top of stack is Adolf
Value popped from stack is Adolf
Value popped from stack is Zachary
Value on top of intStack is 0
Value popped from intStack is 0
Value popped from intStack is 10
Value popped from intStack is 5
*/