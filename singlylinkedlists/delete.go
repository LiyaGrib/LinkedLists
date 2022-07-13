package singlylinkedlists

func (list *List) DeleteFront() {
	if list.head != nil {
		list.head = list.head.next
	}

	list.len--
}

func (list *List) DeleteBack() {
	current := list.head

	if list.head.next == nil {
		list.head = nil
		return
	}

	for current.next.next != nil {
		current = current.next
	}

	current.next = nil

	list.len--
}

func (list *List) DeleteAny(n uint) {
	current := list.head

	if n == 1 {
		list.head = current.next
	} else {
		for i := uint(1); i < n-1; i++ {
			current = current.next
		}

		current.next = current.next.next
	}

	list.len--
}
