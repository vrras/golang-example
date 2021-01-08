package main

import (
	"errors"
	"fmt"
)

// Pembagian is function
func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan NOL")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	result, error := Pembagian(10, 5)
	if error == nil {
		fmt.Println("Hasil", result)
	} else {
		fmt.Println("Error", error.Error())
	}
}
