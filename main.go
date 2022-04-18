package main

import (
	"fmt"
	"datastruct/linkedlist"
	"datastruct/queue"
	"datastruct/stack"
)

func main() {
	ls := linkedlist.NewLinkedList[string]()
	ls.Append("a")
	fmt.Println(ls.AddAt(1, "b"))
	fmt.Println(ls.GetSize())
	fmt.Println(ls.GetAt(1))
	fmt.Println(ls.RemoveAt(1))

	s := stack.NewStack[string]()
	s.Push("a")
	fmt.Println(s.GetSize())
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())

	q := queue.NewQueue[string]()
	q.Enqueue("a")
	fmt.Println(q.GetSize())
	fmt.Println(q.Peek())
	fmt.Println(q.Dequeue())
}