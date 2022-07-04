package main

import (
	"errors"
	"log"
)

type (
	Element struct {
		Body any
		Next *Element
	}

	List struct {
		Len  uint
		Next *Element
	}
)

func ListInit() (l *List) {
	l = new(List)
	l.Len = 0
	l.Next = nil
	return l
}

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

func main() {
	list := ListInit()
	list.InsertFront("First")
	list.InsertBack("Second")
	list.InsertBack("pupa")
	list.Display()
	log.Println("----------------------------")

	data, err := list.GetBody(5)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(*data)
	}
}
