package list

// ListNode

type ListNode struct {
	Val  any
	next *ListNode
}

func (n *ListNode) Next() any {
	return n.next
}

// List

type List struct {
	head, tail *ListNode
	len        int
}

func (l *List) Init() *List {
	l.head, l.tail = nil, nil
	l.len = 0

	return l
}

func New() *List {
	return new(List).Init()
}

func (l *List) Len() int {
	return l.len
}

func (l *List) AddValue(val any) *ListNode {
	return l.insert(&ListNode{val, nil}, l.tail)
}

func (l *List) InsertValue(val any, after *ListNode) *ListNode {
	return l.insert(&ListNode{val, nil}, after)
}

func (l *List) insert(new, after *ListNode) *ListNode {
	if new == nil {
		panic("Can't insert nil")
	}

	// first one
	if after == nil {
		l.head = new
	} else {
		after.next = new
	}

	// last one
	if after == l.tail {
		l.tail = new
	}

	l.len++

	return new
}
