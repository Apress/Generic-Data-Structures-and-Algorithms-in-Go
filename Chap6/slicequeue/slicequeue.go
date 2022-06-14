package slicequeue

type Queue[T any] struct {
	items []T
}

type Iterator[T any] struct {
	next int // index in items
	items []T
}

// Methods
func (queue *Queue[T]) Insert(item T) {
	// item is added to the right-most position in the slice
	queue.items = append(queue.items, item)
}

func (queue *Queue[T]) Remove() T {
	returnValue := queue.items[0]
	queue.items = queue.items[1:]
	return returnValue
}

func (queue Queue[T]) First() T {
	return queue.items[0]
}

func (queue Queue[T]) Size() int {
	return len(queue.items)
}

func (queue *Queue[T]) Range() Iterator[T] {
	return Iterator[T]{0, queue.items}
}

func (iterator *Iterator[T]) Empty() bool {
	return iterator.next == len(iterator.items)
}

func (iterator *Iterator[T]) Next() T {
	returnValue := iterator.items[iterator.next]
	iterator.next++
	return returnValue
}