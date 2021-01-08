package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 32)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println("Error", err.Error())
	}

	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value)

	itoa := strconv.Itoa(1000000)
	if err == nil {
		fmt.Println(itoa)
	} else {
		fmt.Println("Error", err.Error())
	}

	atoi, err := strconv.Atoi("1000000")
	if err == nil {
		fmt.Println(atoi)
	} else {
		fmt.Println("Error", err.Error())
	}
}
