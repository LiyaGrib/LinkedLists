package doublylinkedlists

import "log"

func (head *Head) Display() {
	if head.Next == nil {
		return
	}
	current := head.Next

	for {
		log.Println(current.Body)
		if current.Next == nil {
			break
		}

		current = current.Next
	}
}