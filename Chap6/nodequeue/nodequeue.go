package nodequeue

type Node[T any] struct {
	item T  
	next *Node[T]
}

type Queue[T any] struct {
	first, last *Node[T]
	length int
}

type Iterator[T any] struct {
	next *Node[T]
}

// Methods
func (queue *Queue[T]) Insert(item T) {
	newNode := &Node[T]{item, nil}
	if queue.first == nil {
		queue.first = newNode
		queue.last = queue.first
	} else {
		queue.last.next = newNode
		queue.last = newNode
	}
	queue.length += 1 
}

func (queue *Queue[T]) Remove() T {
	returnValue := queue.first.item
	queue.first = queue.first.next
	if queue.first == nil {
		queue.last = nil
	}
	queue.length -= 1
	return returnValue
}

func (queue Queue[T]) First() T {
	return queue.first.item
}

func (queue Queue[T]) Size() int {
	return queue.length
}

func (queue *Queue[T]) Range() Iterator[T] {
	return Iterator[T]{queue.first}
}

func (iterator *Iterator[T]) Empty() bool {
	return iterator.next == nil
}

func (iterator *Iterator[T]) Next() T {
	returnValue := iterator.next.item
	if iterator.next != nil {
		iterator.next = iterator.next.next
	}
	return returnValue
}

