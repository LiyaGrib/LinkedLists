package singlylinkedlists

func (list *List) DeleteFront() {
	current := list.head
	list.head = current.next

	list.len--
}

func (list *List) DeleteBack() {
	current := list.head
	post := current.next

	for current.next != nil {
		current = current.next
		post = post.next
		if post.next == nil {
			current.next = nil
			break
		}
	}

	list.len--
}
