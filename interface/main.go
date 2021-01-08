package main

import "fmt"

// HasName is ...
type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// Person is ...
type Person struct {
	Name string
}

// GetName is ...
func (person Person) GetName() string {
	return person.Name
}

// Animal is ...
type Animal struct {
	Name string
}

// GetName is ...
func (animal Animal) GetName() string {
	return animal.Name
}

// Interface kosong
func ups(i int, apapun interface{}) interface{} {
	if i == 1 && apapun == true {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}

func main() {
	var firas Person
	firas.Name = "Firas"

	sayHello(firas)

	cat := Animal{Name: "Push"}
	sayHello(cat)

	var kosong interface{} = ups(1, true)
	fmt.Println(kosong)
}
