package linkedlist

import "errors"


/* 連結リスト */
type LinkedList[T any] struct {
	top *Node[T]
	size int
}


type Node[T any] struct {
	next *Node[T]
	data T 
}


func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}


func (ls *LinkedList[T]) GetAt(index int) (T, error) {
	if ls.size < index {
		var ret T
		return ret, errors.New("LinkedList Error: index out of range")
	}
	node := ls.top
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node.data, nil
}


func (ls *LinkedList[T]) GetSize() int {
	return ls.size
}


func (ls *LinkedList[T]) Append(data any) {
	ls.size += 1
	n := Node[T]{}
	n.data = data.(T)
	if ls.top == nil {
		ls.top = &n
		return
	}	
	node := ls.top
	for node.next != nil {
		node = node.next
	}
	node.next = &n
}


func (ls *LinkedList[T]) AddAt(index int, data any) bool {
	if ls.size < index {
		return false
	}
	ls.size += 1
	n := Node[T]{}
	n.data = data.(T)

	if index == 0 {
		n.next = ls.top
		ls.top = &n
		return true
	}

	pre := ls.top
	for i := 0; i < index - 1; i++ {
		pre = pre.next
	}

	n.next = pre.next
	pre.next = &n
	return true
}


func (ls *LinkedList[T]) RemoveAt(index int) error {
	if ls.size < index {
		return errors.New("LinkedList Error: index out of range")
	}
	ls.size -= 1
	pre := ls.top
	for i := 0; i < index - 1; i++ {
		pre = pre.next
	}

	n := pre.next
	pre.next = n.next
	return nil
}
