package Deque

import (
	"container/list"
)

// Deque that extends container/list capabilities just a little
type Deque struct {
	*list.List
}

func New() *Deque {
	return &Deque{list.New()}
}

// Remove de Element in the Front of the Deque and returns it's value
func (q *Deque) PopFront() any {
	return q.Remove(q.Front())
}

// Remove de Element in the Back of the Deque and returns it's value
func (q *Deque) PopBack() any {
	return q.Remove(q.Back())
}
