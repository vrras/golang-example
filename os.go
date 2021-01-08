package main

import (
	"fmt"
	"os"
)

func main() {
	// Args
	args := os.Args
	fmt.Println("Argument :")
	for i := 1; i < len(args); i++ {
		fmt.Println(args[i])
	}

	// Hostname
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", hostname)
	} else {
		fmt.Println("Error", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)

}
