package main

import (
	"fmt"
)

type Node[T any] struct {
	value T 
	next *Node[T]
}

type Stack[T any] struct {
	first *Node[T]
}

// Methods
func (stack *Stack[T]) Push(item T) {
	newNode := Node[T]{item, nil}
	// newNode.value = item
	newNode.next = stack.first
	stack.first = &newNode
}

func (stack *Stack[T]) Top() T {
	return stack.first.value
}

func (stack *Stack[T]) Pop() T {
	result := stack.first.value
	stack.first = stack.first.next
	return result
}

func (stack Stack[T]) IsEmpty() bool {
	return stack.first == nil
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