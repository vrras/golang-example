package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func getFullName() (string, string, string) {
	return "Firas", "Luthfi", "Dwiyansyah"
}

func getCompleteName() (firstName string, middleName, lastName string) {
	firstName = "Firas"
	middleName = "Luthfi"
	lastName = "Dwiyansyah"

	return
}

// Function variadic
func sumAll(name string, numbers ...int) int {
	fmt.Println(name, "sudah masuk")
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Function value
func getGoodBye(name string) string {
	return "Good bye " + name
}

// Function as parameter
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

//Anonymous Function
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// Recursive function
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

// Defer function
func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()

	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("Result", result)
}

// Panic function
func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	sayHello("Firas", "Luthfi")
	fmt.Println(getHello("Firas Luthfi Dwiyansyah"))

	firstName, _, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	firstName2, middleName2, lastName2 := getCompleteName()
	fmt.Println(firstName2)
	fmt.Println(middleName2)
	fmt.Println(lastName2)

	fmt.Println(sumAll("Firas", 20, 10, 10, 20, 1, 2, 3, 2, 1))

	numbers := []int{1, 2, 4, 3, 2, 3, 5}
	fmt.Println(sumAll("Firas", numbers...))

	goodbye := getGoodBye("Firas")
	fmt.Println(goodbye)

	sayHelloWithFilter("Firas", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("firas", blacklist)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})

	factorial := factorialRecursive(5)
	fmt.Println(factorial)

	// Closure
	counter := 0
	name := "Firas"

	increment := func() {
		name := "Luthfi"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	fmt.Println(name)

	increment()
	increment()

	fmt.Println(counter)

	runApplication(10)
	runApp(false)
}
