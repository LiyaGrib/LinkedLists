package singlylinkedlists

func (list *List) DeleteFront() {
	current := list.Next
	list.Next = current.Next

	list.Len--
}

func (list *List) DeleteBack() {
	current := list.Next
	post := current.Next

	for current.Next != nil {
		current = current.Next
		post = post.Next
		if post.Next == nil {
			current.Next = nil
			break
		}
	}

	list.Len--
}
