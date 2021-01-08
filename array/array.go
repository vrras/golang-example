package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Firas"
	names[1] = "Luthfi"
	names[2] = "Dwiyansyah"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var shoesNumber = [3]int{
		31,
		32,
		33,
	}

	fmt.Println(shoesNumber)
	fmt.Println(shoesNumber[0])
	fmt.Println(shoesNumber[1])
	fmt.Println(shoesNumber[2])

	fmt.Println(len(names))
	fmt.Println(len(shoesNumber))
}
