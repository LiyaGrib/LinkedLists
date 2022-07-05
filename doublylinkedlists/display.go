package doublylinkedlists

import "log"

func (list *List) Display() {
	if list.head == nil {
		return
	}
	current := list.head

	for {
		log.Println(current.Body)
		if current.next == nil {
			break
		}

		current = current.next
	}
}
