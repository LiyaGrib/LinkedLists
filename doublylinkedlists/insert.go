package doublylinkedlists

func (list *List) InsertFront(body any) {
	node := Node{
		Body: body,
	}

	if list.head == nil {
		list.head = &node
		list.last = &node
	} else {
		current := list.head
		current.prev = &node

		node.next = current
		list.head = &node
	}

	list.len++
}

func (list *List) InsertFrontEl(node *Node) {
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

func (list *List) InsertBack(body any) {
	node := Node{
		Body: body,
	}

	if list.head == nil {
		list.head = &node
		list.last = &node
	} else {
		current := list.head
		for {
			if current.next == nil {
				current.next = &node
				node.prev = current
				break
			}

			current = current.next
		}
	}

	list.len++
}
