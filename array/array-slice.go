package main

import "fmt"

func main() {
	names := [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"Desember",
	}
	slice1 := names[4:7]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1[0] = "May Again"
	fmt.Println(slice1[0])
	fmt.Println(names)

	var slice2 = names[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Firas")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(names)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Firas"
	newSlice[1] = "Luthfi"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}
