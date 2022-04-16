package list

import "testing"

func TestList(t *testing.T) {
	l := New()
	ln := l.AddValue(5)

	if len := l.Len(); len != 1 {
		t.Errorf("l.Len() = %d, want %d", len, 1)
		// t.Errorf("l.root.next = %p, l.root.prev = %p; both should both be nil or %p", l.root.next, l.root.prev, root)
	}

	if l.head != ln || l.tail != ln {
		t.Errorf("l.head = %p, l.tail = %p; both should both be nil or %p", l.head, l.tail, ln)
	}
}
