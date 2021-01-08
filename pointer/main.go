package main

import "fmt"

// Address is ...
type Address struct {
	City, Province, Country string
}

// Man is ...
type Man struct {
	Name string
}

// Married is ...
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

// ChangeAddressToSingapore is ...
func ChangeAddressToSingapore(address *Address) {
	address.Country = "Singapore"
}

func main() {
	var address1 Address = Address{"Kuningan", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Cirebon"

	*address2 = Address{"Purworejo", "Jawa Tengah", "Indonesia"}

	address3 = &Address{"Purworejo", "Jawa Tengah", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

	var alamat = Address{"Purworejo", "Jawa Tengah", "Indonesia"}
	ChangeAddressToSingapore(&alamat)
	fmt.Println(alamat)

	firas := Man{"Firas"}
	firas.Married()

	fmt.Println(firas.Name)
}
