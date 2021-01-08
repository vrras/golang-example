package main

import "fmt"

func random() interface{} {
	return "Ups"
}

func main() {
	var randomData interface{} = random()
	var resultString string = randomData.(string)
	fmt.Println("String", resultString)

	var result interface{} = random()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	case bool:
		fmt.Println("Boolean", value)
	default:
		fmt.Println("Unknown type")
	}
}
