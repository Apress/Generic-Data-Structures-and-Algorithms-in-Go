package singlylinkedlist 

import (
	"fmt"
)

type Ordered interface {
	~string | ~int | ~float64
}

type Node[T Ordered] struct {
	Item T
	next *Node[T]
}

type List[T Ordered] struct {
	first       *Node[T]
	numberItems int
}

// Methods
func (list *List[T]) Append(item T) {
	// Adds item to a new node at the end of the list
	newNode := Node[T]{item, nil}
	if list.first == nil {
		list.first = &newNode
	} else {
		last := list.first
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = &newNode
	}
	list.numberItems += 1
}

func (list *List[T]) InsertAt(index int, item T) error {
	// Adds item to a new node at position index in the list
	if index < 0 || index > list.numberItems {
		return fmt.Errorf("Index out of bounds error")
	}
	newNode := Node[T]{item, nil}
	if index == 0 {
		newNode.next = list.first
		list.first = &newNode
		list.numberItems += 1
		return nil // No error
	}
	node := list.first
	count := 0
	previous := node
	for count < index {
		previous = node
		count++
		node = node.next
	}
	newNode.next = node
	previous.next = &newNode
	list.numberItems += 1
	return nil // no error
}

func (list *List[T]) RemoveAt(index int) (T, error) {
	if index < 0 || index > list.numberItems {
		var zero T
		return zero, fmt.Errorf("Index out of bounds error")
	}
	node := list.first
	if index == 0 {
		toRemove := node
		list.first = toRemove.next
		list.numberItems -= 1
		return toRemove.Item, nil
	}
	count := 0
	previous := node
	for count < index {
		previous = node
		count++
		node = node.next
	}
	toRemove := node
	previous.next = toRemove.next
	list.numberItems -= 1
	return toRemove.Item, nil
}

func (list *List[T]) IndexOf(item T) int {
	node := list.first
	count := 0
	for {
		if node.Item == item {
			return count
		}
		if node.next == nil {
			return -1
		}
		node = node.next
		count += 1
	}
}

func (list *List[T]) ItemAfter(item T) T {
	// Scan list for the first occurence of item
	node := list.first 
	for {
		if node == nil { // item not found
			var zero T 
			return zero
		}
		if node.Item == item {
			break
		}
		node = node.next
	}
	return node.next.Item
}

func (list *List[T]) Items() []T {
	result := []T{}
	node := list.first
	for i := 0; i < list.numberItems; i++ {
		result = append(result, node.Item)
		node = node.next
	}
	return result
}

func (list *List[T]) First() *Node[T] {
	return list.first
}

func (list *List[T]) Size() int {
	return list.numberItems
}