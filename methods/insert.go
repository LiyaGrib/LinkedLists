package methods

func (list *List) InsertFront(body any) {
	el := Element{
		Body: body,
		Next: list.Next,
	}
	list.Next = &el

	list.Len++
}

func (list *List) InsertBack(body any) {
	el := Element{
		Body: body,
	}
	current := list.Next

	for current.Next != nil {
		current = current.Next
	}

	current.Next = &el

	list.Len++
}
