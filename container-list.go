package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Firas")
	data.PushBack("Luthfi")
	data.PushBack("Dwiyansyah")
	data.PushFront("Putri")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	// Depan ke belakang
	// for element := data.Front(); element != nil; element = element.Next() {
	// 	fmt.Println(element.Value)
	// }

	// Belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
