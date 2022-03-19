package main

import (
	"DataStructureWithGo/LinkedList"
	"fmt"
)

func main() {
	var list = LinkedList.New()
	list.Add(5, 10, 15, 20, 25, 30, 35, 40)
	list.Add(50, 55)
	list.Delete(50)
	list.AddAt(0, 0)
	list.DeleteAt(0)

	for i := 0; i < list.Len(); i++ {
		val, _ := list.At(i)
		fmt.Print(val, " ")
	}
}
