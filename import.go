package main

import "go-example/helpers"

func main() {
	helpers.SayHello(helpers.Name)
	helpers.Name = "Chandra"
	helpers.SayHello(helpers.Name)
}
