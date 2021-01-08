package main

import "fmt"

// Customer is type of customer
type Customer struct {
	FirstName, Lastname string
	Address             string
	Age                 int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.FirstName, customer.Lastname)
}

func main() {
	var firas Customer
	firas.FirstName = "Firas"
	firas.Lastname = "Luthfi"
	firas.Address = "Kuningan"
	firas.Age = 23

	firas.sayHello("Putri")

	fmt.Println(firas.FirstName)
	fmt.Println(firas.Lastname)
	fmt.Println(firas.Address)
	fmt.Println(firas.Age)

	putri := Customer{
		FirstName: "Putri",
		Lastname:  "Indah",
		Address:   "Rumpin",
		Age:       22,
	}

	fmt.Println(putri.FirstName)
	fmt.Println(putri.Lastname)
	fmt.Println(putri.Address)
	fmt.Println(putri.Age)

	iyan := Customer{"Iyan", "Afrilliyan", "Bekasi", 28}
	fmt.Println(iyan.FirstName)
	fmt.Println(iyan.Lastname)
	fmt.Println(iyan.Address)
	fmt.Println(iyan.Age)
}
