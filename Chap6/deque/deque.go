package deque

type Deque[T any] struct {
	items []T
}

func (deque *Deque[T]) InsertFront(item T) {
	deque.items = append(deque.items, item) // Expands deque.items 
	for i := len(deque.items) - 1; i > 0 ; i-- {
		deque.items[i] = deque.items[i - 1]
	}
	deque.items[0] = item
}

func (deque *Deque[T]) InsertBack(item T) {
	deque.items = append(deque.items, item)
}

func (deque *Deque[T]) First() T {
	return deque.items[0]
}

func (deque *Deque[T]) RemoveFirst() T {
	returnValue := deque.items[0]
	deque.items = deque.items[1:]
	return returnValue
}

func (deque *Deque[T]) Last() T {
	return deque.items[len(deque.items) - 1]
}

func (deque *Deque[T]) RemoveLast() T {
	length := len(deque.items)
	returnValue := deque.items[length - 1]
	deque.items = deque.items[:(length - 1)]
	return returnValue
}

func (deque *Deque[T]) Empty() bool {
	return len(deque.items) == 0
}        