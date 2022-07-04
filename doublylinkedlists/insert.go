package doublylinkedlists

func (head *Head) InsertFront(body any) {
	el := Element{
		Body: body,
	}

	if head.Next == nil {
		head.Next = &el
	} else {
		current := head.Next
		current.Prev = &el

		el.Next = current
		head.Next = &el
	}

	head.Len++
}

func (head *Head) InsertBack(body any) {
	el := Element {
		Body: body,
	}

	if head.Next == nil {
		head.Next = &el
	} else {
		current := head.Next
		for {
			if current.Next == nil {
				current.Next = &el
				el.Prev = current
				break
			}

			current = current.Next
		}
	}
}
