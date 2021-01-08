package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var person map[string]string = nil
	fmt.Println(person)

	newPerson := NewMap("")
	if newPerson == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(newPerson)
	}
}
