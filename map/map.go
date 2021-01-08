package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Firas",
		"address": "Kuningan",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	fmt.Println(len(person))

	book := make(map[string]string)
	book["title"] = "Belajar golang"
	book["author"] = "Firas"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}
