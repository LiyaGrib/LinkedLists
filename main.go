package main

import (
	// "linkedLists/singlylinkedlists"
	"linkedLists/doublylinkedlists"
	"log"
)

func main() {
	list := doublylinkedlists.InitDoublyList()
	list.InsertFront("asdasd")	
	list.InsertFront("First")
	list.Display()
	log.Println("------------")


	str := "hehe"
	list.InsertBack(str)
	list.Display()
	



	// list := singlylinkedlists.ListInit()
	// list.InsertFront("First")
	// list.InsertBack("Second")
	// list.InsertBack("pupa")
	// list.Display()
	// log.Println("----------------------------")

	// data, err := list.GetBody(5)
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	log.Println(*data)
	// }
}
