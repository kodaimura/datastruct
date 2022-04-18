package queue

import "errors"


/* キュー */
type Queue[T any] struct {
	top *Node[T]
	size int
}

type Node[T any] struct {
	next *Node[T]
	data T
}


func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}


func (q *Queue[T]) IsEmpty() bool {
	return q.top == nil
}


func (q *Queue[T]) Peek() (T, error) {
	if q.top == nil {
		var ret T
		return ret, errors.New("Queue Error: size 0")
	}
	return q.top.data, nil
}


func (q *Queue[T]) GetSize() int {
	return q.size
}


func (q *Queue[T]) Enqueue(item T) {
	q.size += 1
	if q.top == nil {
		q.top = &Node[T]{data: item}
		return
	}

	node := q.top
	for node.next != nil {
		node = node.next
	}
	node.next = &Node[T]{data: item}
}


func (q *Queue[T]) Dequeue() (T, error) {
	if q.top == nil {
		var ret T
		return ret, errors.New("Queue Error: size 0")
	}
	top := q.top
	q.top = q.top.next
	q.size -= 1
	return top.data, nil
}