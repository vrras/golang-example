package main

import (
	"fmt"
	"reflect"
)

// Sample is ...
type Sample struct {
	Name string `required:"true" max:"10"`
}

type SampleLain struct {
	Name string
}

func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	samples := Sample{"Firas"}
	sampleType := reflect.TypeOf(samples)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	samples.Name = ""
	fmt.Println(isValid(samples))

	sampleLain := SampleLain{""}
	fmt.Println(isValid(sampleLain))
}
