package doublylinkedlists

import (
	"errors"
	"log"
)

func (list *List) InsertFront(body any) {
	node := &Node{
		Body: body,
	}

	if list.head == nil {
		list.head = node
		list.last = node
	} else {
		current := list.head
		current.prev = node

		node.next = current
		list.head = node
	}

	list.len++
}

func (list *List) InsertFrontElements(node *Node) {
	if list.head == nil {
		list.head = node
	} else {
		current := list.head

		current.prev = node
		node.next = current

		list.head = node
	}

	list.len++
}

func (list *List) InsertReplace(body any, n uint) (err error) {
	node := &Node{
		Body: body,
	}

	if list.len < n {
		err = errors.New("{InsertReplace}: n > list.len")
		return
	}

	if list.head == nil {
		list.head = node
		list.last = node
	} else {
		current := list.head

		for i := uint(1); i < n; i++ {
			current = current.next
		}

		if current.prev == nil {
			current.next.prev = node
			node.next = current.next
			list.head = node
		}

		if current.next == nil {
			current.prev.next = node
			node.prev = current.prev
			list.last = node
		}

		if current.prev != nil && current.next != nil {
			current.prev.next = node
			current.next.prev = node

			node.prev = current.prev
			node.next = current.next
		}
	}

	return
}

func (list *List) InsertAny(body any, n uint) {
	node := &Node{
		Body: body,
	}

	if list.len+1 < n {
		log.Println("{InsertAny}: n > list.len + 1")
		return
	}

	nodeRight := list.head

	if n == list.len+1 {
		for {
			if nodeRight.next == nil {
				nodeRight.next = node
				node.prev = nodeRight
				break
			}

			nodeRight = nodeRight.next
		}
	} else {
		for i := uint(1); i < n; i++ {
			nodeRight = nodeRight.next
		}
		if nodeRight.prev == nil {
			node.next = nodeRight
			node.prev = nil
			list.head = node
		} else {

			node.next = nodeRight
			nodeRight.prev.next = node
			node.prev = nodeRight.prev
			nodeRight.prev = node
		}
	}

	list.len++
}

func (list *List) InsertBack(body any) {
	node := &Node{
		Body: body,
	}

	if list.head == nil {
		list.head = node
		list.last = node
	} else {
		current := list.head
		for {
			if current.next == nil {
				current.next = node
				node.prev = current
				break
			}

			current = current.next
		}
	}

	list.len++
}
