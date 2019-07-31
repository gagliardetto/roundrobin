package roundrobin

import (
	"errors"
	"net/url"
	"sync"
)

// ErrNoElements is the error that servers dose not exists
var ErrNoElements = errors.New("no elements provided")

// RoundRobin is an interface for representing round-robin balancing.
type RoundRobin interface {
	Next() interface{}
	Append(elem interface{})
	Replace(elems []interface{})
	Len() int
}

type roundrobin struct {
	elements []interface{}
	next     int32
	mu       *sync.Mutex
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

// Next returns the next address
func (r *roundrobin) Next() interface{} {
	r.mu.Lock()
	sc := r.elements[r.next]
	r.next = (r.next + 1) % int32(len(r.elements))
	r.mu.Unlock()
	return sc
}

// Next returns the next address
func (r *roundrobin) Append(elem interface{}) {
	r.mu.Lock()
	r.elements = append(r.elements, elem)
	r.mu.Unlock()
}

// Next returns the next address
func (r *roundrobin) Replace(elems []interface{}) {
	r.mu.Lock()
	r.elements = elems
	r.mu.Unlock()
}

// Next returns the next address
func (r *roundrobin) Len() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return len(r.elements)
}
