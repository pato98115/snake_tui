package queue

import (
	"container/list"
)

// Queue that extends container/list capabilities just a little
type Queue struct {
	*list.List
}

// Remove de Element in the Front of the Queue and returns it's value
func (q *Queue) PopFront() any {
	return q.Remove(q.Front())
}
