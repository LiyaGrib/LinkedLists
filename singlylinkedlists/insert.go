package singlylinkedlists

func (list *List) InsertFront(body any) {
	node := Node{
		Body: body,
		next: list.head,
	}
	list.head = &node

	list.len++
}

func (list *List) InsertBack(body any) {
	node := Node{
		Body: body,
	}
	current := list.head

	for current.next != nil {
		current = current.next
	}

	current.next = &node

	list.len++
}
