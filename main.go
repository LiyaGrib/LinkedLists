package main

import (
	"singlyLinkedLists/methods"
	"log"
)

func main() {
	list := methods.ListInit()
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
