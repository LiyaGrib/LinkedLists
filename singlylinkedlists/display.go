package singlylinkedlists

import "log"

func (list *List) Display() {
	if list.Next == nil {
		return
	}
	current := list.Next

	for {
		log.Println(current.Body)
		if current.Next == nil {
			break
		}

		current = current.Next
	}
}
