package roundrobin

import (
	"errors"
	"net/url"
	"sync"
)

// New returns a RoundRobin of []interface{}
func New(elements []interface{}) (RoundRobin, error) {
	if elements == nil || len(elements) == 0 {
		return nil, ErrNoElements
	}

	return newRR(elements), nil
}
func newRR(elems []interface{}) RoundRobin {
	return &roundrobin{
		elements: elems,
		mu:       &sync.Mutex{},
	}
}

// New returns a RoundRobin of []*url.URL
func NewURLs(elements []*url.URL) (RoundRobin, error) {
	if elements == nil || len(elements) == 0 {
		return nil, ErrNoElements
	}

	var elems []interface{}
	for _, v := range elements {
		elems = append(elems, v)
	}

	return newRR(elems), nil
}

// New returns a RoundRobin of []string
func NewStrings(elements []string) (RoundRobin, error) {
	if elements == nil || len(elements) == 0 {
		return nil, ErrNoElements
	}

	var elems []interface{}
	for _, v := range elements {
		elems = append(elems, v)
	}

	return newRR(elems), nil
}

// New returns a RoundRobin of []int
func NewInts(elements []int) (RoundRobin, error) {
	if elements == nil || len(elements) == 0 {
		return nil, ErrNoElements
	}

	var elems []interface{}
	for _, v := range elements {
		elems = append(elems, v)
	}

	return newRR(elems), nil
}

// ErrNoElements is the error that servers dose not exists
var ErrNoElements = errors.New("no elements provided")

// RoundRobin is an interface for representing round-robin balancing.
type RoundRobin interface {
	Next() interface{}
	Append(elem interface{})
	Replace(elems []interface{})
	Len() int
	IterateAll(func(interface{}) bool)
}

type roundrobin struct {
	elements []interface{}
	next     int32
	mu       *sync.Mutex
}

// Next returns the next address
func (r *roundrobin) Next() interface{} {
	r.mu.Lock()
	sc := r.elements[r.next]
	r.next = (r.next + 1) % int32(len(r.elements))
	r.mu.Unlock()
	return sc
}

// Append adds an element to the list of elements;
// if the element is nil, it is ignored and not added.
func (r *roundrobin) Append(elem interface{}) {
	if elem == nil {
		return
	}
	r.mu.Lock()
	r.elements = append(r.elements, elem)
	r.mu.Unlock()
}

// Replace replaces the list of elements with the provided one.
func (r *roundrobin) Replace(elems []interface{}) {
	if elems == nil || len(elems) == 0 {
		return
	}
	r.mu.Lock()
	r.next = 0
	r.elements = elems
	r.mu.Unlock()
}

// Len returns the count of how many elements are in the list.
func (r *roundrobin) Len() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return len(r.elements)
}

// IterateAll iterates over all elements
// (WARNING: it's a blocking operation)
func (r *roundrobin) IterateAll(iterator func(interface{}) bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, elem := range r.elements {
		doContinue := iterator(elem)
		if !doContinue {
			break
		}
	}
}
