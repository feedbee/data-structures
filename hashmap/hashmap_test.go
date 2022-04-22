package hashmap

import (
	"math/rand"
	"testing"
)

func TestSeq(t *testing.T) {
	hs := New(1250000)

	for i := 0; i < 1000000; i++ {
		hs.Put(i, i+1)
	}

	for i := 0; i < 1000000; i++ {
		contains(t, &hs, i, true)
		if r, _ := hs.Get(i); r != i+1 {
			t.Errorf("hs.Get(%d) = %d, want %d", i, r, i+1)
		}
	}

	for i := 0; i < 1000000; i++ {
		hs.Remove(i)
	}

	for i := 0; i < 1000000; i++ {
		contains(t, &hs, i, false)
		if _, found := hs.Get(i); found != false {
			t.Errorf("hs.Get(%d) = found=%t, want found=%t", i, found, false)
		}
	}
}

func TestRand(t *testing.T) {
	rand.Seed(33)
	hs := New()
	ss := make([]int, 0, 1000000)

	for i := 0; i < 1000000; i++ {
		rv := rand.Intn(1000000)
		hs.Put(rv, rv)
		ss = append(ss, rv)
	}

	for _, v := range ss {
		contains(t, &hs, v, true)
		if r, _ := hs.Get(v); r != v {
			t.Errorf("hs.Get(%d) = %d, want %d", v, r, v)
		}
	}

	for _, v := range ss {
		hs.Remove(v)
	}

	for _, v := range ss {
		contains(t, &hs, v, false)
	}
}

func TestDuplicates(t *testing.T) {
	hs := New()

	n := 18
	contains(t, &hs, n, false)

	hs.Put(n, 1)
	hs.Put(n, 2)
	hs.Put(n, 4)
	contains(t, &hs, n, true)
	if r, _ := hs.Get(n); r != 4 {
		t.Errorf("hs.Get(%d) = %d, want %d", n, r, 4)
	}

	hs.Remove(n)
	contains(t, &hs, n, false)

	hs.Put(n, 10)
	contains(t, &hs, n, true)
	if r, _ := hs.Get(n); r != 10 {
		t.Errorf("hs.Get(%d) = %d, want %d", n, r, 10)
	}

	hs.Remove(n)
	hs.Remove(n)
	hs.Remove(n)
	contains(t, &hs, n, false)
}

func contains(t *testing.T, hs *HashMap, v int, expectation bool) {
	t.Helper()
	if r := hs.Contains(v); r != expectation {
		t.Errorf("hs.Contains(%d) = %t, want %t", v, r, expectation)
	}
}
