package lists

import (
	"fmt"
)

/**
initial and very basic linked list
*/

func NewListListV1[T any]() *LinkedListV1[T] {
	return &LinkedListV1[T]{}
}

type LinkedListV1[T any] struct {
	head *node[T]
}

func (l *LinkedListV1[T]) Length() int {
	//this could be better, but for now, just walk the list and count
	var length int = 0
	current := l.head
	for current != nil {
		length++
		current = current.next
	}
	return length
}

func (l *LinkedListV1[T]) Get(index int) (T, bool) {
	var t T
	if index >= l.Length() {
		return t, false
	}
	counter := 0
	current := l.head
	for counter < index {
		current = current.next
		counter++
	}

	return current.value, true
}

func (l *LinkedListV1[T]) Append(v T) {
	temp := &node[T]{
		next:  nil,
		value: v,
	}

	if l.head == nil {
		l.head = temp
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = temp
	}
}

func (l *LinkedListV1[T]) String() string {
	buffer := "["
	current := l.head
	if current != nil {
		buffer += fmt.Sprintf("%v", current.value)
		current = current.next
		for current != nil {
			buffer += fmt.Sprintf(",%v", current.value)
			current = current.next
		}
	}

	buffer += "]"

	return buffer
}

func (l *LinkedListV1[T]) Print() {

	fmt.Println(l.String())
}
