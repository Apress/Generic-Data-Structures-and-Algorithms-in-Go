package set

type Ordered interface {
	~string | ~int | ~float64
} 

type Set[T Ordered] struct {
	items map[T]bool
}

// Methods  

// Add item to set
func (set *Set[T]) Insert(item T) {
	if set.items == nil {
		set.items = make(map[T]bool)
	}
	// Prevent duplicate entry 
	_, present := set.items[item]
	if !present {
		set.items[item] = true
	}
}

// Remove item from set
func (set *Set[T]) Delete(item T) {
	_, present := set.items[item]
	if present {
		delete(set.items, item)
	}
}

// Return true if item is in set, otherwise false
func (set *Set[T]) In(item T) bool {
	_, present := set.items[item]
	return present
}

// Return a slice of all the items in set
func (set *Set[T]) Items() []T {
	items := []T{}
	for item := range set.items {
		items = append(items, item)
	}
	return items
}

// Return the number of items in set
func (set *Set[T]) Size() int {
	return len(set.items)
}

// Return a new set containing all the unique items of set and set2
func (set *Set[T]) Union(set2 Set[T]) *Set[T] {
	result := Set[T]{}
	result.items = make(map[T]bool)
	for index := range set.items {
		result.items[index] = true
	}
	for j := range set2.items {
		_, present := result.items[j]
		if !present {
			result.items[j] = true
		}
	}
	return &result
}

// Return a new set containing the items found in both set and set2
func (set *Set[T]) Intersection(set2 Set[T]) *Set[T] {
	result := Set[T]{}
	result.items = make(map[T]bool)
	for i := range set2.items {
		_, present := set.items[i]
		if present {
			result.items[i] = true
		}
	}
	return &result
}

// Return a new set of items in set not found in set2
func (set *Set[T]) Difference(set2 Set[T]) *Set[T] {
	result := Set[T]{}
	result.items = make(map[T]bool)
	for i := range set.items {
		_, present := set2.items[i]
		if !present {
			result.items[i] = true
		}
	}
	return &result
}

// Return true if all items of set2 are in set
func (set *Set[T]) Subset(set2 Set[T]) bool {
	for i := range set.items {
		_, present := set2.items[i]
		if !present {
			return false
		}
	}
	return true
}
