package singlylinkedlists

import "errors"

// Get element's body from 1
//
// If n > list.Len < n returns nil
func (list *List) GetBody(n uint) (body *any, err error) {
	if list.Len < n {
		err = errors.New("{GetBody}: n > list.Len")
		return
	}

	current := list.Next

	for i := uint(1); i < n; i++ {
		current = current.Next
	}
	body = &current.Body
	return
}
