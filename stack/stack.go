package stack

import "errors"


/* スタック */
type Stack[T any] struct {
	top *Node[T]
	size int
}

type Node[T any] struct {
	next *Node[T]
	data T
}


func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil
}


func (s *Stack[T]) Peek() (T, error) {
	if s.top == nil {
		var ret T
		return ret, errors.New("Stack Error: size 0")
	}
	return s.top.data, nil
}


func (s *Stack[T]) GetSize() int {
	return s.size
}


func (s *Stack[T]) Push(item T) {
	n := &Node[T]{data: item}
	n.next = s.top
	s.top = n
	s.size += 1
}


func (s *Stack[T]) Pop() (T, error) {
	if s.top == nil {
		var ret T
		return ret, errors.New("Stack Error: size 0")
	}
	top := s.top
	s.top = s.top.next
	s.size -= 1
	return top.data, nil
}