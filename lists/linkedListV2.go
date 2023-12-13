package lists

import "fmt"

func NewListListV2[T any]() *LinkedListV2[T] {
	return &LinkedListV2[T]{}
}

type LinkedListV2[T any] struct {
	length int
	head   *node[T]
}

func (l *LinkedListV2[T]) Length() int {
	return l.length
}

func (l *LinkedListV2[T]) Append(v T) {
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
	l.length++
}

func (l *LinkedListV2[T]) String() string {
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

func (l *LinkedListV2[T]) Print() {

	fmt.Println(l.String())
}
