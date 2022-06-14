package nodestack

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
