package main

import "fmt"

func main() {
	name := "Luthfi"

	switch name {
	case "Firas":
		fmt.Println("Hello Firas")
	case "Luthfi":
		fmt.Println("Hello Luthfi")
	case "Dwiyansyah":
		fmt.Println("Hello Dwiyansyah")
	default:
		fmt.Println("Kenalan dong")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama kepanjangan")
	case false:
		fmt.Println("Nama bener")
	}

	length2 := len(name)
	switch {
	case length2 > 10:
		fmt.Println("Nama kepanjangan")
	case length2 > 5:
		fmt.Println("Nama lumayan kepanjangan")
	default:
		fmt.Println("Nama sudah benar")
	}
}
