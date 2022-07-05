package main

import (
	"log"
	// "linkedLists/singlylinkedlists"
	"linkedLists/doublylinkedlists"
)

func main() {
	list := doublylinkedlists.InitDoublyList()

	list.InsertFront("asdasd4")
	list.InsertFront("asdasd3")
	list.InsertFront("asdasd2")
	list.InsertFront("asdasd1")
	list.Display()
	log.Println("------------")
	// list.DeleteFront()
	// list.DeleteBack()
	err := list.DeleteElement(3)
	if err != nil {
		log.Println(err)
	}
	list.Display()

	// str := "hehe"
	// list.InsertBack(str)
	// list.Display()

	// var ss = []string {"asd", "dsa", "qwe"}
	// for i := 0; i < 3; i++ {
	// 	el := doublylinkedlists.Element{}
	// 	el.Body = ss[i]
	// 	list.InsertFrontEl(&el)
	// }
	// list.Display()

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
