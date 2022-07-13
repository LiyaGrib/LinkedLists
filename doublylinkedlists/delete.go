package doublylinkedlists

import (
	"errors"
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
		err = errors.New("{DeleteElement}: n > list.len")
		return
	}
	current := list.head

	for i := uint(1); i < n; i++ {
		current = current.next
	}

	if current.next == nil {
		list.last = current.prev
		list.last.next = nil
	}

	if current.prev == nil {
		list.head = current.next
		list.head.prev = nil
	}

	if current.next != nil && current.prev != nil {
		currentRight := current.next
		currentRight.prev = current.prev
		currentLeft := current.prev
		currentLeft.next = current.next
	}
	list.len--

	return
}
