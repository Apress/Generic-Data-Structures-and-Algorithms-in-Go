package slicestack

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