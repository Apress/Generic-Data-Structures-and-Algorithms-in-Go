package heap

type Ordered interface {
	~float64 | ~int | ~string
}

type Heap[T Ordered] struct {
	Items []T
}

// Methods
func (heap *Heap[T]) Swap(index1, index2 int) {
	heap.Items[index1], heap.Items[index2] = heap.Items[index2], heap.Items[index1]
}

func NewHeap[T Ordered](input []T) *Heap[T] {
	heap := &Heap[T]{}
	for i := 0; i < len(input); i++ {
		heap.Insert(input[i])
	}
	return heap
}

func (heap *Heap[T]) Insert(value T) {
	heap.Items = append(heap.Items, value)
	heap.buildHeap(len(heap.Items) - 1)
}

func (heap *Heap[T]) Remove() {
	// Can only remove Items[0], the largest value
	heap.Items[0] = heap.Items[len(heap.Items)-1] 
	heap.Items = heap.Items[:(len(heap.Items) - 1)]
	heap.rebuildHeap(0)
}

func (heap *Heap[T]) Largest() T {
	return heap.Items[0]
}

func (heap *Heap[T]) buildHeap(index int) {
	var parent int
	if index > 0 {
		parent = (index - 1) / 2
		if heap.Items[index] > heap.Items[parent] {
			heap.Swap(index, parent)
		}
		heap.buildHeap(parent)
	}
}

func (heap *Heap[T]) rebuildHeap(index int) {
	length := len(heap.Items)
	if (2 * index + 1) < length {
		left := 2*index + 1
		right := 2*index + 2
		largest := index

		if left < length && right < length && heap.Items[left] >= heap.Items[right] && heap.Items[index] < heap.Items[left] {
			largest = left
		} else if right < length && heap.Items[right] >= heap.Items[left] && heap.Items[index] < heap.Items[right]{
			largest = right
		} else if left < length && right >= length && heap.Items[index] < heap.Items[left] {
			largest = left
		}
		if index != largest {
			heap.Swap(index, largest)
			heap.rebuildHeap(largest)
		}
	}
}