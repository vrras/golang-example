package main

import "fmt"

func main() {
	var name = "Firas"
	var e byte = name[2]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(eString)
}
