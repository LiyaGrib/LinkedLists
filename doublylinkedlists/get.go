package doublylinkedlists

import "errors"

// Get node's body from 1
//
// If n > list.len < n returns nil
func (list *List) GetBody(n uint) (body *any, err error) {
	if list.len < n {
		err = errors.New("{GetBody}: n > list.len")
		return
	}

	current := list.head

	for i := uint(1); i < n; i++ {
		current = current.next
	}
	body = &current.Body
	return
}
