package singlylinkedlists

func (list *List) InsertFront(body any) {
	node := &Node{
		Body: body,
		next: list.head,
	}
	list.head = node

	list.len++
}

func (list *List) InsertBack(body any) {
	node := &Node{
		Body: body,
	}
	current := list.head

	for current.next != nil {
		current = current.next
	}

	current.next = node

	list.len++
}

func (list *List) InsertAny(body any, n uint) {
	node := &Node{
		Body: body,
	}
	current := list.head

	if n == 1 {
		list.head = node
		node.next = current
	} else {
		for i := uint(1); i < n-1; i++ {
			current = current.next
		}
		node.next = current.next
		current.next = node
	}

	list.len++
}
