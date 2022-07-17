package main

import (
	"log"

	"linkedLists/singlylinkedlists"
	// "linkedLists/doublylinkedlists"
)

func main() {
	list := singlylinkedlists.ListInit()
	list.InsertFront("A1")
	list.InsertBack("A2")
	list.InsertBack("A3")
	list.InsertBack("A4")
	list.Display()
	log.Println("----------------------------")

	list.InsertAny("kek", 3)
	list.Display()
	log.Println("----------------------------")

	// data, err := list.GetBody(3)
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	log.Println(*data)
	// }

	// list := doublylinkedlists.InitDoublyList()

	// list.InsertFront("asdasd4")
	// list.InsertFront("asdasd3")
	// list.InsertFront("asdasd2")
	// list.InsertFront("asdasd1")
	// list.Display()
	// log.Println("------------")

	// list.InsertAny("test", 2)
	// list.Display()
	// log.Println("------------")
	// list.InsertBack("pupa")
	// list.Display()

	// err := list.DeleteElements(1)
	// if err != nil {
	// }
	// list.Display()

	// var ss = []string {"asd", "dsa", "qwe"}
	// for i := len(ss) - 1; i >= 0; i-- {
	// 	el := doublylinkedlists.Node{}
	// 	el.Body = ss[i]
	// 	list.InsertFrontElements(&el)
	// }
	// list.Display()
}
