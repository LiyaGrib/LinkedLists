package doublylinkedlists

import (
	"errors"
	"log"
)

func (list *List) DeleteFront() {
	if list.head != nil {
		list.head = list.head.next
		list.head.prev = nil
	}

	list.len--
}

func (list *List) DeleteBack() {
	if list.last != nil {
		list.last = list.last.prev
		list.last.next = nil
	}

	list.len--
}

func (list *List) DeleteElement(n uint) (err error) {
	if list.len < n {
		err = errors.New("{DeleteElement}: n > list.Len")
		return
	}
	current := list.head

	for i := uint(1); i < n; i++ {
		current = current.next
	}
	log.Println(current)

	return
}
