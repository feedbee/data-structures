package hashset

import "math"

type HashSet struct {
	array    []*Node
	capacity int
	len      int
}

func New(initialCapacity ...int) HashSet {
	initialCapacityInt := 64 //default
	if len(initialCapacity) == 1 {
		initialCapacityInt = initialCapacity[0]
	} else if len(initialCapacity) != 0 {
		panic("Too many arguments for hashset.New(initialCapacity?)")
	}

	hs := HashSet{nil, initialCapacityInt, 0}
	hs.Init()
	return hs
}

func (hs *HashSet) Init() {
	hs.array = make([]*Node, hs.capacity)
}

func (hs *HashSet) Len() int {
	return hs.len
}

func (hs *HashSet) Add(key int) {
	if hs.capacity < 1 {
		hs.adjustCapacity(64)
	}

	prev, node, hash := hs.locate(key)
	if node != nil {
		return // duplicate
	}

	newNode := &Node{key, nil}
	if prev != nil {
		prev.Next = newNode // last node
	} else {
		hs.array[hash] = newNode // first=last node
	}
	hs.len++
	hs.adjustCapacityIfNeeded()
}

func (hs *HashSet) Remove(key int) {
	prev, node, hash := hs.locate(key)
	if node == nil {
		return // not found
	}

	// found
	if prev == nil { // first
		hs.array[hash] = node.Next
	} else { // not first
		prev.Next = node.Next
	}
	hs.len--
}

func (hs *HashSet) Contains(key int) bool {
	_, node, _ := hs.locate(key)
	return node != nil
}

func (hs *HashSet) locate(key int) (prev, node *Node, hash int) {
	hash = hs.hash(key)
	if node = hs.array[hash]; node == nil {
		return // not found
	} else {
		for node != nil && node.Key != key {
			prev = node
			node = node.Next
		}
		// found
		return // node != nil -> found, prev - last node if not found
	}
}

// Grow the storage is > 80% capacity is used, or reduce storage if <20% capacity is used
func (hs *HashSet) adjustCapacityIfNeeded() {
	if hs.len > int(hs.capacity*80/100) { // 80%
		hs.adjustCapacity(hs.capacity * 2)
	} else if hs.len < int(hs.capacity*20/100) { // 20%
		hs.adjustCapacity(hs.capacity / 2)
	}
}

func (hs *HashSet) adjustCapacity(newCapacity int) {
	newArray := make([]*Node, newCapacity)
	for _, node := range hs.array {
		for node != nil {
			newKey := hs.hash(node.Key, newCapacity)
			if nd := newArray[newKey]; nd == nil {
				newArray[newKey] = node // first = last
			} else {
				for nd.Next != nil {
					nd = nd.Next
				}
				nd.Next = node // last
			}
			node, node.Next = node.Next, nil
		}
	}

	hs.array = newArray
	hs.capacity = newCapacity
}

type Node struct {
	Key  int
	Next *Node
}

func (hs *HashSet) hash(key int, capacity ...int) int {
	capacityInt := hs.capacity
	if len(capacity) > 0 {
		capacityInt = capacity[0]
	}
	return int(math.Abs(float64(key))) % capacityInt
}
