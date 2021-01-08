package main

import "fmt"

func main() {
	user := "vrras"
	password := "123"

	if user == "vrras" && password == "123" {
		fmt.Println("Hello Firas")
	} else if user == "vrras" && password != "123" {
		fmt.Println("Password kamu salah")
	} else {
		fmt.Println("Maaf kamu bukan Firas")
	}

	if length := len(user); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}
